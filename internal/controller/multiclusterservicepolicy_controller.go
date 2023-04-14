/*
Copyright 2023.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	// cpc "open-cluster-management.io/config-policy-controller/api/v1"

	examplecomv1alpha1 "github.com/david-martin/multicluster-service-policy-controller/api/v1alpha1"
)

// cpc "github.com/open-cluster-management-io/config-policy-controller/api/v1"
//
// MultiClusterServicePolicyReconciler reconciles a MultiClusterServicePolicy object
type MultiClusterServicePolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=example.com.example.com,resources=multiclusterservicepolicies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=example.com.example.com,resources=multiclusterservicepolicies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=example.com.example.com,resources=multiclusterservicepolicies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MultiClusterServicePolicy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *MultiClusterServicePolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	// var configPolicy cpc.ConfigurationPolicy

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MultiClusterServicePolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1alpha1.MultiClusterServicePolicy{}).
		Complete(r)
}
