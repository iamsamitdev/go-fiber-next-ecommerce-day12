package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// สร้างตัวแปร validate สำหรับการตรวจสอบความถูกต้องของข้อมูล
var validate = validator.New()

// สร้างฟังก์ชัน Validate ที่ใช้สำหรับตรวจสอบความถูกต้องของข้อมูล
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// init function จะทำงานเมื่อ package ถูกโหลด
func init() {
	// ลงทะเบียน custom validator สำหรับรหัสผ่านที่ซับซ้อน
	validate.RegisterValidation("password_complex", validatePasswordComplex)
}

// validatePasswordComplex เป็น custom validator สำหรับตรวจสอบความซับซ้อนของรหัสผ่าน
func validatePasswordComplex(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return IsValidPassword(password)
}

// ValidatePassword ตรวจสอบรหัสผ่านและคืนค่า error message ที่เหมาะสม
func ValidatePassword(password string) error {
	if !IsValidPassword(password) {
		if len(password) < 8 {
			return errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัวอักษร")
		}
		return errors.New("รหัสผ่านต้องมีตัวอักษรใหญ่ ตัวอักษรเล็ก ตัวเลข และอักขระพิเศษอย่างน้อยตัวละ 1 ตัว")
	}
	return nil
}
