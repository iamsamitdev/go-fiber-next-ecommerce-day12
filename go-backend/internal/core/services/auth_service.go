package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/services"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type authService struct {
	userRepo repositories.UserRepository
	roleRepo repositories.RoleRepository
}

func NewAuthService(userRepo repositories.UserRepository, roleRepo repositories.RoleRepository) services.AuthService {
	return &authService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *authService) Register(ctx context.Context, req *entities.RegisterRequest) (*entities.User, error) {
	// ตรวจสอบว่าอีเมลมีอยู่แล้วหรือไม่
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, errors.New("อีเมลนี้มีอยู่ในระบบแล้ว")
	}

	// ตรวจสอบความซับซ้อนของรหัสผ่าน
	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// หา role "user" เป็นค่าเริ่มต้น
	userRole, err := s.roleRepo.GetByName(ctx, "user")
	if err != nil {
		return nil, errors.New("ไม่พบบทบาทผู้ใช้")
	}

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address:   req.Address,
		Active:    true,
		RoleID:    userRole.ID,
	}

	if err := s.userRepo.Create(ctx, user, hashedPassword); err != nil {
		return nil, err
	}

	// ดึงข้อมูลผู้ใช้พร้อม role
	return s.userRepo.GetByID(ctx, user.ID)
}

func (s *authService) AdminRegister(ctx context.Context, req *entities.AdminRegisterRequest) (*entities.User, error) {
	// ตรวจสอบว่าอีเมลมีอยู่แล้วหรือไม่
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, errors.New("อีเมลนี้มีอยู่ในระบบแล้ว")
	}

	// ตรวจสอบความซับซ้อนของรหัสผ่าน
	if err := utils.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// แปลง string เป็น UUID
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return nil, errors.New("รูปแบบ role ID ไม่ถูกต้อง")
	}

	// ตรวจสอบว่า role มีอยู่
	if _, err := s.roleRepo.GetByID(ctx, roleID); err != nil {
		return nil, errors.New("ไม่พบบทบาทที่ระบุ")
	}

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address:   req.Address,
		Active:    true,
		RoleID:    roleID,
	}

	if err := s.userRepo.Create(ctx, user, hashedPassword); err != nil {
		return nil, err
	}

	// ดึงข้อมูลผู้ใช้พร้อม role
	return s.userRepo.GetByID(ctx, user.ID)
}

func (s *authService) Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error) {
	// ค้นหาผู้ใช้ตามอีเมล
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("อีเมลหรือรหัสผ่านไม่ถูกต้อง")
	}

	// ตรวจสอบว่าผู้ใช้ยังใช้งานอยู่หรือไม่
	if !user.Active {
		return nil, errors.New("บัญชีผู้ใช้ถูกระงับ")
	}

	// ดึงรหัสผ่านที่เข้ารหัสแล้ว
	hashedPassword, err := s.userRepo.GetPasswordHash(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	// ตรวจสอบรหัสผ่าน
	if !utils.CheckPassword(req.Password, hashedPassword) {
		return nil, errors.New("อีเมลหรือรหัสผ่านไม่ถูกต้อง")
	}

	// สร้าง JWT token
	token, err := utils.GenerateJWT(user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return nil, err
	}

	// สร้าง refresh token
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// บันทึก refresh token
	if err := s.userRepo.SetRefreshToken(ctx, user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &entities.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *entities.RefreshTokenRequest) (*entities.LoginResponse, error) {
	// ค้นหาผู้ใช้ตาม refresh token
	user, err := s.userRepo.GetByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, errors.New("refresh token ไม่ถูกต้อง")
	}

	// ตรวจสอบว่าผู้ใช้ยังใช้งานอยู่หรือไม่
	if !user.Active {
		return nil, errors.New("บัญชีผู้ใช้ถูกระงับ")
	}

	// สร้าง JWT token ใหม่
	token, err := utils.GenerateJWT(user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return nil, err
	}

	// สร้าง refresh token ใหม่
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// บันทึก refresh token ใหม่
	if err := s.userRepo.SetRefreshToken(ctx, user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &entities.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

func (s *authService) Logout(ctx context.Context, userID uuid.UUID) error {
	// ลบ refresh token
	return s.userRepo.SetRefreshToken(ctx, userID, "")
}

func (s *authService) ChangePassword(ctx context.Context, userID uuid.UUID, req *entities.ChangePasswordRequest) error {
	// ดึงรหัสผ่านปัจจุบัน
	hashedPassword, err := s.userRepo.GetPasswordHash(ctx, userID)
	if err != nil {
		return err
	}

	// ตรวจสอบรหัสผ่านเก่า
	if !utils.CheckPassword(req.OldPassword, hashedPassword) {
		return errors.New("รหัสผ่านเก่าไม่ถูกต้อง")
	}

	// ตรวจสอบความซับซ้อนของรหัสผ่านใหม่
	if err := utils.ValidatePassword(req.NewPassword); err != nil {
		return err
	}

	// เข้ารหัสรหัสผ่านใหม่
	newHashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// อัพเดทรหัสผ่าน
	return s.userRepo.UpdatePassword(ctx, userID, newHashedPassword)
}

func (s *authService) ForgotPassword(ctx context.Context, req *entities.ForgotPasswordRequest) error {
	// ตรวจสอบว่าอีเมลมีอยู่ในระบบหรือไม่
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err != nil {
		return errors.New("ไม่พบอีเมลในระบบ")
	}

	// สร้าง reset token
	resetToken, err := s.generateResetToken()
	if err != nil {
		return err
	}

	// บันทึก reset token
	if err := s.userRepo.SetResetToken(ctx, req.Email, resetToken); err != nil {
		return err
	}

	// TODO: ส่งอีเมลพร้อม reset token ให้ผู้ใช้
	// ในการใช้งานจริงควรส่งอีเมลแทนการ return token

	return nil
}

func (s *authService) ResetPassword(ctx context.Context, req *entities.ResetPasswordRequest) error {
	// ค้นหาผู้ใช้ตาม reset token
	user, err := s.userRepo.GetByResetToken(ctx, req.Token)
	if err != nil {
		return errors.New("token ไม่ถูกต้องหรือหมดอายุแล้ว")
	}

	// ตรวจสอบความซับซ้อนของรหัสผ่านใหม่
	if err := utils.ValidatePassword(req.NewPassword); err != nil {
		return err
	}

	// เข้ารหัสรหัสผ่านใหม่
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// อัพเดทรหัสผ่าน
	if err := s.userRepo.UpdatePassword(ctx, user.ID, hashedPassword); err != nil {
		return err
	}

	// ลบ reset token
	return s.userRepo.ClearResetToken(ctx, user.ID)
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*entities.User, error) {
	// ตรวจสอบ JWT token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	// แปลง userID เป็น UUID
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.New("รูปแบบ user ID ไม่ถูกต้อง")
	}

	// ดึงข้อมูลผู้ใช้
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("ไม่พบผู้ใช้")
	}

	// ตรวจสอบว่าผู้ใช้ยังใช้งานอยู่หรือไม่
	if !user.Active {
		return nil, errors.New("บัญชีผู้ใช้ถูกระงับ")
	}

	return user, nil
}

func (s *authService) generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *authService) generateResetToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
