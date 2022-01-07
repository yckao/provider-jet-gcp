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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TenantObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TenantParameters struct {

	// Whether to allow email/password user authentication.
	// +kubebuilder:validation:Optional
	AllowPasswordSignup *bool `json:"allowPasswordSignup,omitempty" tf:"allow_password_signup,omitempty"`

	// Whether authentication is disabled for the tenant. If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	// +kubebuilder:validation:Optional
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// Human friendly display name of the tenant.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Whether to enable email link user authentication.
	// +kubebuilder:validation:Optional
	EnableEmailLinkSignin *bool `json:"enableEmailLinkSignin,omitempty" tf:"enable_email_link_signin,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// TenantSpec defines the desired state of Tenant
type TenantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantParameters `json:"forProvider"`
}

// TenantStatus defines the observed state of Tenant.
type TenantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tenant is the Schema for the Tenants API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantSpec   `json:"spec"`
	Status            TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantList contains a list of Tenants
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

// Repository type metadata.
var (
	Tenant_Kind             = "Tenant"
	Tenant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tenant_Kind}.String()
	Tenant_KindAPIVersion   = Tenant_Kind + "." + CRDGroupVersion.String()
	Tenant_GroupVersionKind = CRDGroupVersion.WithKind(Tenant_Kind)
)

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
