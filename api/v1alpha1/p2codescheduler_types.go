/*
Copyright 2024.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// P2codeSchedulerSpec defines the desired state of P2codeScheduler
type P2codeSchedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of P2codeScheduler. Edit p2codescheduler_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// P2codeSchedulerStatus defines the observed state of P2codeScheduler
type P2codeSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// P2codeScheduler is the Schema for the p2codeschedulers API
type P2codeScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   P2codeSchedulerSpec   `json:"spec,omitempty"`
	Status P2codeSchedulerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// P2codeSchedulerList contains a list of P2codeScheduler
type P2codeSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []P2codeScheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&P2codeScheduler{}, &P2codeSchedulerList{})
}
