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

type DatasetIamMemberConditionObservation struct {
}

type DatasetIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type DatasetIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type DatasetIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []DatasetIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// DatasetIamMemberSpec defines the desired state of DatasetIamMember
type DatasetIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetIamMemberParameters `json:"forProvider"`
}

// DatasetIamMemberStatus defines the observed state of DatasetIamMember.
type DatasetIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIamMember is the Schema for the DatasetIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DatasetIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetIamMemberSpec   `json:"spec"`
	Status            DatasetIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIamMemberList contains a list of DatasetIamMembers
type DatasetIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasetIamMember `json:"items"`
}

// Repository type metadata.
var (
	DatasetIamMember_Kind             = "DatasetIamMember"
	DatasetIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatasetIamMember_Kind}.String()
	DatasetIamMember_KindAPIVersion   = DatasetIamMember_Kind + "." + CRDGroupVersion.String()
	DatasetIamMember_GroupVersionKind = CRDGroupVersion.WithKind(DatasetIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&DatasetIamMember{}, &DatasetIamMemberList{})
}
