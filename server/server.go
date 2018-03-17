package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Tkanos/gonfig"
	"github.com/labstack/echo"
)

// Configuration is the structure of server.config.json
type Configuration struct {
	Port int
}

// Start runs the webserver
func Start() {
	config := Configuration{}
	err := gonfig.GetConf(getFileName(), &config)

	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Port)))
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
