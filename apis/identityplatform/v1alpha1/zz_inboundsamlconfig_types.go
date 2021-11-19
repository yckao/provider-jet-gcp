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

type IdpCertificatesObservation struct {
}

type IdpCertificatesParameters struct {

	// The IdP's x509 certificate.
	// +kubebuilder:validation:Optional
	X509Certificate *string `json:"x509Certificate,omitempty" tf:"x509_certificate,omitempty"`
}

type IdpConfigObservation struct {
}

type IdpConfigParameters struct {

	// The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// +kubebuilder:validation:Required
	IdpCertificates []IdpCertificatesParameters `json:"idpCertificates" tf:"idp_certificates,omitempty"`

	// Unique identifier for all SAML entities
	// +kubebuilder:validation:Required
	IdpEntityID *string `json:"idpEntityId" tf:"idp_entity_id,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	// +kubebuilder:validation:Optional
	SignRequest *bool `json:"signRequest,omitempty" tf:"sign_request,omitempty"`

	// URL to send Authentication request to.
	// +kubebuilder:validation:Required
	SsoURL *string `json:"ssoUrl" tf:"sso_url,omitempty"`
}

type InboundSamlConfigObservation struct {
}

type InboundSamlConfigParameters struct {

	// Human friendly display name.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// SAML IdP configuration when the project acts as the relying party
	// +kubebuilder:validation:Required
	IdpConfig []IdpConfigParameters `json:"idpConfig" tf:"idp_config,omitempty"`

	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// +kubebuilder:validation:Required
	SpConfig []SpConfigParameters `json:"spConfig" tf:"sp_config,omitempty"`
}

type SpCertificatesObservation struct {
	X509Certificate *string `json:"x509Certificate,omitempty" tf:"x509_certificate,omitempty"`
}

type SpCertificatesParameters struct {
}

type SpConfigObservation struct {
	SpCertificates []SpCertificatesObservation `json:"spCertificates,omitempty" tf:"sp_certificates,omitempty"`
}

type SpConfigParameters struct {

	// Callback URI where responses from IDP are handled. Must start with 'https://'.
	// +kubebuilder:validation:Optional
	CallbackURI *string `json:"callbackUri,omitempty" tf:"callback_uri,omitempty"`

	// Unique identifier for all SAML entities.
	// +kubebuilder:validation:Optional
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

// InboundSamlConfigSpec defines the desired state of InboundSamlConfig
type InboundSamlConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InboundSamlConfigParameters `json:"forProvider"`
}

// InboundSamlConfigStatus defines the observed state of InboundSamlConfig.
type InboundSamlConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InboundSamlConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InboundSamlConfig is the Schema for the InboundSamlConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type InboundSamlConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InboundSamlConfigSpec   `json:"spec"`
	Status            InboundSamlConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InboundSamlConfigList contains a list of InboundSamlConfigs
type InboundSamlConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InboundSamlConfig `json:"items"`
}

// Repository type metadata.
var (
	InboundSamlConfig_Kind             = "InboundSamlConfig"
	InboundSamlConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InboundSamlConfig_Kind}.String()
	InboundSamlConfig_KindAPIVersion   = InboundSamlConfig_Kind + "." + CRDGroupVersion.String()
	InboundSamlConfig_GroupVersionKind = CRDGroupVersion.WithKind(InboundSamlConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&InboundSamlConfig{}, &InboundSamlConfigList{})
}
