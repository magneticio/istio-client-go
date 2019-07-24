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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	configv1alpha2 "github.com/magneticio/istio-client-go/pkg/apis/config/v1alpha2"
	versioned "github.com/magneticio/istio-client-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/magneticio/istio-client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/magneticio/istio-client-go/pkg/client/listers/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InstanceInformer provides access to a shared informer and lister for
// Instances.
type InstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.InstanceLister
}

type instanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInstanceInformer constructs a new informer for Instance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInstanceInformer constructs a new informer for Instance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthenticationV1alpha2().Instances(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthenticationV1alpha2().Instances(namespace).Watch(options)
			},
		},
		&configv1alpha2.Instance{},
		resyncPeriod,
		indexers,
	)
}

func (f *instanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *instanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha2.Instance{}, f.defaultInformer)
}

func (f *instanceInformer) Lister() v1alpha2.InstanceLister {
	return v1alpha2.NewInstanceLister(f.Informer().GetIndexer())
}
