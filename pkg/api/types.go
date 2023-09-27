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

type LocalConfigDefinition ConfigStepDefinition
type RemoteConfigDefinition ConfigStepDefinition
type DNSConfigDefinition ConfigStepDefinition

type ConfigStepDefinition struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	Parameters  StepParameter `json:"parameters"`
}

type StepParameter struct {
	In  []Parameter `json:"in"`
	Out []Parameter `json:"out"`
}

type Parameter struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Required    bool        `json:"required"`
	Items       []Parameter `json:"items"`
}

type ConfigurationWorkflow struct {
	Name                     string                      `json:"name"`
	Description              string                      `json:"description"`
	LocalConfigurationSteps  []ConfigurationStep         `json:"localConfigurationSteps"`
	RemoteConfigurationSteps []ConfigurationStep         `json:"remoteConfigurationSteps"`
	DNSConfigurationSteps    []ConfigurationStep         `json:"dnsConfigurationSteps"`
	Status                   ConfigurationWorkflowStatus `json:"status"`
}

type ConfigurationStep struct {
	Name string        `json:"name"`
	Use  string        `json:"use"`
	In   []interface{} `json:"in"`
}

type ConfigurationWorkflowState string

const (
	ConfigurationWorkflowStatePending ConfigurationWorkflowState = "pending"
	ConfigurationWorkflowStateRunning ConfigurationWorkflowState = "running"
	ConfigurationWorkflowStateSuccess ConfigurationWorkflowState = "success"
	ConfigurationWorkflowStateFailed  ConfigurationWorkflowState = "failed"
)

type ConfigurationWorkflowStatus struct {
	State   ConfigurationWorkflowState `json:"state"`
	Message string                     `json:"message"`
	Steps   []ConfigurationStepStatus  `json:"steps"`
}

type ConfigurationStepStatus struct {
	Name    string                     `json:"name"`
	State   ConfigurationWorkflowState `json:"state"`
	Message string                     `json:"message"`
}
