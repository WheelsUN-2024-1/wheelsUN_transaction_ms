package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"wheelsUN_transaction_ms/configs"
	"wheelsUN_transaction_ms/database"

	"gorm.io/gorm"
)

var key = []byte("my16bytekey12345")

// EncryptString encrypts a string using AES encryption
func encrypt(text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptString decrypts an AES-encrypted string
func decrypt(ciphertext string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return string(encrypted), nil
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
	creditcard.Number, err = encrypt(creditcard.Number)
	creditcard.Name, err = encrypt(creditcard.Name)
	creditcard.SecurityCode, err = encrypt(creditcard.SecurityCode)
	creditcard.ExpirationDate, err = encrypt(creditcard.ExpirationDate)

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

func GetCreditCardByID(w http.ResponseWriter, id string) {
	// 1. Verify database connection
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	// 2. Fetch credit card data by ID
	var creditcard database.CreditCardDao
	result := configs.DB.Where("credit_card_id = ?", id).Find(&creditcard)
	if result.Error != nil {
		// Handle potential database errors gracefully (e.g., check for record not found)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Credit card with ID not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	encryptedNumber, err := decrypt(creditcard.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encryptedName, err := decrypt(creditcard.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encryptedSecurityCode, err := decrypt(creditcard.SecurityCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encryptedExpirationDate, err := decrypt(creditcard.ExpirationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	creditcard.Number = encryptedNumber
	creditcard.Name = encryptedName
	creditcard.SecurityCode = encryptedSecurityCode
	creditcard.ExpirationDate = encryptedExpirationDate

	// 3. Marshal decrypted data into JSON response
	responseJSON, err := json.Marshal(creditcard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

func PutCreditCard(w http.ResponseWriter, r *http.Request, id string) {
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
	result := configs.DB.Model(&database.CreditCardDao{}).Where("credit_card_id = ?", id).Updates(&creditcard)
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

func DeleteCreditCard(w http.ResponseWriter, id string) {

	query := "DELETE FROM creditCard WHERE credit_card_id = ?"

	var creditCard database.CreditCardDao
	if configs.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	if err := configs.DB.Where("credit_card_id = ?", id).First(&creditCard).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verificar si configs.DB es nulo

	result := configs.DB.Exec(query, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(creditCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
