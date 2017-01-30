package main

import (
	"encoding/json"
	"fmt"

	"github.com/yosida95/uritemplate"
	"gopkg.in/alecthomas/kingpin.v2"
)

type jsonFlag struct {
	values uritemplate.Values
}

func parseSingleValue(input interface{}) (value uritemplate.Value, err error) {
	switch data := input.(type) {
	case string:
		value = uritemplate.String(data)
	case []interface{}:
		list := make([]string, 0, len(data))
		for _, e := range data {
			s, ok := e.(string)
			if !ok {
				err = fmt.Errorf("list requires all elements to be strings, '%v' is not a string", e)
				return
			}
			list = append(list, s)
		}
		value = uritemplate.List(list...)
	case map[string]interface{}:
		kv := make([]string, 0, len(data)*2)
		for k, v := range data {
			s, ok := v.(string)
			if !ok {
				err = fmt.Errorf("map requires all values to be strings, '%v' is not a string", v)
				return
			}
			kv = append(kv, k, s)
		}
		value = uritemplate.KV(kv...)
	default:
		err = fmt.Errorf("'%v' is not a valid string, key-value pair or array", input)
	}
	return
}

func (v *jsonFlag) Set(value string) error {
	var p interface{}
	err := json.Unmarshal([]byte(value), &p)
	if err != nil {
		return err
	}
	inputMap, ok := p.(map[string]interface{})
	if !ok {
		return fmt.Errorf("'%s' is not a valid JSON object", value)
	}

	for key, val := range inputMap {
		templateValue, err := parseSingleValue(val)
		if err != nil {
			return err
		}

		v.values.Set(key, templateValue)
	}

	return nil
}

func (v *jsonFlag) String() string {
	return ""
}

func (v *jsonFlag) IsCumulative() bool {
	return true
}

func JSONValues(s kingpin.Settings) *uritemplate.Values {
	target := make(uritemplate.Values)
	s.SetValue(&jsonFlag{target})
	return &target
}
