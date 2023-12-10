package dto

type ItemRequest struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}

type ItemRequestID struct {
	ID int `json:"id" validate:"required"`
}

type ItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
