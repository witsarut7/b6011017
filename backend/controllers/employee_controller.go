package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/witsarut7/app/ent"
	"github.com/witsarut7/app/ent/employee"
)

// EmployeeController defines the struct for the employee controller
type EmployeeController struct {
	client *ent.Client
	router gin.IRouter
}

// Employee defines the struct for the employee controller
type Employee struct {
	EMPLOYEENAME     string
	EMPLOYEEPASSWORD string
}

// CreateEmployee handles POST requests for adding employee entities
// @Summary Create employee
// @Description Create employee
// @ID create-employee
// @Accept   json
// @Produce  json
// @Param employee body ent.Employee true "Employee entity"
// @Success 200 {object} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees [post]
func (ctl *EmployeeController) CreateEmployee(c *gin.Context) {
	obj := ent.Employee{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "employee binding failed",
		})
		return
	}

	em, err := ctl.client.Employee.
		Create().
		SetEMPLOYEENAME(obj.EMPLOYEENAME).
		SetEMPLOYEEPASSWORD(obj.EMPLOYEEPASSWORD).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, em)
}

// GetEmployee handles GET requests to retrieve a employee entity
// @Summary Get a employee entity by ID
// @Description get employee by ID
// @ID get-employee
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees/{id} [get]
func (ctl *EmployeeController) GetEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, em)
}

// ListEmployee handles request to get a list of employee entities
// @Summary List employee entities
// @Description list employee entities
// @ID list-employee
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees [get]
func (ctl *EmployeeController) ListEmployee(c *gin.Context) {
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

	employees, err := ctl.client.Employee.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, employees)
}

// NewEmployeeController creates and registers handles for the employee controller
func NewEmployeeController(router gin.IRouter, client *ent.Client) *EmployeeController {
	emc := &EmployeeController{
		client: client,
		router: router,
	}
	emc.register()
	return emc
}

// InitEmployeeController registers routes to the main engine
func (ctl *EmployeeController) register() {
	employees := ctl.router.Group("/employees")

	employees.GET("", ctl.ListEmployee)

	// CRUD
	employees.POST("", ctl.CreateEmployee)
	employees.GET(":id", ctl.GetEmployee)
}
