// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKey) DeepCopyInto(out *APIKey) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(APIKeySource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKey.
func (in *APIKey) DeepCopy() *APIKey {
	if in == nil {
		return nil
	}
	out := new(APIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeySource) DeepCopyInto(out *APIKeySource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeySource.
func (in *APIKeySource) DeepCopy() *APIKeySource {
	if in == nil {
		return nil
	}
	out := new(APIKeySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProject) DeepCopyInto(out *MongoDBAtlasProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProject.
func (in *MongoDBAtlasProject) DeepCopy() *MongoDBAtlasProject {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectList) DeepCopyInto(out *MongoDBAtlasProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBAtlasProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectList.
func (in *MongoDBAtlasProjectList) DeepCopy() *MongoDBAtlasProjectList {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectSpec) DeepCopyInto(out *MongoDBAtlasProjectSpec) {
	*out = *in
	in.OrgID.DeepCopyInto(&out.OrgID)
	in.APIKey.DeepCopyInto(&out.APIKey)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectSpec.
func (in *MongoDBAtlasProjectSpec) DeepCopy() *MongoDBAtlasProjectSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasProjectStatus) DeepCopyInto(out *MongoDBAtlasProjectStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasProjectStatus.
func (in *MongoDBAtlasProjectStatus) DeepCopy() *MongoDBAtlasProjectStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrgID) DeepCopyInto(out *OrgID) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(OrgIDSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrgID.
func (in *OrgID) DeepCopy() *OrgID {
	if in == nil {
		return nil
	}
	out := new(OrgID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrgIDSource) DeepCopyInto(out *OrgIDSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrgIDSource.
func (in *OrgIDSource) DeepCopy() *OrgIDSource {
	if in == nil {
		return nil
	}
	out := new(OrgIDSource)
	in.DeepCopyInto(out)
	return out
}
