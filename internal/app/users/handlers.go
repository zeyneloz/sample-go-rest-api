package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/config"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/core"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/models"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/utils/auth"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
	"github.com/zeyneloz/sample-go-rest-api/pkg/crypt"
)

func registerHandler(w http.ResponseWriter, c *core.Context) error {
	// Read request body.
	body, err := ioutil.ReadAll(c.Body)
	if err != nil {
		return errors.NewInternalError(err, "Body read error")
	}
	// Parse body as json.
	var schema registerSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return errors.NewInternalError(err, "Body parse error")
	}
	// Validate reqeust body.
	if err = schema.Validate(); err != nil {
		return err
	}

	// Create user.
	user := models.User{
		Email: schema.Email,
		Name:  schema.Name,
	}
	if err = user.SetPassword(schema.Password); err != nil {
		return err
	}
	if err = db.Instance.Create(&user).Error; err != nil {
		return errors.NewInternalError(err, "User create error")
	}

	// Return 201 Created.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return nil
}

func loginHandler(w http.ResponseWriter, c *core.Context) error {
	// Read request body.
	body, err := ioutil.ReadAll(c.Body)
	if err != nil {
		return errors.NewInternalError(err, "Body read error")
	}
	// Parse body as json.
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return errors.NewInternalError(err, "Body parse error")
	}
	// Validate reqeust body.
	if err = schema.Validate(); err != nil {
		return err
	}

	// Get the user by email.
	var user models.User
	err = db.Instance.Where("email = ?", schema.Email).First(&user).Error

	// Check if user really exsists with that email.
	if gorm.IsRecordNotFoundError(err) {
		return errors.NewHTTPValidationError(map[string]string{
			"email": "User wtih this email is not found.",
		})
	}

	// If there is any other error, return internal error.
	if err != nil {
		return errors.NewInternalError(err, "GORM error")
	}

	// Check if the password matches.
	if !user.CheckPassword(schema.Password) {
		return errors.NewHTTPValidationError(map[string]string{
			"password": "Wrong password.",
		})
	}

	cnf := config.GetConfig()
	jwtData := auth.UserToJWT(user)

	// Generate new JWT token for the user.
	token, err := crypt.GenerateJWT(cnf.SecretKey, cnf.JwtExpire, jwtData)

	if err != nil {
		return errors.NewInternalError(err, "JWT generate error")
	}

	responseSchema := loginResponseSchema{
		Token: token,
	}
	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}

func userDetailHandler(w http.ResponseWriter, c *core.Context) error {
	userID := chi.URLParam(c.Request, "userID")
	myID := strconv.Itoa(c.User.ID)

	// A user can only get their detail, not other users detail.
	if myID != userID {
		return errors.NewHTTPError(http.StatusUnauthorized, "Unauthorized", map[string]string{})
	}

	responseSchema := userDetailResponseSchema{
		ID:    c.User.ID,
		Email: c.User.Email,
		Name:  c.User.Name,
	}
	response, err := json.Marshal(responseSchema)
	if err != nil {
		return errors.NewInternalError(err, "JSON marshall error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}
