package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/monorepo/fetch/commons/dtos"
	"net/http"
	"strings"
	"time"
)

const SecretKey = "jds"

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var jwtStruct = dtos.TokenUser{}
		token := c.Request().Header.Get("authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dtos.StandardResponse{
				Code:    http.StatusUnauthorized,
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: http.StatusText(http.StatusUnauthorized),
			})
		}

		if !strings.Contains(token, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, dtos.StandardResponse{
				Code:    http.StatusUnauthorized,
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: "Invalid Token",
			})
		}

		split := strings.Split(token, " ")[1]
		tokenJwt, err := jwt.Parse(split, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", t.Header["alg"]))
			}
			return []byte(SecretKey), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dtos.StandardResponse{
				Code:    http.StatusUnauthorized,
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: "Expired",
				Error:   err.Error(),
			})
		}

		claims, ok := tokenJwt.Claims.(jwt.MapClaims)
		if ok && tokenJwt.Valid {
			tmp := claims["sub"].(map[string]interface{})
			tmpByte, _ := json.Marshal(tmp)
			_ = json.Unmarshal(tmpByte, &jwtStruct)
		}
		fmt.Println(jwtStruct)

		if !time.Now().After(jwtStruct.Exp) {
			return c.JSON(http.StatusUnauthorized, dtos.StandardResponse{
				Code:    http.StatusUnauthorized,
				Status:  http.StatusText(http.StatusUnauthorized),
				Message: "Expired",
				Error:   err.Error(),
			})
		}

		c.Set("jwt", jwtStruct)

		return next(c)
	}
}
