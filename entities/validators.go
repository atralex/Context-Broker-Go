package entities

func testNumber(value interface{}) bool {
	if _, ok := value.(float64); !ok {
		return false
	}
	return true
}
func testText(value interface{}) bool {
	if _, ok := value.(string); !ok {
		return false
	}
	return true
}
func testBoolean(value interface{}) bool {
	if _, ok := value.(bool); !ok {
		return false
	}
	return true
}
func testStructuredValue(value interface{}) bool {
	_, isMap := value.(map[string]interface{})
	_, isArray := value.([]interface{})
	if !isMap && !isArray {
		return false
	}
	return true
}

var Validators = map[string]func(param interface{}) bool{
	"Number":          testNumber,
	"Text":            testText,
	"Boolean":         testBoolean,
	"StructuredValue": testStructuredValue,
}
