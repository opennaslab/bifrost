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
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"opennaslab.io/bifrost/cmd/options"
)

var dbConnection *gorm.DB

func DatabaseConnectionInit(opts *options.BifrostDBOptions) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opts.DBUser, opts.DBPassword, opts.DBHost, opts.DBPort, opts.DBSchema)

	var err error
	dbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := InitWorkflow(dbConnection); err != nil {
		return err
	}
	return nil
}
