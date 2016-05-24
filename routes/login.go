package routes

import (
	"../models"
	"../repos"
	"../server/config"
	"../server/status"
	"../services"
	"../services/oauth2"
	"../services/oauth2/providers"
	"github.com/labstack/echo"
	"fmt"
	"crypto/rand"
)

var myProviders map[string]oauth2.IProvider

type LoginRoute interface {
	Get(echo.Context) error
	Callback(echo.Context) error
}

var Login LoginRoute

type loginRoute struct {
	LoginRoute
	usersRepository repos.UsersRepository
	googleAPI       services.GoogleAPI
}

func (this *loginRoute) Get(ctx echo.Context) error {
	provider := ctx.Param("provider")

	if p, ok := myProviders[provider]; ok {
		return ctx.Redirect(status.FOUND, p.RedirectUrl())
	}

	return echo.NewHTTPError(status.NOT_FOUND, "Unknown provider")
}

func (this *loginRoute) Callback(ctx echo.Context) error {
	code := ctx.QueryParam("code")

	if code == "" {
		return echo.NewHTTPError(status.BAD_REQUEST, "Must provide a code")
	}

	provider := ctx.Param("provider")

	if p, ok := myProviders[provider]; ok {
		token, err := p.ExchangeToken(code)

		if err != nil {
			return echo.NewHTTPError(status.INTERNAL_SERVER_ERROR, "Exchange token error")
		}

		profile, err := this.googleAPI.GetProfile(token.Token)
		user, err := this.usersRepository.FindByProviderId(profile.Id)
		randomToken := randToken()

		// TODO - JWT

		if err == nil {
			// update
			user.Token = token.Token
			user.JWT = randomToken
			this.usersRepository.Update(user)
		} else {
			// create
			this.usersRepository.Insert(
				&models.User{
					IdProvider: profile.Id,
					DisplayName: profile.DisplayName,
					ProfilePicture: profile.Image.Url,
					Email: profile.Emails[0].Value,
					Token: token.Token,
					JWT: randomToken,
				},
			)
		}

		cookie := new(echo.Cookie)
		cookie.SetName("gekko_jwt")
		cookie.SetValue(randomToken)

		ctx.SetCookie(cookie)
		return ctx.Redirect(status.FOUND, "/app")
	}

	return echo.NewHTTPError(status.NOT_FOUND, "Unknown provider")
}

func init() {
	myProviders = make(map[string]oauth2.IProvider)

	myProviders["google"] = providers.NewGoogle(
		config.Config.GoogleId,
		config.Config.GoogleSecret,
		"http://localhost:3000/login/google/callback",
	)

	myProviders["dropbox"] = providers.NewDropbox(
		config.Config.DropboxId,
		config.Config.DropboxSecret,
		"http://localhost:3000/login/dropbox/callback",
	)

	Login = &loginRoute{
		usersRepository: repos.NewUsersRepository(config.Config.Database), // TODO - fix this path
		googleAPI: services.NewGoogleAPI(),
	}
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
