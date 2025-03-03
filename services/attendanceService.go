package services

import (
	"errors"

	"github.com/mht77/mahoor/contracts"
	"github.com/mht77/mahoor/models"
	"github.com/mht77/mahoor/repositories"
)

type AttendanceService interface {
	CreateAttendance(Attendance *contracts.AttendanceRequest) (*models.Attendance, error)
	GetAllAttendances() (*[]models.Attendance, error)
	DeleteAttendance(id uint) error
}

type attendanceService struct {
	repo repositories.AttendanceRepository
}

func NewAttendanceService(repo repositories.AttendanceRepository) AttendanceService {
	return &attendanceService{repo: repo}
}

func (s *attendanceService) CreateAttendance(AttendanceRequest *contracts.AttendanceRequest) (*models.Attendance, error) {
	if AttendanceRequest.Name == "" {
		return nil, errors.New("name is required")
	}
	Attendance := &models.Attendance{
		Name:   AttendanceRequest.Name,
		Number: AttendanceRequest.Number,
	}
	return s.repo.CreateAttendance(Attendance)
}

func (s *attendanceService) GetAllAttendances() (*[]models.Attendance, error) {
	return s.repo.GetAllAttendances()
}

func (s *attendanceService) DeleteAttendance(id uint) error {
	return s.repo.DeleteAttendance(id)
}
