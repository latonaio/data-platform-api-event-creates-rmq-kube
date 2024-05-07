package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-event-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Event:                         *data.Event,
		EventType:                     data.EventType,
		EventOwner:                    data.EventOwner,
		EventOwnerBusinessPartnerRole: data.EventOwnerBusinessPartnerRole,
		PersonResponsible:             data.PersonResponsible,
		ValidityStartDate:             data.ValidityStartDate,
		ValidityStartTime:             data.ValidityStartTime,
		ValidityEndDate:               data.ValidityEndDate,
		ValidityEndTime:               data.ValidityEndTime,
		OperationStartDate:			   data.OperationStartDate,
		OperationStartTime:			   data.OperationStartTime,
		OperationEndDate:			   data.OperationEndDate,
		OperationEndTime:			   data.OperationEndTime,
		Description:                   data.Description,
		LongText:                      data.LongText,
		Introduction:                  data.Introduction,
		Site:                          data.Site,
		Project:                       data.Project,
		WBSElement:                    data.WBSElement,
		Tag1:                          data.Tag1,
		Tag2:                          data.Tag2,
		Tag3:                          data.Tag3,
		Tag4:                          data.Tag4,
		DistributionProfile:           data.DistributionProfile,
		PointConditionType:            data.PointConditionType,
		QuestionnaireType:			   data.QuestionnaireType,
		QuestionnaireTemplate:		   data.QuestionnaireTemplate,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		LastChangeUser:				   data.LastChangeUser,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Event:                   *dataHeader.Event,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		Event:          *dataHeader.Event,
		AddressID:      data.AddressID,
		PostalCode:     data.PostalCode,
		LocalSubRegion: data.LocalSubRegion,
		LocalRegion:    data.LocalRegion,
		Country:        data.Country,
		GlobalRegion:   data.GlobalRegion,
		TimeZone:       data.TimeZone,
		District:       data.District,
		StreetName:     data.StreetName,
		CityName:       data.CityName,
		Building:       data.Building,
		Floor:          data.Floor,
		Room:           data.Room,
		XCoordinate:    data.XCoordinate,
		YCoordinate:    data.YCoordinate,
		ZCoordinate:    data.ZCoordinate,
		Site:           data.Site,
	}
}

func ConvertToCampaignUpdates(header dpfm_api_input_reader.Header, campaign dpfm_api_input_reader.Campaign) *CampaignUpdates {
	dataHeader := header
	data := campaign

	return &CampaignUpdates{
		Event:          *dataHeader.Event,
		Campaign:       data.Campaign,
		LastChangeDate: data.LastChangeDate,
	}
}

func ConvertToGameUpdates(header dpfm_api_input_reader.Header, game dpfm_api_input_reader.Game) *GameUpdates {
	dataHeader := header
	data := game

	return &GameUpdates{
		Event:          *dataHeader.Event,
		Game:           data.Game,
		LastChangeDate: data.LastChangeDate,
	}
}

func ConvertToPointConditionElementUpdates(header dpfm_api_input_reader.Header, pointConditionElement dpfm_api_input_reader.PointConditionElement) *PointConditionElementUpdates {
	dataHeader := header
	data := pointConditionElement

	return &PointConditionElementUpdates{
		Event:                          *dataHeader.Event,
		PointConditionRecord:           data.PointConditionRecord,
		PointConditionSequentialNumber: data.PointConditionSequentialNumber,
		PointSymbol:                    data.PointSymbol,
		Sender:                         data.Sender,
		PointTransactionType:           data.PointTransactionType,
		PointConditionType:             data.PointConditionType,
		PointConditionRateValue:        data.PointConditionRateValue,
		PointConditionRatio:            data.PointConditionRatio,
		PlusMinus:                      data.PlusMinus,
		LastChangeDate:                 data.LastChangeDate,
	}
}
