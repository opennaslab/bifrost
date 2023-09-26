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
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"gopkg.in/yaml.v3"

	"opennaslab.io/bifrost/pkg/api"
)

const (
	RegistryGitDir = "local_step"
)

func ListAllLocalConfigActions(refresh bool) ([]api.ConfigStepDefinition, error) {
	cloneDir := os.Getenv("HOME") + "/" + RegistryGitDir
	if refresh {
		os.RemoveAll(cloneDir)
	}
	_, err := os.Stat(cloneDir)
	if err != nil {
		_, err := git.PlainClone(cloneDir, false, &git.CloneOptions{
			URL:             RegistryGitRepo,
			SingleBranch:    true,
			InsecureSkipTLS: true,
		})
		if err != nil {
			return nil, err
		}
	}

	var files []string
	err = filepath.Walk(cloneDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	list := []api.ConfigStepDefinition{}
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		def := api.ConfigStepDefinition{}
		yaml.Unmarshal(data, &def)
		list = append(list, def)
	}

	return list, nil
}
