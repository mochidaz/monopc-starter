package model

type Category struct {
	ID   int
	Name string
	Auditable
}

func NewCategory(
	id int,
	name string,
	createdBy string,
) *Category {
	return &Category{
		ID:        id,
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}
