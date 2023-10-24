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

package options

import (
	"fmt"
	"os"
)

type Option func(*Config) error

type Config struct {
	DB BifrostDBOptions
}

type BifrostDBOptions struct {
	DBDriver   string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBSchema   string
}

func WithDBOptions() Option {

	return func(config *Config) error {
		DBDriver := os.Getenv("DB_DRIVER")
		DBUser := os.Getenv("DB_USER")
		DBPassword := os.Getenv("DB_PASSWORD")
		DBHost := os.Getenv("DB_HOST")
		DBPort := os.Getenv("DB_PORT")
		DBSchema := os.Getenv("DB_SCHEMA")

		// set default value
		if DBDriver == "" {
			DBDriver = "sqlite"
		}

		if DBDriver == "mysql" {

			if DBUser == "" {
				return fmt.Errorf("DB_USER is required")
			}
			if DBPassword == "" {
				return fmt.Errorf("DB_PASSWORD is required")
			}
			if DBHost == "" {
				return fmt.Errorf("DB_HOST is required")
			}
			if DBPort == "" {
				return fmt.Errorf("DB_PORT is required")
			}
			if DBSchema == "" {
				return fmt.Errorf("DB_SCHEMA is required")
			}
		} else if DBDriver == "sqlite" {
		} else {
			return fmt.Errorf("DB_DRIVER is not supported")
		}

		BifrostDBOptions := BifrostDBOptions{
			DBDriver:   DBDriver,
			DBUser:     DBUser,
			DBPassword: DBPassword,
			DBHost:     DBHost,
			DBPort:     DBPort,
			DBSchema:   DBSchema,
		}
		config.DB = BifrostDBOptions
		return nil
	}
}
