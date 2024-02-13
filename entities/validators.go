package entities

import "github.com/gin-gonic/gin"

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

func CheckIDandType(entity Entity) bool {
	if entity.ID == "" {
		return false
	}
	if entity.Type == "" {
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

func ValidateEntity(entity Entity, c *gin.Context) string {
	if err := c.ShouldBindJSON(&entity); err != nil {
		return "Error al leer la entidad"
	}
	if !CheckIDandType(entity) {
		return "ID and Type are required"
	}
	for _, prop := range entity.Properties {
		validationFunc, ok := Validators[string(prop.Type)]
		if !ok {
			return "Invalid property" + string(prop.Type)
		}
		check := validationFunc(prop.Value)
		if !check {
			return "Invalid property" + string(prop.Type)
		}
	}
	return "ok"
}
