package entities

type ReportDataModel struct {
	ReportID  string `json:"report_id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"type:varchar(100);not null"`
	Detail    string `json:"detail" gorm:"type:text;not null"`
	District  string `json:"district" gorm:"type:varchar(100);not null"`
	BeforeImg string `json:"before_img" gorm:"type:text;not null"`
	CreatedAt string `json:"created_at" gorm:"type:varchar(100);not null"`

	Status       string  `json:"status" gorm:"type:varchar(100);not null"`
	SolvedBy     *string `json:"solved_by" gorm:"type:varchar(100);nullable"`
	SolvedDetail *string `json:"solved_detail" gorm:"type:text;nullable"`
	AfterImg     *string `json:"after_img" gorm:"type:text;nullable"`
	SolvedAt     *string `json:"solved_at" gorm:"type:varchar(100);nullable"`

	ReportStar    *int    `json:"report_star" gorm:"nullable"`
	ReportComment *string `json:"report_comment" gorm:"type:text;nullable"`

	ReportLike    int `json:"report_like" gorm:"default:0;not null"`
	ReportDislike int `json:"report_dislike" gorm:"default:0;not null"`
}

type ReportUserModel struct {
	Title    string `json:"title" gorm:"type:varchar(100);not null"`
	Detail   string `json:"detail" gorm:"type:text;not null"`
	District string `json:"district" gorm:"type:varchar(100);not null"`
}

type ReportOrganizeModel struct {
	SolvedBy     string `json:"solved_by" gorm:"type:varchar(100);nullable"`
	SolvedDetail string `json:"solved_detail" gorm:"type:text;nullable"`
}

type ReportUserCommentModel struct {
	ReportStar    int    `json:"report_star" gorm:"nullable"`
	ReportComment string `json:"report_comment" gorm:"type:text;nullable"`
}

type ReactionReport struct {
	ReactionTpye string `json:"reaction_type"`
}

func (ReportDataModel) TableName() string {
	return "reports"
}
