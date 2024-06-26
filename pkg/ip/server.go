package ip

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Server struct {
	storage *Storage
}

func NewServer(storage *Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) CreateRequestIPModel(r *http.Request) (RequestIPModel, error) {
	parts := strings.Split(r.RemoteAddr, ":")
	model := RequestIPModel{
		ID:        uuid.New(),
		IP:        parts[0],
		CreatedAt: time.Now(),
	}

	entity, err := s.storage.CreateRequestIPEntity(RequestIPModelToRequestIPEntity(model))
	if err != nil {
		return RequestIPModel{}, err
	}

	return RequestIPEntityToRequestIPModel(entity), nil
}

func (s *Server) GetStatsIPModels() ([]StatsIPModel, error) {
	var models []StatsIPModel
	entities, err := s.storage.GetStatsIPEntities()
	if err != nil {
		return models, err
	}

	models = StatsIPEntitiesToStatsIPModels(entities)
	return models, nil
}
