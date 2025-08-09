package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"m-banking.info/config"
	"m-banking.info/model"
)

var DB *gorm.DB

func InitPostgres() {
	config.LoadConfig()

	dns := fmt.Sprintf("host=%s password=%s dbname=%s post=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "dummy123$$"),
		config.GetEnv("DB_NAME", "postgres"),
		config.GetEnv("DB_PORT", "5432"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to database", err)
	}

	log.Print("✅ Connected to PostgreSQL")

	//Auto migrate model
	err = DB.AutoMigrate(&model.M_Account{})
	if err != nil {
		log.Fatal("❌ Auto Migration failed:", err)
	}

	log.Print("✅ Auto migration completed")

}
