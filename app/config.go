package app

import "github.com/sirupsen/logrus"

type Config struct {
    PagesRoot string
    TemplatesRoot string
    LoggingLevel logrus.Level
}

var ServerConfig *Config
