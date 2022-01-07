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

type TagTemplateIAMMemberConditionObservation struct {
}

type TagTemplateIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TagTemplateIAMMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TagTemplateIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TagTemplateIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	TagTemplate *string `json:"tagTemplate" tf:"tag_template,omitempty"`
}

// TagTemplateIAMMemberSpec defines the desired state of TagTemplateIAMMember
type TagTemplateIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagTemplateIAMMemberParameters `json:"forProvider"`
}

// TagTemplateIAMMemberStatus defines the observed state of TagTemplateIAMMember.
type TagTemplateIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagTemplateIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIAMMember is the Schema for the TagTemplateIAMMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TagTemplateIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagTemplateIAMMemberSpec   `json:"spec"`
	Status            TagTemplateIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIAMMemberList contains a list of TagTemplateIAMMembers
type TagTemplateIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagTemplateIAMMember `json:"items"`
}

// Repository type metadata.
var (
	TagTemplateIAMMember_Kind             = "TagTemplateIAMMember"
	TagTemplateIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagTemplateIAMMember_Kind}.String()
	TagTemplateIAMMember_KindAPIVersion   = TagTemplateIAMMember_Kind + "." + CRDGroupVersion.String()
	TagTemplateIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(TagTemplateIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TagTemplateIAMMember{}, &TagTemplateIAMMemberList{})
}
