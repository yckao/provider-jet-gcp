/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Database.
func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &DatabaseInstanceList{},
			Managed: &DatabaseInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SSLCert.
func (mg *SSLCert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &DatabaseInstanceList{},
			Managed: &DatabaseInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &DatabaseInstanceList{},
			Managed: &DatabaseInstance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}
