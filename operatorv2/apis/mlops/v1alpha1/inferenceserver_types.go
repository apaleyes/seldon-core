/*
Copyright 2021 The Seldon Authors.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InferenceServerSpec defines the desired state of InferenceServer
type InferenceServerSpec struct {
	// Ability to customize the server with a v1.podSpec
	PodSpec PodSpec `json:"podSpec,omitempty" protobuf:"bytes,1,opt,name=podSpec"`
	// 1 of N semantic to specify a particular type of server
	MLServer *MLServerSpec     `json:"mlserver,omitempty" protobuf:"bytes,2,opt,name=mlserver"`
	Triton   *TritonSpec       `json:"triton,omitempty" protobuf:"bytes,3,opt,name=triton"`
	Custom   *CustomServerSpec `json:"custom,omitempty" protobuf:"bytes,4,opt,name=custom"`
}

type ServerCustomizationSpec struct {
	// Allows container overrides for Server
	// +optional
	Container v1.Container `json:"container,omitempty" protobuf:"bytes,1,opt,name=container"`
	// Runtime version of the predictor docker image
	// +optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" protobuf:"bytes,2,opt,name=runtimeVersion"`
	// extra capabilities for this server. Will be matched to InferenceArtifact requirements
	Capabilities []string `json:"capabilities,omitempty" protobuf:"bytes,3,opt,name=capabilities"`
}

// InferenceServerStatus defines the observed state of InferenceServer
type InferenceServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InferenceServer is the Schema for the inferenceservers API
type InferenceServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InferenceServerSpec   `json:"spec,omitempty"`
	Status InferenceServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InferenceServerList contains a list of InferenceServer
type InferenceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InferenceServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InferenceServer{}, &InferenceServerList{})
}