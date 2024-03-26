package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"wheelsUN_transaction_ms/configs"
	"wheelsUN_transaction_ms/database"

	"gorm.io/gorm"
)

var key = []byte("my32bytekey12345678901234567890") // Clave constante de 32 bytes

func encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func PostCreditCard(w http.ResponseWriter, r *http.Request) {
	var creditcard database.CreditCardDao

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si configs.DB es nulo
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}
	encryptedNumber, err := encrypt([]byte(creditcard.Number))
	encryptedName, err := encrypt([]byte(creditcard.Name))
	encryptedSecurityCode, err := encrypt([]byte(creditcard.SecurityCode))
	encryptedExpirationDate, err := encrypt([]byte(creditcard.ExpirationDate))

	creditcard.Number = string(encryptedNumber)
	creditcard.Name = string(encryptedName)
	creditcard.SecurityCode = string(encryptedSecurityCode)
	creditcard.ExpirationDate = string(encryptedExpirationDate)

	// Crear la entrada en la base de datos
	result := configs.DB.Create(&creditcard)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetCreditCardByID(w http.ResponseWriter, r *http.Request) {

	// 1. Extract ID from request path
	id := r.URL.Query().Get("id") // Assuming ID is passed in the query string
	if id == "" {
		http.Error(w, "Missing required parameter 'id' in query string", http.StatusBadRequest)
		return
	}

	// 2. Verify database connection
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	// 3. Fetch credit card data by ID
	var creditcard database.CreditCardDao
	result := configs.DB.First(&creditcard, id)
	if result.Error != nil {
		// Handle potential database errors gracefully (e.g., check for record not found)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Credit card with ID not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	encryptedNumber, err := decrypt([]byte(creditcard.Number))
	encryptedName, err := decrypt([]byte(creditcard.Name))
	encryptedSecurityCode, err := decrypt([]byte(creditcard.SecurityCode))
	encryptedExpirationDate, err := decrypt([]byte(creditcard.ExpirationDate))

	creditcard.Number = string(encryptedNumber)
	creditcard.Name = string(encryptedName)
	creditcard.SecurityCode = string(encryptedSecurityCode)
	creditcard.ExpirationDate = string(encryptedExpirationDate)

	// 5. Marshal decrypted data into JSON response
	responseJSON, err := json.Marshal(creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 6. Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func PutCreditCard(w http.ResponseWriter, r *http.Request) {
	// Extraer la ID de la tarjeta de crédito de la URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing required parameter 'id' in query string", http.StatusBadRequest)
		return
	}

	var creditcard database.CreditCardDao

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si configs.DB es nulo
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	// Update the entry in the database
	result := configs.DB.Model(&database.CreditCardDao{}).Where("id = ?", id).Updates(&creditcard)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
