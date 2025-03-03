package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/services"
)

type AttendanceController struct {
	attendanceService services.AttendanceService
}

func NewAttendanceController(attendanceService services.AttendanceService) *AttendanceController {
	return &AttendanceController{attendanceService: attendanceService}
}

// CreateAttendance godoc
// @Summary Create a new Attendance
// @Description Create a new Attendance with the given details
// @Tags attendances
// @Accept json
// @Produce json
// @Param Attendance body contracts.AttendanceRequest true "Attendance details"
// @Success 200 {object} contracts.AttendanceRequest
// @Failure 400 {object} string "Bad Request"
// @Router /attendances [post]
// @Security BearerAuth
func (a *AttendanceController) CreateAttendance(c *gin.Context) {
	var AttendanceRequest contracts.AttendanceRequest
	if err := c.ShouldBindJSON(&AttendanceRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Attendance, err := a.attendanceService.CreateAttendance(&AttendanceRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, Attendance)
}

// GetAllattendances godoc
// @Summary Get all attendances
// @Description Get all attendances
// @Tags attendances
// @Produce json
// @Success 200 {array} models.Attendance
// @Failure 500 {object} string "Internal Server Error
// @Router /attendances [get]
// @Security BearerAuth
func (a *AttendanceController) GetAllattendances(c *gin.Context) {
	attendances, err := a.attendanceService.GetAllAttendances()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attendances)
}

// DeleteAttendance godoc
// @Summary Delete an Attendance
// @Description Delete an Attendance by ID
// @Tags attendances
// @Produce json
// @Success 204
// @Failure 400 {object} string
// @Router /attendances/{id} [delete]
// @Security BearerAuth
func (a *AttendanceController) DeleteAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = a.attendanceService.DeleteAttendance(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
