package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	// Text to encode in QR Code
	code := "Aidana"

	// Generate the QR code
	qrCode, err := qrcode.Encode(code, qrcode.Medium, 256)
	if err != nil {
		log.Println("Error generating QR code:", err)
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	w.Write(qrCode)
}

func main() {
	http.HandleFunc("/generate-qrcode", generateQRCode)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
