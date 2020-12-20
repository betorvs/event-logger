package domain

//Event represents the Event
type Event struct {
	Name    string            `json:"name"`
	Kind    string            `json:"kind"`
	Labels  map[string]string `json:"labels"`
	Message string            `json:"message"`
	Source  string            `json:"source"`
	Status  int               `json:"status"`
}
