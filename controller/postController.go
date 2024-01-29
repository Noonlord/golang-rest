package controller

import (
	"api-ent/db"
	"api-ent/ent"
	"api-ent/ent/post"
	"api-ent/ent/user"
	"api-ent/utils"
	"net/http"
	"strconv"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	posts, err := db.Post.Query().Select(
		post.FieldID,
		post.FieldTitle,
		post.FieldBody,
	).WithAuthor(
		func(uq *ent.UserQuery) {
			uq.Select(user.FieldName)
		},
	).All(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	utils.Return(w, http.StatusOK, posts)
}

func GetAllTitles(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	titles, err := db.Post.Query().Select(
		post.FieldTitle,
	).All(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	utils.Return(w, http.StatusOK, titles)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	var post ent.Post
	err := utils.ParseBody(r, &post)
	if utils.HandleErr(err, http.StatusBadRequest, w) {
		return
	}

	userId, err := strconv.Atoi(r.Header.Get("X-UserId"))
	if utils.HandleErr(err, http.StatusBadRequest, w) {
		return
	}

	res, err := db.Post.Create().SetTitle(post.Title).SetBody(post.Body).SetAuthorID(userId).Save(r.Context())
	if utils.HandleErr(err, http.StatusInternalServerError, w) {
		return
	}

	utils.Return(w, http.StatusOK, res)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	id, err := strconv.Atoi(utils.GetParam(r, "id"))
	if utils.HandleErr(err, http.StatusBadRequest, w) {
		return
	}
	post, err := db.Post.Query().Where(post.IDEQ(id)).WithAuthor(
		func(uq *ent.UserQuery) {
			uq.Select(user.FieldName)
		},
	).Only(r.Context())
	if utils.HandleErr(err, http.StatusNotFound, w) {
		return
	}
	utils.Return(w, http.StatusOK, post)
}

func GetPostsByTitle(w http.ResponseWriter, r *http.Request) {
	db := db.GetDb()
	title := utils.GetParam(r, "title")
	posts, err := db.Post.Query().Where(post.TitleEQ(title)).WithAuthor(
		func(uq *ent.UserQuery) {
			uq.Select(user.FieldName)
		},
	).All(r.Context())
	if utils.HandleErr(err, http.StatusNotFound, w) {
		return
	}
	utils.Return(w, http.StatusOK, posts)
}
