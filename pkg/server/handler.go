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

	"opennaslab.io/bifrost/pkg/registry"
)

func ListLocalStepsHandler(ctx *gin.Context) {
	refreshRegistry := ctx.Query("refresh") == "true"
	if refreshRegistry {
		ctx.JSON(http.StatusOK, []byte("{}"))
	}
	steps, err := registry.ListAllLocalSteps(refreshRegistry)
	if err != nil {
		klog.Errorf("List local steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	respData, err := json.Marshal(steps)
	if err != nil {
		klog.Errorf("Marshal local steps failed:%v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	klog.Infof("jw1:%s", respData)
	ctx.JSON(http.StatusOK, respData)
}

func GetLocalStepHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte("OK"))
}
