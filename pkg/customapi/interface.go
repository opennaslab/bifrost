package customapi

import (
	"encoding/json"
	"fmt"
	"reflect"

	"opennaslab.io/bifrost/pkg/api"
	localsteps "opennaslab.io/bifrost/pkg/customapi/localsteps"
	remotesteps "opennaslab.io/bifrost/pkg/customapi/remotesteps"
)

var LocalStepsInfoMap = map[string]StepsInfo{
	"frpc-config": {
		Name:        "frpc-config",
		Image:       "opennaslab/frpc-config:latest",
		Description: "install/config frpc in local host",
	},
}

var RemoteStepsInfoMap = map[string]StepsInfo{
	"docker-config": {
		Name:        "docker-config",
		Image:       "opennaslan/docker-config:latest",
		Description: "install/config docker in remote host",
	},
}

var DNSStepsInfoMap = map[string]StepsInfo{}

type TypedInterface interface {
	Validate() error
}

var LocalStepsStruct = map[string]TypedInterface{
	"frpc-config": localsteps.FrpcParameterIn{Service: []localsteps.FrpService{{}}},
}

var RemoteStepsStruct = map[string]TypedInterface{
	"docker-config": remotesteps.DockerConfigParameterIn{},
}

var DNSStepsStruct = map[string]TypedInterface{}

func GetTypedConfig(stepType string, step *api.ConfigurationStep) (TypedInterface, error) {
	if stepType == api.LocalStepType {
		if _, ok := LocalStepsStruct[step.Use]; !ok {
			return nil, fmt.Errorf("not found")
		}
		// TODO: there is currency problem
		ret := LocalStepsStruct[step.Use]
		if err := json.Unmarshal([]byte(step.In), ret); err != nil {
			return nil, err
		}
		return ret, nil
	}

	return nil, fmt.Errorf("not found")
}

type Documentation struct {
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Required    bool            `json:"required"`
	Items       []Documentation `json:"items,omitempty"`
}

func GenerateDocumentation(obj interface{}) []Documentation {
	doc := []Documentation{}

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		docField := Documentation{
			Name:        field.Tag.Get("json"),
			Type:        field.Type.Name(),
			Description: field.Tag.Get("description"),
			Required:    field.Tag.Get("required") == "true",
		}

		if field.Type.Kind() == reflect.Slice {
			docField.Type = "array"
			docField.Items = append(docField.Items, GenerateDocumentation(value.Index(0).Interface())...)
		}

		doc = append(doc, docField)
	}

	return doc
}

type StepsInfo struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	Parameters  StepParameter `json:"parameters"`
}

type StepParameter struct {
	In []Documentation `json:"in"`
}

func GetLocalStepDefinition(name string) *StepsInfo {
	if _, ok := LocalStepsInfoMap[name]; !ok {
		return nil
	}
	paraInDoc := GenerateDocumentation(LocalStepsStruct[name])
	ret := LocalStepsInfoMap[name]
	ret.Parameters.In = paraInDoc
	return &ret
}

func GetRemoteStepDefinition(name string) *StepsInfo {
	if _, ok := RemoteStepsInfoMap[name]; !ok {
		return nil
	}
	paraInDoc := GenerateDocumentation(RemoteStepsStruct[name])
	ret := RemoteStepsInfoMap[name]
	ret.Parameters.In = paraInDoc
	return &ret
}

func GetDNSStepDefinition(name string) *StepsInfo {
	if _, ok := DNSStepsInfoMap[name]; !ok {
		return nil
	}
	ret := DNSStepsInfoMap[name]
	return &ret
}

func ListLocalStepDefinitions() []StepsInfo {
	ret := []StepsInfo{}
	for name := range LocalStepsInfoMap {
		paraInDoc := GenerateDocumentation(LocalStepsStruct[name])
		ele := LocalStepsInfoMap[name]
		ele.Parameters.In = paraInDoc
		ret = append(ret, ele)
	}
	return ret
}

func ListRemoteStepDefinitions() []StepsInfo {
	ret := []StepsInfo{}
	for name := range RemoteStepsInfoMap {
		paraInDoc := GenerateDocumentation(RemoteStepsStruct[name])
		ele := RemoteStepsInfoMap[name]
		ele.Parameters.In = paraInDoc
		ret = append(ret, ele)
	}
	return ret
}
