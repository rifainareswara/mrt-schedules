package station

import (
	"encoding/json"
	"errors"
	"mrt-schedules/common/client"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
	GetStationById(id string) (StationResponse, error)
	GetStationSchedule(id string) (ScheduleResponse, error)
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
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, err
	}

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response, nil
}

func (s *service) GetStationById(id string) (StationResponse, error) {
	// First, get all stations
	stations, err := s.GetAllStations()
	if err != nil {
		return StationResponse{}, err
	}

	// Find the station with the matching ID
	for _, station := range stations {
		if station.Id == id {
			return station, nil
		}
	}

	return StationResponse{}, errors.New("station not found")
}

func (s *service) GetStationSchedule(id string) (ScheduleResponse, error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns/" + id

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return ScheduleResponse{}, err
	}

	var schedule Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return ScheduleResponse{}, err
	}

	if schedule.StationId == "" {
		return ScheduleResponse{}, errors.New("station not found")
	}

	// Extract next schedule times
	nextBundaranHI, nextLebakBulus, err := GetNextScheduleTimes(schedule)
	if err != nil {
		return ScheduleResponse{}, err
	}

	response := ScheduleResponse{
		StationId:          schedule.StationId,
		StationName:        schedule.StationName,
		BundaranHISchedule: nextBundaranHI,
		LebakBulusSchedule: nextLebakBulus,
	}

	return response, nil
}

// GetNextScheduleTimes returns the next departure times for both directions
func GetNextScheduleTimes(schedule Schedule) (nextBundaranHI, nextLebakBulus string, err error) {
	// Get parsed schedules
	bundaranHITimes, err := ConvertScheduleToTimeFormat(schedule.ScheduleBundaranHI)
	if err != nil {
		return "", "", err
	}

	lebakBulusTimes, err := ConvertScheduleToTimeFormat(schedule.ScheduleLebakBulus)
	if err != nil {
		return "", "", err
	}

	// Get current time
	now := time.Now()
	currentTimeStr := now.Format("15:04")

	// Find next Bundaran HI departure
	nextBundaranHI = ""
	for _, t := range bundaranHITimes {
		if t.Format("15:04") > currentTimeStr {
			nextBundaranHI = t.Format("15:04")
			break
		}
	}

	// Find next Lebak Bulus departure
	nextLebakBulus = ""
	for _, t := range lebakBulusTimes {
		if t.Format("15:04") > currentTimeStr {
			nextLebakBulus = t.Format("15:04")
			break
		}
	}

	return nextBundaranHI, nextLebakBulus, nil
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}

		parsedTime, err = time.Parse("15:04", trimmedTime)
		if err != nil {
			err = errors.New("invalid time format: " + trimmedTime)
			return nil, err
		}

		response = append(response, parsedTime)
	}

	return response, nil
}
