// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CreditGrantsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1CreditGrantsItemsDataSourceModel] `json:"data,computed"`
}

type V1CreditGrantsDataSourceModel struct {
	EffectiveBefore   timetypes.RFC3339                                                `tfsdk:"effective_before" json:"effective_before,optional" format:"date-time"`
	NotExpiringBefore timetypes.RFC3339                                                `tfsdk:"not_expiring_before" json:"not_expiring_before,optional" format:"date-time"`
	CreditGrantIDs    *[]types.String                                                  `tfsdk:"credit_grant_ids" json:"credit_grant_ids,optional"`
	CreditTypeIDs     *[]types.String                                                  `tfsdk:"credit_type_ids" json:"credit_type_ids,optional"`
	CustomerIDs       *[]types.String                                                  `tfsdk:"customer_ids" json:"customer_ids,optional"`
	MaxItems          types.Int64                                                      `tfsdk:"max_items"`
	Items             customfield.NestedObjectList[V1CreditGrantsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1CreditGrantsDataSourceModel) toListParams(_ context.Context) (params metronome.V1CreditGrantListParams, diags diag.Diagnostics) {
	params = metronome.V1CreditGrantListParams{}

	return
}

type V1CreditGrantsItemsDataSourceModel struct {
	ID                types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Balance           customfield.NestedObject[V1CreditGrantsBalanceDataSourceModel]               `tfsdk:"balance" json:"balance,computed"`
	CustomFields      customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerID        types.String                                                                 `tfsdk:"customer_id" json:"customer_id,computed"`
	Deductions        customfield.NestedObjectList[V1CreditGrantsDeductionsDataSourceModel]        `tfsdk:"deductions" json:"deductions,computed"`
	EffectiveAt       timetypes.RFC3339                                                            `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
	ExpiresAt         timetypes.RFC3339                                                            `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	GrantAmount       customfield.NestedObject[V1CreditGrantsGrantAmountDataSourceModel]           `tfsdk:"grant_amount" json:"grant_amount,computed"`
	Name              types.String                                                                 `tfsdk:"name" json:"name,computed"`
	PaidAmount        customfield.NestedObject[V1CreditGrantsPaidAmountDataSourceModel]            `tfsdk:"paid_amount" json:"paid_amount,computed"`
	PendingDeductions customfield.NestedObjectList[V1CreditGrantsPendingDeductionsDataSourceModel] `tfsdk:"pending_deductions" json:"pending_deductions,computed"`
	Priority          types.Float64                                                                `tfsdk:"priority" json:"priority,computed"`
	CreditGrantType   types.String                                                                 `tfsdk:"credit_grant_type" json:"credit_grant_type,computed"`
	InvoiceID         types.String                                                                 `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Products          customfield.NestedObjectList[V1CreditGrantsProductsDataSourceModel]          `tfsdk:"products" json:"products,computed"`
	Reason            types.String                                                                 `tfsdk:"reason" json:"reason,computed"`
	UniquenessKey     types.String                                                                 `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1CreditGrantsBalanceDataSourceModel struct {
	EffectiveAt      timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
	ExcludingPending types.Float64     `tfsdk:"excluding_pending" json:"excluding_pending,computed"`
	IncludingPending types.Float64     `tfsdk:"including_pending" json:"including_pending,computed"`
}

type V1CreditGrantsDeductionsDataSourceModel struct {
	Amount         types.Float64     `tfsdk:"amount" json:"amount,computed"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"created_by,computed"`
	CreditGrantID  types.String      `tfsdk:"credit_grant_id" json:"credit_grant_id,computed"`
	EffectiveAt    timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
	Reason         types.String      `tfsdk:"reason" json:"reason,computed"`
	RunningBalance types.Float64     `tfsdk:"running_balance" json:"running_balance,computed"`
	InvoiceID      types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
}

type V1CreditGrantsGrantAmountDataSourceModel struct {
	Amount     types.Float64                                                                `tfsdk:"amount" json:"amount,computed"`
	CreditType customfield.NestedObject[V1CreditGrantsGrantAmountCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1CreditGrantsGrantAmountCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CreditGrantsPaidAmountDataSourceModel struct {
	Amount     types.Float64                                                               `tfsdk:"amount" json:"amount,computed"`
	CreditType customfield.NestedObject[V1CreditGrantsPaidAmountCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1CreditGrantsPaidAmountCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CreditGrantsPendingDeductionsDataSourceModel struct {
	Amount         types.Float64     `tfsdk:"amount" json:"amount,computed"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"created_by,computed"`
	CreditGrantID  types.String      `tfsdk:"credit_grant_id" json:"credit_grant_id,computed"`
	EffectiveAt    timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
	Reason         types.String      `tfsdk:"reason" json:"reason,computed"`
	RunningBalance types.Float64     `tfsdk:"running_balance" json:"running_balance,computed"`
	InvoiceID      types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
}

type V1CreditGrantsProductsDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
