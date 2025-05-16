// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_credit

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomerCreditModel struct {
	CustomerID              types.String                                        `tfsdk:"customer_id" json:"customer_id,required"`
	Priority                types.Float64                                       `tfsdk:"priority" json:"priority,required"`
	ProductID               types.String                                        `tfsdk:"product_id" json:"product_id,required"`
	AccessSchedule          *V1CustomerCreditAccessScheduleModel                `tfsdk:"access_schedule" json:"access_schedule,required"`
	Description             types.String                                        `tfsdk:"description" json:"description,optional"`
	Name                    types.String                                        `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID    types.String                                        `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	RateType                types.String                                        `tfsdk:"rate_type" json:"rate_type,optional"`
	SalesforceOpportunityID types.String                                        `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,optional"`
	UniquenessKey           types.String                                        `tfsdk:"uniqueness_key" json:"uniqueness_key,optional"`
	ApplicableContractIDs   *[]types.String                                     `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,optional"`
	ApplicableProductIDs    *[]types.String                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags   *[]types.String                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields            *map[string]types.String                            `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Data                    customfield.NestedObject[V1CustomerCreditDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1CustomerCreditModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CustomerCreditModel) MarshalJSONForUpdate(state V1CustomerCreditModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CustomerCreditAccessScheduleModel struct {
	ScheduleItems *[]*V1CustomerCreditAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                         `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V1CustomerCreditAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V1CustomerCreditDataModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
