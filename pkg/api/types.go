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

type ConfigurationStepDefinition struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Image       string                  `json:"image"`
	Parameters  StepParameterDefinition `json:"parametersDefinition"`
}

type StepParameterDefinition struct {
	In  []ParameterDefinition `json:"in,omitempty"`
	Out []ParameterDefinition `json:"out,omitempty"`
}

const (
	ParameterTypeString  = "string"
	ParameterTypeInteger = "integer"
	ParameterTypeBoolean = "boolean"
	ParameterTypeArray   = "array"
)

type ParameterDefinition struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Type        string                `json:"type"`
	Required    bool                  `json:"required"`
	Items       []ParameterDefinition `json:"items,omitempty"`
}

type ConfigurationWorkflow struct {
	Name               string                      `json:"name"`
	Description        string                      `json:"description"`
	ConfigurationSteps []ConfigurationStep         `json:"configurationSteps"`
	Status             ConfigurationWorkflowStatus `json:"status,omitempty"`
}

type ConfigurationStep struct {
	Name string `json:"name"`
	Use  string `json:"use"`
	In   string `json:"in"`
}

type ConfigurationWorkflowState string

const (
	ConfigurationWorkflowStatePending         ConfigurationWorkflowState = "pending"
	ConfigurationWorkflowStateRunning         ConfigurationWorkflowState = "running"
	ConfigurationWorkflowStateRunningSuccess  ConfigurationWorkflowState = "runningSuccess"
	ConfigurationWorkflowStateRunningFailed   ConfigurationWorkflowState = "runningFailed"
	ConfigurationWorkflowStateDeleting        ConfigurationWorkflowState = "deleting"
	ConfigurationWorkflowStateDeletingFailed  ConfigurationWorkflowState = "deletingFailed"
	ConfigurationWorkflowStateDeletingSuccess ConfigurationWorkflowState = "Deletingsuccess"
)

type ConfigurationWorkflowStatus struct {
	State              ConfigurationWorkflowState `json:"state"`
	Message            string                     `json:"message"`
	ConfigurationSteps []ConfigurationStepStatus  `json:"configurationSteps"`
}

type ConfigurationStepStatus struct {
	Name        string                     `json:"name"`
	ContainerId string                     `json:"containerId"`
	State       ConfigurationWorkflowState `json:"state"`
	Message     string                     `json:"message"`
}
