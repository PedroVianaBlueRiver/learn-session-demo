package utilities

import (
	"encoding/json"
	"log"
)

func SerializeJson(v any) ([]byte, error) {
	contentJson, err := json.Marshal(v)
	if err != nil {

		log.Fatal(err)
		return nil, err
	}
	return contentJson, nil
}

func DeserializeJson(content []byte, v any) error {
	model := v
	err := json.Unmarshal([]byte(content), &model)
	if err != nil {

		log.Fatal(err)
		return err
	}
	return nil
}
