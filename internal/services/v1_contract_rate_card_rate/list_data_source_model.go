// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_rate

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractRateCardRatesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1ContractRateCardRatesItemsDataSourceModel] `json:"data,computed"`
}

type V1ContractRateCardRatesDataSourceModel struct {
	At         timetypes.RFC3339                                                         `tfsdk:"at" json:"at,required" format:"date-time"`
	RateCardID types.String                                                              `tfsdk:"rate_card_id" json:"rate_card_id,required"`
	Selectors  *[]*V1ContractRateCardRatesSelectorsDataSourceModel                       `tfsdk:"selectors" json:"selectors,optional"`
	MaxItems   types.Int64                                                               `tfsdk:"max_items"`
	Items      customfield.NestedObjectList[V1ContractRateCardRatesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1ContractRateCardRatesDataSourceModel) toListParams(_ context.Context) (params metronome.V1ContractRateCardRateListParams, diags diag.Diagnostics) {
	params = metronome.V1ContractRateCardRateListParams{}

	return
}

type V1ContractRateCardRatesSelectorsDataSourceModel struct {
	BillingFrequency          types.String             `tfsdk:"billing_frequency" json:"billing_frequency,optional"`
	PartialPricingGroupValues *map[string]types.String `tfsdk:"partial_pricing_group_values" json:"partial_pricing_group_values,optional"`
	PricingGroupValues        *map[string]types.String `tfsdk:"pricing_group_values" json:"pricing_group_values,optional"`
	ProductID                 types.String             `tfsdk:"product_id" json:"product_id,optional"`
	ProductTags               *[]types.String          `tfsdk:"product_tags" json:"product_tags,optional"`
}

type V1ContractRateCardRatesItemsDataSourceModel struct {
	Entitled            types.Bool                                                                 `tfsdk:"entitled" json:"entitled,computed"`
	ProductCustomFields customfield.Map[types.String]                                              `tfsdk:"product_custom_fields" json:"product_custom_fields,computed"`
	ProductID           types.String                                                               `tfsdk:"product_id" json:"product_id,computed"`
	ProductName         types.String                                                               `tfsdk:"product_name" json:"product_name,computed"`
	ProductTags         customfield.List[types.String]                                             `tfsdk:"product_tags" json:"product_tags,computed"`
	Rate                customfield.NestedObject[V1ContractRateCardRatesRateDataSourceModel]       `tfsdk:"rate" json:"rate,computed"`
	StartingAt          timetypes.RFC3339                                                          `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	BillingFrequency    types.String                                                               `tfsdk:"billing_frequency" json:"billing_frequency,computed"`
	CommitRate          customfield.NestedObject[V1ContractRateCardRatesCommitRateDataSourceModel] `tfsdk:"commit_rate" json:"commit_rate,computed"`
	EndingBefore        timetypes.RFC3339                                                          `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	PricingGroupValues  customfield.Map[types.String]                                              `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
}

type V1ContractRateCardRatesRateDataSourceModel struct {
	RateType           types.String                                                                   `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType         customfield.NestedObject[V1ContractRateCardRatesRateCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate         customfield.Map[jsontypes.Normalized]                                          `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated         types.Bool                                                                     `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price              types.Float64                                                                  `tfsdk:"price" json:"price,computed"`
	PricingGroupValues customfield.Map[types.String]                                                  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	Quantity           types.Float64                                                                  `tfsdk:"quantity" json:"quantity,computed"`
	Tiers              customfield.NestedObjectList[V1ContractRateCardRatesRateTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
	UseListPrices      types.Bool                                                                     `tfsdk:"use_list_prices" json:"use_list_prices,computed"`
}

type V1ContractRateCardRatesRateCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractRateCardRatesRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractRateCardRatesCommitRateDataSourceModel struct {
	RateType types.String                                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	Price    types.Float64                                                                       `tfsdk:"price" json:"price,computed"`
	Tiers    customfield.NestedObjectList[V1ContractRateCardRatesCommitRateTiersDataSourceModel] `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractRateCardRatesCommitRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}
