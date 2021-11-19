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

type ArgumentsObservation struct {
}

type ArgumentsParameters struct {

	// Defaults to FIXED_TYPE. Default value: "FIXED_TYPE" Possible values: ["FIXED_TYPE", "ANY_TYPE"]
	// +kubebuilder:validation:Optional
	ArgumentKind *string `json:"argumentKind,omitempty" tf:"argument_kind,omitempty"`

	// A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
	// ~>**NOTE**: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// Specifies whether the argument is input or output. Can be set for procedures only. Possible values: ["IN", "OUT", "INOUT"]
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of this argument. Can be absent for function return argument.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RoutineObservation struct {
	CreationTime *int64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`
}

type RoutineParameters struct {

	// Input/output argument of a function or a stored procedure.
	// +kubebuilder:validation:Optional
	Arguments []ArgumentsParameters `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// The ID of the dataset containing this routine
	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	// +kubebuilder:validation:Required
	DefinitionBody *string `json:"definitionBody" tf:"definition_body,omitempty"`

	// The description of the routine if defined.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"]
	// +kubebuilder:validation:Optional
	DeterminismLevel *string `json:"determinismLevel,omitempty" tf:"determinism_level,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	// +kubebuilder:validation:Optional
	ImportedLibraries []*string `json:"importedLibraries,omitempty" tf:"imported_libraries,omitempty"`

	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT"]
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +kubebuilder:validation:Optional
	ReturnType *string `json:"returnType,omitempty" tf:"return_type,omitempty"`

	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	// +kubebuilder:validation:Required
	RoutineID *string `json:"routineId" tf:"routine_id,omitempty"`

	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE"]
	// +kubebuilder:validation:Optional
	RoutineType *string `json:"routineType,omitempty" tf:"routine_type,omitempty"`
}

// RoutineSpec defines the desired state of Routine
type RoutineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutineParameters `json:"forProvider"`
}

// RoutineStatus defines the observed state of Routine.
type RoutineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Routine is the Schema for the Routines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Routine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoutineSpec   `json:"spec"`
	Status            RoutineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutineList contains a list of Routines
type RoutineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Routine `json:"items"`
}

// Repository type metadata.
var (
	Routine_Kind             = "Routine"
	Routine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Routine_Kind}.String()
	Routine_KindAPIVersion   = Routine_Kind + "." + CRDGroupVersion.String()
	Routine_GroupVersionKind = CRDGroupVersion.WithKind(Routine_Kind)
)

func init() {
	SchemeBuilder.Register(&Routine{}, &RoutineList{})
}
