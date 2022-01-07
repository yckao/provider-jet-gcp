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

type AdditionalExtensionsObservation struct {
}

type AdditionalExtensionsParameters struct {

	// Indicates whether or not this extension is critical (i.e., if the client does not know how to
	// handle this extension, the client should consider this to be an error).
	// +kubebuilder:validation:Required
	Critical *bool `json:"critical" tf:"critical,omitempty"`

	// Describes values that are relevant in a CA certificate.
	// +kubebuilder:validation:Required
	ObjectID []ObjectIDParameters `json:"objectId" tf:"object_id,omitempty"`

	// The value of this X.509 extension. A base64-encoded string.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type AllowedIssuanceModesObservation struct {
}

type AllowedIssuanceModesParameters struct {

	// When true, allows callers to create Certificates by specifying a CertificateConfig.
	// +kubebuilder:validation:Required
	AllowConfigBasedIssuance *bool `json:"allowConfigBasedIssuance" tf:"allow_config_based_issuance,omitempty"`

	// When true, allows callers to create Certificates by specifying a CSR.
	// +kubebuilder:validation:Required
	AllowCsrBasedIssuance *bool `json:"allowCsrBasedIssuance" tf:"allow_csr_based_issuance,omitempty"`
}

type AllowedKeyTypesObservation struct {
}

type AllowedKeyTypesParameters struct {

	// Represents an allowed Elliptic Curve key type.
	// +kubebuilder:validation:Optional
	EllipticCurve []EllipticCurveParameters `json:"ellipticCurve,omitempty" tf:"elliptic_curve,omitempty"`

	// Describes an RSA key that may be used in a Certificate issued from a CaPool.
	// +kubebuilder:validation:Optional
	Rsa []RsaParameters `json:"rsa,omitempty" tf:"rsa,omitempty"`
}

type BaseKeyUsageObservation struct {
}

type BaseKeyUsageParameters struct {

	// The key may be used to sign certificates.
	// +kubebuilder:validation:Optional
	CertSign *bool `json:"certSign,omitempty" tf:"cert_sign,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation".
	// +kubebuilder:validation:Optional
	ContentCommitment *bool `json:"contentCommitment,omitempty" tf:"content_commitment,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kubebuilder:validation:Optional
	CrlSign *bool `json:"crlSign,omitempty" tf:"crl_sign,omitempty"`

	// The key may be used to encipher data.
	// +kubebuilder:validation:Optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty" tf:"data_encipherment,omitempty"`

	// The key may be used to decipher only.
	// +kubebuilder:validation:Optional
	DecipherOnly *bool `json:"decipherOnly,omitempty" tf:"decipher_only,omitempty"`

	// The key may be used for digital signatures.
	// +kubebuilder:validation:Optional
	DigitalSignature *bool `json:"digitalSignature,omitempty" tf:"digital_signature,omitempty"`

	// The key may be used to encipher only.
	// +kubebuilder:validation:Optional
	EncipherOnly *bool `json:"encipherOnly,omitempty" tf:"encipher_only,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kubebuilder:validation:Optional
	KeyAgreement *bool `json:"keyAgreement,omitempty" tf:"key_agreement,omitempty"`

	// The key may be used to encipher other keys.
	// +kubebuilder:validation:Optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty" tf:"key_encipherment,omitempty"`
}

type BaselineValuesObservation struct {
}

type BaselineValuesParameters struct {

	// Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.
	// +kubebuilder:validation:Optional
	AdditionalExtensions []AdditionalExtensionsParameters `json:"additionalExtensions,omitempty" tf:"additional_extensions,omitempty"`

	// Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the
	// "Authority Information Access" extension in the certificate.
	// +kubebuilder:validation:Optional
	AiaOcspServers []*string `json:"aiaOcspServers,omitempty" tf:"aia_ocsp_servers,omitempty"`

	// Describes values that are relevant in a CA certificate.
	// +kubebuilder:validation:Required
	CAOptions []CAOptionsParameters `json:"caOptions" tf:"ca_options,omitempty"`

	// Indicates the intended use for keys that correspond to a certificate.
	// +kubebuilder:validation:Required
	KeyUsage []KeyUsageParameters `json:"keyUsage" tf:"key_usage,omitempty"`

	// Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kubebuilder:validation:Optional
	PolicyIds []PolicyIdsParameters `json:"policyIds,omitempty" tf:"policy_ids,omitempty"`
}

