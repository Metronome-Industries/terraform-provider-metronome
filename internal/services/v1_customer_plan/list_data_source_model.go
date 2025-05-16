// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_plan

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CustomerPlansDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1CustomerPlansItemsDataSourceModel] `json:"data,computed"`
}

type V1CustomerPlansDataSourceModel struct {
	CustomerID types.String                                                      `tfsdk:"customer_id" path:"customer_id,required"`
	MaxItems   types.Int64                                                       `tfsdk:"max_items"`
	Items      customfield.NestedObjectList[V1CustomerPlansItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1CustomerPlansDataSourceModel) toListParams(_ context.Context) (params metronome.V1CustomerPlanListParams, diags diag.Diagnostics) {
	params = metronome.V1CustomerPlanListParams{
		CustomerID: metronome.F(m.CustomerID.ValueString()),
	}

	return
}

type V1CustomerPlansItemsDataSourceModel struct {
	ID                  types.String                                                      `tfsdk:"id" json:"id,computed"`
	CustomFields        customfield.Map[types.String]                                     `tfsdk:"custom_fields" json:"custom_fields,computed"`
	PlanDescription     types.String                                                      `tfsdk:"plan_description" json:"plan_description,computed"`
	PlanID              types.String                                                      `tfsdk:"plan_id" json:"plan_id,computed"`
	PlanName            types.String                                                      `tfsdk:"plan_name" json:"plan_name,computed"`
	StartingOn          timetypes.RFC3339                                                 `tfsdk:"starting_on" json:"starting_on,computed" format:"date-time"`
	EndingBefore        timetypes.RFC3339                                                 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	NetPaymentTermsDays types.Float64                                                     `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	TrialInfo           customfield.NestedObject[V1CustomerPlansTrialInfoDataSourceModel] `tfsdk:"trial_info" json:"trial_info,computed"`
}

type V1CustomerPlansTrialInfoDataSourceModel struct {
	EndingBefore timetypes.RFC3339                                                                 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	SpendingCaps customfield.NestedObjectList[V1CustomerPlansTrialInfoSpendingCapsDataSourceModel] `tfsdk:"spending_caps" json:"spending_caps,computed"`
}

type V1CustomerPlansTrialInfoSpendingCapsDataSourceModel struct {
	Amount          types.Float64                                                                           `tfsdk:"amount" json:"amount,computed"`
	AmountRemaining types.Float64                                                                           `tfsdk:"amount_remaining" json:"amount_remaining,computed"`
	CreditType      customfield.NestedObject[V1CustomerPlansTrialInfoSpendingCapsCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1CustomerPlansTrialInfoSpendingCapsCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
