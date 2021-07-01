package handler

import (
	"net/http"

	"github.com/rpsingh21/checklist-api/repository"
	"go.uber.org/zap"
)

// AuthHandler handels auth related api
type AuthHandler struct {
	logger   *zap.SugaredLogger
	userRepo *repository.UserRepository
}

// NewAuthHandler return new handler object
func NewAuthHandler(logger *zap.SugaredLogger, ur *repository.UserRepository) *AuthHandler {
	return &AuthHandler{logger: logger, userRepo: ur}
}

func (ah *AuthHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ah.logger.Infof("Auth API call with %s methos", r.Method)
	ResponseWriter(rw, http.StatusOK, "Hello golang", nil)
}

func (ah *AuthHandler) helloWold(rw http.ResponseWriter, r *http.Request) {
	ResponseWriter(rw, http.StatusOK, "Hello golang", nil)
}
