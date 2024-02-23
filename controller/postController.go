package controller

import (
	"api-ent/utils"
	"net/http"
)

func Example(w http.ResponseWriter, r *http.Request) {

	posts := []string{"post1", "post2", "post3"}
	utils.Return(w, http.StatusOK, posts)
}
