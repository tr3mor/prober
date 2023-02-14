package probes_test

import (
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"prober/internal/app/prober/probes"
	"testing"
	"time"
)

type HealthcheckTestSuite struct {
	suite.Suite

	logger *zap.SugaredLogger
	c      *http.Client
}

func (s *HealthcheckTestSuite) SetupTest() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	s.logger = logger.Sugar()
	s.c = &http.Client{Timeout: 10 * time.Second}
}

func TestHealthcheckTestSuite(t *testing.T) {
	suite.Run(t, new(HealthcheckTestSuite))
}

func (s *HealthcheckTestSuite) TestName() {
	//arrange
	h := probes.NewHealthcheck("testName", s.logger, "testUrl", s.c)
	//act
	name := h.Name()
	// assert
	assert.Equal(s.T(), "testName", name)
}

func (s *HealthcheckTestSuite) TestSuccessRequest() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	h := probes.NewHealthcheck("test", s.logger, srv.URL, s.c)

	h.Execute()

	status := testutil.CollectAndCount(h.M.Status)
	duration := testutil.CollectAndCount(h.M.Duration)
	assert.Equal(s.T(), 1, status)
	assert.Greater(s.T(), duration, 0)
}

func (s *HealthcheckTestSuite) TestFailedRequest() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()
	h := probes.NewHealthcheck("failedTest", s.logger, srv.URL, s.c)

	h.Execute()

	status := testutil.ToFloat64(h.M.Status)
	duration := testutil.ToFloat64(h.M.Duration)
	assert.Equal(s.T(), 0.0, status)
	assert.Greater(s.T(), duration, 0.0)
}
