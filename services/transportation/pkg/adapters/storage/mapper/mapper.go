package mapper

import (
	"encoding/json"
	"errors"
)

func ConvertTypes[T any](source interface{}, dest *T) error {
	b, err := json.Marshal(source)
	if err != nil {
		return errors.New("failed to marshal source object")
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return errors.New("failed to unmarshal into destination object")
	}

	return nil
}
