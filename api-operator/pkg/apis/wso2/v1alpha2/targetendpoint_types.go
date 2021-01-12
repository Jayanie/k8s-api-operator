// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TargetEndpointSpec defines the desired state of TargetEndpoint
// +k8s:openapi-gen=true
type TargetEndpointSpec struct {
	// Protocol of the application. Supports "http" and "https".
	ApplicationProtocol string `json:"applicationProtocol"`
	// List of optional ports of the target endpoint.
	// First port should be the port of the target endpoint which is referred in swagger definition.
	Ports []Port `json:"ports"`
	// Deployment details.
	Deploy Deploy `json:"deploy"`
	// Mode of the Target Endpoint. Supports "privateJet", "sidecar", "serverless".
	// Default value "privateJet"
	// +optional
	Mode Mode `json:"mode,omitempty"`
}

// TargetEndpointStatus defines the observed state of TargetEndpoint
// +k8s:openapi-gen=true
type TargetEndpointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

type EndpointSecurity struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// Port represents ports of the Target Endpoint
type Port struct {
	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names.
	Name string `json:"name"`
	// The port that will be exposed by this service.
	Port int32 `json:"port"`
	// Port that is targeted to expose.
	TargetPort int32 `json:"targetPort"`
}

//Replica count for create HPA for targetEndPoint
type Deploy struct {
	Name        string `json:"name"`
	DockerImage string `json:"dockerImage"`
	MinReplicas int32  `json:"minReplicas,omitempty"`
	MaxReplicas int32  `json:"maxReplicas,omitempty"`
	ReqCpu      string `json:"requestCPU,omitempty"`
	ReqMemory   string `json:"reqMemory,omitempty"`
	LimitCpu    string `json:"cpuLimit,omitempty"`
	MemoryLimit string `json:"memoryLimit,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TargetEndpoint is the Schema for the targetendpoints API
// +k8s:openapi-gen=true
type TargetEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TargetEndpointSpec   `json:"spec,omitempty"`
	Status TargetEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TargetEndpointList contains a list of TargetEndpoint
type TargetEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TargetEndpoint{}, &TargetEndpointList{})
}
