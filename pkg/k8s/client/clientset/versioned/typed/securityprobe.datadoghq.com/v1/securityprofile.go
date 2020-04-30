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

package v1

import (
	"time"

	v1 "github.com/Gui774ume/network-security-probe/pkg/k8s/apis/securityprobe.datadoghq.com/v1"
	scheme "github.com/Gui774ume/network-security-probe/pkg/k8s/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecurityProfilesGetter has a method to return a SecurityProfileInterface.
// A group's client should implement this interface.
type SecurityProfilesGetter interface {
	SecurityProfiles(namespace string) SecurityProfileInterface
}

// SecurityProfileInterface has methods to work with SecurityProfile resources.
type SecurityProfileInterface interface {
	Create(*v1.SecurityProfile) (*v1.SecurityProfile, error)
	Update(*v1.SecurityProfile) (*v1.SecurityProfile, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SecurityProfile, error)
	List(opts metav1.ListOptions) (*v1.SecurityProfileList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SecurityProfile, err error)
	SecurityProfileExpansion
}

// securityProfiles implements SecurityProfileInterface
type securityProfiles struct {
	client rest.Interface
	ns     string
}

// newSecurityProfiles returns a SecurityProfiles
func newSecurityProfiles(c *SecurityprobeV1Client, namespace string) *securityProfiles {
	return &securityProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the securityProfile, and returns the corresponding securityProfile object, and an error if there is any.
func (c *securityProfiles) Get(name string, options metav1.GetOptions) (result *v1.SecurityProfile, err error) {
	result = &v1.SecurityProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securityprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityProfiles that match those selectors.
func (c *securityProfiles) List(opts metav1.ListOptions) (result *v1.SecurityProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SecurityProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securityprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityProfiles.
func (c *securityProfiles) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securityprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a securityProfile and creates it.  Returns the server's representation of the securityProfile, and an error, if there is any.
func (c *securityProfiles) Create(securityProfile *v1.SecurityProfile) (result *v1.SecurityProfile, err error) {
	result = &v1.SecurityProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securityprofiles").
		Body(securityProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a securityProfile and updates it. Returns the server's representation of the securityProfile, and an error, if there is any.
func (c *securityProfiles) Update(securityProfile *v1.SecurityProfile) (result *v1.SecurityProfile, err error) {
	result = &v1.SecurityProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securityprofiles").
		Name(securityProfile.Name).
		Body(securityProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the securityProfile and deletes it. Returns an error if one occurs.
func (c *securityProfiles) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securityprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityProfiles) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securityprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched securityProfile.
func (c *securityProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SecurityProfile, err error) {
	result = &v1.SecurityProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securityprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}