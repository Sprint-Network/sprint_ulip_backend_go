package vaahan

import (
	"net/http"
	"prechecks/pkg/response"
	vaahanservice "prechecks/pkg/vaahan"
)

func VehicleDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vehicleNumber := r.URL.Query().Get("vehiclenumber")
	if vehicleNumber == "" {
		http.Error(w, "Vehicle number is required", http.StatusBadRequest)
		return
	}

	vehicleDetails, err := vaahanservice.FetchVehicleDetailsFromULIP(vehicleNumber)
	if err != nil {
		http.Error(w, "Failed to fetch vehicle details", http.StatusInternalServerError)
		return
	}

	response.WriteJSONResponse(w, vehicleDetails, http.StatusOK)
}
