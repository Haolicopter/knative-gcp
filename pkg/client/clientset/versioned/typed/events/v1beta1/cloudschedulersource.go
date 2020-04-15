/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/google/knative-gcp/pkg/apis/events/v1beta1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudSchedulerSourcesGetter has a method to return a CloudSchedulerSourceInterface.
// A group's client should implement this interface.
type CloudSchedulerSourcesGetter interface {
	CloudSchedulerSources(namespace string) CloudSchedulerSourceInterface
}

// CloudSchedulerSourceInterface has methods to work with CloudSchedulerSource resources.
type CloudSchedulerSourceInterface interface {
	Create(*v1beta1.CloudSchedulerSource) (*v1beta1.CloudSchedulerSource, error)
	Update(*v1beta1.CloudSchedulerSource) (*v1beta1.CloudSchedulerSource, error)
	UpdateStatus(*v1beta1.CloudSchedulerSource) (*v1beta1.CloudSchedulerSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.CloudSchedulerSource, error)
	List(opts v1.ListOptions) (*v1beta1.CloudSchedulerSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CloudSchedulerSource, err error)
	CloudSchedulerSourceExpansion
}

// cloudSchedulerSources implements CloudSchedulerSourceInterface
type cloudSchedulerSources struct {
	client rest.Interface
	ns     string
}

// newCloudSchedulerSources returns a CloudSchedulerSources
func newCloudSchedulerSources(c *EventsV1beta1Client, namespace string) *cloudSchedulerSources {
	return &cloudSchedulerSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudSchedulerSource, and returns the corresponding cloudSchedulerSource object, and an error if there is any.
func (c *cloudSchedulerSources) Get(name string, options v1.GetOptions) (result *v1beta1.CloudSchedulerSource, err error) {
	result = &v1beta1.CloudSchedulerSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudSchedulerSources that match those selectors.
func (c *cloudSchedulerSources) List(opts v1.ListOptions) (result *v1beta1.CloudSchedulerSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CloudSchedulerSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudSchedulerSources.
func (c *cloudSchedulerSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudSchedulerSource and creates it.  Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *cloudSchedulerSources) Create(cloudSchedulerSource *v1beta1.CloudSchedulerSource) (result *v1beta1.CloudSchedulerSource, err error) {
	result = &v1beta1.CloudSchedulerSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Body(cloudSchedulerSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudSchedulerSource and updates it. Returns the server's representation of the cloudSchedulerSource, and an error, if there is any.
func (c *cloudSchedulerSources) Update(cloudSchedulerSource *v1beta1.CloudSchedulerSource) (result *v1beta1.CloudSchedulerSource, err error) {
	result = &v1beta1.CloudSchedulerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(cloudSchedulerSource.Name).
		Body(cloudSchedulerSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudSchedulerSources) UpdateStatus(cloudSchedulerSource *v1beta1.CloudSchedulerSource) (result *v1beta1.CloudSchedulerSource, err error) {
	result = &v1beta1.CloudSchedulerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(cloudSchedulerSource.Name).
		SubResource("status").
		Body(cloudSchedulerSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudSchedulerSource and deletes it. Returns an error if one occurs.
func (c *cloudSchedulerSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudSchedulerSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudschedulersources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudSchedulerSource.
func (c *cloudSchedulerSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CloudSchedulerSource, err error) {
	result = &v1beta1.CloudSchedulerSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudschedulersources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}