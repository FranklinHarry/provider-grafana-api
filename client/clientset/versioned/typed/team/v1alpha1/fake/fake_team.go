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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/team/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTeams implements TeamInterface
type FakeTeams struct {
	Fake *FakeTeamV1alpha1
	ns   string
}

var teamsResource = schema.GroupVersionResource{Group: "team.grafana.kubeform.com", Version: "v1alpha1", Resource: "teams"}

var teamsKind = schema.GroupVersionKind{Group: "team.grafana.kubeform.com", Version: "v1alpha1", Kind: "Team"}

// Get takes name of the team, and returns the corresponding team object, and an error if there is any.
func (c *FakeTeams) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(teamsResource, c.ns, name), &v1alpha1.Team{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Team), err
}

// List takes label and field selectors, and returns the list of Teams that match those selectors.
func (c *FakeTeams) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TeamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(teamsResource, teamsKind, c.ns, opts), &v1alpha1.TeamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TeamList{ListMeta: obj.(*v1alpha1.TeamList).ListMeta}
	for _, item := range obj.(*v1alpha1.TeamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested teams.
func (c *FakeTeams) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(teamsResource, c.ns, opts))

}

// Create takes the representation of a team and creates it.  Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Create(ctx context.Context, team *v1alpha1.Team, opts v1.CreateOptions) (result *v1alpha1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(teamsResource, c.ns, team), &v1alpha1.Team{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Team), err
}

// Update takes the representation of a team and updates it. Returns the server's representation of the team, and an error, if there is any.
func (c *FakeTeams) Update(ctx context.Context, team *v1alpha1.Team, opts v1.UpdateOptions) (result *v1alpha1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(teamsResource, c.ns, team), &v1alpha1.Team{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Team), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTeams) UpdateStatus(ctx context.Context, team *v1alpha1.Team, opts v1.UpdateOptions) (*v1alpha1.Team, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(teamsResource, "status", c.ns, team), &v1alpha1.Team{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Team), err
}

// Delete takes name of the team and deletes it. Returns an error if one occurs.
func (c *FakeTeams) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(teamsResource, c.ns, name), &v1alpha1.Team{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTeams) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(teamsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TeamList{})
	return err
}

// Patch applies the patch and returns the patched team.
func (c *FakeTeams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Team, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(teamsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Team{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Team), err
}
