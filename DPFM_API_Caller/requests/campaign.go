package requests

type Campaign struct {
	Event				int		`json:"Event"`
	Campaign			int		`json:"Campaign"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsReleased			*bool	`json:"IsReleased"`
	IsCancelled			*bool	`json:"IsCancelled"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
