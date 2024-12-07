// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"istio.io/istio/123/pkg/config"
	"istio.io/istio/123/pkg/config/schema/resource"
	"istio.io/istio/123/pkg/kube/controllers"
)

// PilotConfigToInstance convert from config.Config, which has no associated proto, to MCP Resource proto.
func PilotConfigToInstance(c *config.Config, schema resource.Schema) *Instance {
	return &Instance{
		Metadata: Metadata{
			Schema:      schema,
			FullName:    FullName{Namespace(c.Namespace), LocalName(c.Name)},
			CreateTime:  c.CreationTimestamp,
			Version:     Version(c.ResourceVersion),
			Labels:      c.Labels,
			Annotations: c.Annotations,
		},
		Message: c.Spec,
	}
}

// ObjectToInstance convert from a controller object to MCP Resource proto.
// Note you need to pass the object and its spec
func ObjectToInstance(c controllers.Object, spec config.Spec, schema resource.Schema) *Instance {
	return &Instance{
		Metadata: Metadata{
			Schema:      schema,
			FullName:    FullName{Namespace(c.GetNamespace()), LocalName(c.GetName())},
			CreateTime:  c.GetCreationTimestamp().Time,
			Version:     Version(c.GetResourceVersion()),
			Labels:      c.GetLabels(),
			Annotations: c.GetAnnotations(),
		},
		Message: spec,
	}
}
