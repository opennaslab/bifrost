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

package main

import (
	"k8s.io/klog"

	"opennaslab.io/bifrost/cmd/options"
	"opennaslab.io/bifrost/pkg/database"
	"opennaslab.io/bifrost/pkg/server"
)

//	@title			bifrost API
//	@version		0.1
//	@description	Take you to the land of light, the city of freedom(A unified external service management system for NAS).
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/api/v1
func main() {
	opt := options.NewBifrostDBOptions()
	if err := opt.Validate(); err != nil {
		klog.Errorf("validate bifrost db options failed:%v", err)
		return
	}
	if err := database.DatabaseConnectionInit(opt); err != nil {
		klog.Errorf("init bifrost db failed:%v", err)
	}
	router := server.NewServerRouter()
	if err := router.Run(":8080"); err != nil {
		klog.Errorf("run server failed:%v", err)
	}
}
