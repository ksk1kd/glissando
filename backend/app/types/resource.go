package types

type APIResource struct {
	ID   uint64 `json:"id"`
	Time uint64 `json:"time"`
}

type APIResources struct {
	Quantity int           `json:"quantity"`
	Items    []APIResource `json:"items"`
}
