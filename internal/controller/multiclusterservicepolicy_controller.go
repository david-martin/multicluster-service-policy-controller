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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	configpolicycontrollerapiv1 "github.com/david-martin/multicluster-service-policy-controller/api/config-policy-controller/api/v1"
	examplecomv1alpha1 "github.com/david-martin/multicluster-service-policy-controller/api/v1alpha1"
)

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
	log := log.FromContext(ctx)

	previous := &examplecomv1alpha1.MultiClusterServicePolicy{}
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: req.Name}, previous)
	if err != nil {
		if err := client.IgnoreNotFound(err); err != nil {
			log.Error(err, "Unable to fetch MultiClusterServicePolicy")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	if previous.GetDeletionTimestamp() != nil && !previous.GetDeletionTimestamp().IsZero() {
		log.Info("MultiClusterServicePolicy is deleting")
		return ctrl.Result{}, nil
	}

	configPolicy := &configpolicycontrollerapiv1.ConfigurationPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      previous.Name,
			Namespace: previous.Namespace,
		},
		Spec: &configpolicycontrollerapiv1.ConfigurationPolicySpec{},
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MultiClusterServicePolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1alpha1.MultiClusterServicePolicy{}).
		Complete(r)
}
