/*
Copyright 2019 X Code.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceName is the type of resources
type ResourceName string

const (
	ResourceCPU      ResourceName = "cpu"
	ResourceMemory   ResourceName = "memory"
	ResourceGPU      ResourceName = "nvidia.com/gpu"
	ResourceStorage  ResourceName = "storage"
	ResourceTime     ResourceName = "time"
	ResourcePriority ResourceName = "priority"
)

// ResourceCostList is a set of (resource name, cost) pairs.
type ResourceCostList map[ResourceName]float64

// CostSpec defines the desired state of Cost
type CostSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Resources ResourceCostList `json:"resources,omitempty"`
}

// CostStatus defines the observed state of Cost
type CostStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cost is the Schema for the costs API
// +k8s:openapi-gen=true
type Cost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CostSpec   `json:"spec,omitempty"`
	Status CostStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CostList contains a list of Cost
type CostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cost `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cost{}, &CostList{})
}
