package handlers

import (
	"net/http"

	"boilerplate/core"
)

// Info ...
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GetInfo ...
func GetInfo(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}
	res.SendOK(Info{
		Name:    "boilerplate",
		Version: "0.0.0",
	})
}
