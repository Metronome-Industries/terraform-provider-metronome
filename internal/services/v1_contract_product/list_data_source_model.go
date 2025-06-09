// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_product

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractProductsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1ContractProductsItemsDataSourceModel] `json:"data,computed"`
}

type V1ContractProductsDataSourceModel struct {
	ArchiveFilter types.String                                                         `tfsdk:"archive_filter" json:"archive_filter,optional"`
	MaxItems      types.Int64                                                          `tfsdk:"max_items"`
	Items         customfield.NestedObjectList[V1ContractProductsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1ContractProductsDataSourceModel) toListParams(_ context.Context) (params metronome.V1ContractProductListParams, diags diag.Diagnostics) {
	params = metronome.V1ContractProductListParams{}

	return
}

type V1ContractProductsItemsDataSourceModel struct {
	ID           types.String                                                           `tfsdk:"id" json:"id,computed"`
	Current      customfield.NestedObject[V1ContractProductsCurrentDataSourceModel]     `tfsdk:"current" json:"current,computed"`
	Initial      customfield.NestedObject[V1ContractProductsInitialDataSourceModel]     `tfsdk:"initial" json:"initial,computed"`
	Type         types.String                                                           `tfsdk:"type" json:"type,computed"`
	Updates      customfield.NestedObjectList[V1ContractProductsUpdatesDataSourceModel] `tfsdk:"updates" json:"updates,computed"`
	ArchivedAt   timetypes.RFC3339                                                      `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields customfield.Map[types.String]                                          `tfsdk:"custom_fields" json:"custom_fields,computed"`
}

type V1ContractProductsCurrentDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                         `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                         `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                         `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                       `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                       `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                           `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                           `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                         `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                         `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                       `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                       `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductsCurrentQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductsCurrentQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                       `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductsCurrentQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductsCurrentQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductsInitialDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                         `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                         `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                         `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                       `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                       `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                           `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                           `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                         `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                         `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                       `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                       `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductsInitialQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductsInitialQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                       `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductsInitialQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductsInitialQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductsUpdatesDataSourceModel struct {
	CreatedAt              timetypes.RFC3339                                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                         `tfsdk:"created_by" json:"created_by,computed"`
	BillableMetricID       types.String                                                                         `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                       `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                       `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                           `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                           `tfsdk:"is_refundable" json:"is_refundable,computed"`
	Name                   types.String                                                                         `tfsdk:"name" json:"name,computed"`
	NetsuiteInternalItemID types.String                                                                         `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                         `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                       `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                       `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductsUpdatesQuantityConversionDataSourceModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductsUpdatesQuantityRoundingDataSourceModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                                    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                       `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductsUpdatesQuantityConversionDataSourceModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductsUpdatesQuantityRoundingDataSourceModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}
