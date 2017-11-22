package fake

import (
	template_v1 "github.com/openshift/api/template/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplates implements TemplateResourceInterface
type FakeTemplates struct {
	Fake *FakeTemplateV1
	ns   string
}

var templatesResource = schema.GroupVersionResource{Group: "template.openshift.io", Version: "v1", Resource: "templates"}

var templatesKind = schema.GroupVersionKind{Group: "template.openshift.io", Version: "v1", Kind: "Template"}

// Get takes name of the templateResource, and returns the corresponding templateResource object, and an error if there is any.
func (c *FakeTemplates) Get(name string, options v1.GetOptions) (result *template_v1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templatesResource, c.ns, name), &template_v1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.Template), err
}

// List takes label and field selectors, and returns the list of Templates that match those selectors.
func (c *FakeTemplates) List(opts v1.ListOptions) (result *template_v1.TemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templatesResource, templatesKind, c.ns, opts), &template_v1.TemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &template_v1.TemplateList{}
	for _, item := range obj.(*template_v1.TemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templates.
func (c *FakeTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templatesResource, c.ns, opts))

}

// Create takes the representation of a templateResource and creates it.  Returns the server's representation of the templateResource, and an error, if there is any.
func (c *FakeTemplates) Create(templateResource *template_v1.Template) (result *template_v1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templatesResource, c.ns, templateResource), &template_v1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.Template), err
}

// Update takes the representation of a templateResource and updates it. Returns the server's representation of the templateResource, and an error, if there is any.
func (c *FakeTemplates) Update(templateResource *template_v1.Template) (result *template_v1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templatesResource, c.ns, templateResource), &template_v1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.Template), err
}

// Delete takes name of the templateResource and deletes it. Returns an error if one occurs.
func (c *FakeTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(templatesResource, c.ns, name), &template_v1.Template{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &template_v1.TemplateList{})
	return err
}

// Patch applies the patch and returns the patched templateResource.
func (c *FakeTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *template_v1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templatesResource, c.ns, name, data, subresources...), &template_v1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*template_v1.Template), err
}
