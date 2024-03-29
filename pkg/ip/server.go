package ip

import (
	"net/http"
)

type Server struct {
	storage *Storage
}

func NewServer(storage *Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) GetUserIp(r *http.Request) RequestIPModel {
	userData := RequestIPModel{
		IP: r.RemoteAddr,
	}

	_, err := s.storage.GetStatus()
	if err != nil {
		return RequestIPModel{}
	}

	return userData
}
