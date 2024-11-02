package output

import (
	"encoding/json"
	"fmt"
)

type Renderer interface {
	RenderError(error Error) string
}

var _ Renderer = Json{}

type Json struct {
	PrettyPrint      bool
	DenormalizeTable bool
}

func (r Json) RenderError(error Error) string {
	return r.marshal(map[string]any{
		"type":    "error",
		"message": error.Message,
	})
}

func (r Json) marshal(v any) string {
	var data []byte
	var err error
	if r.PrettyPrint {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		panic(fmt.Sprintf(`BUG: failed to marshal JSON: %s`, err.Error()))
	}
	return string(data)
}
