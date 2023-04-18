//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServicePolicy) DeepCopyInto(out *MultiClusterServicePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServicePolicy.
func (in *MultiClusterServicePolicy) DeepCopy() *MultiClusterServicePolicy {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServicePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterServicePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServicePolicyList) DeepCopyInto(out *MultiClusterServicePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterServicePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServicePolicyList.
func (in *MultiClusterServicePolicyList) DeepCopy() *MultiClusterServicePolicyList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServicePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterServicePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServicePolicySpec) DeepCopyInto(out *MultiClusterServicePolicySpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedExposedResources != nil {
		in, out := &in.AllowedExposedResources, &out.AllowedExposedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedServices != nil {
		in, out := &in.AllowedServices, &out.AllowedServices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.PlacementRef = in.PlacementRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServicePolicySpec.
func (in *MultiClusterServicePolicySpec) DeepCopy() *MultiClusterServicePolicySpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServicePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterServicePolicyStatus) DeepCopyInto(out *MultiClusterServicePolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterServicePolicyStatus.
func (in *MultiClusterServicePolicyStatus) DeepCopy() *MultiClusterServicePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterServicePolicyStatus)
	in.DeepCopyInto(out)
	return out
}
