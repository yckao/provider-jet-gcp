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
	common "github.com/crossplane-contrib/provider-jet-gcp/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Address.
func (mg *Address) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subnetwork),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetworkRef,
		Selector:     mg.Spec.ForProvider.SubnetworkSelector,
		To: reference.To{
			List:    &SubnetworkList{},
			Managed: &Subnetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnetwork")
	}
	mg.Spec.ForProvider.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Firewall.
func (mg *Firewall) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      common.SelfLinkExtractor(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface[i3].Network),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterface[i3].NetworkRef,
			Selector:     mg.Spec.ForProvider.NetworkInterface[i3].NetworkSelector,
			To: reference.To{
				List:    &NetworkList{},
				Managed: &Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface[i3].Network")
		}
		mg.Spec.ForProvider.NetworkInterface[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface[i3].Subnetwork),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterface[i3].SubnetworkRef,
			Selector:     mg.Spec.ForProvider.NetworkInterface[i3].SubnetworkSelector,
			To: reference.To{
				List:    &SubnetworkList{},
				Managed: &Subnetwork{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface[i3].Subnetwork")
		}
		mg.Spec.ForProvider.NetworkInterface[i3].Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterface[i3].SubnetworkRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Router.
func (mg *Router) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      common.SelfLinkExtractor(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouterNAT.
func (mg *RouterNAT) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Router),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouterRef,
		Selector:     mg.Spec.ForProvider.RouterSelector,
		To: reference.To{
			List:    &RouterList{},
			Managed: &Router{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Router")
	}
	mg.Spec.ForProvider.Router = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouterRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Subnetwork); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subnetwork[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Subnetwork[i3].NameRef,
			Selector:     mg.Spec.ForProvider.Subnetwork[i3].NameSelector,
			To: reference.To{
				List:    &SubnetworkList{},
				Managed: &Subnetwork{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Subnetwork[i3].Name")
		}
		mg.Spec.ForProvider.Subnetwork[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Subnetwork[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Subnetwork.
func (mg *Subnetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	return nil
}
