package response

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"prechecks/pkg/logger"
	"prechecks/types"
)

// WriteJSONResponse - Helper to serialize and write a JSON response
func WriteJSONResponse(w http.ResponseWriter, resp interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// Marshal the response into JSON
	data, err := json.Marshal(resp)
	if err != nil {
		// Handle JSON marshalling error gracefully
		http.Error(w, `{"error":"failed to serialize response"}`, http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

// SetErrorResponse - Helper to log errors and populate the error response structure
func SetErrorResponse(ctx context.Context, erresp *types.ErrorResponse, errMsg string, sysErr string, status int) {
	logger.Log.WithFields(logrus.Fields{
		"errMsg": errMsg,
		"sysErr": sysErr,
		"status": status,
	}).Error(errMsg)

	erresp.Status = status
	erresp.ErrorMsg = errMsg
	erresp.SystemErrorMsg = sysErr
}
