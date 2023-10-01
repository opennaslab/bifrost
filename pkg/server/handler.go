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

func CreateOrUpdateWorkflowHandler(ctx *gin.Context) {
	workflow := &api.ConfigurationWorkflow{}
	ctx.BindJSON(workflow)
	for _, step := range workflow.LocalConfigurationSteps {
		typedStep, err := customapi.GetTypedConfig(&step)
		if err != nil {
			klog.Errorf("Get local step %s failed:%v", step.Use, err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if err := typedStep.Validate(); err != nil {
			klog.Errorf("Validate local step %s failed:%v", step.Name, err)
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}
