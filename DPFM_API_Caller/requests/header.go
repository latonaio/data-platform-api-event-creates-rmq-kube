package requests

type Header struct {
	Event							*int	`json:"Event"`
	EventType						string	`json:"EventType"`
	EventOwner						int		`json:"EventOwner"`
	EventOwnerBusinessPartnerRole	string	`json:"EventOwnerBusinessPartnerRole"`
	PersonResponsible				string	`json:"PersonResponsible"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	OperationStartDate			    string  `json:"OperationStartDate"`
	OperationStartTime			    string  `json:"OperationStartTime"`
	OperationEndDate			    string  `json:"OperationEndDate"`
	OperationEndTime			    string  `json:"OperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	Site							int		`json:"Site"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	DistributionProfile				string	`json:"DistributionProfile"`
	PointConditionType				string	`json:"PointConditionType"`
	QuestionnaireType				*string `json:"QuestionnaireType"`
	QuestionnaireTemplate			*string `json:"QuestionnaireTemplate"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	CreateUser						int		`json:"CreateUser"`
	LastChangeUser					int		`json:"LastChangeUser"`
	IsReleased						*bool	`json:"IsReleased"`
	IsCancelled						*bool	`json:"IsCancelled"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
