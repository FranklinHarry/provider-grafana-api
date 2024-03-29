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

type Dashboard struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardSpec   `json:"spec,omitempty"`
	Status            DashboardStatus `json:"status,omitempty"`
}

type DashboardSpec struct {
	State *DashboardSpecResource `json:"state,omitempty" tf:"-"`

	Resource DashboardSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DashboardSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The complete dashboard model JSON.
	ConfigJSON *string `json:"configJSON" tf:"config_json"`
	// The numeric ID of the dashboard computed by Grafana.
	// +optional
	DashboardID *int64 `json:"dashboardID,omitempty" tf:"dashboard_id"`
	// The id of the folder to save the dashboard in.
	// +optional
	Folder *int64 `json:"folder,omitempty" tf:"folder"`
	// Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	// +optional
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite"`
	// URL friendly version of the dashboard title. This field is deprecated, please use `uid` instead.
	// +optional
	// Deprecated
	Slug *string `json:"slug,omitempty" tf:"slug"`
	// The unique identifier of a dashboard. This is used to construct its URL. It’s automatically generated if not provided when creating a dashboard. The uid allows having consistent URLs for accessing dashboards and when syncing dashboards between multiple Grafana installs.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// Whenever you save a version of your dashboard, a copy of that version is saved so that previous versions of your dashboard are not lost.
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type DashboardStatus struct {
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

// DashboardList is a list of Dashboards
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Dashboard CRD objects
	Items []Dashboard `json:"items,omitempty"`
}
