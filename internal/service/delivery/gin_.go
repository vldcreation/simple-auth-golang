package delivery

import (
	"context"
	"net/http"
)

type GinObject struct {
	features
}

type features struct {
}

var (
	text = http.StatusText

	msgSuccess    = map[string]string{"en": "Success", "id": "Sukses"}
	msgFailed     = map[string]string{"en": "Failed", "id": "Gagal"}
	msgUnexpected = map[string]string{
		"en": "An unexpected error occurred. Please try again later.",
		"id": "Terjadi kesalahan tak terduga. Silakan coba lagi nanti.",
	}
)

func NewGinHandler(ctx context.Context, f features) {
	obj := &GinObject{f}
	obj.InitRoutes(ctx)
}
