/*
Copyright 2023 The opennaslab Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicabl e law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

type Config struct {
	Kind        string            `json:"kind"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Data        map[string]string `json:"data"`
}

type Action struct {
	Kind        string            `json:"kind"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	Image       string            `json:"image"`
}

type WorkflowType string

const (
	DispatchWorklow  WorkflowType = "dispatch"
	ScheduleWorkflow WorkflowType = "schedule"
)

type Workflow struct {
	Kind        string       `json:"kind"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	On          WorkflowType `json:"on"`
	// +optional
	Schedule string                  `json:"schedule,omitempty"`
	Actions  []WorkflowActionElement `json:"actions"`
	Status   WorkflowStatus          `json:"status,omitempty"`
}

type WorkflowActionElement struct {
	Use  string            `json:"use"`
	With map[string]string `json:"with"`
}

type WorkflowStatus struct {
	NextExecutionTime string `json:"nextExecutionTime"`
	Finished          bool   `json:"finished"`
	StepIndex         int32  `json:"stepIndex"`
	Message           string `json:"message"`
}
