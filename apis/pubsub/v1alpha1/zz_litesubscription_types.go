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

type DeliveryConfigObservation struct {
}

type DeliveryConfigParameters struct {

	// When this subscription should send messages to subscribers relative to messages persistence in storage. Possible values: ["DELIVER_IMMEDIATELY", "DELIVER_AFTER_STORED", "DELIVERY_REQUIREMENT_UNSPECIFIED"]
	// +kubebuilder:validation:Required
	DeliveryRequirement *string `json:"deliveryRequirement" tf:"delivery_requirement,omitempty"`
}

type LiteSubscriptionObservation struct {
}

type LiteSubscriptionParameters struct {

	// The settings for this subscription's message delivery.
	// +kubebuilder:validation:Optional
	DeliveryConfig []DeliveryConfigParameters `json:"deliveryConfig,omitempty" tf:"delivery_config,omitempty"`

	// Name of the subscription.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the pubsub lite topic.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A reference to a Topic resource.
	// +kubebuilder:validation:Required
	Topic *string `json:"topic" tf:"topic,omitempty"`

	// The zone of the pubsub lite topic.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// LiteSubscriptionSpec defines the desired state of LiteSubscription
type LiteSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiteSubscriptionParameters `json:"forProvider"`
}

// LiteSubscriptionStatus defines the observed state of LiteSubscription.
type LiteSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiteSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiteSubscription is the Schema for the LiteSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type LiteSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiteSubscriptionSpec   `json:"spec"`
	Status            LiteSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiteSubscriptionList contains a list of LiteSubscriptions
type LiteSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiteSubscription `json:"items"`
}

// Repository type metadata.
var (
	LiteSubscription_Kind             = "LiteSubscription"
	LiteSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiteSubscription_Kind}.String()
	LiteSubscription_KindAPIVersion   = LiteSubscription_Kind + "." + CRDGroupVersion.String()
	LiteSubscription_GroupVersionKind = CRDGroupVersion.WithKind(LiteSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&LiteSubscription{}, &LiteSubscriptionList{})
}
