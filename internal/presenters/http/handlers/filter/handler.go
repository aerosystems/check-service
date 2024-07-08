package filter

import (
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/aerosystems/checkmail-service/internal/presenters/http/handlers"
	"time"
)

type Handler struct {
	*handlers.BaseHandler
	filterUsecase handlers.FilterUsecase
}

func NewHandler(
	baseHandler *handlers.BaseHandler,
	filterUsecase handlers.FilterUsecase,
) *Handler {
	return &Handler{
		BaseHandler:   baseHandler,
		filterUsecase: filterUsecase,
	}
}

type Filter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" example:"gmail.com"`
	Type      string    `json:"type" example:"whitelist"`
	Coverage  string    `json:"coverage" example:"equals"`
	CreatedAt time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

func ModelToFilter(filter models.Filter) Filter {
	return Filter{
		Id:        filter.Id,
		Name:      filter.Name,
		Type:      filter.Type,
		Coverage:  filter.Coverage,
		CreatedAt: filter.CreatedAt,
		UpdatedAt: filter.UpdatedAt,
	}
}

func ModelListToFilterList(filters []models.Filter) []Filter {
	filterList := make([]Filter, 0, len(filters))
	for _, model := range filters {
		filterList = append(filterList, ModelToFilter(model))
	}
	return filterList
}
