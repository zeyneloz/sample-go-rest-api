package books

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/core"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
)

func authorCreateHandler(w http.ResponseWriter, c *core.Context) error {
	// Read request body.
	body, err := ioutil.ReadAll(c.Body)
	if err != nil {
		return errors.NewInternalError(err, "Body read error")
	}
	// Parse body as json.
	var schema authorCreateSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return errors.NewInternalError(err, "Body parse error")
	}
	// Validate reqeust body.
	if err = schema.Validate(); err != nil {
		return err
	}

	// Create user.
	author := models.Author{
		Name:     schema.Name,
		LastName: schema.LastName,
	}
	if err = db.Instance.Create(&author).Error; err != nil {
		return errors.NewInternalError(err, "Author create error")
	}

	responseSchema := authorResponseSchema{
		ID:       author.ID,
		Name:     author.Name,
		LastName: author.LastName,
	}
	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}

func authorListHandler(w http.ResponseWriter, c *core.Context) error {
	var authors []models.Author
	err := db.Instance.Find(&authors).Error

	if err != nil {
		return errors.NewInternalError(err, "DB error")
	}

	responseSchema := make([]authorResponseSchema, len(authors))
	for i, author := range authors {
		responseSchema[i] = authorResponseSchema{
			ID:       author.ID,
			Name:     author.Name,
			LastName: author.LastName,
		}
	}

	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}

func bookCreateHandler(w http.ResponseWriter, c *core.Context) error {
	// Read request body.
	body, err := ioutil.ReadAll(c.Body)
	if err != nil {
		return errors.NewInternalError(err, "Body read error")
	}
	// Parse body as json.
	var schema bookCreateSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return errors.NewInternalError(err, "Body parse error")
	}
	// Validate reqeust body.
	if err = schema.Validate(); err != nil {
		return err
	}

	book := models.Book{
		Name:        schema.Name,
		AuthorID:    schema.Author,
		CreatedByID: c.User.ID,
	}

	if err = db.Instance.Create(&book).Error; err != nil {
		return errors.NewInternalError(err, "Book create error")
	}

	// Create response schema
	responseSchema := bookCreateResponseSchema{
		ID:        book.ID,
		Name:      book.Name,
		Author:    book.AuthorID,
		CreatedBy: book.CreatedByID,
	}
	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}

func bookDetailHandler(w http.ResponseWriter, c *core.Context) error {
	bookID := chi.URLParam(c.Request, "bookID")
	var book models.Book
	err := db.Instance.Preload("Author").Preload("CreatedBy").First(&book, bookID).Error

	if gorm.IsRecordNotFoundError(err) {
		return errors.NewHTTPError(http.StatusNotFound, "Book not found", map[string]string{})
	}
	if err != nil {
		return errors.NewInternalError(err, "DB error")
	}

	responseSchema := bookDetailSchema{
		ID:   book.ID,
		Name: book.Name,
		Author: authorResponseSchema{
			ID:       book.Author.ID,
			Name:     book.Author.Name,
			LastName: book.Author.LastName,
		},
		CreatedBy: userDetailSchema{
			ID:   book.CreatedBy.ID,
			Name: book.CreatedBy.Name,
		},
	}
	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}
