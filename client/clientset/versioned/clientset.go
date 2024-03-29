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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	alertv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/alert/v1alpha1"
	builtinv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/builtin/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/dashboard/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/data/v1alpha1"
	folderv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/folder/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/organization/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/role/v1alpha1"
	syntheticv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/synthetic/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-grafana-api/client/clientset/versioned/typed/user/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AlertV1alpha1() alertv1alpha1.AlertV1alpha1Interface
	BuiltinV1alpha1() builtinv1alpha1.BuiltinV1alpha1Interface
	DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface
	DataV1alpha1() datav1alpha1.DataV1alpha1Interface
	FolderV1alpha1() folderv1alpha1.FolderV1alpha1Interface
	OrganizationV1alpha1() organizationv1alpha1.OrganizationV1alpha1Interface
	RoleV1alpha1() rolev1alpha1.RoleV1alpha1Interface
	SyntheticV1alpha1() syntheticv1alpha1.SyntheticV1alpha1Interface
	TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface
	UserV1alpha1() userv1alpha1.UserV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	alertV1alpha1        *alertv1alpha1.AlertV1alpha1Client
	builtinV1alpha1      *builtinv1alpha1.BuiltinV1alpha1Client
	dashboardV1alpha1    *dashboardv1alpha1.DashboardV1alpha1Client
	dataV1alpha1         *datav1alpha1.DataV1alpha1Client
	folderV1alpha1       *folderv1alpha1.FolderV1alpha1Client
	organizationV1alpha1 *organizationv1alpha1.OrganizationV1alpha1Client
	roleV1alpha1         *rolev1alpha1.RoleV1alpha1Client
	syntheticV1alpha1    *syntheticv1alpha1.SyntheticV1alpha1Client
	teamV1alpha1         *teamv1alpha1.TeamV1alpha1Client
	userV1alpha1         *userv1alpha1.UserV1alpha1Client
}

// AlertV1alpha1 retrieves the AlertV1alpha1Client
func (c *Clientset) AlertV1alpha1() alertv1alpha1.AlertV1alpha1Interface {
	return c.alertV1alpha1
}

// BuiltinV1alpha1 retrieves the BuiltinV1alpha1Client
func (c *Clientset) BuiltinV1alpha1() builtinv1alpha1.BuiltinV1alpha1Interface {
	return c.builtinV1alpha1
}

// DashboardV1alpha1 retrieves the DashboardV1alpha1Client
func (c *Clientset) DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface {
	return c.dashboardV1alpha1
}

// DataV1alpha1 retrieves the DataV1alpha1Client
func (c *Clientset) DataV1alpha1() datav1alpha1.DataV1alpha1Interface {
	return c.dataV1alpha1
}

// FolderV1alpha1 retrieves the FolderV1alpha1Client
func (c *Clientset) FolderV1alpha1() folderv1alpha1.FolderV1alpha1Interface {
	return c.folderV1alpha1
}

// OrganizationV1alpha1 retrieves the OrganizationV1alpha1Client
func (c *Clientset) OrganizationV1alpha1() organizationv1alpha1.OrganizationV1alpha1Interface {
	return c.organizationV1alpha1
}

// RoleV1alpha1 retrieves the RoleV1alpha1Client
func (c *Clientset) RoleV1alpha1() rolev1alpha1.RoleV1alpha1Interface {
	return c.roleV1alpha1
}

// SyntheticV1alpha1 retrieves the SyntheticV1alpha1Client
func (c *Clientset) SyntheticV1alpha1() syntheticv1alpha1.SyntheticV1alpha1Interface {
	return c.syntheticV1alpha1
}

// TeamV1alpha1 retrieves the TeamV1alpha1Client
func (c *Clientset) TeamV1alpha1() teamv1alpha1.TeamV1alpha1Interface {
	return c.teamV1alpha1
}

// UserV1alpha1 retrieves the UserV1alpha1Client
func (c *Clientset) UserV1alpha1() userv1alpha1.UserV1alpha1Interface {
	return c.userV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.alertV1alpha1, err = alertv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.builtinV1alpha1, err = builtinv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dashboardV1alpha1, err = dashboardv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dataV1alpha1, err = datav1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.folderV1alpha1, err = folderv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.organizationV1alpha1, err = organizationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.roleV1alpha1, err = rolev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.syntheticV1alpha1, err = syntheticv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.teamV1alpha1, err = teamv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.userV1alpha1, err = userv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.alertV1alpha1 = alertv1alpha1.NewForConfigOrDie(c)
	cs.builtinV1alpha1 = builtinv1alpha1.NewForConfigOrDie(c)
	cs.dashboardV1alpha1 = dashboardv1alpha1.NewForConfigOrDie(c)
	cs.dataV1alpha1 = datav1alpha1.NewForConfigOrDie(c)
	cs.folderV1alpha1 = folderv1alpha1.NewForConfigOrDie(c)
	cs.organizationV1alpha1 = organizationv1alpha1.NewForConfigOrDie(c)
	cs.roleV1alpha1 = rolev1alpha1.NewForConfigOrDie(c)
	cs.syntheticV1alpha1 = syntheticv1alpha1.NewForConfigOrDie(c)
	cs.teamV1alpha1 = teamv1alpha1.NewForConfigOrDie(c)
	cs.userV1alpha1 = userv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.alertV1alpha1 = alertv1alpha1.New(c)
	cs.builtinV1alpha1 = builtinv1alpha1.New(c)
	cs.dashboardV1alpha1 = dashboardv1alpha1.New(c)
	cs.dataV1alpha1 = datav1alpha1.New(c)
	cs.folderV1alpha1 = folderv1alpha1.New(c)
	cs.organizationV1alpha1 = organizationv1alpha1.New(c)
	cs.roleV1alpha1 = rolev1alpha1.New(c)
	cs.syntheticV1alpha1 = syntheticv1alpha1.New(c)
	cs.teamV1alpha1 = teamv1alpha1.New(c)
	cs.userV1alpha1 = userv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
