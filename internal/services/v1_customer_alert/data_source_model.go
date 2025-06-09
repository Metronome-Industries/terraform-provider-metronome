// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_alert

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CustomerAlertDataSourceModel struct {
	AlertID          types.String                                                 `tfsdk:"alert_id" json:"alert_id,required"`
	CustomerID       types.String                                                 `tfsdk:"customer_id" json:"customer_id,required"`
	PlansOrContracts types.String                                                 `tfsdk:"plans_or_contracts" json:"plans_or_contracts,optional"`
	Data             customfield.NestedObject[V1CustomerAlertDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1CustomerAlertDataSourceModel) toReadParams(_ context.Context) (params metronome.V1CustomerAlertGetParams, diags diag.Diagnostics) {
	params = metronome.V1CustomerAlertGetParams{}

	return
}

type V1CustomerAlertDataDataSourceModel struct {
	Alert          customfield.NestedObject[V1CustomerAlertDataAlertDataSourceModel] `tfsdk:"alert" json:"alert,computed"`
	CustomerStatus types.String                                                      `tfsdk:"customer_status" json:"customer_status,computed"`
	TriggeredBy    types.String                                                      `tfsdk:"triggered_by" json:"triggered_by,computed"`
}

type V1CustomerAlertDataAlertDataSourceModel struct {
	ID                     types.String                                                                            `tfsdk:"id" json:"id,computed"`
	Name                   types.String                                                                            `tfsdk:"name" json:"name,computed"`
	Status                 types.String                                                                            `tfsdk:"status" json:"status,computed"`
	Threshold              types.Float64                                                                           `tfsdk:"threshold" json:"threshold,computed"`
	Type                   types.String                                                                            `tfsdk:"type" json:"type,computed"`
	UpdatedAt              timetypes.RFC3339                                                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	CreditGrantTypeFilters customfield.List[types.String]                                                          `tfsdk:"credit_grant_type_filters" json:"credit_grant_type_filters,computed"`
	CreditType             customfield.NestedObject[V1CustomerAlertDataAlertCreditTypeDataSourceModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomFieldFilters     customfield.NestedObjectList[V1CustomerAlertDataAlertCustomFieldFiltersDataSourceModel] `tfsdk:"custom_field_filters" json:"custom_field_filters,computed"`
	GroupKeyFilter         customfield.NestedObject[V1CustomerAlertDataAlertGroupKeyFilterDataSourceModel]         `tfsdk:"group_key_filter" json:"group_key_filter,computed"`
	InvoiceTypesFilter     customfield.List[types.String]                                                          `tfsdk:"invoice_types_filter" json:"invoice_types_filter,computed"`
	UniquenessKey          types.String                                                                            `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1CustomerAlertDataAlertCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CustomerAlertDataAlertCustomFieldFiltersDataSourceModel struct {
	Entity types.String `tfsdk:"entity" json:"entity,computed"`
	Key    types.String `tfsdk:"key" json:"key,computed"`
	Value  types.String `tfsdk:"value" json:"value,computed"`
}

type V1CustomerAlertDataAlertGroupKeyFilterDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
