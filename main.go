package main

import (
	"github.com/fenriz07/EjemploTransbankGoRest/handlers"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)

func main() {

	webpayplus.SetEnvironmentIntegration()

	handlers.Init()

}
