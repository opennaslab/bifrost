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

package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"

	"opennaslab.io/bifrost/pkg/api"
	"opennaslab.io/bifrost/pkg/customapi"
	"opennaslab.io/bifrost/pkg/database"
)

func ListLocalStepsHandler(ctx *gin.Context) {
	steps := customapi.ListLocalStepDefinitions()
	respData, err := json.Marshal(steps)
	if err != nil {
		klog.Errorf("Marshal local steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func ListRemoteStepsHandler(ctx *gin.Context) {
	steps := customapi.ListRemoteStepDefinitions()
	respData, err := json.Marshal(steps)
	if err != nil {
		klog.Errorf("Marshal remote steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func GetLocalStepHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	step := customapi.GetLocalStepDefinition(name)
	if step == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	respData, err := json.Marshal(step)
	if err != nil {
		klog.Errorf("Marshal local steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func GetRemoteStepHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	step := customapi.GetRemoteStepDefinition(name)
	if step == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	respData, err := json.Marshal(step)
	if err != nil {
		klog.Errorf("Marshal remote steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func CreateOrUpdateWorkflowHandler(ctx *gin.Context) {
	workflow := &api.ConfigurationWorkflow{}
	ctx.BindJSON(workflow)
	err := validateSteps(api.LocalStepType, workflow.LocalConfigurationSteps)
	if err != nil {
		klog.Errorf("Validate local steps failed:%v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = validateSteps(api.RemoteStepType, workflow.RemoteConfigurationSteps)
	if err != nil {
		klog.Errorf("Validate remote steps failed:%v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = validateSteps(api.DNSStepType, workflow.DNSConfigurationSteps)
	if err != nil {
		klog.Errorf("Validate dns steps failed:%v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func validateSteps(stepType string, steps []api.ConfigurationStep) error {
	for _, step := range steps {
		typedStep, err := customapi.GetTypedConfig(stepType, &step)
		if err != nil {
			klog.Errorf("Get local step %s failed:%v", step.Use, err)
			return err
		}
		if err := typedStep.Validate(); err != nil {
			klog.Errorf("Validate local step %s failed:%v", step.Name, err)
			return err
		}
	}

	return nil
}

func ListWorkflowsHandler(ctx *gin.Context) {
	db := database.GetWorkflowMode()
	workflows, err := db.ListWorkflows()
	if err != nil {
		klog.Errorf("List workflows failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	respData, err := json.Marshal(workflows)
	if err != nil {
		klog.Errorf("Marshal workflows failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func GetWorkflowHandler(ctx *gin.Context) {
	db := database.GetWorkflowMode()
	workflow, err := db.GetWorkflow(ctx.Param("name"))
	if err != nil {
		klog.Errorf("Get workflow %s failed:%v", ctx.Param("name"), err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	respData, err := json.Marshal(workflow)
	if err != nil {
		klog.Errorf("Marshal workflow %s failed:%v", ctx.Param("name"), err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func DeleteWorkflowHandler(ctx *gin.Context) {
	db := database.GetWorkflowMode()
	wf, err := db.GetWorkflow(ctx.Param("name"))
	if err != nil {
		klog.Errorf("Delete workflow %s failed:%v", ctx.Param("name"), err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	wf.Status.State = api.ConfigurationWorkflowStateDeleting
	if err := db.UpdateWorkflow(wf); err != nil {
		klog.Errorf("Update workflow %s failed:%v", ctx.Param("name"), err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}
