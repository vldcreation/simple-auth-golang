package utils_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vldcreation/simple-auth-golang/internal/usecase/utils"
)

func TestEncryptedPassword(t *testing.T) {
	password := "password"
	encryptedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encrypt password", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, encryptedPassword)
	assert.NotEmpty(t, encryptedPassword)
}

func TestGenerateToken(t *testing.T) {
	var args = []interface{}{
		int64(1),
		"Fullname Test",
		"test",
		"test@gmail.com",
	}

	token, err := utils.GenerateToken(args[0].(int64), args[1].(string), args[2].(string), args[3].(string))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when generate token", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.AccessToken)
}

func TestCompareHashAndPassword(t *testing.T) {
	var args = []interface{}{
		"password",
	}

	encryptedPassword, err := utils.EncryptPassword(args[0].(string))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encrypt password", err)
	}

	args = append(args, encryptedPassword)

	log.Printf("args: %v", args)

	ok := utils.CompareHashAndPassword(args[1].(string), args[0].(string))

	assert.True(t, ok)
}
