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

// ListStepsHandler  godoc
//
//	@Summary		List all bifrost steps
//	@Description	List all supported bifrost steps
//	@Tags			Steps
//	@Produce		json
//	@Success		200			{object}	customapi.StepInfoList
//	@Failure		500			{object}	string
//	@Router			/steps [get]
func ListStepsHandler(ctx *gin.Context) {
	steps := customapi.ListStepDefinitions()
	respData, err := json.Marshal(steps)
	if err != nil {
		klog.Errorf("Marshal steps failed:%v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func GetStepHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	step := customapi.GetStepDefinition(name)
	if step == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	respData, err := json.Marshal(step)
	if err != nil {
		klog.Errorf("Marshal steps failed:%v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func CreateOrUpdateWorkflowHandler(ctx *gin.Context) {
	workflow := &api.ConfigurationWorkflow{}
	err := ctx.BindJSON(workflow)
	if err != nil {
		klog.Errorf("Bind workflow failed:%v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = validateSteps(workflow.ConfigurationSteps)
	if err != nil {
		klog.Errorf("Validate steps failed:%v", err)
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func validateSteps(steps []api.ConfigurationStep) error {
	for _, step := range steps {
		typedStep, err := customapi.GetTypedConfig(&step)
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
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	respData, err := json.Marshal(workflows)
	if err != nil {
		klog.Errorf("Marshal workflows failed:%v", err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func GetWorkflowHandler(ctx *gin.Context) {
	db := database.GetWorkflowMode()
	workflow, err := db.GetWorkflow(ctx.Param("name"))
	if err != nil {
		klog.Errorf("Get workflow %s failed:%v", ctx.Param("name"), err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	respData, err := json.Marshal(workflow)
	if err != nil {
		klog.Errorf("Marshal workflow %s failed:%v", ctx.Param("name"), err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", respData)
}

func DeleteWorkflowHandler(ctx *gin.Context) {
	db := database.GetWorkflowMode()
	wf, err := db.GetWorkflow(ctx.Param("name"))
	if err != nil {
		klog.Errorf("Delete workflow %s failed:%v", ctx.Param("name"), err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	wf.Status.State = api.ConfigurationWorkflowStateDeleting
	if err := db.UpdateWorkflow(wf); err != nil {
		klog.Errorf("Update workflow %s failed:%v", ctx.Param("name"), err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func RunWorkflowHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func StopWorkflowHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func RestartWorkflowHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}

func LogWorkflowHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}
