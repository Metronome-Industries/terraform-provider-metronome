// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_product

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1ContractProductModel struct {
	Type                   types.String                                         `tfsdk:"type" json:"type,required,no_refresh"`
	Name                   types.String                                         `tfsdk:"name" json:"name,required,no_refresh"`
	BillableMetricID       types.String                                         `tfsdk:"billable_metric_id" json:"billable_metric_id,optional,no_refresh"`
	ExcludeFreeUsage       types.Bool                                           `tfsdk:"exclude_free_usage" json:"exclude_free_usage,optional,no_refresh"`
	IsRefundable           types.Bool                                           `tfsdk:"is_refundable" json:"is_refundable,optional,no_refresh"`
	NetsuiteInternalItemID types.String                                         `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,optional,no_refresh"`
	NetsuiteOverageItemID  types.String                                         `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,optional,no_refresh"`
	ProductID              types.String                                         `tfsdk:"product_id" json:"product_id,optional,no_refresh"`
	StartingAt             timetypes.RFC3339                                    `tfsdk:"starting_at" json:"starting_at,optional,no_refresh" format:"date-time"`
	CompositeProductIDs    *[]types.String                                      `tfsdk:"composite_product_ids" json:"composite_product_ids,optional,no_refresh"`
	CompositeTags          *[]types.String                                      `tfsdk:"composite_tags" json:"composite_tags,optional,no_refresh"`
	PresentationGroupKey   *[]types.String                                      `tfsdk:"presentation_group_key" json:"presentation_group_key,optional,no_refresh"`
	PricingGroupKey        *[]types.String                                      `tfsdk:"pricing_group_key" json:"pricing_group_key,optional,no_refresh"`
	Tags                   *[]types.String                                      `tfsdk:"tags" json:"tags,optional,no_refresh"`
	QuantityConversion     *V1ContractProductQuantityConversionModel            `tfsdk:"quantity_conversion" json:"quantity_conversion,optional,no_refresh"`
	QuantityRounding       *V1ContractProductQuantityRoundingModel              `tfsdk:"quantity_rounding" json:"quantity_rounding,optional,no_refresh"`
	Data                   customfield.NestedObject[V1ContractProductDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1ContractProductModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1ContractProductModel) MarshalJSONForUpdate(state V1ContractProductModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1ContractProductQuantityConversionModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,required"`
	Operation        types.String  `tfsdk:"operation" json:"operation,required"`
	Name             types.String  `tfsdk:"name" json:"name,optional"`
}

type V1ContractProductQuantityRoundingModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,required"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,required"`
}

type V1ContractProductDataModel struct {
	ID           types.String                                                    `tfsdk:"id" json:"id,computed"`
	Current      customfield.NestedObject[V1ContractProductDataCurrentModel]     `tfsdk:"current" json:"current,computed"`
	Initial      customfield.NestedObject[V1ContractProductDataInitialModel]     `tfsdk:"initial" json:"initial,computed"`
	Type         types.String                                                    `tfsdk:"type" json:"type,computed"`
	Updates      customfield.NestedObjectList[V1ContractProductDataUpdatesModel] `tfsdk:"updates" json:"updates,computed"`
	ArchivedAt   timetypes.RFC3339                                               `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields customfield.Map[types.String]                                   `tfsdk:"custom_fields" json:"custom_fields,computed"`
}

type V1ContractProductDataCurrentModel struct {
	CreatedAt              timetypes.RFC3339                                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                  `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                  `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                  `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                    `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                    `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                  `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                  `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataCurrentQuantityConversionModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataCurrentQuantityRoundingModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                             `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataCurrentQuantityConversionModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataCurrentQuantityRoundingModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductDataInitialModel struct {
	CreatedAt              timetypes.RFC3339                                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                  `tfsdk:"created_by" json:"created_by,computed"`
	Name                   types.String                                                                  `tfsdk:"name" json:"name,computed"`
	BillableMetricID       types.String                                                                  `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                    `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                    `tfsdk:"is_refundable" json:"is_refundable,computed"`
	NetsuiteInternalItemID types.String                                                                  `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                  `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataInitialQuantityConversionModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataInitialQuantityRoundingModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                             `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataInitialQuantityConversionModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataInitialQuantityRoundingModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}

type V1ContractProductDataUpdatesModel struct {
	CreatedAt              timetypes.RFC3339                                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy              types.String                                                                  `tfsdk:"created_by" json:"created_by,computed"`
	BillableMetricID       types.String                                                                  `tfsdk:"billable_metric_id" json:"billable_metric_id,computed"`
	CompositeProductIDs    customfield.List[types.String]                                                `tfsdk:"composite_product_ids" json:"composite_product_ids,computed"`
	CompositeTags          customfield.List[types.String]                                                `tfsdk:"composite_tags" json:"composite_tags,computed"`
	ExcludeFreeUsage       types.Bool                                                                    `tfsdk:"exclude_free_usage" json:"exclude_free_usage,computed"`
	IsRefundable           types.Bool                                                                    `tfsdk:"is_refundable" json:"is_refundable,computed"`
	Name                   types.String                                                                  `tfsdk:"name" json:"name,computed"`
	NetsuiteInternalItemID types.String                                                                  `tfsdk:"netsuite_internal_item_id" json:"netsuite_internal_item_id,computed"`
	NetsuiteOverageItemID  types.String                                                                  `tfsdk:"netsuite_overage_item_id" json:"netsuite_overage_item_id,computed"`
	PresentationGroupKey   customfield.List[types.String]                                                `tfsdk:"presentation_group_key" json:"presentation_group_key,computed"`
	PricingGroupKey        customfield.List[types.String]                                                `tfsdk:"pricing_group_key" json:"pricing_group_key,computed"`
	QuantityConversion     customfield.NestedObject[V1ContractProductDataUpdatesQuantityConversionModel] `tfsdk:"quantity_conversion" json:"quantity_conversion,computed"`
	QuantityRounding       customfield.NestedObject[V1ContractProductDataUpdatesQuantityRoundingModel]   `tfsdk:"quantity_rounding" json:"quantity_rounding,computed"`
	StartingAt             timetypes.RFC3339                                                             `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Tags                   customfield.List[types.String]                                                `tfsdk:"tags" json:"tags,computed"`
}

type V1ContractProductDataUpdatesQuantityConversionModel struct {
	ConversionFactor types.Float64 `tfsdk:"conversion_factor" json:"conversion_factor,computed"`
	Operation        types.String  `tfsdk:"operation" json:"operation,computed"`
	Name             types.String  `tfsdk:"name" json:"name,computed"`
}

type V1ContractProductDataUpdatesQuantityRoundingModel struct {
	DecimalPlaces  types.Float64 `tfsdk:"decimal_places" json:"decimal_places,computed"`
	RoundingMethod types.String  `tfsdk:"rounding_method" json:"rounding_method,computed"`
}
