package mca

import (
	"net/http"
	"prechecks/pkg/document"
	"prechecks/pkg/response"
	"prechecks/types"

	"prechecks/persistence/mca"
)

func getMCAService() mca.Service {
	return mca.NewService(&http.Client{})
}

func UploadIncorporationCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Response structures
	resp := types.HTTPResponse{}
	erresp := &types.ErrorResponse{}

	// Parse the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		response.SetErrorResponse(ctx, erresp, "Failed to parse file from form data: "+err.Error(), err.Error(), http.StatusBadRequest)
		resp.Error = erresp
		resp.Message = "Failed to parse file"
		resp.StatusCode = erresp.Status
		response.WriteJSONResponse(w, resp, erresp.Status)
		return
	}
	defer file.Close()

	// Extract CIN from the uploaded PDF
	cin, err := document.ExtractCINFromPDF(file)
	if err != nil {
		response.SetErrorResponse(ctx, erresp, "Error extracting CIN: "+err.Error(), err.Error(), http.StatusInternalServerError)
		resp.Error = erresp
		resp.Message = "Error extracting CIN from document"
		resp.StatusCode = erresp.Status
		response.WriteJSONResponse(w, resp, erresp.Status)
		return
	}

	// Call the external API with the extracted CIN
	apiResponse, err := document.FetchEmailFromCIN(cin)
	if err != nil {
		response.SetErrorResponse(ctx, erresp, "Error fetching email from external API: "+err.Error(), err.Error(), http.StatusInternalServerError)
		resp.Error = erresp
		resp.Message = "Error fetching email from external API"
		resp.StatusCode = erresp.Status
		response.WriteJSONResponse(w, resp, erresp.Status)
		return
	}

	resp.Data = map[string]string{"emailAddress": apiResponse, "cin": cin}
	resp.Message = "CIN and email fetched successfully"
	resp.StatusCode = http.StatusOK
	response.WriteJSONResponse(w, resp, resp.StatusCode)
}
