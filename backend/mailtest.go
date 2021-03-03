package main

import (
	"main/controller"
)

func Mailtest() {

	controller.SendMail("test@example.com", "テスト", "verify_email", nil, nil)
}
