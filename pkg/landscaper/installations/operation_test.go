// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package installations_test

import (
	"context"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	lsv1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	lsv1alpha1helper "github.com/gardener/landscaper/apis/core/v1alpha1/helper"
	"github.com/gardener/landscaper/pkg/api"
	"github.com/gardener/landscaper/pkg/landscaper/blueprints"
	"github.com/gardener/landscaper/pkg/landscaper/dataobjects"
	"github.com/gardener/landscaper/pkg/landscaper/installations"
	"github.com/gardener/landscaper/pkg/landscaper/operation"
	"github.com/gardener/landscaper/pkg/utils"
	testutils "github.com/gardener/landscaper/test/utils"
)

var _ = Describe("Operation", func() {

	var (
		kubeClient client.Client
		op         *installations.Operation
	)

	BeforeEach(func() {
		kubeClient = fake.NewClientBuilder().WithScheme(api.LandscaperScheme).Build()
		commonOp := operation.NewOperation(kubeClient, api.LandscaperScheme, record.NewFakeRecorder(1024))
		op = &installations.Operation{
			Inst: &installations.Installation{
				InstallationBase: installations.InstallationBase{Info: &lsv1alpha1.Installation{}},
				Blueprint:        &blueprints.Blueprint{Info: &lsv1alpha1.Blueprint{}},
			},
			Operation: commonOp,
		}
	})

	Context("CreateOrUpdateImports", func() {

		It("should sync a target", func() {
			ctx := context.Background()
			defer ctx.Done()
			target := &lsv1alpha1.Target{}
			target.Annotations = map[string]string{
				"ann": "val1",
			}
			target.Labels = map[string]string{
				"lab": "val2",
			}
			target.Spec.Type = "test-type"
			target.Spec.Configuration = lsv1alpha1.NewAnyJSON([]byte("true"))
			targetObj, err := utils.JSONSerializeToGenericObject(target)
			testutils.ExpectNoError(err)

			op.Inst.Info.Name = "test"
			op.Inst.Info.Namespace = "default"
			op.Inst.Blueprint.Info.Imports = []lsv1alpha1.ImportDefinition{
				{
					FieldValueDefinition: lsv1alpha1.FieldValueDefinition{
						Name:       "my-import",
						TargetType: "test-type",
					},
					Type: lsv1alpha1.ImportTypeTarget,
				},
			}
			op.Inst.Imports = map[string]interface{}{
				"my-import": targetObj,
			}

			testutils.ExpectNoError(op.CreateOrUpdateImports(ctx))

			targetList := &lsv1alpha1.TargetList{}
			testutils.ExpectNoError(kubeClient.List(ctx, targetList))
			Expect(targetList.Items).To(HaveLen(1))
			Expect(targetList.Items[0].Annotations).To(HaveKeyWithValue("ann", "val1"))
			Expect(targetList.Items[0].Labels).To(HaveKeyWithValue("lab", "val2"))
			Expect(targetList.Items[0].Spec.Type).To(Equal(lsv1alpha1.TargetType("test-type")))
			Expect(targetList.Items[0].Spec.Configuration.RawMessage).To(Equal(json.RawMessage("true")))
		})

	})

	Context("CreateOrUpdateExports", func() {

		It("should detect if an to-be-exported dataobject is already owned by another installation", func() {
			ctx := context.Background()
			defer ctx.Done()

			do := &lsv1alpha1.DataObject{
				ObjectMeta: v1.ObjectMeta{
					Name:      lsv1alpha1helper.GenerateDataObjectName("", "myexport"),
					Namespace: "default",
					OwnerReferences: []v1.OwnerReference{
						{
							Name: "owninginst",
							Kind: "Installation",
							UID:  "owning-installation-uid",
						},
					},
				},
				Data: lsv1alpha1.NewAnyJSON([]byte(`"foo"`)),
			}

			testutils.ExpectNoError(kubeClient.Create(ctx, do))

			op.Inst.Info.Name = "test"
			op.Inst.Info.Namespace = "default"
			op.Inst.Info.UID = "new-installation-uid"

			err := op.CreateOrUpdateExports(ctx, []*dataobjects.DataObject{
				{
					Data: lsv1alpha1.NewAnyJSON([]byte(`"foo"`)),
					Metadata: dataobjects.Metadata{
						Namespace:  "default",
						Context:    "",
						Key:        "myexport",
						SourceType: lsv1alpha1.ExportDataObjectSourceType,
						Source:     "test",
					},
				},
			}, nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("dataobject 'default/myexport' for export 'myexport' conflicts with existing dataobject owned by another installation: object 'default/myexport' is already owned by another object with kind 'Installation' (owninginst)"))
		})

		It("should detect if an to-be-exported target is already owned by another installation", func() {
			ctx := context.Background()
			defer ctx.Done()

			target := &lsv1alpha1.Target{
				ObjectMeta: v1.ObjectMeta{
					Name:      lsv1alpha1helper.GenerateDataObjectName("", "myexport"),
					Namespace: "default",
					OwnerReferences: []v1.OwnerReference{
						{
							Name: "owninginst",
							Kind: "Installation",
							UID:  "owning-installation-uid",
						},
					},
				},
				Spec: lsv1alpha1.TargetSpec{
					Type: "target-type",
				},
			}

			testutils.ExpectNoError(kubeClient.Create(ctx, target))

			op.Inst.Info.Name = "test"
			op.Inst.Info.Namespace = "default"
			op.Inst.Info.UID = "new-installation-uid"

			err := op.CreateOrUpdateExports(ctx, nil, []*dataobjects.Target{
				{
					Metadata: dataobjects.Metadata{
						Namespace:  "default",
						Context:    "",
						Key:        "myexport",
						SourceType: lsv1alpha1.ExportDataObjectSourceType,
						Source:     "test",
					},
				},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("target object 'default/myexport' for export 'myexport' conflicts with existing target owned by another installation: object 'default/myexport' is already owned by another object with kind 'Installation' (owninginst)"))
		})

	})

})
