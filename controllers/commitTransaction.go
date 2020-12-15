package controllers

import (
	"fmt"
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

	/*Commit de la transacción y resultado de la misma*/
	resp, err := transaction.Commit(token)

	if err != nil {
		fmt.Println(err)
	}

	log.Println(resp)

	/*Obtención del status de la transacción*/
	resp2, err := transaction.GetStatus(token)

	log.Println(resp2)

	if err != nil {
		log.Println(err)
	}

	/*Anulación*/
	resp3, err := transaction.Refund(token, 1000)

	if err != nil {
		log.Println(err)
	}

	log.Println("Respuesta 3")
	log.Println(resp3)

}
