package dto

type PlayGameRequest struct {
	SelectedInput *int `json:"selected_input" validate:"required"`
}

type PlayGameResponse struct {
	Reward  float64 `json:"reward"`
	Message string  `json:"message,omitempty"`
}
