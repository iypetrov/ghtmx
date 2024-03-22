package server

import (
	"github.com/IliyaYavorovPetrov/ghtmx/app/home/models"
	"net/http"
	"sync"
)

var (
	once sync.Once
	home *Home
)

func Server() *Home {
	once.Do(func() {
		home = &Home{}
	})
	return home
}

type Home struct {
}

func (h *Home) GetUserIp(r *http.Request) models.UserDataModel {
	userData := models.UserDataModel{
		Ip: r.RemoteAddr,
	}

	return userData
}
