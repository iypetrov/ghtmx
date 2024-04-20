package ip

import (
	"net/http"
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

func (s *Server) GetUserIp(r *http.Request) (RequestIPModel, error) {
	model := RequestIPModel{
		ID:        uuid.New(),
		IP:        r.RemoteAddr,
		CreatedAt: time.Now(),
	}

	entity, err := s.storage.CreateRequestIPEntity(RequestIPModelToRequestIPEntity(model))
	if err != nil {
		return RequestIPModel{}, err
	}

	return RequestIPEntityToRequestIPModel(entity), nil
}
