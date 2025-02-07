package station

import (
	"encoding/json"
	"github.com/rifainareswara/mrt-schedules.git/common/client"
	"golang.org/x/net/html/atom"
	"net/http"
	"time"
)

type Service interface {
	GetAllStations() (respons []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (respons []StationResponse, err error) {
	url := "https://wwww.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest("GET", url)
	if err != nil {
		return
	}
	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	for _, station := range stations {
		respons = append(respons, StationResponse{
			Id:   item.Itemid,
			Name: item.Name,
		})
	}

	return
}
