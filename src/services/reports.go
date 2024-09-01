package services

import (
	"errors"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
	"traffy-mock-crud/domain/entities"
	"traffy-mock-crud/domain/repositories"
)

type reportsService struct {
	ReportsRepository repositories.IReportsRepository
	S3Uploader        S3Uploader
}

type IReportsService interface {
	GetAllReports(district string, status string, reportID string) (*[]entities.ReportDataModel, error)
	GetReportByID(id string) (*entities.ReportDataModel, error)
	InsertNewReport(report entities.ReportUserModel, imgFile multipart.File) error
	DeleteReport(reportID string) error
	AcceptReport(reportID string, data entities.ReportOrganizeModel) error
	FinishReport(reportID string, solvedDetail string, imgFile multipart.File) error
	CommentReport(reportID string, data entities.ReportUserCommentModel) error
	AddReaction(reportID string, data entities.ReactionReport) error
	RemoveReaction(reportID string, data entities.ReactionReport) error
}

func NewReportsService(repo0 repositories.IReportsRepository, s3Uploader S3Uploader) IReportsService {
	return &reportsService{
		ReportsRepository: repo0,
		S3Uploader:        s3Uploader,
	}
}

func (r *reportsService) GetAllReports(district string, status string, reportID string) (*[]entities.ReportDataModel, error) {
	return r.ReportsRepository.FindAllReports(district, status, reportID)
}

func (r *reportsService) GetReportByID(id string) (*entities.ReportDataModel, error) {

	return r.ReportsRepository.GetReportByID(id)
}

func (r *reportsService) InsertNewReport(report entities.ReportUserModel, imgFile multipart.File) error {
	reportID := CreateReportID()

	imgFileName := CreateImgFileName(reportID)

	fileURL, err := r.S3Uploader.UploadImage(imgFile, imgFileName)
	if err != nil {
		return err
	}

	reportData := entities.ReportDataModel{
		ReportID:  reportID,
		Title:     report.Title,
		Detail:    report.Detail,
		District:  report.District,
		BeforeImg: fileURL,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),

		Status:       "รอรับเรื่อง",
		SolvedBy:     nil,
		SolvedDetail: nil,

		ReportStar:    nil,
		ReportComment: nil,

		ReportLike:    0,
		ReportDislike: 0,
	}

	err = r.ReportsRepository.InsertNewReport(reportData)
	if err != nil {
		return err
	}
	return nil
}

func (r *reportsService) DeleteReport(reportID string) error {

	err := r.ReportsRepository.DeleteReport(reportID)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportsService) AcceptReport(reportID string, data entities.ReportOrganizeModel) error {

	reportExist, err := r.GetReportByID(reportID)

	if err != nil {
		return err
	}

	if reportExist.Status != "รอรับเรื่อง" {
		return errors.New("report is not waiting")
	}

	reportExist.SolvedBy = &data.SolvedBy
	reportExist.SolvedDetail = &data.SolvedDetail
	reportExist.Status = "กำลังดำเนินการ"

	err = r.ReportsRepository.EditReport(reportID, *reportExist)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportsService) FinishReport(reportID string, solvedDetail string, imgFile multipart.File) error {
	reportExist, err := r.GetReportByID(reportID)

	if err != nil {
		return err
	}

	if reportExist.Status != "กำลังดำเนินการ" {
		return errors.New("report is not in process")
	}

	imgFileName := CreateImgFileName(reportID)

	fileURL, err := r.S3Uploader.UploadImage(imgFile, imgFileName)
	if err != nil {
		return err
	}

	finishTime := time.Now().Format("2006-01-02 15:04:05")

	reportExist.SolvedDetail = &solvedDetail
	reportExist.Status = "เสร็จสิ้น"
	reportExist.AfterImg = &fileURL
	reportExist.SolvedAt = &finishTime

	err = r.ReportsRepository.EditReport(reportID, *reportExist)
	if err != nil {
		return err
	}

	return nil

}

func (r *reportsService) CommentReport(reportID string, data entities.ReportUserCommentModel) error {

	reportExist, err := r.GetReportByID(reportID)

	if err != nil {
		return err
	}

	if reportExist.Status != "เสร็จสิ้น" {
		return errors.New("report is not finished")
	}

	reportExist.ReportComment = &data.ReportComment
	reportExist.ReportStar = &data.ReportStar

	err = r.ReportsRepository.EditReport(reportID, *reportExist)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportsService) AddReaction(reportID string, data entities.ReactionReport) error {

	reportExist, err := r.GetReportByID(reportID)

	if err != nil {
		return err
	}

	if data.ReactionTpye == "like" {
		reportExist.ReportLike = reportExist.ReportLike + 1
	} else if data.ReactionTpye == "dislike" {
		reportExist.ReportDislike = reportExist.ReportDislike + 1
	} else {
		return errors.New("invalid reaction type")
	}

	err = r.ReportsRepository.EditReport(reportID, *reportExist)
	if err != nil {
		return err
	}

	return nil
}

func (r *reportsService) RemoveReaction(reportID string, data entities.ReactionReport) error {

	reportExist, err := r.GetReportByID(reportID)

	if err != nil {
		return err
	}

	if data.ReactionTpye == "like" {
		reportExist.ReportLike = reportExist.ReportLike - 1
	} else if data.ReactionTpye == "dislike" {
		reportExist.ReportDislike = reportExist.ReportDislike - 1
	} else {
		return errors.New("invalid reaction type")
	}

	err = r.ReportsRepository.EditReport(reportID, *reportExist)
	if err != nil {
		return err
	}

	return nil
}

func CreateReportID() string {
	year := time.Now().Year()

	var letterRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, 6)
	for i := range randomString {
		randomString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return strconv.Itoa(year) + "-" + string(randomString)

}

func CreateImgFileName(reportID string) string {
	return reportID + "-" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
}
