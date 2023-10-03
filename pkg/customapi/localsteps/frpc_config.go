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

import (
	"fmt"
)

type FrpcParameterIn struct {
	ServerAddr string       `json:"serverAddr" description:"the frps server address" required:"true"`
	ServerPort int          `json:"serverPort" description:"the frps server port" required:"true"`
	Service    []FrpService `json:"service" description:"the service list" required:"true"`
}

type FrpService struct {
	ServiceName string `json:"serviceName" description:"the service name" required:"true"`
	LocalIP     string `json:"localIP" description:"the local ip address" required:"true"`
	LocalPort   int    `json:"localPort" description:"the local port" required:"true"`
	RemotePort  string `json:"remoteIP" description:"the remote port" required:"true"`
}

func (f FrpcParameterIn) Validate() error {
	if f.ServerAddr == "" {
		return fmt.Errorf("serverAddr is required")
	}
	if f.ServerPort == 0 {
		return fmt.Errorf("serverPort is required")
	}
	if f.Service == nil || len(f.Service) == 0 {
		return fmt.Errorf("service is required")
	}
	return nil
}

func (f FrpcParameterIn) GetExecutionConfig() ([]byte, error) {
	return nil, nil
}
