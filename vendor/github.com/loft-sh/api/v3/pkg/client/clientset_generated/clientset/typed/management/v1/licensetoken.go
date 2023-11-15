// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v3/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LicenseTokensGetter has a method to return a LicenseTokenInterface.
// A group's client should implement this interface.
type LicenseTokensGetter interface {
	LicenseTokens() LicenseTokenInterface
}

// LicenseTokenInterface has methods to work with LicenseToken resources.
type LicenseTokenInterface interface {
	Create(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.CreateOptions) (*v1.LicenseToken, error)
	Update(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.UpdateOptions) (*v1.LicenseToken, error)
	UpdateStatus(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.UpdateOptions) (*v1.LicenseToken, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LicenseToken, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LicenseTokenList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LicenseToken, err error)
	LicenseTokenExpansion
}

// licenseTokens implements LicenseTokenInterface
type licenseTokens struct {
	client rest.Interface
}

// newLicenseTokens returns a LicenseTokens
func newLicenseTokens(c *ManagementV1Client) *licenseTokens {
	return &licenseTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the licenseToken, and returns the corresponding licenseToken object, and an error if there is any.
func (c *licenseTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LicenseToken, err error) {
	result = &v1.LicenseToken{}
	err = c.client.Get().
		Resource("licensetokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LicenseTokens that match those selectors.
func (c *licenseTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LicenseTokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LicenseTokenList{}
	err = c.client.Get().
		Resource("licensetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested licenseTokens.
func (c *licenseTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("licensetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a licenseToken and creates it.  Returns the server's representation of the licenseToken, and an error, if there is any.
func (c *licenseTokens) Create(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.CreateOptions) (result *v1.LicenseToken, err error) {
	result = &v1.LicenseToken{}
	err = c.client.Post().
		Resource("licensetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(licenseToken).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a licenseToken and updates it. Returns the server's representation of the licenseToken, and an error, if there is any.
func (c *licenseTokens) Update(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.UpdateOptions) (result *v1.LicenseToken, err error) {
	result = &v1.LicenseToken{}
	err = c.client.Put().
		Resource("licensetokens").
		Name(licenseToken.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(licenseToken).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *licenseTokens) UpdateStatus(ctx context.Context, licenseToken *v1.LicenseToken, opts metav1.UpdateOptions) (result *v1.LicenseToken, err error) {
	result = &v1.LicenseToken{}
	err = c.client.Put().
		Resource("licensetokens").
		Name(licenseToken.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(licenseToken).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the licenseToken and deletes it. Returns an error if one occurs.
func (c *licenseTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("licensetokens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *licenseTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("licensetokens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched licenseToken.
func (c *licenseTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LicenseToken, err error) {
	result = &v1.LicenseToken{}
	err = c.client.Patch(pt).
		Resource("licensetokens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
