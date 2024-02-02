package usecases

import (
	"github.com/aerosystems/checkmail-service/internal/models"
	"github.com/google/uuid"
)

type DomainRepository interface {
	FindByName(name string) (*models.Domain, error)
	Create(domain *models.Domain) error
	Update(domain *models.Domain) error
	Delete(domain *models.Domain) error
	Count() (map[string]int, error)
	MatchEquals(name string) (*models.Domain, error)
	MatchEnds(name string) (*models.Domain, error)
}

type RootDomainRepository interface {
	FindById(id int) (*models.RootDomain, error)
	FindByName(name string) (*models.RootDomain, error)
	Create(rootDomain *models.RootDomain) error
	Update(rootDomain *models.RootDomain) error
	Delete(rootDomain *models.RootDomain) error
}

type FilterRepository interface {
	FindAll() (*[]models.Filter, error)
	FindById(id int) (*models.Filter, error)
	FindByProjectToken(projectToken string) (*models.Filter, error)
	Create(domain *models.Filter) error
	Update(domain *models.Filter) error
	Delete(domain *models.Filter) error
	MatchEquals(domainName, projectToken string) (*models.Filter, error)
	MatchEnds(domainName, projectToken string) (*models.Filter, error)
}

type ReviewRepository interface {
	Create(domainReview *models.Review) error
}

type ProjectRepository interface {
	GetProjectList(userUuid uuid.UUID) (*[]models.ProjectRPCPayload, error)
	GetProject(projectToken string) (*models.ProjectRPCPayload, error)
}
