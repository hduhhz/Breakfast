// +build !ignore_autogenerated

/*

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

package v1alpha2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bread) DeepCopyInto(out *Bread) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bread.
func (in *Bread) DeepCopy() *Bread {
	if in == nil {
		return nil
	}
	out := new(Bread)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bread) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreadList) DeepCopyInto(out *BreadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bread, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreadList.
func (in *BreadList) DeepCopy() *BreadList {
	if in == nil {
		return nil
	}
	out := new(BreadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BreadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreadSpec) DeepCopyInto(out *BreadSpec) {
	*out = *in
	out.Scv = in.Scv
	out.Framework = in.Framework
	out.Task = in.Task
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreadSpec.
func (in *BreadSpec) DeepCopy() *BreadSpec {
	if in == nil {
		return nil
	}
	out := new(BreadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BreadStatus) DeepCopyInto(out *BreadStatus) {
	*out = *in
	if in.ContainerStatuses != nil {
		in, out := &in.ContainerStatuses, &out.ContainerStatuses
		*out = make([]v1.ContainerStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BreadStatus.
func (in *BreadStatus) DeepCopy() *BreadStatus {
	if in == nil {
		return nil
	}
	out := new(BreadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameworkSpec) DeepCopyInto(out *FrameworkSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameworkSpec.
func (in *FrameworkSpec) DeepCopy() *FrameworkSpec {
	if in == nil {
		return nil
	}
	out := new(FrameworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCVSpec) DeepCopyInto(out *SCVSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCVSpec.
func (in *SCVSpec) DeepCopy() *SCVSpec {
	if in == nil {
		return nil
	}
	out := new(SCVSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}