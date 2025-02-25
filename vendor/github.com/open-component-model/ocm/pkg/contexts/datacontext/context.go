// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package datacontext

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"sync"

	"github.com/mandelsoft/logging"

	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/datacontext/action/handlers"
	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/open-component-model/ocm/pkg/finalizer"
	ocmlog "github.com/open-component-model/ocm/pkg/logging"
	"github.com/open-component-model/ocm/pkg/runtime"
	"github.com/open-component-model/ocm/pkg/utils"
)

const OCM_CONTEXT_SUFFIX = ".context" + common.OCM_TYPE_GROUP_SUFFIX

// BuilderMode controls the handling of unset information in the
// builder configuration when calling the New method.
type BuilderMode int

const (
	// MODE_SHARED uses the default contexts for unset nested context types.
	MODE_SHARED BuilderMode = iota
	// MODE_DEFAULTED uses dedicated context instances configured with the
	// context type specific default registrations.
	MODE_DEFAULTED
	// MODE_EXTENDED uses dedicated context instances configured with
	// context type registrations extending the default registrations.
	MODE_EXTENDED
	// MODE_CONFIGURED uses dedicated context instances configured with the
	// context type registrations configured with the actual state of the
	// default registrations.
	MODE_CONFIGURED
	// MODE_INITIAL uses completely new contexts for unset nested context types
	// and initial registrations.
	MODE_INITIAL
)

func (m BuilderMode) String() string {
	switch m {
	case MODE_SHARED:
		return "shared"
	case MODE_DEFAULTED:
		return "defaulted"
	case MODE_EXTENDED:
		return "extended"
	case MODE_CONFIGURED:
		return "configured"
	case MODE_INITIAL:
		return "initial"
	default:
		return fmt.Sprintf("(invalid %d)", m)
	}
}

func Mode(m ...BuilderMode) BuilderMode {
	return utils.OptionalDefaulted(MODE_EXTENDED, m...)
}

type ContextIdentity = finalizer.ObjectIdentity

type ContextProvider interface {
	// AttributesContext returns the shared attributes
	AttributesContext() AttributesContext
}

// Delegates is the interface for common
// Context features, which might be delegated
// to aggregated contexts.
type Delegates interface {
	ocmlog.LogProvider
	handlers.ActionsProvider
}

type _delegates struct {
	logging logging.Context
	actions handlers.Registry
}

func (d _delegates) LoggingContext() logging.Context {
	return d.logging
}

func (d _delegates) Logger(messageContext ...logging.MessageContext) logging.Logger {
	return d.logging.Logger(messageContext)
}

func (d _delegates) GetActions() handlers.Registry {
	return d.actions
}

func ComposeDelegates(l logging.Context, a handlers.Registry) Delegates {
	return _delegates{l, a}
}

type ContextBinder interface {
	// BindTo binds the context to a context.Context and makes it
	// retrievable by a ForContext method
	BindTo(ctx context.Context) context.Context
}

// Context describes a common interface for a data context used for a dedicated
// purpose.
// Such has a type and always specific attribute store.
// Every Context can be bound to a context.Context.
type Context interface {
	ContextBinder
	ContextProvider
	Delegates

	// GetType returns the context type
	GetType() string
	GetId() ContextIdentity

	GetAttributes() Attributes

	Finalize() error
	Finalizer() *finalizer.Finalizer
}

type InternalContext interface {
	Context
	finalizer.RecorderProvider
	GetKey() interface{}
	Cleanup() error
}

////////////////////////////////////////////////////////////////////////////////

// CONTEXT_TYPE is the global type for an attribute context.
const CONTEXT_TYPE = "attributes" + OCM_CONTEXT_SUFFIX

type AttributesContext interface {
	Context

	BindTo(ctx context.Context) context.Context
}

// AttributeFactory is used to atomicly create a new attribute for a context.
type AttributeFactory func(Context) interface{}

type Attributes interface {
	finalizer.Finalizable

	GetAttribute(name string, def ...interface{}) interface{}
	SetAttribute(name string, value interface{}) error
	SetEncodedAttribute(name string, data []byte, unmarshaller runtime.Unmarshaler) error
	GetOrCreateAttribute(name string, creator AttributeFactory) interface{}
}

// DefaultContext is the default context initialized by init functions.
var DefaultContext = NewWithActions(nil, handlers.DefaultRegistry())

// ForContext returns the Context to use for context.Context.
// This is either an explicit context or the default context.
func ForContext(ctx context.Context) AttributesContext {
	c, _ := ForContextByKey(ctx, key, DefaultContext)
	if c == nil {
		return nil
	}
	return c.(AttributesContext)
}

// WithContext create a new Context bound to a context.Context.
func WithContext(ctx context.Context, parentAttrs Attributes) (Context, context.Context) {
	c := New(parentAttrs)
	return c, c.BindTo(ctx)
}

////////////////////////////////////////////////////////////////////////////////

type Updater interface {
	Update() error
}

type UpdateFunc func() error

func (u UpdateFunc) Update() error {
	return u()
}

type delegates = Delegates

////////////////////////////////////////////////////////////////////////////////

