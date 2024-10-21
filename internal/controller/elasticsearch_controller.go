/*
Copyright 2024.

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

package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ragv1alpha1 "github.com/redhat-et/rag/api/v1alpha1"
)

// ElasticsearchReconciler reconciles a Elasticsearch object
type ElasticsearchReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=rag.opendatahub.io,resources=elasticsearches,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rag.opendatahub.io,resources=elasticsearches/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=rag.opendatahub.io,resources=elasticsearches/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Elasticsearch object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *ElasticsearchReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// check if deployment exists if not create it
	foundDeployment := &appsv1.Deployment{}
	if err := r.Get(ctx, req.NamespacedName, foundDeployment); err != nil {
		// Define a new deployment
		createDeployment := common.createDeployment(req.Name, req.Namespace, "elasticsearch:7.10.0", 9200)
		if err := r.Create(ctx, createDeployment); err != nil {
			logger.Error(err, "Failed to create new Deployment", "Deployment.Namespace", createDeployment.Namespace, "Deployment.Name", createDeployment.Name)
			return ctrl.Result{}, err
		}
	}
	// check if service exists if not create it
	foundService := &corev1.Service{}
	if err := r.Get(ctx, req.NamespacedName, foundService); err != nil {
		// Define a new service
		createService := common.createService(req.Name, req.Namespace, 9200)
		if err := r.Create(ctx, createService); err != nil {
			logger.Error(err, "Failed to create new Service", "Service.Namespace", createService.Namespace, "Service.Name", createService.Name)
			return ctrl.Result{}, err
		}
	}
	// check if pvc exists if not create it
	foundPvc := &corev1.PersistentVolumeClaim{}
	if err := r.Get(ctx, req.NamespacedName, foundPvc); err != nil {
		// Define a new pvc
		createPvc := common.createPvc(req.Name, req.Namespace)
		if err := r.Create(ctx, createPvc); err != nil {
			logger.Error(err, "Failed to create new PVC", "PVC.Namespace", createPvc.Namespace, "PVC.Name", createPvc.Name)
			return ctrl.Result{}, err
		}
	}
	// check if secret exists if not create it
	foundSecret := &corev1.Secret{}
	if err := r.Get(ctx, req.NamespacedName, foundSecret); err != nil {
		// Define a new secret
		createSecret := common.createSecret(req.Name, req.Namespace)
		if err := r.Create(ctx, createSecret); err != nil {
			logger.Error(err, "Failed to create new Secret", "Secret.Namespace", createSecret.Namespace, "Secret.Name", createSecret.Name)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ElasticsearchReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ragv1alpha1.Elasticsearch{}).
		Complete(r)
}
