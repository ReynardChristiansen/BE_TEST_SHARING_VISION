package controllers

import (
	"BE/models"
	"encoding/json"
	"net/http"
	"strings"
)

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ValidateArticle(post *models.Posts) (bool, ResponseError) {
	if len(strings.TrimSpace(post.Title)) < 20 {
		return false, ResponseError{"ERROR", "Title is required and must be at least 20 characters long"}
	}
	if len(strings.TrimSpace(post.Content)) < 200 {
		return false, ResponseError{"ERROR", "Content is required and must be at least 200 characters long"}
	}
	if len(strings.TrimSpace(post.Category)) < 3 {
		return false, ResponseError{"ERROR", "Category is required and must be at least 3 characters long"}
	}
	validStatuses := map[string]bool{"Publish": true, "Draft": true, "Thrash": true}
	if !validStatuses[post.Status] {
		return false, ResponseError{"ERROR", "Status must be either 'Publish', 'Draft', or 'Thrash'"}
	}

	return true, ResponseError{}
}

func SendErrorResponse(w http.ResponseWriter, err ResponseError, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
