package handlers

import (
	"net/http"
	"payment-Api/models"
	"payment-Api/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	 "encoding/json"
)

// PaymentHandler struct
type PaymentHandler struct {
	paymentService services.PaymentService
}

// NewPaymentHandler returns a new PaymentHandler
func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

// CreatePaymentHandler handles the POST /payments endpoint
func (h *PaymentHandler) CreatePaymentHandler(c *gin.Context) {
	var payment models.Payment

	// Bind the incoming JSON to the Payment struct
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the payment using the service layer
	if err := h.paymentService.ProcessPayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save payment"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{"message": "Payment received and saved successfully"})
}


// GetPaymentByReference retrieves a payment by reference
func (h *PaymentHandler) GetPaymentByReference(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    paymentReference := vars["payment_reference"]

    payment, err := h.paymentService.FetchPaymentByReference(paymentReference)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payment)
}