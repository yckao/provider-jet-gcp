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

type EdgeCacheOriginObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EdgeCacheOriginParameters struct {

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	//
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	// +kubebuilder:validation:Optional
	FailoverOrigin *string `json:"failoverOrigin,omitempty" tf:"failover_origin,omitempty"`

	// Set of label tags associated with the EdgeCache resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	//
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	//
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	//
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	//
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	// +kubebuilder:validation:Optional
	MaxAttempts *int64 `json:"maxAttempts,omitempty" tf:"max_attempts,omitempty"`

	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	//
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	//
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	// +kubebuilder:validation:Required
	OriginAddress *string `json:"originAddress" tf:"origin_address,omitempty"`

	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	//
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: ["HTTP2", "HTTPS", "HTTP"]
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Specifies one or more retry conditions for the configured origin.
	//
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	//
	// The default retryCondition is "CONNECT_FAILURE".
	//
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	//
	// Valid values are:
	//
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet. Possible values: ["CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND"]
	// +kubebuilder:validation:Optional
	RetryConditions []*string `json:"retryConditions,omitempty" tf:"retry_conditions,omitempty"`

	// The connection and HTTP timeout configuration for this origin.
	// +kubebuilder:validation:Optional
	Timeout []TimeoutParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type TimeoutObservation struct {
}

type TimeoutParameters struct {

	// The maximum duration to wait for the origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
	// +kubebuilder:validation:Optional
	ConnectTimeout *string `json:"connectTimeout,omitempty" tf:"connect_timeout,omitempty"`

	// The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 503 will be returned if the timeout is reached before a response is returned.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
	// +kubebuilder:validation:Optional
	MaxAttemptsTimeout *string `json:"maxAttemptsTimeout,omitempty" tf:"max_attempts_timeout,omitempty"`

	// The maximum duration to wait for data to arrive when reading from the HTTP connection/stream.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 30s.
	// +kubebuilder:validation:Optional
	ResponseTimeout *string `json:"responseTimeout,omitempty" tf:"response_timeout,omitempty"`
}

// EdgeCacheOriginSpec defines the desired state of EdgeCacheOrigin
type EdgeCacheOriginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeCacheOriginParameters `json:"forProvider"`
}

// EdgeCacheOriginStatus defines the observed state of EdgeCacheOrigin.
type EdgeCacheOriginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeCacheOriginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeCacheOrigin is the Schema for the EdgeCacheOrigins API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type EdgeCacheOrigin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeCacheOriginSpec   `json:"spec"`
	Status            EdgeCacheOriginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeCacheOriginList contains a list of EdgeCacheOrigins
type EdgeCacheOriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeCacheOrigin `json:"items"`
}

// Repository type metadata.
var (
	EdgeCacheOrigin_Kind             = "EdgeCacheOrigin"
	EdgeCacheOrigin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeCacheOrigin_Kind}.String()
	EdgeCacheOrigin_KindAPIVersion   = EdgeCacheOrigin_Kind + "." + CRDGroupVersion.String()
	EdgeCacheOrigin_GroupVersionKind = CRDGroupVersion.WithKind(EdgeCacheOrigin_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeCacheOrigin{}, &EdgeCacheOriginList{})
}
