package env

import (
	"os"
	"reflect"
	"strings"
)

type Environment struct {
	DBConnectionString string
}

var Env Environment = LoadEnv(&Environment{}, ".env")

func LoadEnv[T any](env *T, envFilePath string) T {
	data, error := os.ReadFile(envFilePath)

	if error != nil {
		return *env
	}

	envLines := strings.Split(string(data), "\n")

	envValueOf := reflect.ValueOf(env).Elem()

	for _, line := range envLines {
		lineContent := strings.SplitN(line, "=", 2)
		if len(lineContent) < 2 {
			continue
		}

		field := envValueOf.FieldByName(lineContent[0])
		if field.CanSet() {
			field.SetString(lineContent[1])
		}
	}

	return *env
}
