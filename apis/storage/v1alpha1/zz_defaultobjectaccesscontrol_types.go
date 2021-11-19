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

type DefaultObjectAccessControlObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	Generation *int64 `json:"generation,omitempty" tf:"generation,omitempty"`

	ProjectTeam []ProjectTeamObservation `json:"projectTeam,omitempty" tf:"project_team,omitempty"`
}

type DefaultObjectAccessControlParameters struct {

	// The name of the bucket.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// The entity holding the permission, in one of the following forms:
	// * user-{{userId}}
	// * user-{{email}} (such as "user-liz@example.com")
	// * group-{{groupId}}
	// * group-{{email}} (such as "group-example@googlegroups.com")
	// * domain-{{domain}} (such as "domain-example.com")
	// * project-team-{{projectId}}
	// * allUsers
	// * allAuthenticatedUsers
	// +kubebuilder:validation:Required
	Entity *string `json:"entity" tf:"entity,omitempty"`

	// The name of the object, if applied to an object.
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

type ProjectTeamObservation struct {
}

type ProjectTeamParameters struct {

	// The project team associated with the entity
	// +kubebuilder:validation:Optional
	ProjectNumber *string `json:"projectNumber,omitempty" tf:"project_number,omitempty"`

	// The team. Possible values: ["editors", "owners", "viewers"]
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

// DefaultObjectAccessControlSpec defines the desired state of DefaultObjectAccessControl
type DefaultObjectAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultObjectAccessControlParameters `json:"forProvider"`
}

// DefaultObjectAccessControlStatus defines the observed state of DefaultObjectAccessControl.
type DefaultObjectAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultObjectAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultObjectAccessControl is the Schema for the DefaultObjectAccessControls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultObjectAccessControlSpec   `json:"spec"`
	Status            DefaultObjectAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultObjectAccessControlList contains a list of DefaultObjectAccessControls
type DefaultObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultObjectAccessControl `json:"items"`
}

// Repository type metadata.
var (
	DefaultObjectAccessControl_Kind             = "DefaultObjectAccessControl"
	DefaultObjectAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultObjectAccessControl_Kind}.String()
	DefaultObjectAccessControl_KindAPIVersion   = DefaultObjectAccessControl_Kind + "." + CRDGroupVersion.String()
	DefaultObjectAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(DefaultObjectAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultObjectAccessControl{}, &DefaultObjectAccessControlList{})
}
