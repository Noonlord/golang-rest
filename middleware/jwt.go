package middleware

import (
	"api-ent/constants"
	"api-ent/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(constants.JWT_KEY), nil
		})

		utils.HandleErr(err, http.StatusInternalServerError, w)
		if !token.Valid {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		iat := claims["iat"].(float64)
		// 24 hours
		if int64(iat) < (time.Now().Unix() - 60*60*24) {
			http.Error(w, "Token expired", http.StatusUnauthorized)
		}

		userId := claims["id"].(float64)
		r.Header.Set("X-UserId", strconv.Itoa(int(userId)))
		next.ServeHTTP(w, r)
	})
}
