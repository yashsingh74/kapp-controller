// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "carvel.dev/kapp-controller/pkg/apis/internalpackaging/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInternalPackageMetadatas implements InternalPackageMetadataInterface
type FakeInternalPackageMetadatas struct {
	Fake *FakeInternalV1alpha1
	ns   string
}

var internalpackagemetadatasResource = v1alpha1.SchemeGroupVersion.WithResource("internalpackagemetadatas")

var internalpackagemetadatasKind = v1alpha1.SchemeGroupVersion.WithKind("InternalPackageMetadata")

// Get takes name of the internalPackageMetadata, and returns the corresponding internalPackageMetadata object, and an error if there is any.
func (c *FakeInternalPackageMetadatas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InternalPackageMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(internalpackagemetadatasResource, c.ns, name), &v1alpha1.InternalPackageMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InternalPackageMetadata), err
}

// List takes label and field selectors, and returns the list of InternalPackageMetadatas that match those selectors.
func (c *FakeInternalPackageMetadatas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InternalPackageMetadataList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(internalpackagemetadatasResource, internalpackagemetadatasKind, c.ns, opts), &v1alpha1.InternalPackageMetadataList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InternalPackageMetadataList{ListMeta: obj.(*v1alpha1.InternalPackageMetadataList).ListMeta}
	for _, item := range obj.(*v1alpha1.InternalPackageMetadataList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested internalPackageMetadatas.
func (c *FakeInternalPackageMetadatas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(internalpackagemetadatasResource, c.ns, opts))

}

// Create takes the representation of a internalPackageMetadata and creates it.  Returns the server's representation of the internalPackageMetadata, and an error, if there is any.
func (c *FakeInternalPackageMetadatas) Create(ctx context.Context, internalPackageMetadata *v1alpha1.InternalPackageMetadata, opts v1.CreateOptions) (result *v1alpha1.InternalPackageMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(internalpackagemetadatasResource, c.ns, internalPackageMetadata), &v1alpha1.InternalPackageMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InternalPackageMetadata), err
}

// Update takes the representation of a internalPackageMetadata and updates it. Returns the server's representation of the internalPackageMetadata, and an error, if there is any.
func (c *FakeInternalPackageMetadatas) Update(ctx context.Context, internalPackageMetadata *v1alpha1.InternalPackageMetadata, opts v1.UpdateOptions) (result *v1alpha1.InternalPackageMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(internalpackagemetadatasResource, c.ns, internalPackageMetadata), &v1alpha1.InternalPackageMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InternalPackageMetadata), err
}

// Delete takes name of the internalPackageMetadata and deletes it. Returns an error if one occurs.
func (c *FakeInternalPackageMetadatas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(internalpackagemetadatasResource, c.ns, name, opts), &v1alpha1.InternalPackageMetadata{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInternalPackageMetadatas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(internalpackagemetadatasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InternalPackageMetadataList{})
	return err
}

// Patch applies the patch and returns the patched internalPackageMetadata.
func (c *FakeInternalPackageMetadatas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InternalPackageMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(internalpackagemetadatasResource, c.ns, name, pt, data, subresources...), &v1alpha1.InternalPackageMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InternalPackageMetadata), err
}
