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

type DefaultSupportedIdpConfigObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DefaultSupportedIdpConfigParameters struct {

	// OAuth client ID
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// OAuth client secret
	// +kubebuilder:validation:Required
	ClientSecret *string `json:"clientSecret" tf:"client_secret,omitempty"`

	// If this IDP allows the user to sign in
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// ID of the IDP. Possible values include:
	//
	// * 'apple.com'
	//
	// * 'facebook.com'
	//
	// * 'gc.apple.com'
	//
	// * 'github.com'
	//
	// * 'google.com'
	//
	// * 'linkedin.com'
	//
	// * 'microsoft.com'
	//
	// * 'playgames.google.com'
	//
	// * 'twitter.com'
	//
	// * 'yahoo.com'
	// +kubebuilder:validation:Required
	IdpID *string `json:"idpId" tf:"idp_id,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DefaultSupportedIdpConfigSpec defines the desired state of DefaultSupportedIdpConfig
type DefaultSupportedIdpConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSupportedIdpConfigParameters `json:"forProvider"`
}

// DefaultSupportedIdpConfigStatus defines the observed state of DefaultSupportedIdpConfig.
type DefaultSupportedIdpConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSupportedIdpConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSupportedIdpConfig is the Schema for the DefaultSupportedIdpConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DefaultSupportedIdpConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSupportedIdpConfigSpec   `json:"spec"`
	Status            DefaultSupportedIdpConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSupportedIdpConfigList contains a list of DefaultSupportedIdpConfigs
type DefaultSupportedIdpConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSupportedIdpConfig `json:"items"`
}

// Repository type metadata.
var (
	DefaultSupportedIdpConfig_Kind             = "DefaultSupportedIdpConfig"
	DefaultSupportedIdpConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSupportedIdpConfig_Kind}.String()
	DefaultSupportedIdpConfig_KindAPIVersion   = DefaultSupportedIdpConfig_Kind + "." + CRDGroupVersion.String()
	DefaultSupportedIdpConfig_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSupportedIdpConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSupportedIdpConfig{}, &DefaultSupportedIdpConfigList{})
}
