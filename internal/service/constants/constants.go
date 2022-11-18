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
	EmailNotFoundMsg = map[string]string{
		"en": "Email not found",
		"id": "Email tidak ditemukan",
	}
	EmailNotFound = "Email Not Found"

	UserNotExistMsg = map[string]string{
		"en": "Invalid username / password",
		"id": "Username / password salah",
	}
	UserNotExist = "Invalid username / password"

	PasswordIsWrongMsg = map[string]string{
		"en": "Password is wrong",
		"id": "Password salah",
	}
	PasswordIsWrong = "Password is wrong"
)
