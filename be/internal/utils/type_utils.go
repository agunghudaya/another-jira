package utils

// SafeFloat64 asserts an interface{} to a float64 and returns a default value if the assertion fails.
func SafeFloat64(value interface{}, defaultValue float64) float64 {
	if val, ok := value.(float64); ok {
		return val
	}
	return defaultValue
}
