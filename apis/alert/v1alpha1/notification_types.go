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

type Notification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationSpec   `json:"spec,omitempty"`
	Status            NotificationStatus `json:"status,omitempty"`
}

type NotificationSpec struct {
	State *NotificationSpecResource `json:"state,omitempty" tf:"-"`

	Resource NotificationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type NotificationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to disable sending resolve messages.
	// +optional
	DisableResolveMessage *bool `json:"disableResolveMessage,omitempty" tf:"disable_resolve_message"`
	// Frequency of alert reminders. Frequency must be set if reminders are enabled.
	// +optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency"`
	// Is this the default channel for all your alerts.
	// +optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default"`
	// The name of the alert notification channel.
	Name *string `json:"name" tf:"name"`
	// Additional secure settings, for full reference lookup [Grafana Supported Settings documentation](https://grafana.com/docs/grafana/latest/administration/provisioning/#supported-settings).
	// +optional
	SecureSettings map[string]string `json:"-" sensitive:"true" tf:"secure_settings"`
	// Whether to send reminders for triggered alerts.
	// +optional
	SendReminder *bool `json:"sendReminder,omitempty" tf:"send_reminder"`
	// Additional settings, for full reference see [Grafana HTTP API documentation](https://grafana.com/docs/grafana/latest/http_api/alerting_notification_channels/).
	// +optional
	Settings map[string]string `json:"settings,omitempty" tf:"settings"`
	// The type of the alert notification channel.
	Type *string `json:"type" tf:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
}

type NotificationStatus struct {
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

// NotificationList is a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Notification CRD objects
	Items []Notification `json:"items,omitempty"`
}
