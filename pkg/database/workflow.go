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

package database

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"

	"opennaslab.io/bifrost/pkg/api"
)

type Workflow struct {
	Id       string `gorm:"column:id;type:integer;autoIncrement;primary_key"`
	Name     string `gorm:"column:name;type:text"`
	Manifest string `gorm:"column:manifest;type:text"`
}

func (w Workflow) TableName() string {
	return "workflow"
}

var workflowMode *WorkflowMode

type WorkflowMode struct {
	dbConn *gorm.DB
}

func InitWorkflow(dbConnection *gorm.DB) error {
	if err := dbConnection.AutoMigrate(&Workflow{}); err != nil {
		return err
	}
	workflowMode = &WorkflowMode{dbConn: dbConnection}
	return nil
}

func GetWorkflowMode() *WorkflowMode {
	return workflowMode
}

func (w *WorkflowMode) AddWorkflow(workflow *api.ConfigurationWorkflow) error {
	manifest, err := json.Marshal(&workflow)
	if err != nil {
		return err
	}
	entry := Workflow{
		Name:     workflow.Name,
		Manifest: string(manifest),
	}
	result := w.dbConn.Where("name = ?", entry.Name).Updates(&entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *WorkflowMode) DeleteWorkflow(name string) error {
	result := w.dbConn.Where("name = ?", name).Delete(&Workflow{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}

func (w *WorkflowMode) UpdateWorkflow(workflow *api.ConfigurationWorkflow) error {
	manifest, err := json.Marshal(&workflow)
	if err != nil {
		return err
	}
	entry := Workflow{
		Name:     workflow.Name,
		Manifest: string(manifest),
	}
	result := w.dbConn.Where("name = ?", entry.Name).Updates(&entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *WorkflowMode) GetWorkflow(name string) (*api.ConfigurationWorkflow, error) {
	var entry Workflow
	result := w.dbConn.Where("name = ?", name).First(&entry)
	if result.Error != nil {
		return nil, result.Error
	}
	manifest := []byte(entry.Manifest)
	var workflow api.ConfigurationWorkflow
	err := json.Unmarshal(manifest, &workflow)
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}

func (w *WorkflowMode) ListWorkflows() ([]*api.ConfigurationWorkflow, error) {
	var entries []Workflow
	result := w.dbConn.Find(&entries)
	if result.Error != nil {
		return nil, result.Error
	}
	var workflows []*api.ConfigurationWorkflow
	for _, entry := range entries {
		manifest := []byte(entry.Manifest)
		var workflow api.ConfigurationWorkflow
		err := json.Unmarshal(manifest, &workflow)
		if err != nil {
			return nil, err
		}
		workflows = append(workflows, &workflow)
	}
	return workflows, nil
}
