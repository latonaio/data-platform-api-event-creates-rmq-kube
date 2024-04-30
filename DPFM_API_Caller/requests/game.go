package requests

type Game struct {
	Event				int		`json:"Event"`
	Game				int		`json:"Game"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsReleased			*bool	`json:"IsReleased"`
	IsCancelled			*bool	`json:"IsCancelled"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
