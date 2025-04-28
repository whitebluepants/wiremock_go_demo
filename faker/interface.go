package faker

// Faker 接口定义假数据生成方法
type Faker interface {
	FakerInteger(format string) int64
	FakerFloat(format string) float64
	FakerString(format string) string
	ObjectFaker(schemaID int64) map[string]interface{}
}