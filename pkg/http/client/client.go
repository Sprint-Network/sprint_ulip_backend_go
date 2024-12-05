package client

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cenkalti/backoff/v4"
	"prechecks/constants"
	"prechecks/pkg/logger"
	"prechecks/pkg/retry"
	"prechecks/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

// Interface - Http's Client interface
type Interface interface {
	Do(req *http.Request) (*http.Response, error)
}

// Mocks
type httpClientMock struct {
	mock.Mock
}

// GetMockHTTPClient - Returns mock http client
func GetMockHTTPClient() *httpClientMock {
	return &httpClientMock{}
}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// DoHttpRequest -
func DoHttpRequest(ctx context.Context, client Interface, endpoint, httpMethod string, body ...[]byte) (*http.Response, error) {
	var (
		req    *http.Request
		res    *http.Response
		err    error
		erresp types.ErrorResponse
	)

	if httpMethod == http.MethodPost || httpMethod == http.MethodPut {
		req, err = http.NewRequestWithContext(ctx, httpMethod, endpoint, bytes.NewBuffer(body[0]))
		if err != nil {
			logger.Log.WithFields(logrus.Fields{"error": err}).Error("Failed to create request")
			return nil, erresp
		}
	} else if httpMethod == http.MethodGet {
		req, err = http.NewRequestWithContext(ctx, httpMethod, endpoint, nil)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{"error": err}).Error("Failed to create request")
			return nil, erresp
		}
	}

	req.Header.Set(constants.RequestHeaderContentType, constants.ApplicationJSON)

	res, err = client.Do(req)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"error": err}).Error("Failed to execute api request")
		return nil, erresp
	}
	if res.StatusCode == http.StatusNotFound {
		logger.Log.Error("404: Invalid request endpoint")
		return nil, erresp
	}
	erresp.Status = http.StatusOK
	return res, erresp
}

// RetryWithBackoff is used to retry the API
// Needs a well formed request object and any struct implementing Interface
func RetryWithBackoff(req *http.Request, client Interface) (*http.Response, error) {

	var res *http.Response
	var err error
	retryerr := backoff.Retry(func() error {
		res, err = client.Do(req)
		if err != nil {
			return err
		}
		// Don't retry if unauthorised or endpoint not found or bad request
		if res.StatusCode == http.StatusUnauthorized ||
			res.StatusCode == http.StatusNotFound ||
			res.StatusCode == http.StatusBadRequest {
			return nil
		}
		return nil
	}, backoff.WithMaxRetries(retry.GetExponentialBackoff(), constants.RetryCount))
	if retryerr != nil {
		return nil, retryerr
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
