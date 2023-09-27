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
	"k8s.io/klog"
)

const (
	RegistryGitRepo = "https://github.com/opennaslab/bifrost-registry"
)

var (
	RegistryCacheDir = os.Getenv("HOME") + "/birefrost-registry"
)

func CloneRegistry(refresh bool, homeDir, stepDir string) ([]string, error) {
	klog.Infof("1")
	if refresh {
		os.RemoveAll(homeDir)
	}
	klog.Infof("2")
	_, err := os.Stat(homeDir)
	if err != nil {
		klog.Infof("5")
		_, err := git.PlainClone(homeDir, false, &git.CloneOptions{
			URL:             RegistryGitRepo,
			SingleBranch:    true,
			InsecureSkipTLS: true,
		})
		if err != nil {
			klog.Infof("3")
			return nil, err
		}
	}

	var files []string
	err = filepath.Walk(stepDir, func(path string, info os.FileInfo, err error) error {
		if path == stepDir {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		klog.Infof("4")
		return nil, err
	}
	return files, nil
}
