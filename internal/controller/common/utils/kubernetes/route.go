package common

import (
	route "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createRoute(name string, namespace string) *route.Route {
	r := &route.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: route.RouteSpec{
			To: route.RouteTargetReference{
				Kind: "Service",
				Name: "my-service",
			},
		},
	}
	return r
}
