// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_commit

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CustomerCommitModel struct {
	CustomerID              types.String                                        `tfsdk:"customer_id" json:"customer_id,required"`
	Priority                types.Float64                                       `tfsdk:"priority" json:"priority,required"`
	ProductID               types.String                                        `tfsdk:"product_id" json:"product_id,required"`
	Type                    types.String                                        `tfsdk:"type" json:"type,required"`
	AccessSchedule          *V1CustomerCommitAccessScheduleModel                `tfsdk:"access_schedule" json:"access_schedule,required"`
	Description             types.String                                        `tfsdk:"description" json:"description,optional"`
	InvoiceContractID       types.String                                        `tfsdk:"invoice_contract_id" json:"invoice_contract_id,optional"`
	Name                    types.String                                        `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID    types.String                                        `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	RateType                types.String                                        `tfsdk:"rate_type" json:"rate_type,optional"`
	SalesforceOpportunityID types.String                                        `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,optional"`
	UniquenessKey           types.String                                        `tfsdk:"uniqueness_key" json:"uniqueness_key,optional"`
	ApplicableContractIDs   *[]types.String                                     `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,optional"`
	ApplicableProductIDs    *[]types.String                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags   *[]types.String                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields            *map[string]types.String                            `tfsdk:"custom_fields" json:"custom_fields,optional"`
	InvoiceSchedule         *V1CustomerCommitInvoiceScheduleModel               `tfsdk:"invoice_schedule" json:"invoice_schedule,optional"`
	Data                    customfield.NestedObject[V1CustomerCommitDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1CustomerCommitModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CustomerCommitModel) MarshalJSONForUpdate(state V1CustomerCommitModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CustomerCommitAccessScheduleModel struct {
	ScheduleItems *[]*V1CustomerCommitAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                         `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V1CustomerCommitAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V1CustomerCommitInvoiceScheduleModel struct {
	CreditTypeID      types.String                                           `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V1CustomerCommitInvoiceScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V1CustomerCommitInvoiceScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V1CustomerCommitInvoiceScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1CustomerCommitInvoiceScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1CustomerCommitDataModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
