package todos

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Completed bool   `json:"completed"`
	Order     *int   `json:"order"`
}

type TodoForCreate struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Completed *bool  `json:"completed"`
	Order     *int   `json:"order"`
}
