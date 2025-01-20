package types

type APIMember struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type APIMembers struct {
	Quantity int         `json:"quantity"`
	Items    []APIMember `json:"items"`
}
