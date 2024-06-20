package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var defaultPort int = 1203
var defaultImagesDir string = "images"

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
}

// Config holds configuration parameters for this application
type Config struct {
	Port             int
	VerifyToken      string
	PageAccessToken  string
	AppSecret        string
	PlantnetAPIKey   string
	ImagesDir        string
	LanguageFile     string
	LobeEndpoint     string
	PushbulletAPIKey string
	PushbulletDevice string
}

func readConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := Config{
		Port:             viper.GetInt("port"),
		VerifyToken:      viper.GetString("verify_token"),
		PageAccessToken:  viper.GetString("page_access_token"),
		AppSecret:        viper.GetString("app_secret"),
		PlantnetAPIKey:   viper.GetString("plantnet_api_key"),
		ImagesDir:        viper.GetString("images_dir"),
		LanguageFile:     viper.GetString("language_file"),
		LobeEndpoint:     viper.GetString("lobe_endpoint"),
		PushbulletAPIKey: viper.GetString("pushbullet_apikey"),
		PushbulletDevice: viper.GetString("pushbullet_device"),
	}

	if cfg.Port == 0 {
		log.Debugf("port is not set, use the default: %d", defaultPort)
		cfg.Port = defaultPort
	}
	if cfg.VerifyToken == "" {
		return nil, fmt.Errorf("verify_token is required but not set")
	}
	if cfg.PageAccessToken == "" {
		return nil, fmt.Errorf("page_access_token is required but not set")
	}
	if cfg.AppSecret == "" {
		return nil, fmt.Errorf("app_secret is required but not set")
	}
	if cfg.PlantnetAPIKey == "" {
		return nil, fmt.Errorf("plantnet_api_key is required but not set")
	}
	if cfg.ImagesDir == "" {
		log.Debugf("images_dir is not set, use the default: %s", defaultImagesDir)
		cfg.ImagesDir = defaultImagesDir
	}
	if cfg.LanguageFile == "" {
		return nil, fmt.Errorf("language_file is required but not set")
	}
	if cfg.LobeEndpoint == "" {
		return nil, fmt.Errorf("lobeEndpoint is required")
	}
	if cfg.PushbulletAPIKey == "" {
		return nil, fmt.Errorf("pushbullet_apikey is required")
	}
	if cfg.PushbulletDevice == "" {
		return nil, fmt.Errorf("pushbullet_device is required")
	}

	return &cfg, nil
}
