package module

type Student struct {
	ID uint `json:"id" bind:"required"`
	Name string `json:"name" form:"name" bind:"required"`
	Password string `json:"password" form:"password" bind:"required"`
}
