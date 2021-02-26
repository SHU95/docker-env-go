package infrastructure

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Login() echo.HandlerFunc {
	return func(context echo.Context) error {
		username := context.FormValue("username")
		password := context.FormValue("password")

		if username == "test" && password == "test" {
			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "test" //username
			claims["admin"] = true
			claims["iat"] = time.Now().Unix()                     //発行日時
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
