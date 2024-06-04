package user

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"time"
)

func (u *Service) validateVerifyAccount(req domain.VerifyRequest) error {
	if req.UserName == "" || req.Code == "" {
		return myerror.ErrAuthInvalidDataRequest(nil)
	}
	return nil
}

func (u *Service) VerifyAccount(req domain.VerifyRequest) error {
	err := u.validateVerifyAccount(req)
	if err != nil {
		return err
	}

	user, err := u.userStorage().Get("user_name = ?", req.UserName)
	if err != nil {
		return myerror.ErrAuthNotFoundAccount(err)
	}

	if user.IsActive {
		return myerror.ErrAuthAccountHasVerified(nil)
	}

	if user.ExpCode.Before(time.Now()) {
		return myerror.ErrAuthExpiredCode(nil)
	}

	if user.ActiveCode != req.Code {
		return myerror.ErrAuthWrongCode(nil)
	}

	user.IsActive = true

	if err := u.userStorage().Update(user); err != nil {
		return myerror.ErrAuthUpdateAccount(err)
	}

	return nil
}
