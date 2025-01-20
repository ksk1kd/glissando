package types

type APIProject struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type APIProjects struct {
	Quantity int          `json:"quantity"`
	Items    []APIProject `json:"items"`
}
