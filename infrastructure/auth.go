package infrastructure

import (
	"fmt"
	"text/template/parse"
	"time"

	"github.com/SHU95/docker-env-go/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

const (
	secret  = "tukareteru"
	idKey  = "id"
	iatKey = "iat"
	expTime = "exp"
)

//token 発行
func CreateToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		idKey:      user.ID,
		iatKey: user.CreatedAt.Unix(),
		expTime: user.CreatedAt.Add(time.Hour * 1000).Unix(),
	})

	return token.SignedString([]byte(secret))
}

//バリデーションチェック

func VerifyToken(tokenStr string)(domain.User error){

	parseToken, err := jwt.Parse(tokenStr, func(token *jwt.Token)(interface{}, error){

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method")
			return nil ,err
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.Claims)

	if !ok {
		return nil fmt.Errorf("claims ヾ(｡>﹏<｡)ﾉ")
	}
	return domain.User, nil

	id, ok := claims[idKey].(float64)
	if !ok {
		return nil fmt.Errorf("id ヾ(｡>﹏<｡)ﾉ")
	}

	iatTime, ok := claims[iatKey].(float64)
	if !ok{
		return nil fmt.Errorf("iat ヾ(｡>﹏<｡)ﾉ")
	}
	return (domain.User{
		ID: uint(id),
		CreatedAt: time.Unix(int64(iatTime), 0),
	}, nil);
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
