package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	AppPort        string
	AppURL         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	JWTSecret      string
	JWTExpiresIn   string
	AdminEmail     string
	AdminPassword  string
	AdminFirstName string
	AdminLastName  string
}

func LoadConfig() (*Config, error) {

	// โหลดไฟล์ .env ถ้ามี
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	config := &Config{

		// ค่าที่ปลอดภัยสำหรับ default
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		// ค่าที่ไม่ปลอดภัยสำหรับ default - ต้องกำหนดใน env
		DBPassword:     getEnv("DB_PASS", ""),
		DBName:         getEnv("DB_NAME", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		AdminEmail:     getEnv("ADMIN_EMAIL", ""),
		AdminPassword:  getEnv("ADMIN_PASSWORD", ""),
		AdminFirstName: getEnv("ADMIN_FIRST_NAME", ""),
		AdminLastName:  getEnv("ADMIN_LAST_NAME", ""),
	}

	// ตรวจสอบค่าที่จำเป็นต้องมี
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// ฟังก์ชันตรวจสอบค่าที่จำเป็นใน environment variables
func validateConfig(config *Config) error {
	// ตรวจสอบค่าที่จำเป็นสำหรับ production
	if config.AppEnv == "production" {
		if config.DBPassword == "" {
			return fmt.Errorf("DB_PASS is required for production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET is required for production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWT_SECRET must be at least 32 characters long for production")
		}
		if config.DBSSLMode == "disable" {
			log.Println("Warning: SSL is disabled for database connection in production")
		}
		if config.AdminEmail == "" {
			return fmt.Errorf("ADMIN_EMAIL is required for production environment")
		}
		if config.AdminPassword == "" {
			return fmt.Errorf("ADMIN_PASSWORD is required for production environment")
		}
		if config.AdminFirstName == "" {
			return fmt.Errorf("ADMIN_FIRST_NAME is required for production environment")
		}
		if config.AdminLastName == "" {
			return fmt.Errorf("ADMIN_LAST_NAME is required for production environment")
		}
	}

	// ตรวจสอบรูปแบบ email (เฉพาะเมื่อมีค่า)
	if config.AdminEmail != "" && !isValidEmail(config.AdminEmail) {
		return errors.New("ADMIN_EMAIL must be a valid email address")
	}

	// ตรวจสอบค่าพื้นฐานที่ต้องมีเสมอ
	if config.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	return nil
}

// ฟังก์ชันช่วยสำหรับดึงค่า environment variable หรือค่าดีฟอลต์
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ฟังก์ชันตรวจสอบอีเมลว่าถูกต้องหรือไม่
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// ตรวจสอบพื้นฐาน - ต้องมี @ และ . และไม่เริ่มหรือจบด้วย @
	return len(email) > 0 &&
		len(email) <= 254 &&
		strings.Contains(email, "@") &&
		strings.Contains(email, ".") &&
		!strings.HasPrefix(email, "@") &&
		!strings.HasSuffix(email, "@")
}