type CAOptionsObservation struct {
}

type CAOptionsParameters struct {

	// Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing,
	// the extension will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	IsCA *bool `json:"isCa,omitempty" tf:"is_ca,omitempty"`

	// Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of
	// subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this
	// value is missing, the max path length will be omitted from the CA certificate.
	// +kubebuilder:validation:Optional
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength,omitempty" tf:"max_issuer_path_length,omitempty"`
}

type CAPoolObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CAPoolParameters struct {

	// The IssuancePolicy to control how Certificates will be issued from this CaPool.
	// +kubebuilder:validation:Optional
	IssuancePolicy []IssuancePolicyParameters `json:"issuancePolicy,omitempty" tf:"issuance_policy,omitempty"`

	// Labels with user-defined metadata.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location of the CaPool. A full list of valid locations can be found by
	// running 'gcloud privateca locations list'.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name for this CaPool.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	// +kubebuilder:validation:Optional
	PublishingOptions []PublishingOptionsParameters `json:"publishingOptions,omitempty" tf:"publishing_options,omitempty"`

	// The Tier of this CaPool. Possible values: ["ENTERPRISE", "DEVOPS"]
	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

type CelExpressionObservation struct {
}

type CelExpressionParameters struct {

	// Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type EllipticCurveObservation struct {
}

type EllipticCurveParameters struct {

	// The algorithm used. Possible values: ["ECDSA_P256", "ECDSA_P384", "EDDSA_25519"]
	// +kubebuilder:validation:Required
	SignatureAlgorithm *string `json:"signatureAlgorithm" tf:"signature_algorithm,omitempty"`
}

type ExtendedKeyUsageObservation struct {
}

type ExtendedKeyUsageParameters struct {

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ClientAuth *bool `json:"clientAuth,omitempty" tf:"client_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	// +kubebuilder:validation:Optional
	CodeSigning *bool `json:"codeSigning,omitempty" tf:"code_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	// +kubebuilder:validation:Optional
	EmailProtection *bool `json:"emailProtection,omitempty" tf:"email_protection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	// +kubebuilder:validation:Optional
	OcspSigning *bool `json:"ocspSigning,omitempty" tf:"ocsp_signing,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	// +kubebuilder:validation:Optional
	ServerAuth *bool `json:"serverAuth,omitempty" tf:"server_auth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	// +kubebuilder:validation:Optional
	TimeStamping *bool `json:"timeStamping,omitempty" tf:"time_stamping,omitempty"`
}

type IdentityConstraintsObservation struct {
}

type IdentityConstraintsParameters struct {

	// If this is set, the SubjectAltNames extension may be copied from a certificate request into the signed certificate.
	// Otherwise, the requested SubjectAltNames will be discarded.
	// +kubebuilder:validation:Required
	AllowSubjectAltNamesPassthrough *bool `json:"allowSubjectAltNamesPassthrough" tf:"allow_subject_alt_names_passthrough,omitempty"`

	// If this is set, the Subject field may be copied from a certificate request into the signed certificate.
	// Otherwise, the requested Subject will be discarded.
	// +kubebuilder:validation:Required
	AllowSubjectPassthrough *bool `json:"allowSubjectPassthrough" tf:"allow_subject_passthrough,omitempty"`

	// A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a
	// certificate is signed. To see the full allowed syntax and some examples,
	// see https://cloud.google.com/certificate-authority-service/docs/cel-guide
	// +kubebuilder:validation:Optional
	CelExpression []CelExpressionParameters `json:"celExpression,omitempty" tf:"cel_expression,omitempty"`
}

type IssuancePolicyObservation struct {
}

type IssuancePolicyParameters struct {

	// IssuanceModes specifies the allowed ways in which Certificates may be requested from this CaPool.
	// +kubebuilder:validation:Optional
	AllowedIssuanceModes []AllowedIssuanceModesParameters `json:"allowedIssuanceModes,omitempty" tf:"allowed_issuance_modes,omitempty"`

	// If any AllowedKeyType is specified, then the certificate request's public key must match one of the key types listed here.
	// Otherwise, any key may be used.
	// +kubebuilder:validation:Optional
	AllowedKeyTypes []AllowedKeyTypesParameters `json:"allowedKeyTypes,omitempty" tf:"allowed_key_types,omitempty"`

	// A set of X.509 values that will be applied to all certificates issued through this CaPool. If a certificate request
	// includes conflicting values for the same properties, they will be overwritten by the values defined here. If a certificate
	// request uses a CertificateTemplate that defines conflicting predefinedValues for the same properties, the certificate
	// issuance request will fail.
	// +kubebuilder:validation:Optional
	BaselineValues []BaselineValuesParameters `json:"baselineValues,omitempty" tf:"baseline_values,omitempty"`

	// Describes constraints on identities that may appear in Certificates issued through this CaPool.
	// If this is omitted, then this CaPool will not add restrictions on a certificate's identity.
	// +kubebuilder:validation:Optional
	IdentityConstraints []IdentityConstraintsParameters `json:"identityConstraints,omitempty" tf:"identity_constraints,omitempty"`

	// The maximum lifetime allowed for issued Certificates. Note that if the issuing CertificateAuthority
	// expires before a Certificate's requested maximumLifetime, the effective lifetime will be explicitly truncated to match it.
	// +kubebuilder:validation:Optional
	MaximumLifetime *string `json:"maximumLifetime,omitempty" tf:"maximum_lifetime,omitempty"`
}

type KeyUsageObservation struct {
}

type KeyUsageParameters struct {

	// Describes high-level ways in which a key may be used.
	// +kubebuilder:validation:Required
	BaseKeyUsage []BaseKeyUsageParameters `json:"baseKeyUsage" tf:"base_key_usage,omitempty"`

	// Describes high-level ways in which a key may be used.
	// +kubebuilder:validation:Required
	ExtendedKeyUsage []ExtendedKeyUsageParameters `json:"extendedKeyUsage" tf:"extended_key_usage,omitempty"`

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Optional
	UnknownExtendedKeyUsages []UnknownExtendedKeyUsagesParameters `json:"unknownExtendedKeyUsages,omitempty" tf:"unknown_extended_key_usages,omitempty"`
}

type ObjectIDObservation struct {
}

type ObjectIDParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type PolicyIdsObservation struct {
}

type PolicyIdsParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

type PublishingOptionsObservation struct {
}

type PublishingOptionsParameters struct {

	// When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access"
	// X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding
	// X.509 extension will not be written in issued certificates.
	// +kubebuilder:validation:Required
	PublishCACert *bool `json:"publishCaCert" tf:"publish_ca_cert,omitempty"`

	// When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension
	// in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not
	// be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are
	// also rebuilt shortly after a certificate is revoked.
	// +kubebuilder:validation:Required
	PublishCrl *bool `json:"publishCrl" tf:"publish_crl,omitempty"`
}

type RsaObservation struct {
}

type RsaParameters struct {

	// The maximum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
	// service will not enforce an explicit upper bound on RSA modulus sizes.
	// +kubebuilder:validation:Optional
	MaxModulusSize *string `json:"maxModulusSize,omitempty" tf:"max_modulus_size,omitempty"`

	// The minimum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the
	// service-level min RSA modulus size will continue to apply.
	// +kubebuilder:validation:Optional
	MinModulusSize *string `json:"minModulusSize,omitempty" tf:"min_modulus_size,omitempty"`
}

type UnknownExtendedKeyUsagesObservation struct {
}

type UnknownExtendedKeyUsagesParameters struct {

	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	// +kubebuilder:validation:Required
	ObjectIDPath []*int64 `json:"objectIdPath" tf:"object_id_path,omitempty"`
}

// CAPoolSpec defines the desired state of CAPool
type CAPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CAPoolParameters `json:"forProvider"`
}

// CAPoolStatus defines the observed state of CAPool.
type CAPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CAPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CAPool is the Schema for the CAPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type CAPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CAPoolSpec   `json:"spec"`
	Status            CAPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CAPoolList contains a list of CAPools
type CAPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CAPool `json:"items"`
}

// Repository type metadata.
var (
	CAPool_Kind             = "CAPool"
	CAPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CAPool_Kind}.String()
	CAPool_KindAPIVersion   = CAPool_Kind + "." + CRDGroupVersion.String()
	CAPool_GroupVersionKind = CRDGroupVersion.WithKind(CAPool_Kind)
)

func init() {
	SchemeBuilder.Register(&CAPool{}, &CAPoolList{})
}
