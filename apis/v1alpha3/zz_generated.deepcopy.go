//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadResourceMapping) DeepCopyInto(out *ClusterWorkloadResourceMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadResourceMapping.
func (in *ClusterWorkloadResourceMapping) DeepCopy() *ClusterWorkloadResourceMapping {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadResourceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkloadResourceMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadResourceMappingContainer) DeepCopyInto(out *ClusterWorkloadResourceMappingContainer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadResourceMappingContainer.
func (in *ClusterWorkloadResourceMappingContainer) DeepCopy() *ClusterWorkloadResourceMappingContainer {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadResourceMappingContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadResourceMappingList) DeepCopyInto(out *ClusterWorkloadResourceMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterWorkloadResourceMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadResourceMappingList.
func (in *ClusterWorkloadResourceMappingList) DeepCopy() *ClusterWorkloadResourceMappingList {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadResourceMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterWorkloadResourceMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadResourceMappingSpec) DeepCopyInto(out *ClusterWorkloadResourceMappingSpec) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]ClusterWorkloadResourceMappingTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadResourceMappingSpec.
func (in *ClusterWorkloadResourceMappingSpec) DeepCopy() *ClusterWorkloadResourceMappingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadResourceMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadResourceMappingTemplate) DeepCopyInto(out *ClusterWorkloadResourceMappingTemplate) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ClusterWorkloadResourceMappingContainer, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadResourceMappingTemplate.
func (in *ClusterWorkloadResourceMappingTemplate) DeepCopy() *ClusterWorkloadResourceMappingTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadResourceMappingTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvMapping) DeepCopyInto(out *EnvMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvMapping.
func (in *EnvMapping) DeepCopy() *EnvMapping {
	if in == nil {
		return nil
	}
	out := new(EnvMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBinding) DeepCopyInto(out *ServiceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBinding.
func (in *ServiceBinding) DeepCopy() *ServiceBinding {
	if in == nil {
		return nil
	}
	out := new(ServiceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingList) DeepCopyInto(out *ServiceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingList.
func (in *ServiceBindingList) DeepCopy() *ServiceBindingList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSecretReference) DeepCopyInto(out *ServiceBindingSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSecretReference.
func (in *ServiceBindingSecretReference) DeepCopy() *ServiceBindingSecretReference {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingServiceReference) DeepCopyInto(out *ServiceBindingServiceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingServiceReference.
func (in *ServiceBindingServiceReference) DeepCopy() *ServiceBindingServiceReference {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingServiceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSpec) DeepCopyInto(out *ServiceBindingSpec) {
	*out = *in
	in.Workload.DeepCopyInto(&out.Workload)
	out.Service = in.Service
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSpec.
func (in *ServiceBindingSpec) DeepCopy() *ServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingStatus) DeepCopyInto(out *ServiceBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(ServiceBindingSecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingStatus.
func (in *ServiceBindingStatus) DeepCopy() *ServiceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingWorkloadReference) DeepCopyInto(out *ServiceBindingWorkloadReference) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingWorkloadReference.
func (in *ServiceBindingWorkloadReference) DeepCopy() *ServiceBindingWorkloadReference {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingWorkloadReference)
	in.DeepCopyInto(out)
	return out
}
