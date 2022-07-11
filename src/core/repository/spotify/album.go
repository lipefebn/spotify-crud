package repository

type Album struct {
	ID  			string
	Name 			string
	Album_group 	string
	Album_type 		string
	Release_date	int64
	Popularity		int
}

func (a Album) GetId() string{
	return a.ID
}