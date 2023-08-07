package yaml

import (
	"bytes"
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
)

type YamlUnstructured struct {
	Unstructured *unstructured.Unstructured
	Gvk          *schema.GroupVersionKind
}

func GetMatchLabel(_unstructed *unstructured.Unstructured) (string, error) {
	matchLabels, b, err := unstructured.NestedFieldNoCopy(_unstructed.Object, "spec", "selector", "matchLabels")
	if err != nil {
		return "", err
	}
	var bt bytes.Buffer
	if b {
		matchLabelsMap, ok := matchLabels.(map[string]interface{})
		if ok {
			size := len(matchLabelsMap)
			i := 0
			for k, v := range matchLabelsMap {
				bt.WriteString(fmt.Sprintf("%s=%s", k, v))
				if i < size-1 {
					bt.WriteString(",")
				}
				i++
			}
		} else {
			return "", errors.New("matchLabels类型不对")
		}
	} else {
		return "", errors.New("匹配不到matchLabels")
	}
	return bt.String(), nil
}

func ParseYaml(yamlContent []byte) ([]*YamlUnstructured, error) {
	var yamlList []*YamlUnstructured
	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(yamlContent), 100)
	for {
		var rawObj runtime.RawExtension
		if err := decoder.Decode(&rawObj); err != nil {
			break
		}
		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			return nil, err
		}
		unstructuredObj, err := ConvObjToUnstructed(obj)
		if err != nil {
			return nil, err
		}

		var yamlUtil YamlUnstructured
		yamlUtil.Unstructured = unstructuredObj
		yamlUtil.Gvk = gvk
		yamlList = append(yamlList, &yamlUtil)
	}
	return yamlList, nil
}
