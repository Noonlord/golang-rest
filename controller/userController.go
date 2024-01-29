package controller

import (
	"api-ent/db"
	"api-ent/ent"
	"api-ent/ent/user"
	"api-ent/utils"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	users, err := db.User.Query().Select(
		user.FieldID,
		user.FieldName,
	).All(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	utils.Return(w, http.StatusOK, users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	var user ent.User
	err := utils.ParseBody(r, &user)
	if utils.HandleErr(err, http.StatusBadRequest, w) {
		return
	}

	passwordBytes := []byte(user.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(passwordBytes, 10)
	if utils.HandleErr(err, http.StatusInternalServerError, w) {
		return
	}

	user.Password = string(hashedPass)
	fmt.Println(user.Password)

	res, err := db.User.Create().SetName(user.Name).SetPassword(user.Password).Save(r.Context())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "This username is taken!", http.StatusBadRequest)
		return
	}
	utils.Return(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	id, err := strconv.Atoi(utils.GetParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// get user and their posts
	user, err := db.User.Query().Where(user.IDEQ(id)).WithPosts().Only(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	utils.Return(w, http.StatusOK, user)
}
