package handlers

import (
	"BE/config"
	"BE/controllers"
	"BE/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticleResponse struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var post models.Posts
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		controllers.SendErrorResponse(w, controllers.ResponseError{"ERROR", "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	if valid, err := controllers.ValidateArticle(&post); !valid {
		controllers.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if result := config.DB.Create(&post); result.Error != nil {
		controllers.SendErrorResponse(w, controllers.ResponseError{"ERROR", result.Error.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "SUCCESS",
		"message": "Article created successfully",
	})
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	var posts []models.Posts
	vars := mux.Vars(r)

	limit, err1 := strconv.Atoi(vars["limit"])
	offset, err2 := strconv.Atoi(vars["offset"])

	if err1 != nil || err2 != nil {
		response := map[string]interface{}{
			"status":  "ERROR",
			"message": "Invalid limit or offset",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	config.DB.Limit(limit).Offset(offset).Find(&posts)

	if len(posts) == 0 {
		response := map[string]interface{}{
			"status":  "ERROR",
			"message": "No articles found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	var articles []ArticleResponse
	for _, post := range posts {
		articles = append(articles, ArticleResponse{
			Title:    post.Title,
			Content:  post.Content,
			Category: post.Category,
			Status:   post.Status,
		})
	}

	// Kirim response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post models.Posts

	if result := config.DB.First(&post, vars["id"]); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ERROR",
			"message": "Article not found",
		})
		return
	}

	response := ArticleResponse{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post models.Posts

	if result := config.DB.First(&post, vars["id"]); result.Error != nil {
		controllers.SendErrorResponse(w, controllers.ResponseError{"ERROR", "Article not found"}, http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		controllers.SendErrorResponse(w, controllers.ResponseError{"ERROR", "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	if valid, err := controllers.ValidateArticle(&post); !valid {
		controllers.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	config.DB.Save(&post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "SUCCESS",
		"message": "Article updated successfully",
	})
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post models.Posts

	if result := config.DB.First(&post, vars["id"]); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ERROR",
			"message": "Article not found",
		})
		return
	}

	config.DB.Delete(&post)

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "SUCCESS",
		"message": "Article deleted successfully",
	})
}

