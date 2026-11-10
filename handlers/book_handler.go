package handlers

import (
	"encoding/json"
	"errors"
	"go-postgres-api/models"
	"go-postgres-api/repositories"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type BookHandler struct {
	Repo *repositories.BookRepository
}

func NewBookHandler(repository *repositories.BookRepository) *BookHandler {
	return &BookHandler{
		Repo: repository,
	}
}

func (h *BookHandler) KitoblarOlish(w http.ResponseWriter, r *http.Request) {
	books, err := h.Repo.GetAll(r.Context())
	if err != nil {
		log.Printf("GET /books: %v", err)
		writeError(w, http.StatusInternalServerError, "Kitoblarni olishda server xatosi❗️")
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func (h *BookHandler) IdBoyichaOlish(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id musbat bolish kk ")
		return
	}

	book, err := h.Repo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "Kitob topilmadi")
			return
		}

		log.Printf("GET /books/%d xatolik: %v", id, err)
		writeError(w, http.StatusInternalServerError, "Kitobni olishda server xatosi")
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func (h *BookHandler) YangiKitob(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Price  float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "Noto'g'ri JSON ma'lumotlari yuborildi")
		return
	}

	if strings.TrimSpace(input.Title) == "" ||
		strings.TrimSpace(input.Author) == "" ||
		input.Price <= 0 {
		writeError(w, http.StatusBadRequest, "Title, Author va musbat Price bo'lishi shart")
		return
	}

	newBook := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Price:  input.Price,
	}

	createdBook, err := h.Repo.Create(r.Context(), newBook)
	if err != nil {
		log.Printf("POST /books xatolik: %v", err)
		writeError(w, http.StatusInternalServerError, "Kitobni saqlashda server xatosi❗️")
		return
	}

	writeJSON(w, http.StatusCreated, createdBook)
}

func (h *BookHandler) KitobYangilash(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "id ni musbat kiriting")
		return
	}

	var input struct {
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Price  float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "notogri malumot kiritdngiz")
		return
	}

	if strings.TrimSpace(input.Title) == "" ||
		strings.TrimSpace(input.Author) == "" ||
		input.Price <= 0 {
		writeError(w, http.StatusBadRequest, "price musbatmas yoki qolgani toliqmas")
		return
	}

	bookToUpdate := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Price:  input.Price,
	}

	updatedBook, err := h.Repo.Update(r.Context(), id, bookToUpdate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "kitob yoq")
			return
		}

		log.Printf("PUT /books/%d xatolik: %v", id, err)
		writeError(w, http.StatusInternalServerError, "notogri malumot kiritdingiz")
		return
	}

	writeJSON(w, http.StatusOK, updatedBook)
}

func (h *BookHandler) KItobOchirish(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		if err != nil || id <= 0 {
			writeError(w, http.StatusBadRequest, "id + son bolsn")
			return
		}

		log.Printf("DELETE /books/%d xatolik: %v", id, err)
		writeError(w, http.StatusInternalServerError, "Kitobni o'chirish server xatosi ❗️")
		return
	}

	err = h.Repo.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "Kitob topilmadi")
			return
		}

		log.Printf("DELETE /books/%d xatolik: %v", id, err)
		writeError(w, http.StatusInternalServerError, "Kitobni o'chirishda server xatosi ❗️")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
