package sarathi

import (
	"encoding/json"
	"net/http"
	"prechecks/pkg/response"
	sarathiservice "prechecks/pkg/sarathi"
	"prechecks/types"
)

type DLDetails struct {
	DLNumber string `json:"dlnumber"`
	DOB      string `json:"dob"`
	Name     string `json:"name"`
}

func DLDetailsHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody DLDetails

	resp := types.HTTPResponse{}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		response.SetErrorResponse(r.Context(), &types.ErrorResponse{}, "Invalid request body "+err.Error(), err.Error(), http.StatusBadRequest)
		resp.Error = &types.ErrorResponse{}
		resp.Message = "Invalid request body " + err.Error()
		resp.StatusCode = http.StatusBadRequest
		response.WriteJSONResponse(w, resp, http.StatusBadRequest)
		return
	}

	if reqBody.DLNumber == "" || reqBody.DOB == "" {
		response.SetErrorResponse(r.Context(), &types.ErrorResponse{}, "DL Number and DOB is required", "", http.StatusBadRequest)
		resp.Error = &types.ErrorResponse{}
		resp.Message = "DL Number and DOB is required"
		resp.StatusCode = http.StatusBadRequest
		response.WriteJSONResponse(w, resp, http.StatusBadRequest)
		return
	}

	isValid, err := sarathiservice.ValidateDLWithUIP(reqBody.DLNumber, reqBody.DOB, reqBody.Name)
	if err != nil {
		response.SetErrorResponse(r.Context(), &types.ErrorResponse{}, "Invalid Driving License, Name does not match with ULIP Data", "", http.StatusOK)
		resp.Error = &types.ErrorResponse{}
		resp.Message = "Invalid Driving License, Name does not match with ULIP Data"
		resp.StatusCode = http.StatusOK
		response.WriteJSONResponse(w, resp, http.StatusOK)
		return
	}

	response.WriteJSONResponse(w, map[string]any{
		"DLNumber":          reqBody.DLNumber,
		"validation_status": isValid,
	}, http.StatusOK)
}
