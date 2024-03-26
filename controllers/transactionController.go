package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wheelsUN_transaction_ms/configs"
	"wheelsUN_transaction_ms/database"
	"wheelsUN_transaction_ms/models"

	"gorm.io/gorm"
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

	// Convertir la estructura de respuesta en XML
	var resp models.Response
	err = xml.NewDecoder(strings.NewReader(response)).Decode(&resp)
	if err != nil {
		fmt.Println("Error al decodificar XML:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear el objeto de base de datos
	var dataBase database.TransactionDao
	dataBase.ReferenceCode = payment.Transaction.Order.ReferenceCode
	dataBase.Description = payment.Transaction.Order.Description
	dataBase.Value = payment.Transaction.Order.AdditionalValues.TX_VALUE.Value
	dataBase.PaymentMethods = payment.Transaction.PaymentMethod
	dataBase.State = resp.TransactionResponse.State
	dataBase.TransactionIdPay = resp.TransactionResponse.TransactionId
	dataBase.OrderId = resp.TransactionResponse.Order
	dataBase.TripId = 1
	dataBase.CreditCardId = 1

	// Guardar en la base de datos
	result := configs.DB.Create(&dataBase)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir el objeto de base de datos a JSON
	jsonData, err := json.Marshal(dataBase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer encabezado de respuesta para indicar que se devuelve JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetTransactionReferenceCode(w http.ResponseWriter, r *http.Request, referenceCode string) {
	// 2. Verify database connection
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	// 3. Fetch transaction data by reference code
	var transaction database.TransactionDao
	result := configs.DB.First(&transaction, "reference_code = ?", referenceCode)
	if result.Error != nil {
		// Handle potential database errors gracefully (e.g., check for record not found)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Transaction with referenceCode not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	// 5. Marshal decrypted data into JSON response
	responseJSON, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 6. Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
