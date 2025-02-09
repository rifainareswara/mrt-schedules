package station

import (
	"encoding/json"
	"github.com/rifainareswara/mrt-schedules.git/common/client"
	"net/http"
	"time"
)

type Service interface {
	GetAllStations() ([]StationResponse, error)
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

func (s *service) GetAllStations() ([]StationResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	response := make([]StationResponse, 0, len(stations))
	for _, station := range stations {
		response = append(response, StationResponse{
			Id:   station.Id, // Changed to match the existing struct field name
			Name: station.Name,
		})
	}

	return response, nil
}
