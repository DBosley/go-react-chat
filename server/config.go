package server

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Tkanos/gonfig"
)

// Configuration is the structure of server.config.json
type Configuration struct {
	Port int
}

// LoadConfig loads the configuration
func LoadConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf(getFileName(), &config)
	handleError(err)

	return config
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"server.config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname)+"/config", strings.Join(filename, ""))
	println(filePath)

	return filePath
}
