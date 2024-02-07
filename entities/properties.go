package entities

type DataType string

const (
	Number          DataType = "Number"
	Text            DataType = "Text"
	Boolean         DataType = "Boolean"
	StructuredValue DataType = "StructuredValue"
)

type Property struct {
	Type  DataType    `json:"type"`
	Value interface{} `json:"value"`
}

type Entity struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
}
