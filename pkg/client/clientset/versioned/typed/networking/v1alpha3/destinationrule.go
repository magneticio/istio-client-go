/*
Portions Copyright 2019 The Kubernetes Authors.
Portions Copyright 2019 Aspen Mesh Authors.
Portions Copyright 2019 Vamp Authors.

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

package v1alpha3

import (
	v1alpha3 "github.com/magneticio/istio-client-go/pkg/apis/networking/v1alpha3"
	scheme "github.com/magneticio/istio-client-go/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DestinationRulesGetter has a method to return a DestinationRuleInterface.
// A group's client should implement this interface.
type DestinationRulesGetter interface {
	DestinationRules(namespace string) DestinationRuleInterface
}

// DestinationRuleInterface has methods to work with DestinationRule resources.
type DestinationRuleInterface interface {
	Create(*v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error)
	Update(*v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha3.DestinationRule, error)
	List(opts v1.ListOptions) (*v1alpha3.DestinationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.DestinationRule, err error)
	DestinationRuleExpansion
}

// destinationRules implements DestinationRuleInterface
type destinationRules struct {
	client rest.Interface
	ns     string
}

// newDestinationRules returns a DestinationRules
func newDestinationRules(c *NetworkingV1alpha3Client, namespace string) *destinationRules {
	return &destinationRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the destinationRule, and returns the corresponding destinationRule object, and an error if there is any.
func (c *destinationRules) Get(name string, options v1.GetOptions) (result *v1alpha3.DestinationRule, err error) {
	result = &v1alpha3.DestinationRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("destinationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DestinationRules that match those selectors.
func (c *destinationRules) List(opts v1.ListOptions) (result *v1alpha3.DestinationRuleList, err error) {
	result = &v1alpha3.DestinationRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("destinationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested destinationRules.
func (c *destinationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("destinationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a destinationRule and creates it.  Returns the server's representation of the destinationRule, and an error, if there is any.
func (c *destinationRules) Create(destinationRule *v1alpha3.DestinationRule) (result *v1alpha3.DestinationRule, err error) {
	result = &v1alpha3.DestinationRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("destinationrules").
		Body(destinationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a destinationRule and updates it. Returns the server's representation of the destinationRule, and an error, if there is any.
func (c *destinationRules) Update(destinationRule *v1alpha3.DestinationRule) (result *v1alpha3.DestinationRule, err error) {
	result = &v1alpha3.DestinationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("destinationrules").
		Name(destinationRule.Name).
		Body(destinationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the destinationRule and deletes it. Returns an error if one occurs.
func (c *destinationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("destinationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *destinationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("destinationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched destinationRule.
func (c *destinationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.DestinationRule, err error) {
	result = &v1alpha3.DestinationRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("destinationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
