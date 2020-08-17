package vo

type Demo struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
