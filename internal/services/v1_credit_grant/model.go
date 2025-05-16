// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CreditGrantModel struct {
	CustomerID       types.String                                     `tfsdk:"customer_id" json:"customer_id,required"`
	Priority         types.Float64                                    `tfsdk:"priority" json:"priority,required"`
	GrantAmount      *V1CreditGrantGrantAmountModel                   `tfsdk:"grant_amount" json:"grant_amount,required"`
	PaidAmount       *V1CreditGrantPaidAmountModel                    `tfsdk:"paid_amount" json:"paid_amount,required"`
	EffectiveAt      timetypes.RFC3339                                `tfsdk:"effective_at" json:"effective_at,optional" format:"date-time"`
	InvoiceDate      timetypes.RFC3339                                `tfsdk:"invoice_date" json:"invoice_date,optional" format:"date-time"`
	Reason           types.String                                     `tfsdk:"reason" json:"reason,optional"`
	UniquenessKey    types.String                                     `tfsdk:"uniqueness_key" json:"uniqueness_key,optional"`
	CustomFields     *map[string]types.String                         `tfsdk:"custom_fields" json:"custom_fields,optional"`
	ProductIDs       *[]types.String                                  `tfsdk:"product_ids" json:"product_ids,optional"`
	RolloverSettings *V1CreditGrantRolloverSettingsModel              `tfsdk:"rollover_settings" json:"rollover_settings,optional"`
	ExpiresAt        timetypes.RFC3339                                `tfsdk:"expires_at" json:"expires_at,required" format:"date-time"`
	Name             types.String                                     `tfsdk:"name" json:"name,required"`
	CreditGrantType  types.String                                     `tfsdk:"credit_grant_type" json:"credit_grant_type,optional"`
	ID               types.String                                     `tfsdk:"id" json:"id,optional"`
	Data             customfield.NestedObject[V1CreditGrantDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1CreditGrantModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CreditGrantModel) MarshalJSONForUpdate(state V1CreditGrantModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CreditGrantGrantAmountModel struct {
	Amount       types.Float64 `tfsdk:"amount" json:"amount,required"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
}

type V1CreditGrantPaidAmountModel struct {
	Amount       types.Float64 `tfsdk:"amount" json:"amount,required"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
}

type V1CreditGrantRolloverSettingsModel struct {
	ExpiresAt      timetypes.RFC3339                                 `tfsdk:"expires_at" json:"expires_at,required" format:"date-time"`
	Priority       types.Float64                                     `tfsdk:"priority" json:"priority,required"`
	RolloverAmount *V1CreditGrantRolloverSettingsRolloverAmountModel `tfsdk:"rollover_amount" json:"rollover_amount,required"`
}

type V1CreditGrantRolloverSettingsRolloverAmountModel struct {
	Type  types.String  `tfsdk:"type" json:"type,required"`
	Value types.Float64 `tfsdk:"value" json:"value,required"`
}

type V1CreditGrantDataModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}
