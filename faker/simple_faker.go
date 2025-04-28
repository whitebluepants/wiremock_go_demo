package faker

import "wiremock_go_demo/models"

// SimpleFaker 简单实现
type SimpleFaker struct{}

func (f *SimpleFaker) FakerInteger(format string) int64 {
	return 42
}

func (f *SimpleFaker) FakerFloat(format string) float64 {
	return 3.14
}

func (f *SimpleFaker) FakerString(format string) string {
	return "mock_string"
}

func (f *SimpleFaker) ObjectFaker(schemaID int64) map[string]interface{} {
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