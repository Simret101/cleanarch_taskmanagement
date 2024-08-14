package tests

import (
	"task/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordService(t *testing.T) {
	passwordService := Infrastructure.NewPasswordService()
	password := "securepassword"

	hashedPassword, err := passwordService.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	isValid := passwordService.ComparePasswords(hashedPassword, password)
	assert.True(t, isValid)

	isValid = passwordService.ComparePasswords(hashedPassword, "wrongpassword")
	assert.False(t, isValid)
}
