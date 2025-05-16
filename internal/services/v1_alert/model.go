// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_alert

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1AlertModel struct {
	AlertType              types.String                               `tfsdk:"alert_type" json:"alert_type,required"`
	Name                   types.String                               `tfsdk:"name" json:"name,required"`
	Threshold              types.Float64                              `tfsdk:"threshold" json:"threshold,required"`
	BillableMetricID       types.String                               `tfsdk:"billable_metric_id" json:"billable_metric_id,optional"`
	CreditTypeID           types.String                               `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	CustomerID             types.String                               `tfsdk:"customer_id" json:"customer_id,optional"`
	EvaluateOnCreate       types.Bool                                 `tfsdk:"evaluate_on_create" json:"evaluate_on_create,optional"`
	PlanID                 types.String                               `tfsdk:"plan_id" json:"plan_id,optional"`
	UniquenessKey          types.String                               `tfsdk:"uniqueness_key" json:"uniqueness_key,optional"`
	CreditGrantTypeFilters *[]types.String                            `tfsdk:"credit_grant_type_filters" json:"credit_grant_type_filters,optional"`
	InvoiceTypesFilter     *[]types.String                            `tfsdk:"invoice_types_filter" json:"invoice_types_filter,optional"`
	CustomFieldFilters     *[]*V1AlertCustomFieldFiltersModel         `tfsdk:"custom_field_filters" json:"custom_field_filters,optional"`
	GroupKeyFilter         *V1AlertGroupKeyFilterModel                `tfsdk:"group_key_filter" json:"group_key_filter,optional"`
	Data                   customfield.NestedObject[V1AlertDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1AlertModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1AlertModel) MarshalJSONForUpdate(state V1AlertModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1AlertCustomFieldFiltersModel struct {
	Entity types.String `tfsdk:"entity" json:"entity,required"`
	Key    types.String `tfsdk:"key" json:"key,required"`
	Value  types.String `tfsdk:"value" json:"value,required"`
}

type V1AlertGroupKeyFilterModel struct {
	Key   types.String `tfsdk:"key" json:"key,required"`
	Value types.String `tfsdk:"value" json:"value,required"`
}

type V1AlertDataModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
