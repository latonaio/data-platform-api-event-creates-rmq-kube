package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Event							int		`json:"Event"`
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
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
}

type PartnerUpdates struct {
	Event                 	int     `json:"Event"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type AddressUpdates struct {
	Event     		int     	`json:"Event"`
	AddressID   	int     	`json:"AddressID"`
	PostalCode  	string 		`json:"PostalCode"`
	LocalSubRegion 	string 		`json:"LocalSubRegion"`
	LocalRegion 	string 		`json:"LocalRegion"`
	Country     	string 		`json:"Country"`
	GlobalRegion   	string 		`json:"GlobalRegion"`
	TimeZone   		string 		`json:"TimeZone"`
	District    	*string 	`json:"District"`
	StreetName  	*string 	`json:"StreetName"`
	CityName    	*string 	`json:"CityName"`
	Building    	*string 	`json:"Building"`
	Floor       	*int		`json:"Floor"`
	Room        	*int		`json:"Room"`
	XCoordinate 	*float32	`json:"XCoordinate"`
	YCoordinate 	*float32	`json:"YCoordinate"`
	ZCoordinate 	*float32	`json:"ZCoordinate"`
	Site			*int		`json:"Site"`
}

type CampaignUpdates struct {
	Event				int		`json:"Event"`
	Campaign			int		`json:"Campaign"`
	LastChangeDate		string	`json:"LastChangeDate"`
}

type GameUpdates struct {
	Event				int		`json:"Event"`
	Game				int		`json:"Game"`
	LastChangeDate		string	`json:"LastChangeDate"`
}

type PointConditionElementUpdates struct {
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
	LastChangeDate					string	`json:"LastChangeDate"`
}

