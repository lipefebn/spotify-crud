package repository


type Artist struct {
	ID				string
	Name			string
	Popularity		int
	Followers		int	
}

func (a Artist) GetId() string{
	return a.ID
}