package config

import (
	"backend-api/models/product/entity"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, *sql.DB, error) {

	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`,
		CONFIG["POSTGRES_HOST"],
		CONFIG["POSTGRES_USER"],
		CONFIG["POSTGRES_PASS"],
		CONFIG["POSTGRES_SCHEMA"],
		CONFIG["POSTGRES_PORT"],
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	dbConn.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Println("Error connect to PostgreSQL: ", err.Error())
		return nil, nil, err
	}

	sqlDb, errDb := dbConn.DB()
	if errDb != nil {
		log.Println(errDb)
	} else {
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(1000)
	}

	log.Println("Postgres connection success")
	return dbConn, sqlDb, nil
}
