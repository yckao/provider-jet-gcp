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

type InstanceIamBindingConditionObservation struct {
}

type InstanceIamBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type InstanceIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type InstanceIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []InstanceIamBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// InstanceIamBindingSpec defines the desired state of InstanceIamBinding
type InstanceIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIamBindingParameters `json:"forProvider"`
}

// InstanceIamBindingStatus defines the observed state of InstanceIamBinding.
type InstanceIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIamBinding is the Schema for the InstanceIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type InstanceIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceIamBindingSpec   `json:"spec"`
	Status            InstanceIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIamBindingList contains a list of InstanceIamBindings
type InstanceIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIamBinding `json:"items"`
}

// Repository type metadata.
var (
	InstanceIamBinding_Kind             = "InstanceIamBinding"
	InstanceIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceIamBinding_Kind}.String()
	InstanceIamBinding_KindAPIVersion   = InstanceIamBinding_Kind + "." + CRDGroupVersion.String()
	InstanceIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(InstanceIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceIamBinding{}, &InstanceIamBindingList{})
}
