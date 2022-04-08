package test

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestABC(t *testing.T) {
	type StructA struct {
		Process string `json:"p1" jsonb:"p2"`
	}
	a1 := StructA{"abc"}

	json := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "jsonb",
	}.Froze()

	aj, _ := json.Marshal(a1)
	t.Log(string(aj))

	json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "json",
		OnlyTaggedField:        true,
	}.Froze()

	aj2, _ := json.Marshal(a1)
	t.Log(string(aj2))

}

type StructA struct {
	Process string `json:"p1" jsonb:"p2"`
}

func (a StructA) MarshalJSON() ([]byte, error) {
	json2 := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "jsonb",
		OnlyTaggedField:        true,
	}.Froze()
	type structA StructA
	a1 := structA(a)
	return json2.Marshal(a1)
}

func TestMarshalJSON(t *testing.T) {
	a1 := StructA{"abcd"}
	ar, _ := json.Marshal(a1)
	//ar, _ := a1.MarshalJSON2()
	t.Log(string(ar))
}
