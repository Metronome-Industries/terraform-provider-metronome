// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1ContractRateCardModel struct {
	FiatCreditTypeID      types.String                                          `tfsdk:"fiat_credit_type_id" json:"fiat_credit_type_id,optional,no_refresh"`
	CustomFields          *map[string]types.String                              `tfsdk:"custom_fields" json:"custom_fields,optional,no_refresh"`
	CreditTypeConversions *[]*V1ContractRateCardCreditTypeConversionsModel      `tfsdk:"credit_type_conversions" json:"credit_type_conversions,optional,no_refresh"`
	Name                  types.String                                          `tfsdk:"name" json:"name,required,no_refresh"`
	Description           types.String                                          `tfsdk:"description" json:"description,optional,no_refresh"`
	RateCardID            types.String                                          `tfsdk:"rate_card_id" json:"rate_card_id,optional,no_refresh"`
	Aliases               *[]*V1ContractRateCardAliasesModel                    `tfsdk:"aliases" json:"aliases,optional,no_refresh"`
	Data                  customfield.NestedObject[V1ContractRateCardDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1ContractRateCardModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1ContractRateCardModel) MarshalJSONForUpdate(state V1ContractRateCardModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1ContractRateCardCreditTypeConversionsModel struct {
	CustomCreditTypeID  types.String  `tfsdk:"custom_credit_type_id" json:"custom_credit_type_id,required"`
	FiatPerCustomCredit types.Float64 `tfsdk:"fiat_per_custom_credit" json:"fiat_per_custom_credit,required"`
}

type V1ContractRateCardAliasesModel struct {
	Name         types.String      `tfsdk:"name" json:"name,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,optional" format:"date-time"`
}

type V1ContractRateCardDataModel struct {
	ID                    types.String                                                                   `tfsdk:"id" json:"id,computed"`
	CreatedAt             timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy             types.String                                                                   `tfsdk:"created_by" json:"created_by,computed"`
	Name                  types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Aliases               customfield.NestedObjectList[V1ContractRateCardDataAliasesModel]               `tfsdk:"aliases" json:"aliases,computed"`
	CreditTypeConversions customfield.NestedObjectList[V1ContractRateCardDataCreditTypeConversionsModel] `tfsdk:"credit_type_conversions" json:"credit_type_conversions,computed"`
	CustomFields          customfield.Map[types.String]                                                  `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description           types.String                                                                   `tfsdk:"description" json:"description,computed"`
	FiatCreditType        customfield.NestedObject[V1ContractRateCardDataFiatCreditTypeModel]            `tfsdk:"fiat_credit_type" json:"fiat_credit_type,computed"`
}

type V1ContractRateCardDataAliasesModel struct {
	Name         types.String      `tfsdk:"name" json:"name,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractRateCardDataCreditTypeConversionsModel struct {
	CustomCreditType    customfield.NestedObject[V1ContractRateCardDataCreditTypeConversionsCustomCreditTypeModel] `tfsdk:"custom_credit_type" json:"custom_credit_type,computed"`
	FiatPerCustomCredit types.String                                                                               `tfsdk:"fiat_per_custom_credit" json:"fiat_per_custom_credit,computed"`
}

type V1ContractRateCardDataCreditTypeConversionsCustomCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractRateCardDataFiatCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
