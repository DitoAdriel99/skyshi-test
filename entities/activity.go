package entities

type ActivityPayload struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type Activity struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}
