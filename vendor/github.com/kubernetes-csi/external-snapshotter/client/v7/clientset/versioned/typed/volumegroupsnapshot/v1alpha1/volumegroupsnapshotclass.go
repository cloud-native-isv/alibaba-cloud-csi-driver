/*
Copyright 2024 The Kubernetes Authors.

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

	v1alpha1 "github.com/kubernetes-csi/external-snapshotter/client/v7/apis/volumegroupsnapshot/v1alpha1"
	scheme "github.com/kubernetes-csi/external-snapshotter/client/v7/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeGroupSnapshotClassesGetter has a method to return a VolumeGroupSnapshotClassInterface.
// A group's client should implement this interface.
type VolumeGroupSnapshotClassesGetter interface {
	VolumeGroupSnapshotClasses() VolumeGroupSnapshotClassInterface
}

// VolumeGroupSnapshotClassInterface has methods to work with VolumeGroupSnapshotClass resources.
type VolumeGroupSnapshotClassInterface interface {
	Create(ctx context.Context, volumeGroupSnapshotClass *v1alpha1.VolumeGroupSnapshotClass, opts v1.CreateOptions) (*v1alpha1.VolumeGroupSnapshotClass, error)
	Update(ctx context.Context, volumeGroupSnapshotClass *v1alpha1.VolumeGroupSnapshotClass, opts v1.UpdateOptions) (*v1alpha1.VolumeGroupSnapshotClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumeGroupSnapshotClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeGroupSnapshotClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeGroupSnapshotClass, err error)
	VolumeGroupSnapshotClassExpansion
}

// volumeGroupSnapshotClasses implements VolumeGroupSnapshotClassInterface
type volumeGroupSnapshotClasses struct {
	client rest.Interface
}

// newVolumeGroupSnapshotClasses returns a VolumeGroupSnapshotClasses
func newVolumeGroupSnapshotClasses(c *GroupsnapshotV1alpha1Client) *volumeGroupSnapshotClasses {
	return &volumeGroupSnapshotClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeGroupSnapshotClass, and returns the corresponding volumeGroupSnapshotClass object, and an error if there is any.
func (c *volumeGroupSnapshotClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeGroupSnapshotClass, err error) {
	result = &v1alpha1.VolumeGroupSnapshotClass{}
	err = c.client.Get().
		Resource("volumegroupsnapshotclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeGroupSnapshotClasses that match those selectors.
func (c *volumeGroupSnapshotClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeGroupSnapshotClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeGroupSnapshotClassList{}
	err = c.client.Get().
		Resource("volumegroupsnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeGroupSnapshotClasses.
func (c *volumeGroupSnapshotClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumegroupsnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeGroupSnapshotClass and creates it.  Returns the server's representation of the volumeGroupSnapshotClass, and an error, if there is any.
func (c *volumeGroupSnapshotClasses) Create(ctx context.Context, volumeGroupSnapshotClass *v1alpha1.VolumeGroupSnapshotClass, opts v1.CreateOptions) (result *v1alpha1.VolumeGroupSnapshotClass, err error) {
	result = &v1alpha1.VolumeGroupSnapshotClass{}
	err = c.client.Post().
		Resource("volumegroupsnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeGroupSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeGroupSnapshotClass and updates it. Returns the server's representation of the volumeGroupSnapshotClass, and an error, if there is any.
func (c *volumeGroupSnapshotClasses) Update(ctx context.Context, volumeGroupSnapshotClass *v1alpha1.VolumeGroupSnapshotClass, opts v1.UpdateOptions) (result *v1alpha1.VolumeGroupSnapshotClass, err error) {
	result = &v1alpha1.VolumeGroupSnapshotClass{}
	err = c.client.Put().
		Resource("volumegroupsnapshotclasses").
		Name(volumeGroupSnapshotClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeGroupSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeGroupSnapshotClass and deletes it. Returns an error if one occurs.
func (c *volumeGroupSnapshotClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumegroupsnapshotclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeGroupSnapshotClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumegroupsnapshotclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeGroupSnapshotClass.
func (c *volumeGroupSnapshotClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeGroupSnapshotClass, err error) {
	result = &v1alpha1.VolumeGroupSnapshotClass{}
	err = c.client.Patch(pt).
		Resource("volumegroupsnapshotclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
