// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_named_schedule

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractNamedScheduleModel struct {
	RateCardID   types.String                                                   `tfsdk:"rate_card_id" json:"rate_card_id,required,no_refresh"`
	ScheduleName types.String                                                   `tfsdk:"schedule_name" json:"schedule_name,required,no_refresh"`
	StartingAt   timetypes.RFC3339                                              `tfsdk:"starting_at" json:"starting_at,required,no_refresh" format:"date-time"`
	Value        jsontypes.Normalized                                           `tfsdk:"value" json:"value,required,no_refresh"`
	EndingBefore timetypes.RFC3339                                              `tfsdk:"ending_before" json:"ending_before,optional,no_refresh" format:"date-time"`
	Data         customfield.NestedObjectList[V1ContractNamedScheduleDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1ContractNamedScheduleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1ContractNamedScheduleModel) MarshalJSONForUpdate(state V1ContractNamedScheduleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1ContractNamedScheduleDataModel struct {
	StartingAt   timetypes.RFC3339    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Value        jsontypes.Normalized `tfsdk:"value" json:"value,computed"`
	EndingBefore timetypes.RFC3339    `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}
