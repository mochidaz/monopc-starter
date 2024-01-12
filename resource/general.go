package resource

type QueryRequest struct {
	Query  string `json:"query"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
