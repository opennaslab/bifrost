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

	return router
}

func initConfigRouter(router *gin.Engine, corsHandler gin.HandlerFunc) {
	configGroup := router.Group("/api/v1/localsteps")
	configGroup.GET("", ListLocalStepsHandler)
	// Get specific step
	configGroup.GET("/:name", GetLocalStepHandler)
}
