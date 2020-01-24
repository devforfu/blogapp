package app

import "github.com/sirupsen/logrus"

type Config struct {
    TemplatesRoot string
    LoggingLevel logrus.Level
}

var ServerConfig *Config
