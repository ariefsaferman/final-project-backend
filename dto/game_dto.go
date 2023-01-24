package dto

type PlayGameRequest struct {
	SelectedInput *int `json:"selected_input" validate:"required"`
}

type PlayGameResponse struct {
	Amount float64 `json:"amount"`
}
