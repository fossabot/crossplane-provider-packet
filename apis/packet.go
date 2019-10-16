/*
Copyright 2019 The Crossplane Authors.

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

// Generate deepcopy methodsets
//go:generate ${CONTROLLERGEN} object:headerFile=../hack/boilerplate.go.txt paths=./...

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate ${CROSSPLANETOOLS_ANGRYJET} generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

// Package api contains Kubernetes API groups for Packet cloud provider.
package api

import (
	"k8s.io/apimachinery/pkg/runtime"

	serverv1alpha1 "github.com/hasheddan/stack-packet/apis/server/v1alpha1"
	packetv1alpha1 "github.com/hasheddan/stack-packet/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		packetv1alpha1.SchemeBuilder.AddToScheme,
		serverv1alpha1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}