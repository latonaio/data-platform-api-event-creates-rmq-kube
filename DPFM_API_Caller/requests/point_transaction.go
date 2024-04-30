package requests

type PointTransaction struct {
	Event							int		`json:"Event"`
	Sender							int		`json:"Sender"`
	Receiver						int		`json:"Receiver"`
	PointConditionRecord			int		`json:"PointConditionRecord"`
	PointConditionSequentialNumber	int		`json:"PointConditionSequentialNumber"`
	PointTransaction				int		`json:"PointTransaction"`
	PointSymbol						string	`json:"PointSymbol"`
	PointTransactionType			string	`json:"PointTransactionType"`
	PointConditionType				string	`json:"PointConditionType"`
	PointConditionRateValue			float32	`json:"PointConditionRateValue"`
	PointConditionRatio				float32	`json:"PointConditionRatio"`
	PlusMinus						string	`json:"PlusMinus"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsCancelled						*bool	`json:"IsCancelled"`
}
