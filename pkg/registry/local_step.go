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

package registry

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"

	"opennaslab.io/bifrost/pkg/api"
)

const (
	LocalStepDir = "local-step"
)

func ListAllLocalSteps(refresh bool) ([]*api.LocalConfigDefinition, error) {
	stepDir := RegistryCacheDir + "/" + LocalStepDir
	files, err := CloneRegistry(refresh, RegistryCacheDir, stepDir)
	if err != nil {
		return nil, err
	}

	list := []*api.LocalConfigDefinition{}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		def := &api.LocalConfigDefinition{}
		if err := yaml.Unmarshal(data, def); err != nil {
			return nil, err
		}
		list = append(list, def)
	}

	return list, nil
}

func GetLocalStep(name string) (*api.LocalConfigDefinition, error) {
	stepDir := RegistryCacheDir + "/" + LocalStepDir
	files, err := CloneRegistry(false, RegistryCacheDir, stepDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if strings.Split(path.Base(file), ".")[0] == name {
			data, err := os.ReadFile(file)
			if err != nil {
				return nil, err
			}
			def := &api.LocalConfigDefinition{}
			if err := yaml.Unmarshal(data, def); err != nil {
				return nil, err
			}
			return def, nil
		}
	}
	return nil, fmt.Errorf("step %s not found", name)
}
