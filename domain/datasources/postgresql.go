package datasources

import (
	"log"
	"os"
	"traffy-mock-crud/domain/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	DB *gorm.DB
}

func NewPostgreSQL(maxPoolSize int) *PostgreSQL {
	dsn := os.Getenv("POSTGRESQL_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("error connecting to PostgreSQL: ", err)
	}

	db.AutoMigrate(&entities.ReportDataModel{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error getting database instance: ", err)
	}

	sqlDB.SetMaxOpenConns(maxPoolSize)

	return &PostgreSQL{
		DB: db,
	}
}
