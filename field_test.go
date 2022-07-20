package learn_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Lingga").Info("Hello World")

	logger.WithField("username", "Wahyu").
		WithField("name", "Lingga Wahyu Rochim").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Lingga",
		"name":     "Lingga Wahyu Rochim",
	}).Info("Hello World")
}
