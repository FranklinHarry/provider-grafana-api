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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/alert/v1alpha1"
	builtinv1alpha1 "kubeform.dev/provider-grafana-api/apis/builtin/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-grafana-api/apis/dashboard/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-grafana-api/apis/data/v1alpha1"
	folderv1alpha1 "kubeform.dev/provider-grafana-api/apis/folder/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-grafana-api/apis/organization/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-grafana-api/apis/role/v1alpha1"
	syntheticv1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-grafana-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-grafana-api/apis/user/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=alert.grafana.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("notifications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().Notifications().Informer()}, nil

		// Group=builtin.grafana.kubeform.com, Version=v1alpha1
	case builtinv1alpha1.SchemeGroupVersion.WithResource("roleassignments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Builtin().V1alpha1().RoleAssignments().Informer()}, nil

		// Group=dashboard.grafana.kubeform.com, Version=v1alpha1
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("dashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Dashboards().Informer()}, nil
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("permissions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Permissions().Informer()}, nil

		// Group=data.grafana.kubeform.com, Version=v1alpha1
	case datav1alpha1.SchemeGroupVersion.WithResource("sources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Data().V1alpha1().Sources().Informer()}, nil

		// Group=folder.grafana.kubeform.com, Version=v1alpha1
	case folderv1alpha1.SchemeGroupVersion.WithResource("folders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Folder().V1alpha1().Folders().Informer()}, nil
	case folderv1alpha1.SchemeGroupVersion.WithResource("permissions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Folder().V1alpha1().Permissions().Informer()}, nil

		// Group=organization.grafana.kubeform.com, Version=v1alpha1
	case organizationv1alpha1.SchemeGroupVersion.WithResource("organizations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Organization().V1alpha1().Organizations().Informer()}, nil

		// Group=role.grafana.kubeform.com, Version=v1alpha1
	case rolev1alpha1.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Role().V1alpha1().Roles().Informer()}, nil

		// Group=synthetic.grafana.kubeform.com, Version=v1alpha1
	case syntheticv1alpha1.SchemeGroupVersion.WithResource("monitoringchecks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetic().V1alpha1().MonitoringChecks().Informer()}, nil
	case syntheticv1alpha1.SchemeGroupVersion.WithResource("monitoringprobes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetic().V1alpha1().MonitoringProbes().Informer()}, nil

		// Group=team.grafana.kubeform.com, Version=v1alpha1
	case teamv1alpha1.SchemeGroupVersion.WithResource("externalgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Team().V1alpha1().ExternalGroups().Informer()}, nil
	case teamv1alpha1.SchemeGroupVersion.WithResource("preferenceses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Team().V1alpha1().Preferenceses().Informer()}, nil
	case teamv1alpha1.SchemeGroupVersion.WithResource("teams"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Team().V1alpha1().Teams().Informer()}, nil

		// Group=user.grafana.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().Users().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
