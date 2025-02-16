/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var _ = logf.Log.WithName("awsclusterroleidentity-resource")

func (r *AWSClusterRoleIdentity) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha4-awsclusterroleidentity,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=awsclusterroleidentities,versions=v1alpha4,name=validation.awsclusterroleidentity.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1alpha4-awsclusterroleidentity,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=awsclusterroleidentities,versions=v1alpha4,name=default.awsclusterroleidentity.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var (
	_ webhook.Validator = &AWSClusterRoleIdentity{}
	_ webhook.Defaulter = &AWSClusterRoleIdentity{}
)

// ValidateCreate will do any extra validation when creating an AWSClusterRoleIdentity.
func (r *AWSClusterRoleIdentity) ValidateCreate() error {
	if r.Spec.SourceIdentityRef == nil {
		return field.Invalid(field.NewPath("spec", "sourceIdentityRef"),
			r.Spec.SourceIdentityRef, "field cannot be set to nil")
	}

	// Validate selector parses as Selector
	if r.Spec.AllowedNamespaces != nil {
		_, err := metav1.LabelSelectorAsSelector(&r.Spec.AllowedNamespaces.Selector)
		if err != nil {
			return field.Invalid(field.NewPath("spec", "allowedNamespaces", "selector"), r.Spec.AllowedNamespaces.Selector, err.Error())
		}
	}

	return nil
}

// ValidateDelete allows you to add any extra validation when deleting an AWSClusterRoleIdentity.
func (r *AWSClusterRoleIdentity) ValidateDelete() error {
	return nil
}

// ValidateUpdate will do any extra validation when updating an AWSClusterRoleIdentity.
func (r *AWSClusterRoleIdentity) ValidateUpdate(old runtime.Object) error {
	oldP, ok := old.(*AWSClusterRoleIdentity)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected an AWSClusterRoleIdentity but got a %T", old))
	}

	// If a SourceIdentityRef is set, do not allow removal of it.
	if oldP.Spec.SourceIdentityRef != nil && r.Spec.SourceIdentityRef == nil {
		return field.Invalid(field.NewPath("spec", "sourceIdentityRef"),
			r.Spec.SourceIdentityRef, "field cannot be set to nil")
	}

	// Validate selector parses as Selector
	if r.Spec.AllowedNamespaces != nil {
		_, err := metav1.LabelSelectorAsSelector(&r.Spec.AllowedNamespaces.Selector)
		if err != nil {
			return field.Invalid(field.NewPath("spec", "allowedNamespaces", "selector"), r.Spec.AllowedNamespaces.Selector, err.Error())
		}
	}

	return nil
}

// Default will set default values for the AWSClusterRoleIdentity.
func (r *AWSClusterRoleIdentity) Default() {
}
