package request

type CreateRecipeParam struct {
	Title       string `json:"title" binding:"required"`
	MakingTime  string `json:"making_time" binding:"required"`
	Serves      string `json:"serves" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
	Cost        int    `json:"cost" binding:"required"`
}

type UpdateRecipeParam struct {
	Title       string `json:"title" binding:"omitempty"`
	MakingTime  string `json:"making_time" binding:"omitempty"`
	Serves      string `json:"serves" binding:"omitempty"`
	Ingredients string `json:"ingredients" binding:"omitempty"`
	Cost        int    `json:"cost" binding:"omitempty"`
}
