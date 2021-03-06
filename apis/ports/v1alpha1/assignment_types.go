/*
Copyright 2020 The Crossplane Authors.

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
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AssignmentSpec defines the desired state of Assignment
type AssignmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AssignmentParameters `json:"forProvider"`
}

// AssignmentStatus defines the observed state of Assignment
type AssignmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// Assignment is a managed resource that represents a Packet Assignment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,packet}
type Assignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssignmentSpec   `json:"spec,omitempty"`
	Status AssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssignmentList contains a list of Assignments
type AssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Assignment `json:"items"`
}

// AssignmentParameters define the desired state of a Packet Virtual Network.
// https://www.packet.com/developers/api/vlans/#create-an-virtual-network
//
// Reference values are used for optional parameters to determine if
// LateInitialization should update the parameter after creation.
type AssignmentParameters struct {
	// +immutable
	DeviceID string `json:"deviceId,omitempty"`

	// +optional
	// +immutable
	DeviceIDRef *runtimev1alpha1.Reference `json:"deviceIdRef,omitempty"`

	// +optional
	DeviceIDSelector *runtimev1alpha1.Selector `json:"deviceIdSelector,omitempty"`

	// +immutable
	Name string `json:"name"`

	// +immutable
	VirtualNetworkID string `json:"virtualNetworkId,omitempty"`

	// +optional
	// +immutable
	VirtualNetworkIDRef *runtimev1alpha1.Reference `json:"virtualNetworkIdRef,omitempty"`

	// +optional
	VirtualNetworkIDSelector *runtimev1alpha1.Selector `json:"virtualNetworkIdSelector,omitempty"`
}
