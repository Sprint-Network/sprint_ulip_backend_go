package mca

import (
	"context"

	httpclient "prechecks/pkg/http/client"
)

type serviceHandler struct {
	httpClient httpclient.Interface
}

type Service interface {
	UploadIncorporationCertificate(ctx context.Context) error
}

// Implements user service methods
func NewService(c httpclient.Interface) Service {
	return &serviceHandler{
		httpClient: c,
	}
}

func (handler *serviceHandler) UploadIncorporationCertificate(ctx context.Context) error {
	return nil
}
