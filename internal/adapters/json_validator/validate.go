package json_validator

import (
	"encoding/json"
	"fmt"
)

func (v validator) Validate(data []byte) error {
	var d interface{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return fmt.Errorf("can't unmarshall json: %v", err)
	}

	err = v.sch.Validate(d)
	if err != nil {
		return fmt.Errorf("data not valid: %v", err)
	}

	return nil
}
