/*
Copyright 2022 The Kubernetes Authors.
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

package removeduplicates

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_RemoveDuplicatesArgs
// TODO: the final default values would be discussed in community
func SetDefaults_RemoveDuplicatesArgs(obj *RemoveDuplicatesArgs) {
	if obj.Namespaces == nil {
		obj.Namespaces = nil
	}
	if obj.ExcludeOwnerKinds == nil {
		obj.ExcludeOwnerKinds = nil
	}
}
