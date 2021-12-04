package service

type AuthService interface {
	Login(username string, password string) bool
	GetToken(username string, isAdmin bool) string
}

type authService struct {
	loginService LoginService
	jwtService   JWTService
}

func NewAuthService(loginService LoginService, jwtService JWTService) AuthService {
	return &authService{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (service *authService) Login(username string, password string) bool {
	isAuth := service.loginService.Login(username, password)

	return isAuth
}

func (service *authService) GetToken(username string, isAdmin bool) string {
	return service.jwtService.GenerateToken(username, isAdmin)
}
