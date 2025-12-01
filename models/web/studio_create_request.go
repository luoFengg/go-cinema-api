package web

type StudioCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required,min=1,max=1000"`
}