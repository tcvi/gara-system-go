package user

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"time"
)

func (u *Service) validateResendCode(req domain.ResendCodeRequest) error {
	if req.UserName == "" {
		return myerror.ErrInvalidVerify(nil)
	}

	return nil
}

func (u *Service) ResendCode(req domain.ResendCodeRequest) error {
	if err := u.validateResendCode(req); err != nil {
		return err
	}

	user, err := u.userStorage().Get("user_name = ?", req.UserName)
	if err != nil {
		return myerror.ErrAuthNotFoundAccount(err)
	}

	if user.IsActive {
		return myerror.ErrAuthAccountHasVerified(nil)
	}

	now := time.Now()

	if user.ExpCode.Sub(now) > util.LIMIT_RESEND_CODE_DURATION {
		return myerror.ErrAuthResendCodeLater(nil)
	}

	user.ActiveCode = util.GenerateVerificationCode()
	user.ExpCode = now.Add(util.EXPIRED_CODE_DURATION)

	err = u.awsService.SendSMS(user.PhoneNumber, user.ActiveCode)
	if err != nil {
		return myerror.ErrAuthResendCode(err)
	}

	return u.userStorage().Update(user)
}
