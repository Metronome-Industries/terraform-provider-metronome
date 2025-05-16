// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_product

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1ContractProductDataSourceModel struct {
	Data customfield.NestedObject[V1ContractProductDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1ContractProductDataSourceModel) toReadParams(_ context.Context) (params metronome.V1ContractProductGetParams, diags diag.Diagnostics) {
	params = metronome.V1ContractProductGetParams{}

	return
}

type V1ContractProductDataDataSourceModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	Current      customfield.NestedObject[V1ContractProductDataCurrentDataSourceModel]     `tfsdk:"current" json:"current,computed"`
	Initial      customfield.NestedObject[V1ContractProductDataInitialDataSourceModel]     `tfsdk:"initial" json:"initial,computed"`
	Type         types.String                                                              `tfsdk:"type" json:"type,computed"`
	Updates      customfield.NestedObjectList[V1ContractProductDataUpdatesDataSourceModel] `tfsdk:"updates" json:"updates,computed"`
	ArchivedAt   timetypes.RFC3339                                                         `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields customfield.Map[types.String]                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
}

type V1ContractProductDataCurrentDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                            `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                            `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                            `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                          `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                          `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                              `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                              `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                            `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                            `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                          `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                          `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataCurrentQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataCurrentQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                       `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                          `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataCurrentQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataCurrentQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductDataInitialDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                            `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                            `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                            `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                          `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                          `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                              `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                              `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                            `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                            `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                          `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                          `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataInitialQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataInitialQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                       `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                          `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataInitialQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataInitialQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductDataUpdatesDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                            `tfsdk:"created_by" json:"created_by,computed"`
	BillableMetricID       types.String                                                                            `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                          `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                          `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                              `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                              `tfsdk:"is_refundable" json:"is_refundable,computed"`
	Name                   types.String                                                                            `tfsdk:"name" json:"name,computed"`
	NetsuiteInternalItemID types.String                                                                            `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                            `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                          `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                          `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataUpdatesQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataUpdatesQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                       `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                          `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataUpdatesQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataUpdatesQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}
