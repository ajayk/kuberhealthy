// Copyright 2018 Comcast Cable Communications Management, LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package khcheckcrd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// KuberhealthyCheckClient holds client data for talking to Kubernetes about
// the khstate custom resource
type KuberhealthyCheckClient struct {
	restClient rest.Interface
	ns         string
}

// Create creates a new resource for this CRD
func (c *KuberhealthyCheckClient) Create(check *KuberhealthyCheck, resource string) (*KuberhealthyCheck, error) {
	mu.Lock()
	defer mu.Unlock()
	result := KuberhealthyCheck{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource(resource).
		Body(check).
		Do().
		Into(&result)
	return &result, err
}

// Delete deletes a resource for this CRD
func (c *KuberhealthyCheckClient) Delete(resource string, name string) (*KuberhealthyCheck, error) {
	mu.Lock()
	defer mu.Unlock()
	result := KuberhealthyCheck{}
	err := c.restClient.
		Delete().
		Namespace(c.ns).
		Resource(resource).
		Name(name).
		Do().
		Into(&result)
	return &result, err
}

// Update updates a resource for this CRD
func (c *KuberhealthyCheckClient) Update(check *KuberhealthyCheck, resource string, name string) (*KuberhealthyCheck, error) {
	mu.Lock()
	defer mu.Unlock()
	result := KuberhealthyCheck{}
	// err := c.restClient.Verb("update").Namespace(c.ns).Resource(resource).Name(name).Do().Into(&result)
	err := c.restClient.
		Put().
		Namespace(c.ns).
		Resource(resource).
		Body(check).
		Name(name).
		Do().
		Into(&result)
	return &result, err
}

// Get fetches a resource of this CRD
func (c *KuberhealthyCheckClient) Get(opts metav1.GetOptions, resource string, name string) (*KuberhealthyCheck, error) {
	mu.Lock()
	defer mu.Unlock()
	result := KuberhealthyCheck{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(resource).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)
	return &result, err
}

// List lists resources for this CRD
func (c *KuberhealthyCheckClient) List(opts metav1.ListOptions, resource string) (*KuberhealthyCheckList, error) {
	mu.Lock()
	defer mu.Unlock()
	result := KuberhealthyCheckList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(resource).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)
	return &result, err
}

// TODO - implement watches?
