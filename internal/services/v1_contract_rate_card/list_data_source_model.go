// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractRateCardsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1ContractRateCardsItemsDataSourceModel] `json:"data,computed"`
}

type V1ContractRateCardsDataSourceModel struct {
	Body     jsontypes.Normalized                                                  `tfsdk:"body" json:"body,optional"`
	MaxItems types.Int64                                                           `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[V1ContractRateCardsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1ContractRateCardsDataSourceModel) toListParams(_ context.Context) (params metronome.V1ContractRateCardListParams, diags diag.Diagnostics) {
	params = metronome.V1ContractRateCardListParams{}

	return
}

type V1ContractRateCardsItemsDataSourceModel struct {
	ID                    types.String                                                                          `tfsdk:"id" json:"id,computed"`
	CreatedAt             timetypes.RFC3339                                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy             types.String                                                                          `tfsdk:"created_by" json:"created_by,computed"`
	Name                  types.String                                                                          `tfsdk:"name" json:"name,computed"`
	Aliases               customfield.NestedObjectList[V1ContractRateCardsAliasesDataSourceModel]               `tfsdk:"aliases" json:"aliases,computed"`
	CreditTypeConversions customfield.NestedObjectList[V1ContractRateCardsCreditTypeConversionsDataSourceModel] `tfsdk:"credit_type_conversions" json:"credit_type_conversions,computed"`
	CustomFields          customfield.Map[types.String]                                                         `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description           types.String                                                                          `tfsdk:"description" json:"description,computed"`
	FiatCreditType        customfield.NestedObject[V1ContractRateCardsFiatCreditTypeDataSourceModel]            `tfsdk:"fiat_credit_type" json:"fiat_credit_type,computed"`
}

type V1ContractRateCardsAliasesDataSourceModel struct {
	Name         types.String      `tfsdk:"name" json:"name,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractRateCardsCreditTypeConversionsDataSourceModel struct {
	CustomCreditType    customfield.NestedObject[V1ContractRateCardsCreditTypeConversionsCustomCreditTypeDataSourceModel] `tfsdk:"custom_credit_type" json:"custom_credit_type,computed"`
	FiatPerCustomCredit types.String                                                                                      `tfsdk:"fiat_per_custom_credit" json:"fiat_per_custom_credit,computed"`
}

type V1ContractRateCardsCreditTypeConversionsCustomCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractRateCardsFiatCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
