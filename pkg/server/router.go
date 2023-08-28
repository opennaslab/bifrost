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
)

func NewServerRouter() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	corsHandler := cors.New(config)

	initConfigRouter(router, corsHandler)
	initActionRouter(router, corsHandler)
	initWorkflowRouter(router, corsHandler)

	return router
}

func initConfigRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	configGroup := router.Group("/api/v1/configs")
	// List all configs
	configGroup.GET("", ListConfigHandler)
	// Get specific config
	configGroup.GET("/:config_name", nil)
	// Create or update a new config
	configGroup.POST("/:config_name", nil)
}

func initActionRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	actionGroup := router.Group("/api/v1/actions")
	// List all actions
	actionGroup.GET("", nil)
	// Get specific action
	actionGroup.GET("/:action_name", nil)
	// Create or update a new action
}

func initWorkflowRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	workflowGroup := router.Group("/api/v1/workflows")
	// List all workflows
	workflowGroup.GET("", nil)
	// Get specific workflow
	workflowGroup.GET("/:workflow_name", nil)
	// Create or update a new workflow
	workflowGroup.POST("/:workflow_name", nil)
}
