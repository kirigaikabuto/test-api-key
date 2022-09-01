package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/kirigaikabuto/test-api-key/api_key"
	"github.com/kirigaikabuto/test-api-key/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configName           = "main"
	configPath           = "/config/"
	version              = "0.0.1"
	postgresUser         = ""
	postgresPassword     = ""
	postgresDatabaseName = ""
	postgresHost         = ""
	postgresPort         = 5432
	postgresParams       = ""
	port                 = "5000"
	flags                = []cli.Flag{
		&cli.StringFlag{
			Name:        "config, c",
			Usage:       "path to .env config file",
			Value:       "main",
			Destination: &configName,
		},
	}
)

func parseEnvFile() {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.AddConfigPath(filepath + configPath)
	viper.SetConfigName(configName)
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	postgresUser = viper.GetString("db.primary.user")
	postgresPassword = viper.GetString("db.primary.pass")
	postgresDatabaseName = viper.GetString("db.primary.name")
	postgresParams = viper.GetString("db.primary.param")
	postgresPort = viper.GetInt("db.primary.port")
	postgresHost = viper.GetString("db.primary.host")
}

func run(c *cli.Context) error {
	parseEnvFile()
	gin.SetMode(gin.ReleaseMode)
	cfg := common.PostgresConfig{
		Host:     postgresHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		Database: postgresDatabaseName,
		Params:   postgresParams,
	}

	apiKeyStore, err := api_key.NewPostgresStore(cfg)
	apiKeyService := api_key.NewService(apiKeyStore)
	apiKeyHttpEndpoints := api_key.NewHttpEndpoints(setdata_common.NewCommandHandler(apiKeyService))
	r := gin.Default()
	apiKeyGroup := r.Group("/api-key")
	{
		apiKeyGroup.POST("/", apiKeyHttpEndpoints.MakeCreate())
		apiKeyGroup.GET("/id", apiKeyHttpEndpoints.MakeGetById())
		apiKeyGroup.GET("/key", apiKeyHttpEndpoints.MakeGetByKey())
		apiKeyGroup.GET("/", apiKeyHttpEndpoints.MakeList())
	}
	log.Info().Msg("app is running on port:" + port)
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting.")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "city api"
	app.Description = ""
	app.Usage = "city api"
	app.UsageText = "city api"
	app.Version = version
	app.Flags = flags
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Err(err)
	}
}
