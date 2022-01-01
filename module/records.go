package module

type Record struct {
	ID uint `json:"id" bind:"required"`
	Name string `json:"name" form:"name" bind:"required"`
	Course string `json:"course" form:"course" bind:"required"`
	Grade string `json:"grade" form:"grade" bind:"required"`
}

