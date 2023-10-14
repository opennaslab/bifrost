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
	Workflow map[string]api.ConfigurationWorkflowState
	mutex    sync.Mutex
}

var WFQueue *WorkflowQueue

func InitWorkflowQueue() {
	WFQueue = &WorkflowQueue{
		Workflow: make(map[string]api.ConfigurationWorkflowState),
		mutex:    sync.Mutex{},
	}
}

func (w *WorkflowQueue) Add(name string, op api.ConfigurationWorkflowState) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.Workflow[name] = op
}

func (w *WorkflowQueue) Run() {
	for {
		for name, op := range w.Workflow {
			w.mutex.Lock()
			delete(w.Workflow, name)
			w.mutex.Unlock()

			requeue, err := w.Reconcile(name, op)
			if err == nil && !requeue {
				if _, ok := w.Workflow[name]; ok {
					w.Add(name, op)
				}
			}
		}
		if len(w.Workflow) == 0 {
			time.Sleep(time.Second * 3)
		}
	}
}

func (w *WorkflowQueue) Reconcile(name string, op api.ConfigurationWorkflowState) (requeue bool, err error) {
	db := database.GetWorkflowMode()
	wf, err := db.GetWorkflow(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	defer func() {
		_ = database.GetWorkflowMode().UpdateWorkflow(wf)
	}()
	switch op {
	case api.ConfigurationWorkflowStateDeleting:
		return w.DeleteWorkflow(wf)
	case api.ConfigurationWorkflowStateRunning:
		return false, nil
	case api.ConfigurationWorkflowStateStopping:
		return false, nil
	case api.ConfigurationWorkflowStateRestarting:
		return false, nil
	}

	return false, nil
}

func (w *WorkflowQueue) DeleteWorkflow(wf *api.ConfigurationWorkflow) (requeue bool, err error) {
	for index := len(wf.ConfigurationSteps) - 1; index >= 0; index-- {
		stepState := &(wf.Status.ConfigurationSteps[index])

		switch stepState.State {
		case api.ConfigurationWorkflowStateSubmitted, api.ConfigurationWorkflowStateDeletingSuccess:
			continue
		case api.ConfigurationWorkflowStateRunning, api.ConfigurationWorkflowStateRunningSuccess:
			err := container.DeleteContainer(stepState.ContainerId)
			if err != nil {
				wf.Status.State = api.ConfigurationWorkflowStateDeletingFailed
				return false, err
			}
			stepState.State = api.ConfigurationWorkflowStateDeletingSuccess
			continue
		case api.ConfigurationWorkflowStateDeleting:
			containerState, exitCode, err := container.GetContainer(stepState.ContainerId)
			if err != nil {
				wf.Status.State = api.ConfigurationWorkflowStateDeletingFailed
				return false, err
			}
			if containerState == "exited" && exitCode == 0 {
				stepState.State = api.ConfigurationWorkflowStateDeletingSuccess
				continue
			}
		}
		stepDef := customapi.GetStepDefinition(wf.ConfigurationSteps[index].Use)
		if stepDef == nil {
			return false, fmt.Errorf("dns step definition %s not found", wf.ConfigurationSteps[index].Use)
		}
		id, err := container.CreateContainer(wf.Name, wf.ConfigurationSteps[index].Name, stepDef.Image)
		if err != nil {
			return false, err
		}
		stepState.ContainerId = id
		stepState.State = api.ConfigurationWorkflowStateDeleting
	}
	wf.Status.State = api.ConfigurationWorkflowStateDeletingSuccess

	return false, nil
}

func (w *WorkflowQueue) StopWorkflow(wf *api.ConfigurationWorkflow) (requeue bool, err error) {
	for index := len(wf.ConfigurationSteps) - 1; index >= 0; index-- {
		return false, nil
	}
	return false, nil
}
