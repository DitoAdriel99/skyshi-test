package entities

type TodoPayload struct {
	Title      string `json:"title"`
	ActivityID int64  `json:"activity_group_id"`
	IsActive   bool   `json:"isActive"`
	Priority   string `json:"priority"`
}

type Todo struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	ActivityID int64  `json:"activity_group_id"`
	IsActive   bool   `json:"isActive"`
	Priority   string `json:"priority"`
	UpdatedAt  string `json:"updatedAt"`
	CreatedAt  string `json:"createdAt"`
}
