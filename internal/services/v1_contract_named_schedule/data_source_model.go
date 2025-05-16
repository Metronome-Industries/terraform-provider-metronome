// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_named_schedule

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1ContractNamedScheduleDataSourceModel struct {
	Data customfield.NestedObjectList[V1ContractNamedScheduleDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
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
