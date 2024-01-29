package controller

import (
	"api-ent/constants"
	"api-ent/db"
	"api-ent/dto"
	"api-ent/ent/user"
	"api-ent/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := utils.ParseBody(r, &loginRequest)
	if utils.HandleErr(err, http.StatusBadRequest, w) {
		return
	}

	user, err := db.GetDb().User.Query().
		Where(user.NameEQ(loginRequest.Username)).
		Select(user.FieldPassword, user.FieldName).
		Only(r.Context())
	if utils.HandleErr(err, http.StatusNotFound, w) {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if utils.HandleErr(err, http.StatusUnauthorized, w) {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
	})

	tokenStr, err := token.SignedString([]byte(constants.JWT_KEY))
	if utils.HandleErr(err, http.StatusInternalServerError, w) {
		return
	}

	utils.Return(w, http.StatusOK, tokenStr)
}
