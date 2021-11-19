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

type OrganizationExclusionObservation struct {
}

type OrganizationExclusionParameters struct {

	// A human-readable description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to false.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// The name of the logging exclusion.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	OrgID *string `json:"orgId" tf:"org_id,omitempty"`
}

// OrganizationExclusionSpec defines the desired state of OrganizationExclusion
type OrganizationExclusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationExclusionParameters `json:"forProvider"`
}

// OrganizationExclusionStatus defines the observed state of OrganizationExclusion.
type OrganizationExclusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationExclusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationExclusion is the Schema for the OrganizationExclusions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type OrganizationExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationExclusionSpec   `json:"spec"`
	Status            OrganizationExclusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationExclusionList contains a list of OrganizationExclusions
type OrganizationExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationExclusion `json:"items"`
}

// Repository type metadata.
var (
	OrganizationExclusion_Kind             = "OrganizationExclusion"
	OrganizationExclusion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationExclusion_Kind}.String()
	OrganizationExclusion_KindAPIVersion   = OrganizationExclusion_Kind + "." + CRDGroupVersion.String()
	OrganizationExclusion_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationExclusion_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationExclusion{}, &OrganizationExclusionList{})
}
