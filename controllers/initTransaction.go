package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
)

/*InitTransaction retorna la vista*/
func InitTransaction(w http.ResponseWriter, r *http.Request) {

	/*
		transaction.Create(nOrden,session,monto,urlRetorno)
		Creamos la transacción
		Parametros:
		 - Id de compra
		 - Id session
		 - Monto
		 - Url de retorno
	*/

	nOrder := r.FormValue("order")

	order := fmt.Sprintf("goFenriz%s", nOrder)

	transaction, err := transaction.Create(order, "sesion1234557545", 1000, "http://localhost:8080/commit")

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	views := template.Must(template.ParseGlob("resources/views/*"))

	data := map[string]interface{}{
		"url":   transaction.URL,
		"token": transaction.Token,
	}

	err = views.ExecuteTemplate(w, "init.html", data)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}
