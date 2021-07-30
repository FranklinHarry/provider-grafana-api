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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"
	scheme "kubeform.dev/provider-grafana-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MonitoringProbesGetter has a method to return a MonitoringProbeInterface.
// A group's client should implement this interface.
type MonitoringProbesGetter interface {
	MonitoringProbes(namespace string) MonitoringProbeInterface
}

// MonitoringProbeInterface has methods to work with MonitoringProbe resources.
type MonitoringProbeInterface interface {
	Create(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.CreateOptions) (*v1alpha1.MonitoringProbe, error)
	Update(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (*v1alpha1.MonitoringProbe, error)
	UpdateStatus(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (*v1alpha1.MonitoringProbe, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MonitoringProbe, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MonitoringProbeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitoringProbe, err error)
	MonitoringProbeExpansion
}

// monitoringProbes implements MonitoringProbeInterface
type monitoringProbes struct {
	client rest.Interface
	ns     string
}

// newMonitoringProbes returns a MonitoringProbes
func newMonitoringProbes(c *SyntheticV1alpha1Client, namespace string) *monitoringProbes {
	return &monitoringProbes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the monitoringProbe, and returns the corresponding monitoringProbe object, and an error if there is any.
func (c *monitoringProbes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitoringProbe, err error) {
	result = &v1alpha1.MonitoringProbe{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("monitoringprobes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MonitoringProbes that match those selectors.
func (c *monitoringProbes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitoringProbeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MonitoringProbeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("monitoringprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested monitoringProbes.
func (c *monitoringProbes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("monitoringprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a monitoringProbe and creates it.  Returns the server's representation of the monitoringProbe, and an error, if there is any.
func (c *monitoringProbes) Create(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.CreateOptions) (result *v1alpha1.MonitoringProbe, err error) {
	result = &v1alpha1.MonitoringProbe{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("monitoringprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(monitoringProbe).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a monitoringProbe and updates it. Returns the server's representation of the monitoringProbe, and an error, if there is any.
func (c *monitoringProbes) Update(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (result *v1alpha1.MonitoringProbe, err error) {
	result = &v1alpha1.MonitoringProbe{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("monitoringprobes").
		Name(monitoringProbe.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(monitoringProbe).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *monitoringProbes) UpdateStatus(ctx context.Context, monitoringProbe *v1alpha1.MonitoringProbe, opts v1.UpdateOptions) (result *v1alpha1.MonitoringProbe, err error) {
	result = &v1alpha1.MonitoringProbe{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("monitoringprobes").
		Name(monitoringProbe.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(monitoringProbe).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the monitoringProbe and deletes it. Returns an error if one occurs.
func (c *monitoringProbes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("monitoringprobes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *monitoringProbes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("monitoringprobes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched monitoringProbe.
func (c *monitoringProbes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitoringProbe, err error) {
	result = &v1alpha1.MonitoringProbe{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("monitoringprobes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}