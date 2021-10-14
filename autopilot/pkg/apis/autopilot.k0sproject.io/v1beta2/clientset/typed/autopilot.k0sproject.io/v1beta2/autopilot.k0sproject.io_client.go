/*
Copyright 2021 Mirantis
*/
// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	"net/http"

	v1beta2 "github.com/k0sproject/autopilot/pkg/apis/autopilot.k0sproject.io/v1beta2"
	"github.com/k0sproject/autopilot/pkg/apis/autopilot.k0sproject.io/v1beta2/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type AutopilotV1beta2Interface interface {
	RESTClient() rest.Interface
	ControlNodesGetter
	ControlNodeListsGetter
	PlansGetter
	PlanListsGetter
	UpdateConfigsGetter
	UpdateConfigListsGetter
}

// AutopilotV1beta2Client is used to interact with features provided by the autopilot.k0sproject.io group.
type AutopilotV1beta2Client struct {
	restClient rest.Interface
}

func (c *AutopilotV1beta2Client) ControlNodes() ControlNodeInterface {
	return newControlNodes(c)
}

func (c *AutopilotV1beta2Client) ControlNodeLists() ControlNodeListInterface {
	return newControlNodeLists(c)
}

func (c *AutopilotV1beta2Client) Plans() PlanInterface {
	return newPlans(c)
}

func (c *AutopilotV1beta2Client) PlanLists() PlanListInterface {
	return newPlanLists(c)
}

func (c *AutopilotV1beta2Client) UpdateConfigs() UpdateConfigInterface {
	return newUpdateConfigs(c)
}

func (c *AutopilotV1beta2Client) UpdateConfigLists() UpdateConfigListInterface {
	return newUpdateConfigLists(c)
}

// NewForConfig creates a new AutopilotV1beta2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*AutopilotV1beta2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new AutopilotV1beta2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AutopilotV1beta2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &AutopilotV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new AutopilotV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AutopilotV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AutopilotV1beta2Client for the given RESTClient.
func New(c rest.Interface) *AutopilotV1beta2Client {
	return &AutopilotV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AutopilotV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
