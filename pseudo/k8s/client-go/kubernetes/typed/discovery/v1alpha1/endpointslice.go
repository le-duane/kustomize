/*
Copyright The Kubernetes Authors.

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
	"time"

	v1alpha1 "sigs.k8s.io/kustomize/pseudo/k8s/api/discovery/v1alpha1"
	v1 "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/apis/meta/v1"
	types "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/types"
	watch "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/watch"
	scheme "sigs.k8s.io/kustomize/pseudo/k8s/client-go/kubernetes/scheme"
	rest "sigs.k8s.io/kustomize/pseudo/k8s/client-go/rest"
)

// EndpointSlicesGetter has a method to return a EndpointSliceInterface.
// A group's client should implement this interface.
type EndpointSlicesGetter interface {
	EndpointSlices(namespace string) EndpointSliceInterface
}

// EndpointSliceInterface has methods to work with EndpointSlice resources.
type EndpointSliceInterface interface {
	Create(*v1alpha1.EndpointSlice) (*v1alpha1.EndpointSlice, error)
	Update(*v1alpha1.EndpointSlice) (*v1alpha1.EndpointSlice, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EndpointSlice, error)
	List(opts v1.ListOptions) (*v1alpha1.EndpointSliceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EndpointSlice, err error)
	EndpointSliceExpansion
}

// endpointSlices implements EndpointSliceInterface
type endpointSlices struct {
	client rest.Interface
	ns     string
}

// newEndpointSlices returns a EndpointSlices
func newEndpointSlices(c *DiscoveryV1alpha1Client, namespace string) *endpointSlices {
	return &endpointSlices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpointSlice, and returns the corresponding endpointSlice object, and an error if there is any.
func (c *endpointSlices) Get(name string, options v1.GetOptions) (result *v1alpha1.EndpointSlice, err error) {
	result = &v1alpha1.EndpointSlice{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors.
func (c *endpointSlices) List(opts v1.ListOptions) (result *v1alpha1.EndpointSliceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EndpointSliceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpointSlices.
func (c *endpointSlices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a endpointSlice and creates it.  Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *endpointSlices) Create(endpointSlice *v1alpha1.EndpointSlice) (result *v1alpha1.EndpointSlice, err error) {
	result = &v1alpha1.EndpointSlice{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpointslices").
		Body(endpointSlice).
		Do().
		Into(result)
	return
}

// Update takes the representation of a endpointSlice and updates it. Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *endpointSlices) Update(endpointSlice *v1alpha1.EndpointSlice) (result *v1alpha1.EndpointSlice, err error) {
	result = &v1alpha1.EndpointSlice{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(endpointSlice.Name).
		Body(endpointSlice).
		Do().
		Into(result)
	return
}

// Delete takes name of the endpointSlice and deletes it. Returns an error if one occurs.
func (c *endpointSlices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpointSlices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched endpointSlice.
func (c *endpointSlices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EndpointSlice, err error) {
	result = &v1alpha1.EndpointSlice{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpointslices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
