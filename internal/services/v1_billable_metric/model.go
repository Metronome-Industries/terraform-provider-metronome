// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1BillableMetricModel struct {
	BillableMetricID types.String                                        `tfsdk:"billable_metric_id" path:"billable_metric_id,optional"`
	Name             types.String                                        `tfsdk:"name" json:"name,required,no_refresh"`
	AggregationKey   types.String                                        `tfsdk:"aggregation_key" json:"aggregation_key,optional,no_refresh"`
	AggregationType  types.String                                        `tfsdk:"aggregation_type" json:"aggregation_type,optional,no_refresh"`
	Sql              types.String                                        `tfsdk:"sql" json:"sql,optional,no_refresh"`
	CustomFields     *map[string]types.String                            `tfsdk:"custom_fields" json:"custom_fields,optional,no_refresh"`
	GroupKeys        *[]*[]types.String                                  `tfsdk:"group_keys" json:"group_keys,optional,no_refresh"`
	EventTypeFilter  *V1BillableMetricEventTypeFilterModel               `tfsdk:"event_type_filter" json:"event_type_filter,optional,no_refresh"`
	PropertyFilters  *[]*V1BillableMetricPropertyFiltersModel            `tfsdk:"property_filters" json:"property_filters,optional,no_refresh"`
	Data             customfield.NestedObject[V1BillableMetricDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1BillableMetricModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1BillableMetricModel) MarshalJSONForUpdate(state V1BillableMetricModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1BillableMetricEventTypeFilterModel struct {
	InValues    *[]types.String `tfsdk:"in_values" json:"in_values,optional"`
	NotInValues *[]types.String `tfsdk:"not_in_values" json:"not_in_values,optional"`
}

type V1BillableMetricPropertyFiltersModel struct {
	Name        types.String    `tfsdk:"name" json:"name,required"`
	Exists      types.Bool      `tfsdk:"exists" json:"exists,optional"`
	InValues    *[]types.String `tfsdk:"in_values" json:"in_values,optional"`
	NotInValues *[]types.String `tfsdk:"not_in_values" json:"not_in_values,optional"`
}

type V1BillableMetricDataModel struct {
	ID              types.String                                                           `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                           `tfsdk:"name" json:"name,computed"`
	AggregationKey  types.String                                                           `tfsdk:"aggregation_key" json:"aggregation_key,computed"`
	AggregationType types.String                                                           `tfsdk:"aggregation_type" json:"aggregation_type,computed"`
	ArchivedAt      timetypes.RFC3339                                                      `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields    customfield.Map[types.String]                                          `tfsdk:"custom_fields" json:"custom_fields,computed"`
	EventTypeFilter customfield.NestedObject[V1BillableMetricDataEventTypeFilterModel]     `tfsdk:"event_type_filter" json:"event_type_filter,computed"`
	GroupKeys       customfield.List[customfield.List[types.String]]                       `tfsdk:"group_keys" json:"group_keys,computed"`
	PropertyFilters customfield.NestedObjectList[V1BillableMetricDataPropertyFiltersModel] `tfsdk:"property_filters" json:"property_filters,computed"`
	Sql             types.String                                                           `tfsdk:"sql" json:"sql,computed"`
}

type V1BillableMetricDataEventTypeFilterModel struct {
	InValues    customfield.List[types.String] `tfsdk:"in_values" json:"in_values,computed"`
	NotInValues customfield.List[types.String] `tfsdk:"not_in_values" json:"not_in_values,computed"`
}

type V1BillableMetricDataPropertyFiltersModel struct {
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Exists      types.Bool                     `tfsdk:"exists" json:"exists,computed"`
	InValues    customfield.List[types.String] `tfsdk:"in_values" json:"in_values,computed"`
	NotInValues customfield.List[types.String] `tfsdk:"not_in_values" json:"not_in_values,computed"`
}
