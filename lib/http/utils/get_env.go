package modularHTTP

import (
	"os"
	"strconv"
)

func GetEnvString(key string, optionalValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return optionalValue
	}

	return value
}

func GetEnvInt(key string, optionalValue int) int {
	value, exists := os.LookupEnv(key)

	if !exists {
		return optionalValue

	}
	result, err := strconv.Atoi(value)

	if err != nil {
		return optionalValue
	}

	return result

}

func GetEnvBool(key string, optionalValue bool) bool {
	value, exists := os.LookupEnv(key)

	if !exists {
		return optionalValue
	}

	result, err := strconv.ParseBool(value)

	if err != nil {
		return optionalValue
	}

	return result
}
