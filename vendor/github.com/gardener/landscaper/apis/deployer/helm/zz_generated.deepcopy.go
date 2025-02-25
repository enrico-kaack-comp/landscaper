//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package helm

import (
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"

	config "github.com/gardener/landscaper/apis/config"
	core "github.com/gardener/landscaper/apis/core"
	v1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	continuousreconcile "github.com/gardener/landscaper/apis/deployer/utils/continuousreconcile"
	managedresource "github.com/gardener/landscaper/apis/deployer/utils/managedresource"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveAccess) DeepCopyInto(out *ArchiveAccess) {
	*out = *in
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(RemoteArchiveAccess)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveAccess.
func (in *ArchiveAccess) DeepCopy() *ArchiveAccess {
	if in == nil {
		return nil
	}
	out := new(ArchiveAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1alpha1.LocalSecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	if in.FromResource != nil {
		in, out := &in.FromResource, &out.FromResource
		*out = new(RemoteChartReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Archive != nil {
		in, out := &in.Archive, &out.Archive
		*out = new(ArchiveAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmChartRepo != nil {
		in, out := &in.HelmChartRepo, &out.HelmChartRepo
		*out = new(HelmChartRepo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(config.OCIConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetSelector != nil {
		in, out := &in.TargetSelector, &out.TargetSelector
		*out = make([]v1alpha1.TargetSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Export.DeepCopyInto(&out.Export)
	if in.HPAConfiguration != nil {
		in, out := &in.HPAConfiguration, &out.HPAConfiguration
		*out = new(HPAConfiguration)
		**out = **in
	}
	in.Controller.DeepCopyInto(&out.Controller)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Controller) DeepCopyInto(out *Controller) {
	*out = *in
	in.CommonControllerConfig.DeepCopyInto(&out.CommonControllerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Controller.
func (in *Controller) DeepCopy() *Controller {
	if in == nil {
		return nil
	}
	out := new(Controller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportConfiguration) DeepCopyInto(out *ExportConfiguration) {
	*out = *in
	if in.DefaultTimeout != nil {
		in, out := &in.DefaultTimeout, &out.DefaultTimeout
		*out = new(v1alpha1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportConfiguration.
func (in *ExportConfiguration) DeepCopy() *ExportConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExportConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAConfiguration) DeepCopyInto(out *HPAConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAConfiguration.
func (in *HPAConfiguration) DeepCopy() *HPAConfiguration {
	if in == nil {
		return nil
	}
	out := new(HPAConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepo) DeepCopyInto(out *HelmChartRepo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepo.
func (in *HelmChartRepo) DeepCopy() *HelmChartRepo {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartRepoCredentials) DeepCopyInto(out *HelmChartRepoCredentials) {
	*out = *in
	if in.Auths != nil {
		in, out := &in.Auths, &out.Auths
		*out = make([]Auth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartRepoCredentials.
func (in *HelmChartRepoCredentials) DeepCopy() *HelmChartRepoCredentials {
	if in == nil {
		return nil
	}
	out := new(HelmChartRepoCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmDeploymentConfiguration) DeepCopyInto(out *HelmDeploymentConfiguration) {
	*out = *in
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		*out = make(map[string]core.AnyJSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = make(map[string]core.AnyJSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Uninstall != nil {
		in, out := &in.Uninstall, &out.Uninstall
		*out = make(map[string]core.AnyJSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmDeploymentConfiguration.
func (in *HelmDeploymentConfiguration) DeepCopy() *HelmDeploymentConfiguration {
	if in == nil {
		return nil
	}
	out := new(HelmDeploymentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInstallConfiguration) DeepCopyInto(out *HelmInstallConfiguration) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1alpha1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInstallConfiguration.
func (in *HelmInstallConfiguration) DeepCopy() *HelmInstallConfiguration {
	if in == nil {
		return nil
	}
	out := new(HelmInstallConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmUninstallConfiguration) DeepCopyInto(out *HelmUninstallConfiguration) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1alpha1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmUninstallConfiguration.
func (in *HelmUninstallConfiguration) DeepCopy() *HelmUninstallConfiguration {
	if in == nil {
		return nil
	}
	out := new(HelmUninstallConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfiguration) DeepCopyInto(out *ProviderConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ReadinessChecks.DeepCopyInto(&out.ReadinessChecks)
	in.Chart.DeepCopyInto(&out.Chart)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.ExportsFromManifests != nil {
		in, out := &in.ExportsFromManifests, &out.ExportsFromManifests
		*out = make([]managedresource.Export, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Exports != nil {
		in, out := &in.Exports, &out.Exports
		*out = new(managedresource.Exports)
		(*in).DeepCopyInto(*out)
	}
	if in.ContinuousReconcile != nil {
		in, out := &in.ContinuousReconcile, &out.ContinuousReconcile
		*out = new(continuousreconcile.ContinuousReconcileSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmDeployment != nil {
		in, out := &in.HelmDeployment, &out.HelmDeployment
		*out = new(bool)
		**out = **in
	}
	if in.HelmDeploymentConfig != nil {
		in, out := &in.HelmDeploymentConfig, &out.HelmDeploymentConfig
		*out = new(HelmDeploymentConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionGroups != nil {
		in, out := &in.DeletionGroups, &out.DeletionGroups
		*out = make([]managedresource.DeletionGroupDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeletionGroupsDuringUpdate != nil {
		in, out := &in.DeletionGroupsDuringUpdate, &out.DeletionGroupsDuringUpdate
		*out = make([]managedresource.DeletionGroupDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfiguration.
func (in *ProviderConfiguration) DeepCopy() *ProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderStatus) DeepCopyInto(out *ProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make(managedresource.ManagedResourceStatusList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderStatus.
func (in *ProviderStatus) DeepCopy() *ProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteArchiveAccess) DeepCopyInto(out *RemoteArchiveAccess) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteArchiveAccess.
func (in *RemoteArchiveAccess) DeepCopy() *RemoteArchiveAccess {
	if in == nil {
		return nil
	}
	out := new(RemoteArchiveAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteChartReference) DeepCopyInto(out *RemoteChartReference) {
	*out = *in
	in.ComponentDescriptorDefinition.DeepCopyInto(&out.ComponentDescriptorDefinition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteChartReference.
func (in *RemoteChartReference) DeepCopy() *RemoteChartReference {
	if in == nil {
		return nil
	}
	out := new(RemoteChartReference)
	in.DeepCopyInto(out)
	return out
}
