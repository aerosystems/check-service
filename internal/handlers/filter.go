package handlers

import (
	"errors"
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/aerosystems/checkmail-service/pkg/validators"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type TopDomainRequest struct {
	Name     string `json:"name" example:"gmail.com"`
	Type     string `json:"type" example:"whitelist"`
	Coverage string `json:"coverage" example:"equals"`
}

func (h *BaseHandler) CreateFilter(w http.ResponseWriter, r *http.Request) {
	return
}

// CreateFilterReview godoc
// @Summary create top domain
// @Tags topDomains
// @Accept  json
// @Produce application/json
// @Param comment body TopDomainRequest true "raw request body"
// @Success 201 {object} Response{data=models.Filter}
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/filter [post]
func (h *BaseHandler) CreateFilterReview(w http.ResponseWriter, r *http.Request) {
	xApiKey := r.Header.Get("X-Api-Key")
	var requestPayload TopDomainRequest
	if err := ReadRequest(w, r, &requestPayload); err != nil {
		WriteResponse(w, http.StatusUnprocessableEntity, NewErrorPayload(422201, "could not read request body", err))
		return
	}

	if requestPayload.Name == "" {
		err := errors.New("name does not exists or empty")
		_ = WriteResponse(w, http.StatusUnprocessableEntity, NewErrorPayload(422204, err.Error(), err))
		return
	}

	if requestPayload.Type == "" {
		err := errors.New("type does not exists or empty")
		_ = WriteResponse(w, http.StatusUnprocessableEntity, NewErrorPayload(422205, err.Error(), err))
		return
	}

	if requestPayload.Coverage == "" {
		err := errors.New("coverage does not exists or empty")
		_ = WriteResponse(w, http.StatusUnprocessableEntity, NewErrorPayload(422206, err.Error(), err))
		return
	}

	if err := validators.ValidateDomainTypes(requestPayload.Type); err != nil {
		_ = WriteResponse(w, http.StatusBadRequest, NewErrorPayload(400203, err.Error(), err))
		return
	}

	if err := validators.ValidateDomainCoverage(requestPayload.Coverage); err != nil {
		_ = WriteResponse(w, http.StatusBadRequest, NewErrorPayload(400203, err.Error(), err))
		return
	}

	if err := validators.ValidateDomainName(requestPayload.Name); err != nil {
		_ = WriteResponse(w, http.StatusBadRequest, NewErrorPayload(400204, err.Error(), err))
		return
	}

	arrDomain := strings.Split(requestPayload.Name, ".")
	root := arrDomain[len(arrDomain)-1]
	rootDomain, _ := h.rootDomainRepo.FindByName(root)
	if rootDomain == nil {
		err := errors.New("root domain does not exist")
		_ = WriteResponse(w, http.StatusBadRequest, NewErrorPayload(400205, err.Error(), err))
		return
	}

	newTopDomain := models.Filter{
		Name:         requestPayload.Name,
		Type:         requestPayload.Type,
		Coverage:     requestPayload.Coverage,
		ProjectToken: xApiKey,
	}

	if err := h.topDomainRepo.Create(&newTopDomain); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			_ = WriteResponse(w, http.StatusConflict, NewErrorPayload(409201, "top domain already exists", err))
			return
		}
		_ = WriteResponse(w, http.StatusInternalServerError, NewErrorPayload(500202, "could not create top domain", err))
		return
	}

	_ = WriteResponse(w, http.StatusCreated, NewResponsePayload("top domain created", newTopDomain))
	return
}
