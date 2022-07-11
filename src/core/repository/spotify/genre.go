package repository

type Genre struct {
	ID string
}

func (g Genre) GetId() string{
	return g.ID
}