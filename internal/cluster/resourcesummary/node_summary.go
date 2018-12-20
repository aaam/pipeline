// Copyright © 2018 Banzai Cloud
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

package resourcesummary

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// NodeSummary describes a node's resource summary with CPU and Memory capacity/request/limit/allocatable.
type NodeSummary struct {
	Summary

	Status string
}

// GetNodeSummary returns NodeSummary type with the given data.
func GetNodeSummary(node *v1.Node, requests, limits map[v1.ResourceName]resource.Quantity) NodeSummary {
	capacity := map[v1.ResourceName]resource.Quantity{
		v1.ResourceCPU:    {},
		v1.ResourceMemory: {},
	}

	allocatable := map[v1.ResourceName]resource.Quantity{
		v1.ResourceCPU:    {},
		v1.ResourceMemory: {},
	}

	if cpu := node.Status.Capacity.Cpu(); cpu != nil {
		capacity[v1.ResourceCPU] = *cpu
	}

	if cpu := node.Status.Allocatable.Cpu(); cpu != nil {
		allocatable[v1.ResourceCPU] = *cpu
	}

	if mem := node.Status.Capacity.Memory(); mem != nil {
		capacity[v1.ResourceMemory] = *mem
	}

	if mem := node.Status.Allocatable.Memory(); mem != nil {
		allocatable[v1.ResourceMemory] = *mem
	}

	summary := NodeSummary{
		Summary: GetSummary(capacity, allocatable, requests, limits),
		Status:  GetNodeStatus(node),
	}

	return summary
}