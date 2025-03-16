package station

import (
	"encoding/json"
	"mrt-schedules/common/client"
	"net/http"
	"time"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
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

func (s *service) GetAllStations() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns/"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}
