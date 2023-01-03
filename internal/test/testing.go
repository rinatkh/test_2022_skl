package test

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/sirupsen/logrus"
	"testing"
)

// TestLogger ...
func TestLogger(t *testing.T) *logrus.Entry {
	t.Helper()
	logger := logrus.New()
	entry := logrus.NewEntry(logger)
	return entry
}

// TestLogger ...
func TestConfig(t *testing.T) *config.Config {
	t.Helper()
	return &config.Config{}
}
