// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	in.Definition.DeepCopyInto(&out.Definition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bandwidth) DeepCopyInto(out *Bandwidth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bandwidth.
func (in *Bandwidth) DeepCopy() *Bandwidth {
	if in == nil {
		return nil
	}
	out := new(Bandwidth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conditions) DeepCopyInto(out *Conditions) {
	*out = *in
	out.HeaderCondition = in.HeaderCondition
	out.IPCondition = in.IPCondition
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in *Conditions) DeepCopy() *Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
	*out = *in
	out.ConfigurationStatusFields = in.ConfigurationStatusFields
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatusFields) DeepCopyInto(out *ConfigurationStatusFields) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatusFields.
func (in *ConfigurationStatusFields) DeepCopy() *ConfigurationStatusFields {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatusFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Definition) DeepCopyInto(out *Definition) {
	*out = *in
	if in.SwaggerConfigmapNames != nil {
		in, out := &in.SwaggerConfigmapNames, &out.SwaggerConfigmapNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Interceptors.DeepCopyInto(&out.Interceptors)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Definition.
func (in *Definition) DeepCopy() *Definition {
	if in == nil {
		return nil
	}
	out := new(Definition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deploy) DeepCopyInto(out *Deploy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deploy.
func (in *Deploy) DeepCopy() *Deploy {
	if in == nil {
		return nil
	}
	out := new(Deploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSecurity) DeepCopyInto(out *EndpointSecurity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSecurity.
func (in *EndpointSecurity) DeepCopy() *EndpointSecurity {
	if in == nil {
		return nil
	}
	out := new(EndpointSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderCondition) DeepCopyInto(out *HeaderCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderCondition.
func (in *HeaderCondition) DeepCopy() *HeaderCondition {
	if in == nil {
		return nil
	}
	out := new(HeaderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPCondition) DeepCopyInto(out *IPCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPCondition.
func (in *IPCondition) DeepCopy() *IPCondition {
	if in == nil {
		return nil
	}
	out := new(IPCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interceptors) DeepCopyInto(out *Interceptors) {
	*out = *in
	if in.Ballerina != nil {
		in, out := &in.Ballerina, &out.Ballerina
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Java != nil {
		in, out := &in.Java, &out.Java
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interceptors.
func (in *Interceptors) DeepCopy() *Interceptors {
	if in == nil {
		return nil
	}
	out := new(Interceptors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualType) DeepCopyInto(out *ManualType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualType.
func (in *ManualType) DeepCopy() *ManualType {
	if in == nil {
		return nil
	}
	out := new(ManualType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedType) DeepCopyInto(out *PinnedType) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedType.
func (in *PinnedType) DeepCopy() *PinnedType {
	if in == nil {
		return nil
	}
	out := new(PinnedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimiting) DeepCopyInto(out *RateLimiting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimiting.
func (in *RateLimiting) DeepCopy() *RateLimiting {
	if in == nil {
		return nil
	}
	out := new(RateLimiting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimiting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitingList) DeepCopyInto(out *RateLimitingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimiting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitingList.
func (in *RateLimitingList) DeepCopy() *RateLimitingList {
	if in == nil {
		return nil
	}
	out := new(RateLimitingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitingSpec) DeepCopyInto(out *RateLimitingSpec) {
	*out = *in
	out.RequestCount = in.RequestCount
	out.Bandwidth = in.Bandwidth
	out.Conditions = in.Conditions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitingSpec.
func (in *RateLimitingSpec) DeepCopy() *RateLimitingSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitingStatus) DeepCopyInto(out *RateLimitingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitingStatus.
func (in *RateLimitingStatus) DeepCopy() *RateLimitingStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseType) DeepCopyInto(out *ReleaseType) {
	*out = *in
	if in.Revisions != nil {
		in, out := &in.Revisions, &out.Revisions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseType.
func (in *ReleaseType) DeepCopy() *ReleaseType {
	if in == nil {
		return nil
	}
	out := new(ReleaseType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestCount) DeepCopyInto(out *RequestCount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestCount.
func (in *RequestCount) DeepCopy() *RequestCount {
	if in == nil {
		return nil
	}
	out := new(RequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionSpec) DeepCopyInto(out *RevisionSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionSpec.
func (in *RevisionSpec) DeepCopy() *RevisionSpec {
	if in == nil {
		return nil
	}
	out := new(RevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTemplateSpec) DeepCopyInto(out *RevisionTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTemplateSpec.
func (in *RevisionTemplateSpec) DeepCopy() *RevisionTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(RevisionTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunLatestType) DeepCopyInto(out *RunLatestType) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunLatestType.
func (in *RunLatestType) DeepCopy() *RunLatestType {
	if in == nil {
		return nil
	}
	out := new(RunLatestType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Security) DeepCopyInto(out *Security) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Security.
func (in *Security) DeepCopy() *Security {
	if in == nil {
		return nil
	}
	out := new(Security)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Security) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityList) DeepCopyInto(out *SecurityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Security, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityList.
func (in *SecurityList) DeepCopy() *SecurityList {
	if in == nil {
		return nil
	}
	out := new(SecurityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (in *SecuritySpec) DeepCopy() *SecuritySpec {
	if in == nil {
		return nil
	}
	out := new(SecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityStatus) DeepCopyInto(out *SecurityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityStatus.
func (in *SecurityStatus) DeepCopy() *SecurityStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil {
		return nil
	}
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.DeprecatedRunLatest != nil {
		in, out := &in.DeprecatedRunLatest, &out.DeprecatedRunLatest
		*out = new(RunLatestType)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedPinned != nil {
		in, out := &in.DeprecatedPinned, &out.DeprecatedPinned
		*out = new(PinnedType)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedManual != nil {
		in, out := &in.DeprecatedManual, &out.DeprecatedManual
		*out = new(ManualType)
		**out = **in
	}
	if in.DeprecatedRelease != nil {
		in, out := &in.DeprecatedRelease, &out.DeprecatedRelease
		*out = new(ReleaseType)
		(*in).DeepCopyInto(*out)
	}
	in.ConfigurationSpec.DeepCopyInto(&out.ConfigurationSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	out.ConfigurationStatusFields = in.ConfigurationStatusFields
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetEndpoint) DeepCopyInto(out *TargetEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetEndpoint.
func (in *TargetEndpoint) DeepCopy() *TargetEndpoint {
	if in == nil {
		return nil
	}
	out := new(TargetEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetEndpointList) DeepCopyInto(out *TargetEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TargetEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetEndpointList.
func (in *TargetEndpointList) DeepCopy() *TargetEndpointList {
	if in == nil {
		return nil
	}
	out := new(TargetEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetEndpointSpec) DeepCopyInto(out *TargetEndpointSpec) {
	*out = *in
	out.Deploy = in.Deploy
	out.EndpointSecurity = in.EndpointSecurity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetEndpointSpec.
func (in *TargetEndpointSpec) DeepCopy() *TargetEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(TargetEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetEndpointStatus) DeepCopyInto(out *TargetEndpointStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetEndpointStatus.
func (in *TargetEndpointStatus) DeepCopy() *TargetEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(TargetEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
