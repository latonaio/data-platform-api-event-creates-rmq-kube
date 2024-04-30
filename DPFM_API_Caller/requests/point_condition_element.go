package requests

type PointConditionElement struct {
	Event							int		`json:"Event"`
	PointConditionRecord			int		`json:"PointConditionRecord"`
	PointConditionSequentialNumber	int		`json:"PointConditionSequentialNumber"`
	PointSymbol						string	`json:"PointSymbol"`
	Sender							int		`json:"Sender"`
	PointTransactionType			string	`json:"PointTransactionType"`
	PointConditionType				string	`json:"PointConditionType"`
	PointConditionRateValue			float32	`json:"PointConditionRateValue"`
	PointConditionRatio				float32	`json:"PointConditionRatio"`
	PlusMinus						string	`json:"PlusMinus"`
	CreationDate					string	`json:"CreationDate"`
	LastChangeDate					string	`json:"LastChangeDate"`
	IsReleased						*bool	`json:"IsReleased"`
	IsCancelled						*bool	`json:"IsCancelled"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
