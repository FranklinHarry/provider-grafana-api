/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MonitoringProbe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringProbeSpec   `json:"spec,omitempty"`
	Status            MonitoringProbeStatus `json:"status,omitempty"`
}

type MonitoringProbeSpec struct {
	State *MonitoringProbeSpecResource `json:"state,omitempty" tf:"-"`

	Resource MonitoringProbeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type MonitoringProbeSpecResource struct {
	// The probe authentication token. Your probe must use this to authenticate with Grafana Cloud.
	// +optional
	AuthToken *string `json:"-" sensitive:"true" tf:"auth_token"`
	// The ID of the probe.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Custom labels to be included with collected metrics and logs.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Latitude coordinates.
	Latitude *float64 `json:"latitude" tf:"latitude"`
	// Longitude coordinates.
	Longitude *float64 `json:"longitude" tf:"longitude"`
	// Name of the probe.
	Name *string `json:"name" tf:"name"`
	// Public probes are run by Grafana Labs and can be used by all users. You must be an admin to set this to `true`.
	// +optional
	Public *bool `json:"public,omitempty" tf:"public"`
	// Region of the probe.
	Region *string `json:"region" tf:"region"`
	// The tenant ID of the probe.
	// +optional
	TenantID *int64 `json:"tenantID,omitempty" tf:"tenant_id"`
}

type MonitoringProbeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitoringProbeList is a list of MonitoringProbes
type MonitoringProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitoringProbe CRD objects
	Items []MonitoringProbe `json:"items,omitempty"`
}
