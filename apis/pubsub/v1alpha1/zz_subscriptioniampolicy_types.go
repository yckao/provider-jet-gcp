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

type SubscriptionIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type SubscriptionIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Subscription *string `json:"subscription" tf:"subscription,omitempty"`
}

// SubscriptionIamPolicySpec defines the desired state of SubscriptionIamPolicy
type SubscriptionIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionIamPolicyParameters `json:"forProvider"`
}

// SubscriptionIamPolicyStatus defines the observed state of SubscriptionIamPolicy.
type SubscriptionIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIamPolicy is the Schema for the SubscriptionIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type SubscriptionIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionIamPolicySpec   `json:"spec"`
	Status            SubscriptionIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIamPolicyList contains a list of SubscriptionIamPolicys
type SubscriptionIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionIamPolicy_Kind             = "SubscriptionIamPolicy"
	SubscriptionIamPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionIamPolicy_Kind}.String()
	SubscriptionIamPolicy_KindAPIVersion   = SubscriptionIamPolicy_Kind + "." + CRDGroupVersion.String()
	SubscriptionIamPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionIamPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionIamPolicy{}, &SubscriptionIamPolicyList{})
}
