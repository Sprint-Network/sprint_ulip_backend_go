package v1

import (
	"net/http"
	"prechecks/api/v1/mca"
	"prechecks/api/v1/sarathi"
	"prechecks/api/v1/vaahan"

	"github.com/go-chi/chi"
)

// Router - Return chi router's http handler
func Router() http.Handler {

	r := chi.NewRouter()
	r.Get("/upload-incorporation-certificate", mca.UploadIncorporationCertificate)
	r.Get("/vehicleDetails", vaahan.VehicleDetailsHandler)
	r.Post("/validateDL", sarathi.DLDetailsHandler)
	return r
}
