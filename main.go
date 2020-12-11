package main

import (
	"github.com/fenriz07/EjemploTransbankGoRest/handlers"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/environment"
)

func main() {

	environmentIntegration := environment.IntegrationEnviroment{}

	environment.SetInstance(environmentIntegration)

	client.SetInstance()

	handlers.Init()

}
