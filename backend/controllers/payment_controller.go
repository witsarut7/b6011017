package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/witsarut7/app/ent/payment"

	"github.com/gin-gonic/gin"
	"github.com/witsarut7/app/ent"
	"github.com/witsarut7/app/ent/customer"
	"github.com/witsarut7/app/ent/employee"
	"github.com/witsarut7/app/ent/paymenttype"
	"github.com/witsarut7/app/ent/roomtype"
)

// PaymentController defines the struct for the payment controller
type PaymentController struct {
	client *ent.Client
	router gin.IRouter
}

// Payment defines the struct for the payment controller
type Payment struct {
	CUSTOMER    int
	ROOMTYPE    int
	PAYMENTTYPE int
	EMPLOYEE    int
	PAYDAY      string
}

// CreatePayment handles POST requests for adding payment entities
// @Summary Create payment
// @Description Create payment
// @ID create-payment
// @Accept   json
// @Produce  json
// @Param payment body ent.Payment true "Payment entity"
// @Success 200 {object} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [post]
func (ctl *PaymentController) CreatePayment(c *gin.Context) {
	obj := Payment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "payment video binding failed",
		})
		return
	}
	
	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.CUSTOMER))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "customer not found",
		})
		return
	}

	rt, err := ctl.client.Roomtype.
		Query().
		Where(roomtype.IDEQ(int(obj.ROOMTYPE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "roomtype not found",
		})
		return
	}

	pt, err := ctl.client.Paymenttype.
		Query().
		Where(paymenttype.IDEQ(int(obj.PAYMENTTYPE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "paymenttype not found",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.EMPLOYEE))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}
	
	time, err := time.Parse(time.RFC3339, obj.PAYDAY)
	pm, err := ctl.client.Payment.
		Create().
		SetCustomer(cus).
		SetRoomtype(rt).
		SetPaymenttype(pt).
		SetEmployee(em).
		SetPAYDAY(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pm)
}

// GetPayment handles GET requests to retrieve a payment entity
// @Summary Get a payment entity by ID
// @Description get payment by ID
// @ID get-payment
// @Produce  json
// @Param id path int true "Payment ID"
// @Success 200 {object} Payment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments/{id} [get]
func (ctl *PaymentController) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pm, err := ctl.client.Payment.
		Query().
		WithCustomer().
		WithRoomtype().
		WithPaymenttype().
		WithEmployee().
		Where(payment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pm)
}

// ListPayment handles request to get a list of payment entities
// @Summary List payment entities
// @Description list payment entities
// @ID list-payment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Payment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /payments [get]
func (ctl *PaymentController) ListPayment(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	payments, err := ctl.client.Payment.
		Query().
		WithCustomer().
		WithRoomtype().
		WithPaymenttype().
		WithEmployee().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, payments)
}

// NewPaymentController creates and registers handles for the payment controller
func NewPaymentController(router gin.IRouter, client *ent.Client) *PaymentController {
	pmc := &PaymentController{
		client: client,
		router: router,
	}
	pmc.register()
	return pmc

}

// InitPaymentController registers routes to the main engine
func (ctl *PaymentController) register() {
	payments := ctl.router.Group("/payments")

	payments.GET("", ctl.ListPayment)

	// CRUD
	payments.POST("", ctl.CreatePayment)
	payments.GET(":id", ctl.GetPayment)
}
