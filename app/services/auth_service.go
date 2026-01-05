package services

import (
	"context"
	"donasitamanzakattest/app/enums"
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/app/web"
	"donasitamanzakattest/config"
	pkgjwt "donasitamanzakattest/pkg/jwt"
	"donasitamanzakattest/pkg/util"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type IAuthService interface {
	RegistrationService(req *web.RegistrationRequest) (*web.Session, error)
	LoginService(req *web.LoginRequest) (*web.Session, error)
	PrivateGetProfileService() (*models.WpUserProfile, error)
}

type AuthService struct {
	*Service
}

func NewAuthService(ctx context.Context, repo *repositories.RepositoryContext, cfg *config.Config) IAuthService {
	return &AuthService{
		Service: NewService(ctx, repo, cfg),
	}
}

func (s *AuthService) RegistrationService(req *web.RegistrationRequest) (*web.Session, error) {
	var session *web.Session
	var err error

	var register *models.WpDjaRegister
	var wpUser *models.WpUsers
	var wpDjaUser *models.WpDjaUsers
	var wpDjaVerificationDetail *models.WpDjaVerificationDetails

	timeNow := time.Now()
	rCode := util.GenerateRandomString(12)
	status := 0

	hashPw, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "hash password").WithStatusCode(http.StatusInternalServerError)
	}

	hashPw = "$wp" + hashPw

	register = &models.WpDjaRegister{
		RNamaLengkap: &req.FullName,
		REmail:       &req.Email,
		RWhatsapp:    &req.NoHp,
		RPassword:    &hashPw,
		RCode:        &rCode,
		RStatus:      &status,
		CreatedAt:    &timeNow,
	}

	userLogin := strings.ToLower(strings.Join(strings.Split(strings.Trim(*register.RNamaLengkap, " "), " "), ""))

	wpUser = &models.WpUsers{
		UserLogin:         userLogin,
		UserPass:          hashPw,
		UserNicename:      userLogin,
		UserEmail:         *register.REmail,
		UserURL:           "",
		UserRegistered:    timeNow,
		UserActivationKey: *register.RCode,
		UserStatus:        0,
		DisplayName:       *register.RNamaLengkap,
	}

	randId := "u_" + util.GenerateRandomString(8)
	userVerification := enums.USER_VERIFICATION_NO_STATUS
	userBio := ""

	wpDjaUser = &models.WpDjaUsers{
		UserID:           nil,
		UserRandid:       &randId,
		UserType:         nil,
		UserVerification: &userVerification,
		UserBio:          &userBio,
		UserWa:           register.RWhatsapp,
		UserProvinsi:     nil,
		UserKabkota:      nil,
		UserKecamatan:    nil,
		UserProvinsiID:   nil,
		UserKabkotaID:    nil,
		UserKecamatanID:  nil,
		UserAlamat:       nil,
		UserBankName:     nil,
		UserBankNo:       nil,
		UserBankAn:       nil,
		UserPpImg:        nil,
		UserCoverImg:     nil,
		CreatedAt:        nil,
		UserAnonimF:      nil,
		UserCommissionF:  nil,
		UserSapaan:       nil,
	}

	wpDjaVerificationDetail = &models.WpDjaVerificationDetails{
		CreatedAt: &timeNow,
	}

	err = s.repository.RegistrationRepository(s.ctx, register, wpUser, wpDjaUser, wpDjaVerificationDetail)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "registration").WithStatusCode(http.StatusInternalServerError)
	}

	payloadJwt := pkgjwt.IssueJwtPayload{
		Id:       0,
		Subject:  wpUser.UserLogin,
		Role:     "",
		Lifetime: 3600,
		Email:    *register.REmail,
		FullName: *register.RNamaLengkap,
		Code:     *register.RCode,
		Status:   enums.GetUserVerification(status),
	}

	jwtAdapter := pkgjwt.NewJwtAdapter("issue_jwt", "secret_jwt")

	session, err = jwtAdapter.IssueJwt(&payloadJwt)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "issue jwt").WithStatusCode(http.StatusInternalServerError)
	}

	return session, err
}

func (s *AuthService) LoginService(req *web.LoginRequest) (*web.Session, error) {
	var session *web.Session
	var err error

	var wpUser *models.WpUsersLogin

	wpUser, err = s.repository.LoginRepository(s.ctx, req)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "login").WithStatusCode(http.StatusInternalServerError)
	}
	log.Debug().Interface("wpUser", wpUser).Msg("data wp user login repository")

	if wpUser == nil {
		return nil, helpers.NewErrorTrace(fmt.Errorf("user not found"), "login").WithStatusCode(http.StatusUnauthorized)
	}

	validPW := util.CheckPassword(wpUser.Password, req.Password)
	if !validPW {
		return nil, helpers.NewErrorTrace(fmt.Errorf("email or password incorrect"), "login").WithStatusCode(http.StatusUnauthorized)
	}

	payloadJwt := pkgjwt.IssueJwtPayload{
		Id:       0,
		Subject:  wpUser.UserLogin,
		Role:     "",
		Lifetime: 3600,
		Email:    wpUser.Email,
		FullName: wpUser.FullName,
		Code:     "",
		Status:   enums.GetUserVerification(wpUser.Status),
	}

	jwtAdapter := pkgjwt.NewJwtAdapter("issue_jwt", "secret_jwt")

	session, err = jwtAdapter.IssueJwt(&payloadJwt)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "issue jwt").WithStatusCode(http.StatusInternalServerError)
	}

	return session, err
}

func (s *AuthService) PrivateGetProfileService() (*models.WpUserProfile, error) {
	var wpUserProfile *models.WpUserProfile
	var wpUserSession *models.WpUserSession
	var err error

	wpUserSession, err = s.GetSessionService()
	if err != nil {
		return nil, helpers.NewErrorTrace(fmt.Errorf("harap melakukan login terlebih dahulu"), "session").WithStatusCode(http.StatusUnauthorized)
	}

	wpUserProfile, err = s.repository.PrivateGetProfileRepository(wpUserSession)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "profile").WithStatusCode(http.StatusInternalServerError)
	}

	return wpUserProfile, err
}
