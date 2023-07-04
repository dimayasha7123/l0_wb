package json_validator

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

const schName = "schema.json"

type validator struct {
	sch *jsonschema.Schema
}

func New(schema string) (validator, error) {
	sch, err := jsonschema.CompileString(schName, schema)
	if err != nil {
		return validator{}, fmt.Errorf("can't create json schema validator: %v", err)
	}

	return validator{sch: sch}, nil
}
