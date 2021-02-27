package infrastructure

import (
	"time"

	"github.com/SHU95/docker-env-go/domain"
	"github.com/dgrijalva/jwt-go"
)

const (
	secret  = "tukareteru"
	id      = "id"
	iatTime = "iat"
	expTime = "exp"
)

func createToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		id:      user.ID,
		iatTime: user.CreatedAt.Unix(),
		expTime: user.CreatedAt.Add(time.Hour * 1000).Unix(),
	})

	return token.SignedString([]byte(secret))
}

/*
func Login() echo.HandlerFunc {
	return func(context echo.Context) error {
		username := context.FormValue("username")
		password := context.FormValue("password")

		if username == "test" && password == "test" {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "test" //username
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix() //発行してからの有効期限

			t, err := token.SignedString([]byte("secret"))

			if err != nil {
				return err
			}
			return context.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}

		return echo.ErrUnauthorized

	}
}

func Restricted() echo.HandlerFunc {
	return func(context echo.Context) error {
		user := context.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		return context.String(http.StatusOK, "Welcome"+name+"!")
	}
}
*/
