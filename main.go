package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/my-gin-server/base/appconfig"
	"github.com/my-gin-server/base/applog"
	"github.com/my-gin-server/base/db"

	"github.com/gin-gonic/gin"
)

type CliFlags struct {
	ConfigPath string
	LogLevel   string
}

func ParseFlags() *CliFlags {
	var cliFlags CliFlags
	flag.StringVar(&cliFlags.ConfigPath, "config", "conf/app_cfg.yml", "path to config file")
	flag.StringVar(&cliFlags.LogLevel, "log-level", "INFO", "default log level")

	flag.Parse()
	return &cliFlags
}

func main() {
	cliFlags := ParseFlags()

	applog.SetupLoggers(cliFlags.LogLevel)

	config, err := appconfig.NewConfig(cliFlags.ConfigPath)
	if err != nil {
		applog.Error.Panicf("Cannot load config %s, error: %s", cliFlags.ConfigPath, err)
	}
	applog.Info.Printf("Server initializing with config: %+v\n", *config)

	if strings.ToLower(config.Server.Mode) == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	db.Init(config)
	defer db.Close()

	InitWorld(server)

	err = server.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err != nil {
		applog.Error.Panicf("Server startup failed: %s", err)
	}
}
