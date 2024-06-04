package user

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/internal/logger"
	"garasystem/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (u *Service) validateCreateUser(req domain.RegisterUserReq) error {
	if req.UserName == "" || req.Phone == "" || req.Password == "" {
		return myerror.ErrInvalidRegister(nil)
	}

	err := u.userStorage().IsExisted(req.UserName, "", req.Phone)
	if err != nil {
		return myerror.ErrUserExisted(err)
	}

	return nil
}

func (u *Service) CreateUser(req domain.RegisterUserReq) error {
	if err := u.validateCreateUser(req); err != nil {
		return err
	}

	var (
		now              = time.Now()
		verificationCode = util.GenerateVerificationCode()
		expCode          = now.Add(util.EXPIRED_CODE_DURATION)
	)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return myerror.ErrCreateUserHashPassword(err)
	}

	user := domain.User{
		PhoneNumber: req.Phone,
		Password:    string(hashPassword),
		UserName:    req.UserName,
		ActiveCode:  verificationCode,
		ExpCode:     expCode,
	}

	err = u.userStorage().Create(&user)
	if err != nil {
		return myerror.ErrCreateUser(err)
	}

	go func() {
		err := u.awsService.SendSMS(user.PhoneNumber, user.ActiveCode)
		if err != nil {
			logger.Log.Error(err)
		}
	}()

	return err
}
