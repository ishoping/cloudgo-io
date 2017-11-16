package service

import (
	"net/http"

	"github.com/unrolled/render"
)

// InfoHandler .
func InfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		var sex string

		if req.Form["sex"][0] == "1" {
			sex = "男"
		} else if req.Form["sex"][0] == "2" {
			sex = "女"
		}

		formatter.HTML(w, http.StatusOK, "table", struct {
			Name   string
			Sex    string
			Height string
			Weight string
		}{Name: req.Form["name"][0], Sex: sex, Height: req.Form["height"][0], Weight: req.Form["weight"][0]})
	}
}
