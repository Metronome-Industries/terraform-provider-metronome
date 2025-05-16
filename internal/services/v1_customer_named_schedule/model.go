// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_named_schedule

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomerNamedScheduleModel struct {
	CustomerID   types.String                                                   `tfsdk:"customer_id" json:"customer_id,required,no_refresh"`
	ScheduleName types.String                                                   `tfsdk:"schedule_name" json:"schedule_name,required,no_refresh"`
	StartingAt   timetypes.RFC3339                                              `tfsdk:"starting_at" json:"starting_at,required,no_refresh" format:"date-time"`
	Value        jsontypes.Normalized                                           `tfsdk:"value" json:"value,required,no_refresh"`
	EndingBefore timetypes.RFC3339                                              `tfsdk:"ending_before" json:"ending_before,optional,no_refresh" format:"date-time"`
	Data         customfield.NestedObjectList[V1CustomerNamedScheduleDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1CustomerNamedScheduleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CustomerNamedScheduleModel) MarshalJSONForUpdate(state V1CustomerNamedScheduleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CustomerNamedScheduleDataModel struct {
	StartingAt   timetypes.RFC3339    `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Value        jsontypes.Normalized `tfsdk:"value" json:"value,computed"`
	EndingBefore timetypes.RFC3339    `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}
