package utils

import "github.com/ghodss/yaml"

// LoadYaml - Parses the YAML structure returns map[string]interface{}
func LoadYaml(yml []byte) map[string]interface{} {
	var ymlStruct interface{}
	// Convert YAML into client-go k8s data structure
	err := yaml.Unmarshal(yml, &ymlStruct)
	CheckError(err)
	if ymlStruct != nil {
		return ymlStruct.(map[string]interface{})
	}
	return nil
}
