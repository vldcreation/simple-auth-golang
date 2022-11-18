package constants

import "net/http"

const FailedToUnmarshall = "failed to unmarshall"

var (
	Text = http.StatusText

	MsgSuccess    = map[string]string{"en": "Success", "id": "Sukses"}
	MsgFailed     = map[string]string{"en": "Failed", "id": "Gagal"}
	MsgUnexpected = map[string]string{
		"en": "An unexpected error occurred. Please try again later.",
		"id": "Terjadi kesalahan tak terduga. Silakan coba lagi nanti.",
	}
	MsgHopeForMiracle = map[string]string{
		"en": "We hope for a miracle",
		"id": "Kami berharap ada keajaiban",
	}
)
