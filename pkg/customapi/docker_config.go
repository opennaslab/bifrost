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

package customapi

import "fmt"

type DockerConfigParameterIn struct {
	ServerAddr     string `json:"serverAddr" description:"the vps address" required:"true"`
	ServerPort     int    `json:"serverPort" description:"the vps port" required:"true"`
	ServerUser     string `json:"serverUser" description:"the vps user" required:"true"`
	ServerPassword string `json:"serverPassword" description:"the vps password" required:"true"`
}

func (d DockerConfigParameterIn) Validate() error {
	if d.ServerAddr == "" {
		return fmt.Errorf("serverAddr is required")
	}
	if d.ServerPort == 0 {
		return fmt.Errorf("serverPort is required")
	}
	if d.ServerUser == "" {
		return fmt.Errorf("serverUser is required")
	}
	if d.ServerPassword == "" {
		return fmt.Errorf("serverPassword is required")
	}
	return nil
}
