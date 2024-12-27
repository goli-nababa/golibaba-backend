package services

import (
	"context"
	"errors"
	"user_service/api/http/handlers/helpers"
	"user_service/api/http/types"
	"user_service/app"
	"user_service/config"
)

var (
	/*	ErrUserOnCreate      = userService.ErrUserOnCreate
		ErrUserNotFound      = userService.ErrUserNotFound
		ErrUserAlreadyExists = userService.ErrUserAlreadyExists*/
	ErrCreatingToken   = errors.New("cannot create token")
	ErrBirthdayInvalid = errors.New("birthday is invalid")
)

type AccountService struct {
	/*	svc                              userPort.Service
		authCache                        *cache.ObjectCache[*presenter.LoginCacheSession]
		emailService                     email.Adapter*/
	authSecret                       string
	expMin, refreshExpMin, otpTtlMin uint
}

func NewAccountService(
	/*	svc userPort.Service,
		cacheService cache.Provider,
		emailService email.Adapter,*/
	authSecret string,
	expMin, refreshExpMin, otpTtlMin uint,
) *AccountService {
	return &AccountService{
		/*		svc:           svc,
				authCache:     cache.NewJsonObjectCache[*presenter.LoginCacheSession](cacheService, "auth."),
				emailService:  emailService,*/
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
		otpTtlMin:     otpTtlMin,
	}
}

func AccountServiceGetter(appContainer app.App, cfg config.ServerConfig) helpers.ServiceGetter[*AccountService] {
	return func(ctx context.Context) *AccountService {
		return NewAccountService(
			/*			appContainer.UserService(ctx),
						appContainer.Cache(),
						appContainer.MailService(),*/
			cfg.Secret,
			cfg.AuthExpirationMinutes,
			cfg.AuthRefreshMinutes,
			cfg.OtpTtlMinutes,
		)
	}
}

func (as *AccountService) Login(c context.Context, req types.LoginRequest) (*types.LoginResponse, error) {
	/*	user, err := as.svc.GetUserByUsernamePassword(c, req.Email, req.Password)

		if err != nil {
			return nil, err
		}

		code, err := helpers.GenerateOTP()

		if err != nil {
			return nil, errors.New("error generating OTP")
		}

		log.Println("OTP sent for user", user.ID, "code:", code)

		err = as.emailService.SendText(
			req.Email,
			fmt.Sprintf("GoliPors OTP code for %s", req.Email),
			fmt.Sprintf("GoliPors OTP code: %s", code),
		)

		if err != nil {
			log.Println("Error while sending otp:", err)
		}

		reqUUID := uuid.New()

		err = as.authCache.Set(
			c, strconv.Itoa(int(user.ID)),
			time.Minute*time.Duration(as.otpTtlMin),
			&presenter.LoginCacheSession{
				SessionID: reqUUID,
				UserID:    user.ID,
				Code:      code,
			},
		)

		if err != nil {
			return nil, err
		}

		return &types.LoginResponse{
			Code:      code,
			SessionId: reqUUID,
		}, nil*/

	return nil, nil
}

func (as *AccountService) VerifyOtp(c context.Context, req types.VerifyOTPRequest) (*types.VerifyOTPResponse, error) {
	/*user, err := as.svc.GetUserByEmail(c, req.Email)

	if err != nil {
		return nil, err
	}

	authSession, err := as.authCache.Get(c, strconv.Itoa(int(user.ID)))

	if err != nil {
		return nil, err
	}

	if authSession == nil ||
		authSession.UserID <= 0 ||
		authSession.Code != req.Code ||
		authSession.SessionID != req.SessionId ||
		authSession.UserID != user.ID {
		return nil, ErrUserNotFound
	}

	err = as.authCache.Del(c, strconv.Itoa(int(user.ID)))

	if err != nil {
		return nil, err
	}

	var (
		authExp    = time.Now().Add(time.Minute * time.Duration(as.expMin))
		refreshExp = time.Now().Add(time.Minute * time.Duration(as.refreshExpMin))
	)

	accessToken, err := jwt2.CreateToken([]byte(as.authSecret), jwt2.GenerateUserClaims(user, authExp))
	refreshToken, err := jwt2.CreateToken([]byte(as.authSecret), jwt2.GenerateUserClaims(user, refreshExp))

	if err != nil {
		return nil, ErrCreatingToken
	}

	return &types.VerifyOTPResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil*/

	return nil, nil
}

func (as *AccountService) Register(c context.Context, req types.RegisterRequest) error {
	/*newU, err := presenter.RegisterRequestToUserDomain(req)

	if err != nil {
		return ErrBirthdayInvalid
	}

	_, err = as.svc.CreateUser(c, newU)

	if err != nil {
		return err
	}*/

	return nil
}
