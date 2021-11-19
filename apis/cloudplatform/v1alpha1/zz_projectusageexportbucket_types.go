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

type ProjectUsageExportBucketObservation struct {
}

type ProjectUsageExportBucketParameters struct {

	// The bucket to store reports in.
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// A prefix for the reports, for instance, the project name.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The project to set the export bucket on. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ProjectUsageExportBucketSpec defines the desired state of ProjectUsageExportBucket
type ProjectUsageExportBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectUsageExportBucketParameters `json:"forProvider"`
}

// ProjectUsageExportBucketStatus defines the observed state of ProjectUsageExportBucket.
type ProjectUsageExportBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectUsageExportBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectUsageExportBucket is the Schema for the ProjectUsageExportBuckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectUsageExportBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectUsageExportBucketSpec   `json:"spec"`
	Status            ProjectUsageExportBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectUsageExportBucketList contains a list of ProjectUsageExportBuckets
type ProjectUsageExportBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectUsageExportBucket `json:"items"`
}

// Repository type metadata.
var (
	ProjectUsageExportBucket_Kind             = "ProjectUsageExportBucket"
	ProjectUsageExportBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectUsageExportBucket_Kind}.String()
	ProjectUsageExportBucket_KindAPIVersion   = ProjectUsageExportBucket_Kind + "." + CRDGroupVersion.String()
	ProjectUsageExportBucket_GroupVersionKind = CRDGroupVersion.WithKind(ProjectUsageExportBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectUsageExportBucket{}, &ProjectUsageExportBucketList{})
}
