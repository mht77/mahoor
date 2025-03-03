package repositories

import (
	"github.com/mht77/mahoor/models"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	CreateAttendance(Attendance *models.Attendance) (*models.Attendance, error)
	GetAllAttendances() (*[]models.Attendance, error)
	DeleteAttendance(id uint) error
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) CreateAttendance(Attendance *models.Attendance) (*models.Attendance, error) {
	err := r.db.Create(Attendance).Error
	if err != nil {
		return nil, err
	}
	return Attendance, nil
}

func (r *attendanceRepository) GetAllAttendances() (*[]models.Attendance, error) {
	var Attendances *[]models.Attendance
	err := r.db.Find(&Attendances).Error
	if err != nil {
		return nil, err
	}
	return Attendances, nil
}

func (r *attendanceRepository) DeleteAttendance(id uint) error {
	return r.db.Delete(&models.Attendance{}, id).Error
}
