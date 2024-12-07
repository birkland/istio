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

package xds

import (
	v3 "istio.io/istio/123/pilot/pkg/xds/v3"
	"istio.io/istio/123/pkg/util/sets"
)

// EventType represents the type of object we are tracking, mapping to envoy TypeUrl.
type EventType = string

var AllTrackingEventTypes = sets.New[EventType](
	v3.ClusterType,
	v3.ListenerType,
	v3.RouteType,
	v3.EndpointType,
)
