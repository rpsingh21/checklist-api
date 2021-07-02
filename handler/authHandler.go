package handler

import (
	"net/http"

	"github.com/rpsingh21/checklist-api/model"
	"github.com/rpsingh21/checklist-api/repository"
	"go.uber.org/zap"
)

// AuthHandler handels auth related api
type AuthHandler struct {
	logger   *zap.SugaredLogger
	userRopo *repository.UserRepository
}

// NewAuthHandler return new handler object
func NewAuthHandler(logger *zap.SugaredLogger, ur *repository.UserRepository) *AuthHandler {
	return &AuthHandler{logger: logger, userRopo: ur}
}

// Create New user
func (ah *AuthHandler) Create(rw http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := user.FromJSON(r.Body); err != nil {
		ah.logger.Errorf("Error to parse request body %v ", err)
		ErrorResponseWriter(rw, http.StatusInternalServerError, err)
		return
	}
	if err := ah.userRopo.Create(user); err != nil {
		ah.logger.Errorf("Faild to create new user %v", user)
		ErrorResponseWriter(rw, http.StatusInternalServerError, err)
		return
	}

	ResponseWriter(rw, http.StatusCreated, "", *user)
}

// Get List of all user
func (ah *AuthHandler) Get(rw http.ResponseWriter, r *http.Request) {

	ResponseWriter(rw, http.StatusOK, "Hello golang", nil)
}
