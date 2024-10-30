package handlers

import (
	"net/http"
	"payment-Api/models"
	"payment-Api/services"

	"github.com/gin-gonic/gin"
	"fmt"
	//"github.com/gorilla/mux"
	// "encoding/json"
)

// PaymentHandler struct
type PaymentHandler struct {
	paymentService services.PaymentService
	//paymentService services.FetchPaymentByReference
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
func (h *PaymentHandler) GetPaymentByReference(c *gin.Context) {
    // Extract query parameter using Gin's context
	paymentReference := c.DefaultQuery("payment_reference", "")
	fmt.Println("PAYMENT ",paymentReference)

    // Use the service layer to fetch the payment by reference
    payment, err := h.paymentService.FetchPaymentByReference(paymentReference)
    if err != nil {
        // Respond with a 404 if the payment is not found
		fmt.Println("PAYMENT ",err)

        c.JSON(404, gin.H{"error": "Payment not found"})
        return
    }

    // Respond with the payment object
    c.JSON(200, payment)
}




//Getpayment details by sender_acc_ref
func(h *PaymentHandler) FetchPaymentsBySenderRef(c *gin.Context){
	senderAccref := c.DefaultQuery("sender_account_ref", "")

	payments,err := h.paymentService.FetchPaymentsBySender(senderAccref)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, payments)
}

//Get Payment details by receiver_acc_ref
func(h *PaymentHandler) FetchPaymentsByReceiverRef(c *gin.Context){
	receiverAccref := c.DefaultQuery("receiver_account_ref","")
	payments,err := h.paymentService.FetchPaymentsByReceiver(receiverAccref)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK,payments)
}

//Get total payments by receiver_acc_ref
func(h *PaymentHandler) FetchTotalAmountByReceiver(c *gin.Context){
	receiverAccountRef := c.DefaultQuery("receiver_account_ref","")

	totalPayments,err := h.paymentService.FetchTotalAmountByReceiver(receiverAccountRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{"totalPayments": totalPayments})
}