package conf

import (
	"os"
	"strconv"
	"sync"

	"prechecks/constants"
	"prechecks/types"
)

var conf *types.Conf
var once sync.Once

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	env := getEnv(key, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return result
}

// New - Initialize Configuration
func New() *types.Conf {
	once.Do(func() {
		conf = &types.Conf{
			AppEnv:  getEnv("APP_ENV", constants.LocalEnvironment),
			AppPort: getEnvAsInt("APP_PORT", 80),
			AppHost: getEnv("APP_HOST", "http://localhost"),
		}
	})

	return conf
}
