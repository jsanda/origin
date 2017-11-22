package v1

import (
	v1 "github.com/openshift/api/route/v1"
	"github.com/openshift/origin/pkg/route/generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type RouteV1Interface interface {
	RESTClient() rest.Interface
	RoutesGetter
}

// RouteV1Client is used to interact with features provided by the route.openshift.io group.
type RouteV1Client struct {
	restClient rest.Interface
}

func (c *RouteV1Client) Routes(namespace string) RouteResourceInterface {
	return newRoutes(c, namespace)
}

// NewForConfig creates a new RouteV1Client for the given config.
func NewForConfig(c *rest.Config) (*RouteV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &RouteV1Client{client}, nil
}

// NewForConfigOrDie creates a new RouteV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *RouteV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new RouteV1Client for the given RESTClient.
func New(c rest.Interface) *RouteV1Client {
	return &RouteV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *RouteV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
