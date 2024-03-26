package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wheelsUN_transaction_ms/models"
)

func GetData(url string, payment interface{}) (string, error) {
	// Serializar el objeto payment a JSON
	jsonData, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}

	// Crear un nuevo cliente HTTP
	client := &http.Client{}

	// Crear la solicitud POST con el contenido JSON
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud y obtener la respuesta
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func PostCardPaymentS(payment models.RequestPayment) (string, error) {
	// Asignar la firma a la solicitud

	// Realizar la petición GET
	response, err := GetData("https://sandbox.api.payulatam.com/payments-api/4.0/service.cgi", payment)
	if err != nil {
		return "", err
	}

	// (Ignorar la parte del response != null)

	return response, nil
}

func GetMD5(payment models.RequestPayment) string {

	// Extract values and handle potential type mismatches
	apiKey := payment.Merchant.ApiKey
	taxValue := payment.Transaction.Order.AdditionalValues.TX_VALUE.Value
	reference := payment.Transaction.Order.ReferenceCode
	currency := payment.Transaction.Order.AdditionalValues.TX_TAX.Currency

	// Concatenate fields with proper conversions
	signature := fmt.Sprintf("%s~%s~%s~%d~%s", apiKey, "508029", reference, taxValue, currency)

	// Create MD5 hasher and calculate hash
	hasher := md5.New()
	hasher.Write([]byte(signature))
	hashBytes := hasher.Sum(nil)

	// Convert hash to lowercase hexadecimal string
	encodedHash := hex.EncodeToString(hashBytes)

	return encodedHash
}

func PostCardPayment(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en una estructura RequestPayment
	var payment models.RequestPayment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("JSON recibido: %+v\n", payment)
	// Realizar el pago con tarjeta de crédito

	payment.Transaction.Order.Signature = GetMD5(payment)

	response, err := PostCardPaymentS(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta al cliente
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
