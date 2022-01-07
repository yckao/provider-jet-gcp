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

type BigqueryOptionsObservation struct {
}

type BigqueryOptionsParameters struct {

	// Whether to use BigQuery's partition tables. By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone.
	// +kubebuilder:validation:Required
	UsePartitionedTables *bool `json:"usePartitionedTables" tf:"use_partitioned_tables,omitempty"`
}

type BillingAccountSinkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	WriterIdentity *string `json:"writerIdentity,omitempty" tf:"writer_identity,omitempty"`
}

type BillingAccountSinkParameters struct {

	// Options that affect sinks exporting data to BigQuery.
	// +kubebuilder:validation:Optional
	BigqueryOptions []BigqueryOptionsParameters `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// The billing account exported to the sink.
	// +kubebuilder:validation:Required
	BillingAccount *string `json:"billingAccount" tf:"billing_account,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The writer associated with the sink must have access to write to the above resource.
	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
	// +kubebuilder:validation:Optional
	Exclusions []ExclusionsParameters `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The name of the logging sink.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ExclusionsObservation struct {
}

type ExclusionsParameters struct {

	// A description of this exclusion.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// BillingAccountSinkSpec defines the desired state of BillingAccountSink
type BillingAccountSinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BillingAccountSinkParameters `json:"forProvider"`
}

// BillingAccountSinkStatus defines the observed state of BillingAccountSink.
type BillingAccountSinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BillingAccountSinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BillingAccountSink is the Schema for the BillingAccountSinks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type BillingAccountSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountSinkSpec   `json:"spec"`
	Status            BillingAccountSinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BillingAccountSinkList contains a list of BillingAccountSinks
type BillingAccountSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BillingAccountSink `json:"items"`
}

// Repository type metadata.
var (
	BillingAccountSink_Kind             = "BillingAccountSink"
	BillingAccountSink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BillingAccountSink_Kind}.String()
	BillingAccountSink_KindAPIVersion   = BillingAccountSink_Kind + "." + CRDGroupVersion.String()
	BillingAccountSink_GroupVersionKind = CRDGroupVersion.WithKind(BillingAccountSink_Kind)
)

func init() {
	SchemeBuilder.Register(&BillingAccountSink{}, &BillingAccountSinkList{})
}
