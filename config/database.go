package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file .env:", err)
	}
}

func CreateDatabase() {
	LoadEnv()

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke MySQL:", err)
	}

	var exists int
	db.Raw("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", os.Getenv("DB_NAME")).Scan(&exists)

	if exists == 0 {
		err = db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME")).Error
		if err != nil {
			log.Fatal("Gagal membuat database:", err)
		}
		log.Println("Database `" + os.Getenv("DB_NAME") + "` berhasil dibuat!")
	} else {
		log.Println("Database `" + os.Getenv("DB_NAME") + "` sudah ada!")
	}
}

func ConnectDB() {
	LoadEnv()

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("Database `" + os.Getenv("DB_NAME") + "` Connected")
	DB = db
}
