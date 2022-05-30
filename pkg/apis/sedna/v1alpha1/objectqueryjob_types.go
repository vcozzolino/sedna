/*
Copyright 2021 The KubeEdge Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// ObjectQueryJob describes the Object query jobs.
type ObjectQueryJob struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:",metadata,omitempty"`

	Spec   ObjectQueryJobSpec   `json:"spec"`
	Status ObjectQueryJobStatus `json:"status,omitempty"`
}

// ObjectQueryJobSpec is a description of ObjectQueryJobSpec
type ObjectQueryJobSpec struct {
	SearchRule    SearchRule     `json:"searchRule"`
	SearchTargets []SearchTarget `json:"searchTargets"`

	// The following are optional fields for the model
	Description string `json:"description,omitmepty"`
}

type SearchRule struct {
	WithMask bool `json:"withMask"`
}

type SearchTarget struct {
	TargetImagesDir string   `json:"targetImagesDir"`
	SearchInCameras []string `json:"searchInCameras"`
}

// ObjectQueryJobStatus is a description of ObjectQueryJobStatus
type ObjectQueryJobStatus struct {
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// Represents time when the job was completed. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ObjectQueryJobList is a list of ObjectQueryJob
type ObjectQueryJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ObjectQueryJob `json:"items"`
}
