package handler

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	DefaultMaxBodyBytes   = int64(1024)
	DefaultTimeoutSeconds = int64(10)
	DefaultPort           = "3000"
	DefaultVersion        = "v1"
)

// Defines the configuration settings for the Web service
type Config struct {
	MaxBodyBytes    int64
	TimeoutDuration time.Duration
	Port            string
	BaseURL         string
}

// Extract the configurations from the environment variables for the Web service. Provides default
// settings if some environments are not defined
func ExtractConfig() *Config {
	return &Config{
		MaxBodyBytes:    maxBodyBytes(),
		TimeoutDuration: timeout(),
		Port:            port(),
		BaseURL:         baseURL(),
	}
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.WithFields(log.Fields{
			"default": DefaultPort,
		}).Warn("environment PORT is not defined")
		port = DefaultPort
	}
	return ":" + port
}

func timeout() time.Duration {
	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"default": DefaultTimeoutSeconds,
		}).Warn("environment HANDLER_TIMEOUT is not defined correctly")
		ht = DefaultTimeoutSeconds
	}
	return time.Duration(time.Duration(ht) * time.Second)
}

func maxBodyBytes() int64 {
	maxBodyBytes := os.Getenv("MAX_BODY_BYTES")
	mbb, err := strconv.ParseInt(maxBodyBytes, 0, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"default": DefaultMaxBodyBytes,
		}).Warn("environment MAX_BODY_BYTES is not defined correctly")
		mbb = DefaultMaxBodyBytes
	}
	return mbb
}

func baseURL() string {
	version := os.Getenv("VERSION")
	if len(version) == 0 {
		log.WithField("default", "v1").Warn("environment VERSION is not defined")
		version = DefaultVersion
	}
	return "/" + version
}
