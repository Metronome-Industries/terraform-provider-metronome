// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_audit_log

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1AuditLogsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1AuditLogsItemsDataSourceModel] `json:"data,computed"`
}

type V1AuditLogsDataSourceModel struct {
	EndingBefore timetypes.RFC3339                                             `tfsdk:"ending_before" query:"ending_before,optional" format:"date-time"`
	ResourceID   types.String                                                  `tfsdk:"resource_id" query:"resource_id,optional"`
	ResourceType types.String                                                  `tfsdk:"resource_type" query:"resource_type,optional"`
	Sort         types.String                                                  `tfsdk:"sort" query:"sort,optional"`
	StartingOn   timetypes.RFC3339                                             `tfsdk:"starting_on" query:"starting_on,optional" format:"date-time"`
	MaxItems     types.Int64                                                   `tfsdk:"max_items"`
	Items        customfield.NestedObjectList[V1AuditLogsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1AuditLogsDataSourceModel) toListParams(_ context.Context) (params metronome.V1AuditLogListParams, diags diag.Diagnostics) {
	mEndingBefore, errs := m.EndingBefore.ValueRFC3339Time()
	diags.Append(errs...)
	mStartingOn, errs := m.StartingOn.ValueRFC3339Time()
	diags.Append(errs...)

	params = metronome.V1AuditLogListParams{}

	if !m.EndingBefore.IsNull() {
		params.EndingBefore = metronome.F(mEndingBefore)
	}
	if !m.ResourceID.IsNull() {
		params.ResourceID = metronome.F(m.ResourceID.ValueString())
	}
	if !m.ResourceType.IsNull() {
		params.ResourceType = metronome.F(m.ResourceType.ValueString())
	}
	if !m.Sort.IsNull() {
		params.Sort = metronome.F(metronome.V1AuditLogListParamsSort(m.Sort.ValueString()))
	}
	if !m.StartingOn.IsNull() {
		params.StartingOn = metronome.F(mStartingOn)
	}

	return
}

type V1AuditLogsItemsDataSourceModel struct {
	ID           types.String                                                `tfsdk:"id" json:"id,computed"`
	Request      customfield.NestedObject[V1AuditLogsRequestDataSourceModel] `tfsdk:"request" json:"request,computed"`
	Timestamp    timetypes.RFC3339                                           `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Action       types.String                                                `tfsdk:"action" json:"action,computed"`
	Actor        customfield.NestedObject[V1AuditLogsActorDataSourceModel]   `tfsdk:"actor" json:"actor,computed"`
	Description  types.String                                                `tfsdk:"description" json:"description,computed"`
	ResourceID   types.String                                                `tfsdk:"resource_id" json:"resource_id,computed"`
	ResourceType types.String                                                `tfsdk:"resource_type" json:"resource_type,computed"`
	Status       types.String                                                `tfsdk:"status" json:"status,computed"`
}

type V1AuditLogsRequestDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	IP        types.String `tfsdk:"ip" json:"ip,computed"`
	UserAgent types.String `tfsdk:"user_agent" json:"user_agent,computed"`
}

type V1AuditLogsActorDataSourceModel struct {
	ID    types.String `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Email types.String `tfsdk:"email" json:"email,computed"`
}
