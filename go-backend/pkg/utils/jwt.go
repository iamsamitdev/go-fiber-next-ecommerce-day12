package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// สร้างฟังก์ชันสำหรับสร้าง JWT Token
func GenerateJWT(userID, email, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// ฟังก์ชันสำหรับตรวจสอบ JWT Token
func ValidateJWT(tokenString string) (*Claims, error) {

	secret := os.Getenv("JWT_SECRET")

	// ตรวจสอบว่า secret ถูกตั้งค่าใน environment variable หรือไม่
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// ตรวจสอบว่ามีข้อผิดพลาดในการวิเคราะห์ token หรือไม่
	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่า token ถูกต้องและ claims เป็นประเภทที่คาดหวังหรือไม่
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	// ถ้า token ไม่ถูกต้องหรือ claims ไม่ตรงตามที่คาดหวัง
	return nil, jwt.ErrSignatureInvalid
}
