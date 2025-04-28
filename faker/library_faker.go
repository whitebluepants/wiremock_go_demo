package faker

import (
	"github.com/go-faker/faker/v4"
	"wiremock_go_demo/models"
)

// LibraryFaker 使用开源库实现
type LibraryFaker struct{}

func (f *LibraryFaker) FakerInteger(format string) int64 {
	return faker.UnixTime()
}

func (f *LibraryFaker) FakerFloat(format string) float64 {
	return faker.Latitude()
}

func (f *LibraryFaker) FakerString(format string) string {
	return faker.Word()
}

func (f *LibraryFaker) ObjectFaker(schemaID int64) map[string]interface{} {
	propertyList, _ := models.GetCcamSchemaPropertiesBySchemaID(schemaID)
	var objectJsonMap = make(map[string]interface{})

	for _, property := range propertyList {
		key := property.PropertyName
		t := property.PropertyType
		if t == "integer" {
			objectJsonMap[key] = f.FakerInteger(property.PropertyFormat)
		} else if t == "float" {
			objectJsonMap[key] = f.FakerFloat(property.PropertyFormat)
		} else if t == "string" {
			objectJsonMap[key] = f.FakerString(property.PropertyFormat)
		} else if t == "ref" {
			objectJsonMap[key] = f.ObjectFaker(property.RefSchemaID)
		}
	}

	return objectJsonMap
}