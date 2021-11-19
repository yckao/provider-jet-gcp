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

type EnvgroupObservation struct {
}

type EnvgroupParameters struct {

	// Hostnames of the environment group.
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// The resource ID of the environment group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Apigee Organization associated with the Apigee environment group,
	// in the format 'organizations/{{org_name}}'.
	// +kubebuilder:validation:Required
	OrgID *string `json:"orgId" tf:"org_id,omitempty"`
}

// EnvgroupSpec defines the desired state of Envgroup
type EnvgroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvgroupParameters `json:"forProvider"`
}

// EnvgroupStatus defines the observed state of Envgroup.
type EnvgroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvgroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Envgroup is the Schema for the Envgroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Envgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvgroupSpec   `json:"spec"`
	Status            EnvgroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvgroupList contains a list of Envgroups
type EnvgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Envgroup `json:"items"`
}

// Repository type metadata.
var (
	Envgroup_Kind             = "Envgroup"
	Envgroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Envgroup_Kind}.String()
	Envgroup_KindAPIVersion   = Envgroup_Kind + "." + CRDGroupVersion.String()
	Envgroup_GroupVersionKind = CRDGroupVersion.WithKind(Envgroup_Kind)
)

func init() {
	SchemeBuilder.Register(&Envgroup{}, &EnvgroupList{})
}
