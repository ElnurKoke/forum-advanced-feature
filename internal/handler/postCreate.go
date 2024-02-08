package handler

import (
	"forum/internal/models"
	"log"
	"net/http"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		h.ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userValue := r.Context().Value("user")
	if userValue == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	user, ok := userValue.(models.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if !user.IsAuth {
		h.ErrorPage(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	categories, err := h.Service.ServicePostIR.GetCategories()
	if err != nil {
		h.ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		title := r.FormValue("title")
		description := r.FormValue("description")
		categories := r.Form["category"]
		if len(categories) == 0 {
			h.ErrorPage(w, "INVALID CATEGORY, please select existing categories ", http.StatusBadRequest)
			return
		}
		for _, category := range categories {
			if len(category) == 0 || len(category) >= 40 {
				h.ErrorPage(w, "INVALID CATEGORY, category should be shorter than 35 symbols and not empty", http.StatusBadRequest)
				return
			}
		}
		if len(description) > 600 || len(description) == 0 {
			h.ErrorPage(w, "description should be shorter than 400 symbols and not empty", http.StatusBadRequest)
			return
		}
		if len(title) == 0 || len(title) >= 80 {
			h.ErrorPage(w, "INVALID TITLE, title should be shorter than 35 symbols and not empty", http.StatusBadRequest)
			return
		}
		if err := h.Service.ServicePostIR.CreatePost(models.Post{
			Title:       title,
			Description: description,
			Category:    categories,
			Author:      user.Username,
		}); err != nil {
			h.ErrorPage(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case http.MethodGet:
		if err := h.Temp.ExecuteTemplate(w, "postCreate.html", categories); err != nil {
			h.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		h.ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
