package types

type ProjectAdd struct {
	Name string `json:"name" binding:"required"`
}
