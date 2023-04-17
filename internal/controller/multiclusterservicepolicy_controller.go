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
	"encoding/json"
	"fmt"

	governancepolicypropagatorapiv1 "github.com/david-martin/multicluster-service-policy-controller/api/governance-policy-propagator/api/v1"
	skupperapiv1alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ocmclusterv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
	controllerruntime "sigs.k8s.io/controller-runtime"
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

	skupperClusterPolicy := &skupperapiv1alpha1.SkupperClusterPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       "SkupperClusterPolicy",
			APIVersion: fmt.Sprintf("%s/%s", skupperapiv1alpha1.SchemeGroupVersion.Group, skupperapiv1alpha1.SchemeGroupVersion.Version),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: previous.Name,
		},
		Spec: skupperapiv1alpha1.SkupperClusterPolicySpec{
			Namespaces:              previous.Spec.Namespaces,
			AllowedExposedResources: previous.Spec.AllowedExposedResources,
			AllowedServices:         previous.Spec.AllowedServices,
		},
	}

	skupperClusterPolicyBytes, err := json.Marshal(skupperClusterPolicy)
	if err != nil {
		return ctrl.Result{}, err
	}

	configPolicy := &configpolicycontrollerapiv1.ConfigurationPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigurationPolicy",
			APIVersion: fmt.Sprintf("%s/%s", configpolicycontrollerapiv1.GroupVersion.Group, configpolicycontrollerapiv1.GroupVersion.Version),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: previous.Name,
		},
		Spec: &configpolicycontrollerapiv1.ConfigurationPolicySpec{
			RemediationAction: "enforce",
			Severity:          "low",
			ObjectTemplates: []*configpolicycontrollerapiv1.ObjectTemplate{
				{
					ComplianceType: configpolicycontrollerapiv1.MustHave,
					ObjectDefinition: runtime.RawExtension{
						Raw: skupperClusterPolicyBytes,
					},
				},
			},
		},
	}

	// Marshal to JSON bytes
	configPolicyBytes, err := json.Marshal(configPolicy)
	if err != nil {
		return ctrl.Result{}, err
	}

	policyTemplates := []*governancepolicypropagatorapiv1.PolicyTemplate{
		{
			ObjectDefinition: runtime.RawExtension{
				Raw: configPolicyBytes,
			},
		},
	}

	policy := &governancepolicypropagatorapiv1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      previous.Name,
			Namespace: previous.Namespace,
		},
	}
	_, err = controllerruntime.CreateOrUpdate(ctx, r.Client, policy, func() error {
		policy.Annotations = map[string]string{
			"policy.open-cluster-management.io/standards":  "Example Standard",
			"policy.open-cluster-management.io/categories": "Example Category",
			"policy.open-cluster-management.io/controls":   "Example Control",
		}

		policy.Spec = governancepolicypropagatorapiv1.PolicySpec{
			RemediationAction: "enforce",
			Disabled:          false,
			PolicyTemplates:   policyTemplates,
		}

		return nil
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	placement := &ocmclusterv1beta1.Placement{
		ObjectMeta: metav1.ObjectMeta{
			Name:      previous.Name,
			Namespace: previous.Namespace,
		},
	}
	_, err = controllerruntime.CreateOrUpdate(ctx, r.Client, placement, func() error {
		placement.Spec = ocmclusterv1beta1.PlacementSpec{
			Predicates: []ocmclusterv1beta1.ClusterPredicate{
				{
					RequiredClusterSelector: ocmclusterv1beta1.ClusterSelector{
						LabelSelector: metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{},
						},
					},
				},
			},
		}

		return nil
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	placementBinding := &governancepolicypropagatorapiv1.PlacementBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      previous.Name,
			Namespace: previous.Namespace,
		},
	}
	_, err = controllerruntime.CreateOrUpdate(ctx, r.Client, placementBinding, func() error {
		placementBinding.PlacementRef = governancepolicypropagatorapiv1.PlacementSubject{
			Name:     placement.Name,
			Kind:     "Placement",
			APIGroup: ocmclusterv1beta1.GroupVersion.Group,
		}
		placementBinding.Subjects = []governancepolicypropagatorapiv1.Subject{
			{
				Name:     policy.Name,
				Kind:     "Policy",
				APIGroup: governancepolicypropagatorapiv1.GroupVersion.Group,
			},
		}

		return nil
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MultiClusterServicePolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&examplecomv1alpha1.MultiClusterServicePolicy{}).
		Complete(r)
}
