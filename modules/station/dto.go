package station

type Station struct {
	Id   int    `json:"itemid"`
	Name string `json:"name"`
}

type StationResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	StationId          int    `json:"nid"`
	StationName        string `json:"title"`
	SceduleBundaranHI  string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}
