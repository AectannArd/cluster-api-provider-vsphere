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

package feature

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	// Every capv-specific feature gate should add method here following this template:
	//
	// // owner: @username
	// // alpha: v1.X
	// MyFeature featuregate.Feature = "MyFeature".

	// NodeAntiAffinity is a feature gate for the NodeAntiAffinity functionality.
	//
	// alpha: v1.4
	NodeAntiAffinity featuregate.Feature = "NodeAntiAffinity"

	// NodeLabeling is a feature gate for the functionality to propagate Machine labels
	// with the prefix to the Node objects.
	// This is a stop-gap measure which will be removed when we have this functionality
	// present in CAPI.
	// See https://github.com/kubernetes-sigs/cluster-api/pull/6255
	//
	// alpha: v1.4
	NodeLabeling featuregate.Feature = "NodeLabeling"
)

func init() {
	runtime.Must(MutableGates.Add(defaultCAPVFeatureGates))
}

// defaultCAPVFeatureGates consists of all known capv-specific feature keys.
// To add a new feature, define a key for it above and add it here.
var defaultCAPVFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	// Every feature should be initiated here:
	NodeAntiAffinity: {Default: false, PreRelease: featuregate.Alpha},
	NodeLabeling:     {Default: false, PreRelease: featuregate.Alpha},
}
