package customapi

import (
	"encoding/json"
	"fmt"
	"reflect"

	"opennaslab.io/bifrost/pkg/api"
)

var StepsInfoMap = map[string]StepsInfo{
	"frpc-config": {
		Name:        "frpc-config",
		Image:       "opennaslab/frpc-config:latest",
		Description: "install/config frpc in local host",
	},
}

var StepsDocStruct = map[string]interface{}{
	"frpc-config": FrpcParameterIn{Service: []FrpService{{}}},
}

var TypedInterfaceMap = map[string]TypedInterface{
	"frpc-config": &FrpcParameterIn{},
}

type TypedInterface interface {
	Validate() error
	GetExecutionConfig() ([]byte, error)
}

func GetTypedConfig(step *api.ConfigurationStep) (TypedInterface, error) {
	if _, ok := TypedInterfaceMap[step.Use]; !ok {
		return nil, fmt.Errorf("not found")
	}
	// TODO: there is currency problem
	ret := TypedInterfaceMap[step.Use]
	if err := json.Unmarshal([]byte(step.In), ret); err != nil {
		return nil, err
	}
	return ret, nil
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
	if _, ok := StepsInfoMap[name]; !ok {
		return nil
	}
	paraInDoc := GenerateDocumentation(StepsDocStruct[name])
	ret := StepsInfoMap[name]
	ret.Parameters.In = paraInDoc
	return &ret
}

func ListLocalStepDefinitions() []StepsInfo {
	ret := []StepsInfo{}
	for name := range StepsInfoMap {
		paraInDoc := GenerateDocumentation(StepsDocStruct[name])
		ele := StepsInfoMap[name]
		ele.Parameters.In = paraInDoc
		ret = append(ret, ele)
	}
	return ret
}