var key = reflect.TypeOf(_context{})

type _context struct {
	*contextBase
	updater Updater
}

// gcWrapper is used as garbage collectable
// wrapper for a context implementation
// to establish a runtime finalizer.
type gcWrapper struct {
	GCWrapper
	*_context
}

func (w *gcWrapper) SetContext(c *_context) {
	w._context = c
}

// New provides a root attribute context.
func New(parentAttrs ...Attributes) AttributesContext {
	return NewWithActions(utils.Optional(parentAttrs...), handlers.NewRegistry(nil, handlers.DefaultRegistry()))
}

func NewWithActions(parentAttrs Attributes, actions handlers.Registry) AttributesContext {
	return newWithActions(MODE_DEFAULTED, parentAttrs, actions)
}

func newWithActions(mode BuilderMode, parentAttrs Attributes, actions handlers.Registry) AttributesContext {
	c := &_context{}
	c.contextBase = newContextBase(c, CONTEXT_TYPE, key, parentAttrs, &c.updater,
		ComposeDelegates(logging.NewWithBase(ocmlog.Context()), handlers.NewRegistry(nil, actions)),
	)
	return SetupContext(mode, FinalizedContext[gcWrapper](c))
}

func (c *_context) Actions() handlers.Registry {
	if c.updater != nil {
		c.updater.Update()
	}
	return c.contextBase.GetActions()
}

func (c *_context) LoggingContext() logging.Context {
	if c.updater != nil {
		c.updater.Update()
	}
	return c.contextBase.LoggingContext()
}

func (c *_context) Logger(messageContext ...logging.MessageContext) logging.Logger {
	if c.updater != nil {
		c.updater.Update()
	}
	return c.contextBase.Logger(messageContext...)
}

////////////////////////////////////////////////////////////////////////////////

var contextrange, attrsrange = finalizer.NumberRange{}, finalizer.NumberRange{}

type _attributes struct {
	sync.RWMutex
	id         uint64
	ctx        Context
	parent     Attributes
	updater    *Updater
	attributes map[string]interface{}
}

var _ Attributes = &_attributes{}

func NewAttributes(ctx Context, parent Attributes, updater *Updater) Attributes {
	return newAttributes(ctx, parent, updater)
}

func newAttributes(ctx Context, parent Attributes, updater *Updater) *_attributes {
	return &_attributes{
		id:         attrsrange.NextId(),
		ctx:        ctx,
		parent:     parent,
		updater:    updater,
		attributes: map[string]interface{}{},
	}
}

func (c *_attributes) Finalize() error {
	list := errors.ErrListf("finalizing attributes")
	for n, a := range c.attributes {
		if f, ok := a.(finalizer.Finalizable); ok {
			list.Addf(nil, f.Finalize(), "attribute %s", n)
		}
	}
	return list.Result()
}

func (c *_attributes) GetAttribute(name string, def ...interface{}) interface{} {
	if *c.updater != nil {
		(*c.updater).Update()
	}
	c.RLock()
	defer c.RUnlock()
	if a := c.attributes[name]; a != nil {
		return a
	}
	if c.parent != nil {
		if a := c.parent.GetAttribute(name); a != nil {
			return a
		}
	}
	return utils.Optional(def...)
}

func (c *_attributes) SetEncodedAttribute(name string, data []byte, unmarshaller runtime.Unmarshaler) error {
	s := DefaultAttributeScheme.Shortcuts()[name]
	if s != "" {
		name = s
	}
	v, err := DefaultAttributeScheme.Decode(name, data, unmarshaller)
	if err != nil {
		return err
	}
	c.SetAttribute(name, v)
	return nil
}

func (c *_attributes) setAttribute(name string, value interface{}) error {
	c.Lock()
	defer c.Unlock()

	_, err := DefaultAttributeScheme.Encode(name, value, nil)
	if err != nil && !errors.IsErrUnknownKind(err, "attribute") {
		return err
	}
	old := c.attributes[name]
	if old != nil && old != value {
		if c, ok := old.(io.Closer); ok {
			c.Close()
		}
	}
	value, err = DefaultAttributeScheme.Convert(name, value)
	if err != nil && !errors.IsErrUnknownKind(err, "attribute") {
		return err
	}
	c.attributes[name] = value
	return nil
}

func (c *_attributes) SetAttribute(name string, value interface{}) error {
	err := c.setAttribute(name, value)
	if err == nil {
		if *c.updater != nil {
			(*c.updater).Update()
		}
	}
	return err
}

func (c *_attributes) getOrCreateAttribute(name string, creator AttributeFactory) interface{} {
	c.Lock()
	defer c.Unlock()
	if v := c.attributes[name]; v != nil {
		return v
	}
	if c.parent != nil {
		if v := c.parent.GetAttribute(name); v != nil {
			return v
		}
	}
	v := creator(c.ctx)
	c.attributes[name] = v
	return v
}

func (c *_attributes) GetOrCreateAttribute(name string, creator AttributeFactory) interface{} {
	r := c.getOrCreateAttribute(name, creator)
	if *c.updater != nil {
		(*c.updater).Update()
	}
	return r
}
