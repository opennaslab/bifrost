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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "opennaslab.io/bifrost/api"
)

func NewServerRouter() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	corsHandler := cors.New(config)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(corsHandler)

	initStepRouter(router, corsHandler)
	initWorkflowRouter(router, corsHandler)

	return router
}

func initStepRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	localStepGroup := router.Group("/api/v1/steps")
	localStepGroup.GET("", ListStepsHandler)
	// Get specific step
	localStepGroup.GET("/:step_name", GetStepHandler)
}

func initWorkflowRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	workflowGroup := router.Group("/api/v1/workflows")
	// List all workflow
	workflowGroup.GET("", ListWorkflowsHandler)
	// Create/Update workflow
	workflowGroup.POST("", CreateOrUpdateWorkflowHandler)
	// Get specific workflow
	workflowGroup.GET("/:workflow_name", GetWorkflowHandler)
	// Delete workflow
	workflowGroup.DELETE("/:workflow_name", DeleteWorkflowHandler)
	// Run workflow
	workflowGroup.PUT("/:workflow_name/run", nil)
	// Stop workflow
	workflowGroup.PUT("/:workflow_name/stop", nil)
	// Restart workflow
	workflowGroup.PUT("/:workflow_name/restart", nil)
	// logs workflow's step
	workflowGroup.GET("/:workflow_name/:step_name/logs", nil)
}
