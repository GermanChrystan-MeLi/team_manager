package domain

type Position struct {
	ID           string `json:"id"`
	PositionName string `json:"position_name"`
	XLocation    int    `json:"x_location"`
	YLocation    int    `json:"y_location"`
}
