package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-event-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-event-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-event-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToCampaignCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Campaign, error) {
	campaigns := make([]Campaign, 0)

	for _, data := range *subfuncSDC.Message.Campaign {
		campaign, err := TypeConverter[*Campaign](data)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, *campaign)
	}

	return &campaigns, nil
}

func ConvertToGameCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Game, error) {
	games := make([]Game, 0)

	for _, data := range *subfuncSDC.Message.Game {
		game, err := TypeConverter[*Game](data)
		if err != nil {
			return nil, err
		}

		games = append(games, *game)
	}

	return &games, nil
}

func ConvertToPointTransactionCreates(subfuncSDC *sub_func_complementer.SDC) (*[]PointTransaction, error) {
	pointTransactions := make([]PointTransaction, 0)

	for _, data := range *subfuncSDC.Message.PointTransaction {
		pointTransaction, err := TypeConverter[*PointTransaction](data)
		if err != nil {
			return nil, err
		}

		pointTransactions = append(pointTransactions, *pointTransaction)
	}

	return &pointTransactions, nil
}

func ConvertToPointConditionElementCreates(subfuncSDC *sub_func_complementer.SDC) (*[]PointConditionElement, error) {
	pointConditionElements := make([]PointConditionElement, 0)

	for _, data := range *subfuncSDC.Message.PointConditionElement {
		pointConditionElement, err := TypeConverter[*PointConditionElement](data)
		if err != nil {
			return nil, err
		}

		pointConditionElements = append(pointConditionElements, *pointConditionElement)
	}

	return &pointConditionElements, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToCampaignUpdates(campaignUpdates *[]dpfm_api_processing_formatter.CampaignUpdates) (*[]Campaign, error) {
	campaigns := make([]Campaign, 0)

	for _, data := range *campaignUpdates {
		campaign, err := TypeConverter[*Campaign](data)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, *campaign)
	}

	return &campaigns, nil
}

func ConvertToGameUpdates(gameUpdates *[]dpfm_api_processing_formatter.GameUpdates) (*[]Game, error) {
	games := make([]Game, 0)

	for _, data := range *gameUpdates {
		game, err := TypeConverter[*Game](data)
		if err != nil {
			return nil, err
		}

		games = append(games, *game)
	}

	return &games, nil
}

