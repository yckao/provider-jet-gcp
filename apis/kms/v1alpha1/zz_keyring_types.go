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

type KeyRingObservation struct {
}

type KeyRingParameters struct {

	// The location for the KeyRing.
	// A full list of valid locations can be found by running 'gcloud kms locations list'.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The resource name for the KeyRing.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// KeyRingSpec defines the desired state of KeyRing
type KeyRingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyRingParameters `json:"forProvider"`
}

// KeyRingStatus defines the observed state of KeyRing.
type KeyRingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyRingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyRing is the Schema for the KeyRings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type KeyRing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyRingSpec   `json:"spec"`
	Status            KeyRingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyRingList contains a list of KeyRings
type KeyRingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyRing `json:"items"`
}

// Repository type metadata.
var (
	KeyRing_Kind             = "KeyRing"
	KeyRing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyRing_Kind}.String()
	KeyRing_KindAPIVersion   = KeyRing_Kind + "." + CRDGroupVersion.String()
	KeyRing_GroupVersionKind = CRDGroupVersion.WithKind(KeyRing_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyRing{}, &KeyRingList{})
}
