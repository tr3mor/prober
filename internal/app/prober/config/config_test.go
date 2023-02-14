package config_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"prober/internal/app/prober/config"
	"testing"
	"time"
)

type ConfigTestSuite struct {
	suite.Suite

	logger *zap.SugaredLogger
}

func (s *ConfigTestSuite) SetupTest() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	s.logger = logger.Sugar()
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestParseConfig() {
	// arrange
	targets := []config.Target{{"yandex", "https://yandex.ru"}, {"github", "https://github.com"}}
	expectedCfg := &config.Config{Period: 1 * time.Minute, Timeout: 10 * time.Second, PrometheusPort: 2112, Targets: targets}
	//act
	cfg := config.ParseConfig("../../../../config/config.example.yaml", s.logger)
	// assert
	assert.Equal(s.T(), expectedCfg, cfg)
}
