/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
)

// UncountedTerminatedPodsApplyConfiguration represents a declarative configuration of the UncountedTerminatedPods type for use
// with apply.
type UncountedTerminatedPodsApplyConfiguration struct {
	Succeeded []types.UID `json:"succeeded,omitempty"`
	Failed    []types.UID `json:"failed,omitempty"`
}

// UncountedTerminatedPodsApplyConfiguration constructs a declarative configuration of the UncountedTerminatedPods type for use with
// apply.
func UncountedTerminatedPods() *UncountedTerminatedPodsApplyConfiguration {
	return &UncountedTerminatedPodsApplyConfiguration{}
}

// WithSucceeded adds the given value to the Succeeded field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Succeeded field.
func (b *UncountedTerminatedPodsApplyConfiguration) WithSucceeded(values ...types.UID) *UncountedTerminatedPodsApplyConfiguration {
	for i := range values {
		b.Succeeded = append(b.Succeeded, values[i])
	}
	return b
}

// WithFailed adds the given value to the Failed field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Failed field.
func (b *UncountedTerminatedPodsApplyConfiguration) WithFailed(values ...types.UID) *UncountedTerminatedPodsApplyConfiguration {
	for i := range values {
		b.Failed = append(b.Failed, values[i])
	}
	return b
}
