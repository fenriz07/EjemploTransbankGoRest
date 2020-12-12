package controllers

import (
	"log"
	"net/http"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction"
)

/*CommitTransaction verifica el estado de una transaccion*/
func CommitTransaction(w http.ResponseWriter, r *http.Request) {

	/*
	 En caso del que pago sea anulado, comprobar si existe el parametro TBK_TOKEN.
	  Si existe el pago fue anulado por el usuario y debe comprobarse su estado con el Commit,
	  Si fue anulado adicionalmente tenemos los parametros TBK_ORDEN_COMPRA || TBK_ID_SESION
	*/

	var token string = ""
	var numberOrder string = ""
	var idSession string = ""

	canceledToken := r.FormValue("TBK_TOKEN")

	if len(canceledToken) != 0 {
		token = canceledToken
		numberOrder = r.FormValue("TBK_ORDEN_COMPRA")
		idSession = r.FormValue("TBK_ID_SESION")

		log.Printf("Number Order: %s\n Id Session: %s\n", numberOrder, idSession)
	} else {
		token = r.FormValue("token_ws")
	}

	transaction.Commit(token)
}