func ConvertToPointConditionElementUpdates(pointConditionElementUpdates *[]dpfm_api_processing_formatter.PointConditionElementUpdates) (*[]PointConditionElement, error) {
	pointConditionElements := make([]PointConditionElement, 0)

	for _, data := range *pointConditionElementUpdates {
		pointConditionElementUpdate, err := TypeConverter[*PointConditionElement](data)
		if err != nil {
			return nil, err
		}

		pointConditionElements = append(pointConditionElements, *pointConditionElementUpdate)
	}

	return &pointConditionElements, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Event:                         *input.Header.Event,
		EventType:                     input.Header.EventType,
		EventOwner:                    input.Header.EventOwner,
		EventOwnerBusinessPartnerRole: input.Header.EventOwnerBusinessPartnerRole,
		PersonResponsible:             input.Header.PersonResponsible,
		ValidityStartDate:             input.Header.ValidityStartDate,
		ValidityStartTime:             input.Header.ValidityStartTime,
		ValidityEndDate:               input.Header.ValidityEndDate,
		ValidityEndTime:               input.Header.ValidityEndTime,
		OperationStartDate:            input.Header.OperationStartDate,
		OperationStartTime:            input.Header.OperationStartTime,
		OperationEndDate:              input.Header.OperationEndDate,
		OperationEndTime:              input.Header.OperationEndTime,
		Description:                   input.Header.Description,
		LongText:                      input.Header.LongText,
		Introduction:                  input.Header.Introduction,
		Site:                          input.Header.Site,
		Project:                       input.Header.Project,
		WBSElement:                    input.Header.WBSElement,
		Tag1:                          input.Header.Tag1,
		Tag2:                          input.Header.Tag2,
		Tag3:                          input.Header.Tag3,
		Tag4:                          input.Header.Tag4,
		DistributionProfile:           input.Header.DistributionProfile,
		PointConditionType:            input.Header.PointConditionType,
		QuestionnaireType:			   input.Header.QuestionnaireType,
		QuestionnaireTemplate:		   input.Header.QuestionnaireTemplate,
		CreationDate:                  input.Header.CreationDate,
		CreationTime:                  input.Header.CreationTime,
		LastChangeDate:                input.Header.LastChangeDate,
		LastChangeTime:                input.Header.LastChangeTime,
		CreateUser:					   input.Header.CreateUser,
		LastChangeUser:				   input.Header.LastChangeUser,
		IsReleased:                    input.Header.IsReleased,
		IsCancelled:                   input.Header.IsCancelled,
		IsMarkedForDeletion:           input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToAddress(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var addresses []sub_func_complementer.Address

	addresses = append(
		addresses,
		sub_func_complementer.Address{
			Event:          *input.Header.Event,
			AddressID:      input.Header.Address[0].AddressID,
			PostalCode:     input.Header.Address[0].PostalCode,
			LocalSubRegion: input.Header.Address[0].LocalSubRegion,
			LocalRegion:    input.Header.Address[0].LocalRegion,
			Country:        input.Header.Address[0].Country,
			GlobalRegion:   input.Header.Address[0].GlobalRegion,
			TimeZone:       input.Header.Address[0].TimeZone,
			District:       input.Header.Address[0].District,
			StreetName:     input.Header.Address[0].StreetName,
			CityName:       input.Header.Address[0].CityName,
			Building:       input.Header.Address[0].Building,
			Floor:          input.Header.Address[0].Floor,
			Room:           input.Header.Address[0].Room,
			XCoordinate:    input.Header.Address[0].XCoordinate,
			YCoordinate:    input.Header.Address[0].YCoordinate,
			ZCoordinate:    input.Header.Address[0].ZCoordinate,
			Site:           input.Header.Address[0].Site,
		},
	)

	subfuncSDC.Message.Address = &addresses

	return subfuncSDC
}

func ConvertToPartner(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var partners []sub_func_complementer.Partner

	partners = append(
		partners,
		sub_func_complementer.Partner{
			Event:          		*input.Header.Event,
			PartnerFunction:         input.Header.Partner[0].PartnerFunction,
			BusinessPartner:         input.Header.Partner[0].BusinessPartner,
			BusinessPartnerFullName: input.Header.Partner[0].BusinessPartnerFullName,
			BusinessPartnerName:     input.Header.Partner[0].BusinessPartnerName,
			Organization:            input.Header.Partner[0].Organization,
			Country:                 input.Header.Partner[0].Country,
			Language:                input.Header.Partner[0].Language,
			Currency:                input.Header.Partner[0].Currency,
			ExternalDocumentID:      input.Header.Partner[0].ExternalDocumentID,
			AddressID:               input.Header.Partner[0].AddressID,
			EmailAddress:            input.Header.Partner[0].EmailAddress,
		},
	)

	subfuncSDC.Message.Partner = &partners

	return subfuncSDC
}

func ConvertToPointTransaction(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var pointTransactions []sub_func_complementer.PointTransaction

	pointTransactions = append(
		pointTransactions,
		sub_func_complementer.PointTransaction{
			Event:								*input.Header.Event,
			Sender:								input.Header.PointTransaction[0].Sender,
			Receiver:							input.Header.PointTransaction[0].Receiver,
			PointConditionRecord:				input.Header.PointTransaction[0].PointConditionRecord,
			PointConditionSequentialNumber:		input.Header.PointTransaction[0].PointConditionSequentialNumber,
			PointTransaction:					input.Header.PointTransaction[0].PointTransaction,
			PointSymbol:						input.Header.PointTransaction[0].PointSymbol,
			PointTransactionType:				input.Header.PointTransaction[0].PointTransactionType,
			PointConditionType:					input.Header.PointTransaction[0].PointConditionType,
			PointConditionRateValue:			input.Header.PointTransaction[0].PointConditionRateValue,
			PointConditionRatio:				input.Header.PointTransaction[0].PointConditionRatio,
			PlusMinus:							input.Header.PointTransaction[0].PlusMinus,
			CreationDate:						input.Header.PointTransaction[0].CreationDate,
			CreationTime:						input.Header.PointTransaction[0].CreationTime,
			LastChangeDate:						input.Header.PointTransaction[0].LastChangeDate,
			LastChangeTime:						input.Header.PointTransaction[0].LastChangeTime,
			IsCancelled:						input.Header.PointTransaction[0].IsCancelled,
		},
	)

	subfuncSDC.Message.PointTransaction = &pointTransactions

	return subfuncSDC
}

func ConvertToPointConditionElement(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var pointConditionElements []sub_func_complementer.PointConditionElement

	pointConditionElements = append(
		pointConditionElements,
		sub_func_complementer.PointConditionElement{
			Event:								*input.Header.Event,
			PointConditionRecord:				input.Header.PointConditionElement[0].PointConditionRecord,
			PointConditionSequentialNumber:		input.Header.PointConditionElement[0].PointConditionSequentialNumber,
			PointSymbol:						input.Header.PointConditionElement[0].PointSymbol,
			Sender:								input.Header.PointConditionElement[0].Sender,
			PointTransactionType:				input.Header.PointConditionElement[0].PointTransactionType,
			PointConditionType:					input.Header.PointConditionElement[0].PointConditionType,
			PointConditionRateValue:			input.Header.PointConditionElement[0].PointConditionRateValue,
			PointConditionRatio:				input.Header.PointConditionElement[0].PointConditionRatio,
			PlusMinus:							input.Header.PointConditionElement[0].PlusMinus,
			CreationDate:						input.Header.PointConditionElement[0].CreationDate,
			LastChangeDate:						input.Header.PointConditionElement[0].LastChangeDate,
			IsReleased:							input.Header.PointConditionElement[0].IsReleased,
			IsCancelled:						input.Header.PointConditionElement[0].IsCancelled,
			IsMarkedForDeletion:				input.Header.PointConditionElement[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.PointConditionElement = &pointConditionElements

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
