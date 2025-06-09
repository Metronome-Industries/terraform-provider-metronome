// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_named_schedule

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractNamedScheduleDataSourceModel struct {
	RateCardID   types.String                                                             `tfsdk:"rate_card_id" json:"rate_card_id,required"`
	ScheduleName types.String                                                             `tfsdk:"schedule_name" json:"schedule_name,required"`
	CoveringDate timetypes.RFC3339                                                        `tfsdk:"covering_date" json:"covering_date,optional" format:"date-time"`
	Data         customfield.NestedObjectList[V1ContractNamedScheduleDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1ContractNamedScheduleDataSourceModel) toReadParams(_ context.Context) (params metronome.V1ContractNamedScheduleGetParams, diags diag.Diagnostics) {
	params = metronome.V1ContractNamedScheduleGetParams{}

	return
}

type V1ContractNamedScheduleDataDataSourceModel struct {
	StartingAt   timetypes.RFC3339    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Value        jsontypes.Normalized `tfsdk:"value" json:"value,computed"`
	EndingBefore timetypes.RFC3339    `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}
