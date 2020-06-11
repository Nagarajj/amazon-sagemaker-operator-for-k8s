/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

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

package internalversion

import (
	v1 "github.com/aws/amazon-sagemaker-operator-for-k8s/api/endpointconfig/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointConfigLister helps list EndpointConfigs.
type EndpointConfigLister interface {
	// List lists all EndpointConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.EndpointConfig, err error)
	// EndpointConfigs returns an object that can list and get EndpointConfigs.
	EndpointConfigs(namespace string) EndpointConfigNamespaceLister
	EndpointConfigListerExpansion
}

// endpointConfigLister implements the EndpointConfigLister interface.
type endpointConfigLister struct {
	indexer cache.Indexer
}

// NewEndpointConfigLister returns a new EndpointConfigLister.
func NewEndpointConfigLister(indexer cache.Indexer) EndpointConfigLister {
	return &endpointConfigLister{indexer: indexer}
}

// List lists all EndpointConfigs in the indexer.
func (s *endpointConfigLister) List(selector labels.Selector) (ret []*v1.EndpointConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EndpointConfig))
	})
	return ret, err
}

// EndpointConfigs returns an object that can list and get EndpointConfigs.
func (s *endpointConfigLister) EndpointConfigs(namespace string) EndpointConfigNamespaceLister {
	return endpointConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointConfigNamespaceLister helps list and get EndpointConfigs.
type EndpointConfigNamespaceLister interface {
	// List lists all EndpointConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.EndpointConfig, err error)
	// Get retrieves the EndpointConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.EndpointConfig, error)
	EndpointConfigNamespaceListerExpansion
}

// endpointConfigNamespaceLister implements the EndpointConfigNamespaceLister
// interface.
type endpointConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointConfigs in the indexer for a given namespace.
func (s endpointConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.EndpointConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EndpointConfig))
	})
	return ret, err
}

// Get retrieves the EndpointConfig from the indexer for a given namespace and name.
func (s endpointConfigNamespaceLister) Get(name string) (*v1.EndpointConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("endpointconfig"), name)
	}
	return obj.(*v1.EndpointConfig), nil
}