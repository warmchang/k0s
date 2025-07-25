// SPDX-FileCopyrightText: k0s authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/k0sproject/k0s/pkg/client/clientset/typed/k0s/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK0sV1beta1 struct {
	*testing.Fake
}

func (c *FakeK0sV1beta1) ClusterConfigs(namespace string) v1beta1.ClusterConfigInterface {
	return newFakeClusterConfigs(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeK0sV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
