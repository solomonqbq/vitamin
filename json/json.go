package common

import json "github.com/bitly/go-simplejson"

func GetJsonToArray(body, key string) (result []map[string]interface{}, err error) {
	j, _ := json.NewJson([]byte(body))
	x, _ := j.Get(key).Array()

	result = make([]map[string]interface{}, 0)
	for _, v := range x {
		ss := v.(map[string]interface{})
		result = append(result, ss)
	}
	return
}

func GetJsonToMap(body, key string) (result map[string]interface{}, err error) {
	j, err := json.NewJson([]byte(body))
	if err != nil {
		return nil, err
	}
	result = j.Get(key).Interface().(map[string]interface{})
	return result, err
}

func GetJsonToValue(body, key string) (result string, err error) {
	j, err := json.NewJson([]byte(body))
	if err != nil {
		return "", err
	}
	result, err = j.Get(key).String()
	return result, err
}
