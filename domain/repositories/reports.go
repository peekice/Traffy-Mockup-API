package repositories

import (
	"log"
	"traffy-mock-crud/domain/datasources"
	"traffy-mock-crud/domain/entities"

	"gorm.io/gorm"
)

type reportsRepository struct {
	DB *gorm.DB
}

type IReportsRepository interface {
	FindAllReports(district string, status string, reportID string) (*[]entities.ReportDataModel, error)
	GetReportByID(id string) (*entities.ReportDataModel, error)
	InsertNewReport(data entities.ReportDataModel) error
	EditReport(reportID string, report entities.ReportDataModel) error
	DeleteReport(reportID string) error
}

func NewReportsRepository(db *datasources.PostgreSQL) IReportsRepository {
	return &reportsRepository{
		DB: db.DB.Table("reports"),
	}
}
func (r *reportsRepository) FindAllReports(district string, status string, reportID string) (*[]entities.ReportDataModel, error) {
	var reports []entities.ReportDataModel

	r.DB = r.DB.Session(&gorm.Session{NewDB: true})

	query := r.DB

	if district != "" {
		query = query.Where("district = ?", district)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if reportID != "" {
		query = query.Where("report_id = ?", reportID)
	}

	err := query.Find(&reports).Error
	if err != nil {
		return nil, err
	}

	return &reports, nil
}

func (r *reportsRepository) GetReportByID(id string) (*entities.ReportDataModel, error) {
	r.DB = r.DB.Session(&gorm.Session{NewDB: true})

	var report entities.ReportDataModel
	err := r.DB.Where("report_id = ?", id).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *reportsRepository) InsertNewReport(data entities.ReportDataModel) error {
	r.DB = r.DB.Session(&gorm.Session{NewDB: true})

	if err := r.DB.Create(&data).Error; err != nil {
		log.Printf("Reports -> InsertNewReport: %s \n", err)
		return err
	}
	return nil
}

func (r *reportsRepository) EditReport(reportID string, report entities.ReportDataModel) error {
	r.DB = r.DB.Session(&gorm.Session{NewDB: true})

	err := r.DB.Model(&entities.ReportDataModel{}).
		Where("report_id = ?", reportID).
		Updates(report).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *reportsRepository) DeleteReport(reportID string) error {
	r.DB = r.DB.Session(&gorm.Session{NewDB: true})

	err := r.DB.Where("report_id = ?", reportID).Delete(&entities.ReportDataModel{}).Error

	if err != nil {
		return err
	}

	return nil

}
