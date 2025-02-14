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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/magneticio/istio-client-go/pkg/apis/authentication/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolicyLister helps list Policies.
type PolicyLister interface {
	// List lists all Policies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Policy, err error)
	// Policies returns an object that can list and get Policies.
	Policies(namespace string) PolicyNamespaceLister
	PolicyListerExpansion
}

// policyLister implements the PolicyLister interface.
type policyLister struct {
	indexer cache.Indexer
}

// NewPolicyLister returns a new PolicyLister.
func NewPolicyLister(indexer cache.Indexer) PolicyLister {
	return &policyLister{indexer: indexer}
}

// List lists all Policies in the indexer.
func (s *policyLister) List(selector labels.Selector) (ret []*v1alpha1.Policy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Policy))
	})
	return ret, err
}

// Policies returns an object that can list and get Policies.
func (s *policyLister) Policies(namespace string) PolicyNamespaceLister {
	return policyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyNamespaceLister helps list and get Policies.
type PolicyNamespaceLister interface {
	// List lists all Policies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Policy, err error)
	// Get retrieves the Policy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Policy, error)
	PolicyNamespaceListerExpansion
}

// policyNamespaceLister implements the PolicyNamespaceLister
// interface.
type policyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Policies in the indexer for a given namespace.
func (s policyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Policy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Policy))
	})
	return ret, err
}

// Get retrieves the Policy from the indexer for a given namespace and name.
func (s policyNamespaceLister) Get(name string) (*v1alpha1.Policy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("policy"), name)
	}
	return obj.(*v1alpha1.Policy), nil
}
