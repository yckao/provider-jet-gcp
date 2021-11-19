// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
func (in *ArtifactsObservation) DeepCopyInto(out *ArtifactsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactsObservation.
func (in *ArtifactsObservation) DeepCopy() *ArtifactsObservation {
	if in == nil {
		return nil
	}
	out := new(ArtifactsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactsParameters) DeepCopyInto(out *ArtifactsParameters) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactsParameters.
func (in *ArtifactsParameters) DeepCopy() *ArtifactsParameters {
	if in == nil {
		return nil
	}
	out := new(ArtifactsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildObservation) DeepCopyInto(out *BuildObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildObservation.
func (in *BuildObservation) DeepCopy() *BuildObservation {
	if in == nil {
		return nil
	}
	out := new(BuildObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildParameters) DeepCopyInto(out *BuildParameters) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]ArtifactsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LogsBucket != nil {
		in, out := &in.LogsBucket, &out.LogsBucket
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueueTTL != nil {
		in, out := &in.QueueTTL, &out.QueueTTL
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = make([]SecretParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = make([]StepParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildParameters.
func (in *BuildParameters) DeepCopy() *BuildParameters {
	if in == nil {
		return nil
	}
	out := new(BuildParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubObservation) DeepCopyInto(out *GithubObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubObservation.
func (in *GithubObservation) DeepCopy() *GithubObservation {
	if in == nil {
		return nil
	}
	out := new(GithubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubParameters) DeepCopyInto(out *GithubParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = make([]PullRequestParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = make([]PushParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubParameters.
func (in *GithubParameters) DeepCopy() *GithubParameters {
	if in == nil {
		return nil
	}
	out := new(GithubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectsObservation) DeepCopyInto(out *ObjectsObservation) {
	*out = *in
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = make([]TimingObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectsObservation.
func (in *ObjectsObservation) DeepCopy() *ObjectsObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectsParameters) DeepCopyInto(out *ObjectsParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectsParameters.
func (in *ObjectsParameters) DeepCopy() *ObjectsParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsObservation) DeepCopyInto(out *OptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsObservation.
func (in *OptionsObservation) DeepCopy() *OptionsObservation {
	if in == nil {
		return nil
	}
	out := new(OptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsParameters) DeepCopyInto(out *OptionsParameters) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.DynamicSubstitutions != nil {
		in, out := &in.DynamicSubstitutions, &out.DynamicSubstitutions
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LogStreamingOption != nil {
		in, out := &in.LogStreamingOption, &out.LogStreamingOption
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.RequestedVerifyOption != nil {
		in, out := &in.RequestedVerifyOption, &out.RequestedVerifyOption
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SourceProvenanceHash != nil {
		in, out := &in.SourceProvenanceHash, &out.SourceProvenanceHash
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubstitutionOption != nil {
		in, out := &in.SubstitutionOption, &out.SubstitutionOption
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerPool != nil {
		in, out := &in.WorkerPool, &out.WorkerPool
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsParameters.
func (in *OptionsParameters) DeepCopy() *OptionsParameters {
	if in == nil {
		return nil
	}
	out := new(OptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubsubConfigObservation) DeepCopyInto(out *PubsubConfigObservation) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubConfigObservation.
func (in *PubsubConfigObservation) DeepCopy() *PubsubConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PubsubConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubsubConfigParameters) DeepCopyInto(out *PubsubConfigParameters) {
	*out = *in
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubsubConfigParameters.
func (in *PubsubConfigParameters) DeepCopy() *PubsubConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PubsubConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestObservation) DeepCopyInto(out *PullRequestObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestObservation.
func (in *PullRequestObservation) DeepCopy() *PullRequestObservation {
	if in == nil {
		return nil
	}
	out := new(PullRequestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestParameters) DeepCopyInto(out *PullRequestParameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.CommentControl != nil {
		in, out := &in.CommentControl, &out.CommentControl
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestParameters.
func (in *PullRequestParameters) DeepCopy() *PullRequestParameters {
	if in == nil {
		return nil
	}
	out := new(PullRequestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushObservation) DeepCopyInto(out *PushObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushObservation.
func (in *PushObservation) DeepCopy() *PushObservation {
	if in == nil {
		return nil
	}
	out := new(PushObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushParameters) DeepCopyInto(out *PushParameters) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushParameters.
func (in *PushParameters) DeepCopy() *PushParameters {
	if in == nil {
		return nil
	}
	out := new(PushParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoSourceObservation) DeepCopyInto(out *RepoSourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoSourceObservation.
func (in *RepoSourceObservation) DeepCopy() *RepoSourceObservation {
	if in == nil {
		return nil
	}
	out := new(RepoSourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoSourceParameters) DeepCopyInto(out *RepoSourceParameters) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RepoName != nil {
		in, out := &in.RepoName, &out.RepoName
		*out = new(string)
		**out = **in
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoSourceParameters.
func (in *RepoSourceParameters) DeepCopy() *RepoSourceParameters {
	if in == nil {
		return nil
	}
	out := new(RepoSourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObservation) DeepCopyInto(out *SecretObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObservation.
func (in *SecretObservation) DeepCopy() *SecretObservation {
	if in == nil {
		return nil
	}
	out := new(SecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretParameters) DeepCopyInto(out *SecretParameters) {
	*out = *in
	if in.KmsKeyName != nil {
		in, out := &in.KmsKeyName, &out.KmsKeyName
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretParameters.
func (in *SecretParameters) DeepCopy() *SecretParameters {
	if in == nil {
		return nil
	}
	out := new(SecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObservation) DeepCopyInto(out *SourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObservation.
func (in *SourceObservation) DeepCopy() *SourceObservation {
	if in == nil {
		return nil
	}
	out := new(SourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceParameters) DeepCopyInto(out *SourceParameters) {
	*out = *in
	if in.RepoSource != nil {
		in, out := &in.RepoSource, &out.RepoSource
		*out = make([]RepoSourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageSource != nil {
		in, out := &in.StorageSource, &out.StorageSource
		*out = make([]StorageSourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceParameters.
func (in *SourceParameters) DeepCopy() *SourceParameters {
	if in == nil {
		return nil
	}
	out := new(SourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepObservation) DeepCopyInto(out *StepObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepObservation.
func (in *StepObservation) DeepCopy() *StepObservation {
	if in == nil {
		return nil
	}
	out := new(StepObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepParameters) DeepCopyInto(out *StepParameters) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]StepVolumesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WaitFor != nil {
		in, out := &in.WaitFor, &out.WaitFor
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepParameters.
func (in *StepParameters) DeepCopy() *StepParameters {
	if in == nil {
		return nil
	}
	out := new(StepParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepVolumesObservation) DeepCopyInto(out *StepVolumesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepVolumesObservation.
func (in *StepVolumesObservation) DeepCopy() *StepVolumesObservation {
	if in == nil {
		return nil
	}
	out := new(StepVolumesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepVolumesParameters) DeepCopyInto(out *StepVolumesParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepVolumesParameters.
func (in *StepVolumesParameters) DeepCopy() *StepVolumesParameters {
	if in == nil {
		return nil
	}
	out := new(StepVolumesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSourceObservation) DeepCopyInto(out *StorageSourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSourceObservation.
func (in *StorageSourceObservation) DeepCopy() *StorageSourceObservation {
	if in == nil {
		return nil
	}
	out := new(StorageSourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSourceParameters) DeepCopyInto(out *StorageSourceParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Generation != nil {
		in, out := &in.Generation, &out.Generation
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSourceParameters.
func (in *StorageSourceParameters) DeepCopy() *StorageSourceParameters {
	if in == nil {
		return nil
	}
	out := new(StorageSourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimingObservation) DeepCopyInto(out *TimingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimingObservation.
func (in *TimingObservation) DeepCopy() *TimingObservation {
	if in == nil {
		return nil
	}
	out := new(TimingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimingParameters) DeepCopyInto(out *TimingParameters) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimingParameters.
func (in *TimingParameters) DeepCopy() *TimingParameters {
	if in == nil {
		return nil
	}
	out := new(TimingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerObservation) DeepCopyInto(out *TriggerObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.TriggerID != nil {
		in, out := &in.TriggerID, &out.TriggerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerObservation.
func (in *TriggerObservation) DeepCopy() *TriggerObservation {
	if in == nil {
		return nil
	}
	out := new(TriggerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerParameters) DeepCopyInto(out *TriggerParameters) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = make([]BuildParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = make([]GithubParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IgnoredFiles != nil {
		in, out := &in.IgnoredFiles, &out.IgnoredFiles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludedFiles != nil {
		in, out := &in.IncludedFiles, &out.IncludedFiles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.PubsubConfig != nil {
		in, out := &in.PubsubConfig, &out.PubsubConfig
		*out = make([]PubsubConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TriggerTemplate != nil {
		in, out := &in.TriggerTemplate, &out.TriggerTemplate
		*out = make([]TriggerTemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebhookConfig != nil {
		in, out := &in.WebhookConfig, &out.WebhookConfig
		*out = make([]WebhookConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerParameters.
func (in *TriggerParameters) DeepCopy() *TriggerParameters {
	if in == nil {
		return nil
	}
	out := new(TriggerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTemplateObservation) DeepCopyInto(out *TriggerTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTemplateObservation.
func (in *TriggerTemplateObservation) DeepCopy() *TriggerTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(TriggerTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTemplateParameters) DeepCopyInto(out *TriggerTemplateParameters) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RepoName != nil {
		in, out := &in.RepoName, &out.RepoName
		*out = new(string)
		**out = **in
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTemplateParameters.
func (in *TriggerTemplateParameters) DeepCopy() *TriggerTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(TriggerTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumesObservation) DeepCopyInto(out *VolumesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumesObservation.
func (in *VolumesObservation) DeepCopy() *VolumesObservation {
	if in == nil {
		return nil
	}
	out := new(VolumesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumesParameters) DeepCopyInto(out *VolumesParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumesParameters.
func (in *VolumesParameters) DeepCopy() *VolumesParameters {
	if in == nil {
		return nil
	}
	out := new(VolumesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigObservation) DeepCopyInto(out *WebhookConfigObservation) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigObservation.
func (in *WebhookConfigObservation) DeepCopy() *WebhookConfigObservation {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigParameters) DeepCopyInto(out *WebhookConfigParameters) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigParameters.
func (in *WebhookConfigParameters) DeepCopy() *WebhookConfigParameters {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigParameters)
	in.DeepCopyInto(out)
	return out
}
