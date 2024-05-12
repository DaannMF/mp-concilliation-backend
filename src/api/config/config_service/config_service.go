package config_service

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
)

func SetupConfig() {
	if os.Getenv("SCOPE") == "" || os.Getenv("SCOPE") == "test" {
		_, filename, _, _ := runtime.Caller(0)
		re := regexp.MustCompile("(?P<mainPath>.*/src/api)")
		match := re.FindStringSubmatch(filename)
		result := make(map[string]string)

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		os.Setenv("configFileName", fmt.Sprintf("%s%s", result["mainPath"], "/config/config_service/latest/application.properties"))
		os.Setenv("checksumEnabled", "false")
		os.Setenv("IS_PROD_SCOPE", "false")
		return
	}
}
