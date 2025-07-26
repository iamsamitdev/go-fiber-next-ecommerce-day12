package utils

import (
	"errors"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// สร้างฟังก์ชันสำหรับ เข้ารหัสรหัสผ่าน
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// สร้างฟังก์ชันสำหรับ ตรวจสอบรหัสผ่าน
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePasswordComplexity ตรวจสอบความซับซ้อนของรหัสผ่าน
// ต้องมีอย่างน้อย 8 ตัวอักษร, ตัวเลข, ตัวอักษร และอักขระพิเศษ
func ValidatePasswordComplexity(password string) error {
	if len(password) < 8 {
		return errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัวอักษร")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	// ตรวจสอบแต่ละตัวอักษรในรหัสผ่าน
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// ตรวจสอบว่ามีครบทุกประเภทหรือไม่
	if !hasUpper {
		return errors.New("รหัสผ่านต้องมีตัวอักษรใหญ่อย่างน้อย 1 ตัว")
	}
	if !hasLower {
		return errors.New("รหัสผ่านต้องมีตัวอักษรเล็กอย่างน้อย 1 ตัว")
	}
	if !hasNumber {
		return errors.New("รหัสผ่านต้องมีตัวเลขอย่างน้อย 1 ตัว")
	}
	if !hasSpecial {
		return errors.New("รหัสผ่านต้องมีอักขระพิเศษอย่างน้อย 1 ตัว (!@#$%^&*)")
	}

	return nil
}

// IsValidPassword ตรวจสอบรหัสผ่านด้วย regex pattern
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	// ตรวจสอบแต่ละเงื่อนไขด้วย regex แยกกัน
	hasLower, _ := regexp.MatchString(`[a-z]`, password)
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	hasNumber, _ := regexp.MatchString(`\d`, password)
	hasSpecial, _ := regexp.MatchString(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~`+"`"+`]`, password)

	return hasLower && hasUpper && hasNumber && hasSpecial
}
