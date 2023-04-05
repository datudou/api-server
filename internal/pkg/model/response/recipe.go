package response

type Recipe struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        int    `json:"cost"`
}
