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

	"opennaslab.io/bifrost/cmd/options"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func DatabaseConnectionInit(config *options.Config) error {
	dbConfig := config.DB
	// extract different db driver
	if dbConfig.DBDriver == "sqlite" {
		db, err := gorm.Open(sqlite.Open("bifrost.db"), &gorm.Config{})
		if err != nil {
			return err
		}
		dbConnection = db
	} else if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBSchema)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		dbConnection = db
	}
	// init table
	if err := InitWorkflow(dbConnection); err != nil {
		return err
	}
	return nil
}
