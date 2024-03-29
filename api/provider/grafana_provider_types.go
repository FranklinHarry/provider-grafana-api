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

package provider

type GrafanaSpec struct {
	// API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
	Auth *string `json:"-" sensitive:"true" tf:"auth"`
	// Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the `GRAFANA_CA_CERT` environment variable.
	// +optional
	CaCert *string `json:"caCert,omitempty" tf:"ca_cert"`
	// Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
	// +optional
	InsecureSkipVerify *bool `json:"insecureSkipVerify,omitempty" tf:"insecure_skip_verify"`
	// The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment variable.
	OrgID *int64 `json:"orgID" tf:"org_id"`
	// A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
	// +optional
	SmAccessToken *string `json:"-" sensitive:"true" tf:"sm_access_token"`
	// Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable.
	// +optional
	SmURL *string `json:"smURL,omitempty" tf:"sm_url"`
	// Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT` environment variable.
	// +optional
	TlsCert *string `json:"tlsCert,omitempty" tf:"tls_cert"`
	// Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY` environment variable.
	// +optional
	TlsKey *string `json:"tlsKey,omitempty" tf:"tls_key"`
	// The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
	Url *string `json:"url" tf:"url"`
}
