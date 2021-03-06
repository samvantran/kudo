/*

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
	v1alpha1 "github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PlanExecutionLister helps list PlanExecutions.
type PlanExecutionLister interface {
	// List lists all PlanExecutions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PlanExecution, err error)
	// PlanExecutions returns an object that can list and get PlanExecutions.
	PlanExecutions(namespace string) PlanExecutionNamespaceLister
	PlanExecutionListerExpansion
}

// planExecutionLister implements the PlanExecutionLister interface.
type planExecutionLister struct {
	indexer cache.Indexer
}

// NewPlanExecutionLister returns a new PlanExecutionLister.
func NewPlanExecutionLister(indexer cache.Indexer) PlanExecutionLister {
	return &planExecutionLister{indexer: indexer}
}

// List lists all PlanExecutions in the indexer.
func (s *planExecutionLister) List(selector labels.Selector) (ret []*v1alpha1.PlanExecution, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlanExecution))
	})
	return ret, err
}

// PlanExecutions returns an object that can list and get PlanExecutions.
func (s *planExecutionLister) PlanExecutions(namespace string) PlanExecutionNamespaceLister {
	return planExecutionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlanExecutionNamespaceLister helps list and get PlanExecutions.
type PlanExecutionNamespaceLister interface {
	// List lists all PlanExecutions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PlanExecution, err error)
	// Get retrieves the PlanExecution from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PlanExecution, error)
	PlanExecutionNamespaceListerExpansion
}

// planExecutionNamespaceLister implements the PlanExecutionNamespaceLister
// interface.
type planExecutionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlanExecutions in the indexer for a given namespace.
func (s planExecutionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlanExecution, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlanExecution))
	})
	return ret, err
}

// Get retrieves the PlanExecution from the indexer for a given namespace and name.
func (s planExecutionNamespaceLister) Get(name string) (*v1alpha1.PlanExecution, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("planexecution"), name)
	}
	return obj.(*v1alpha1.PlanExecution), nil
}
