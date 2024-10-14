package utils

import (
	"fmt"
	"math/rand"
	"time"
	"net/smtp"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOTP(email, otp string) error {
	from := "ammardzafran22@gmail.com"
	password := "pvao jgxk dzwi mdsd"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body := fmt.Sprintf("Your OTP is %s", otp)
	msg := []byte("Subject: Email Verification\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
}
