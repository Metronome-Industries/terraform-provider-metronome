// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1BillableMetricDataSourceModel struct {
	BillableMetricID types.String                                                  `tfsdk:"billable_metric_id" path:"billable_metric_id,required"`
	Data             customfield.NestedObject[V1BillableMetricDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1BillableMetricDataSourceModel) toReadParams(_ context.Context) (params metronome.V1BillableMetricGetParams, diags diag.Diagnostics) {
	params = metronome.V1BillableMetricGetParams{
		BillableMetricID: metronome.F(m.BillableMetricID.ValueString()),
	}

	return
}

type V1BillableMetricDataDataSourceModel struct {
	ID              types.String                                                                     `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                                     `tfsdk:"name" json:"name,computed"`
	AggregationKey  types.String                                                                     `tfsdk:"aggregation_key" json:"aggregation_key,computed"`
	AggregationType types.String                                                                     `tfsdk:"aggregation_type" json:"aggregation_type,computed"`
	ArchivedAt      timetypes.RFC3339                                                                `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields    customfield.Map[types.String]                                                    `tfsdk:"custom_fields" json:"custom_fields,computed"`
	EventTypeFilter customfield.NestedObject[V1BillableMetricDataEventTypeFilterDataSourceModel]     `tfsdk:"event_type_filter" json:"event_type_filter,computed"`
	GroupKeys       customfield.List[customfield.List[types.String]]                                 `tfsdk:"group_keys" json:"group_keys,computed"`
	PropertyFilters customfield.NestedObjectList[V1BillableMetricDataPropertyFiltersDataSourceModel] `tfsdk:"property_filters" json:"property_filters,computed"`
	Sql             types.String                                                                     `tfsdk:"sql" json:"sql,computed"`
}

type V1BillableMetricDataEventTypeFilterDataSourceModel struct {
	InValues    customfield.List[types.String] `tfsdk:"in_values" json:"in_values,computed"`
	NotInValues customfield.List[types.String] `tfsdk:"not_in_values" json:"not_in_values,computed"`
}

type V1BillableMetricDataPropertyFiltersDataSourceModel struct {
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Exists      types.Bool                     `tfsdk:"exists" json:"exists,computed"`
	InValues    customfield.List[types.String] `tfsdk:"in_values" json:"in_values,computed"`
	NotInValues customfield.List[types.String] `tfsdk:"not_in_values" json:"not_in_values,computed"`
}
