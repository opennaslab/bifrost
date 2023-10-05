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

package controller

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"

	"opennaslab.io/bifrost/pkg/api"
	"opennaslab.io/bifrost/pkg/container"
	"opennaslab.io/bifrost/pkg/customapi"
	"opennaslab.io/bifrost/pkg/database"
)

type WorkflowQueue struct {
	Workflow map[string]interface{}
	mutex    sync.Mutex
}

var WFQueue *WorkflowQueue

func InitWorkflowQueue() {
	WFQueue = &WorkflowQueue{
		Workflow: make(map[string]interface{}),
		mutex:    sync.Mutex{},
	}
}

func (w *WorkflowQueue) AddWorkflow(name string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.Workflow[name] = nil
}

func (w *WorkflowQueue) Run() {
	for {
		w.mutex.Lock()
		defer w.mutex.Unlock()
		for name, _ := range w.Workflow {
			requeue, err := w.Reconcile(name)
			if err == nil && !requeue {
				delete(w.Workflow, name)
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func (w *WorkflowQueue) Reconcile(name string) (requeue bool, err error) {
	db := database.GetWorkflowMode()
	wf, err := db.GetWorkflow(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if wf.Status.State == api.ConfigurationWorkflowStateDeleting {
		requeue, err := w.DeleteWorkflow(wf)
		if !requeue && err != nil {
			err := db.DeleteWorkflow(name)
			if err != nil {
				return false, err
			}
		}
	}
	return false, nil
}

func (w *WorkflowQueue) DeleteWorkflow(wf *api.ConfigurationWorkflow) (requeue bool, err error) {
	for index, step := range wf.DNSConfigurationSteps {
		stepState := wf.Status.DNSConfigurationSteps[index]

		if stepState.State == api.ConfigurationWorkflowStateRunning || stepState.State == api.ConfigurationWorkflowStateRunningSuccess {
			err := container.DeleteContainer(stepState.ContainerId)
			if err != nil {
				return false, err
			}
		}

		if stepState.State == api.ConfigurationWorkflowStateDeletingSuccess {
			continue
		}

		if stepState.State == api.ConfigurationWorkflowStateDeleting {
			containerState, exitCode, err := container.GetContainer(stepState.ContainerId)
			if err != nil {
				return false, err
			}
			if containerState == "exited" && exitCode == 0 {
				stepState.State = api.ConfigurationWorkflowStateDeletingSuccess
				continue
			}
		}

		stepDef := customapi.GetDNSStepDefinition(step.Use)
		if stepDef == nil {
			return false, fmt.Errorf("dns step definition %s not found", step.Use)
		}
		id, err := container.CreateContainer(wf.Name, step.Name, stepDef.Image)
		if err != nil {
			return false, err
		}
		stepState.ContainerId = id
		stepState.State = api.ConfigurationWorkflowStateDeleting
	}

	return false, nil
}

func (w *WorkflowQueue) UpdateWorkflow(name string) (*api.ConfigurationWorkflow, error) {
	return nil, nil
}
