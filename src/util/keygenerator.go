package util

import (
	"crypto/rand"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"gitlab.com/abhishek.k8/crud/src/config"
)

const otpChars = "1234567890"

//TokenGenerator to generate auth token or referral code or anyother token
func TokenGenerator(bytevalue int) string {
	b := make([]byte, bytevalue)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//GenerateOTP - Generate OTP
func GenerateOTP(length int) (string, error) {
	log.Info("Env for OTP: ", config.AppConfig.Environment)
	if strings.ToLower(config.AppConfig.Environment) == "development" || strings.ToLower(config.AppConfig.Environment) == "staging" {
		return "123456", nil
	}
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
