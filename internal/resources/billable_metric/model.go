// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billable_metric

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BillableMetricModel struct {
	BillableMetricID types.String                           `tfsdk:"billable_metric_id" path:"billable_metric_id"`
	AggregationType  types.String                           `tfsdk:"aggregation_type" json:"aggregation_type"`
	Name             types.String                           `tfsdk:"name" json:"name"`
	AggregationKey   types.String                           `tfsdk:"aggregation_key" json:"aggregation_key"`
	CustomFields     map[string]types.String                `tfsdk:"custom_fields" json:"custom_fields"`
	EventTypeFilter  *BillableMetricEventTypeFilterModel    `tfsdk:"event_type_filter" json:"event_type_filter"`
	GroupKeys        *[]*[]types.String                     `tfsdk:"group_keys" json:"group_keys"`
	PropertyFilters  *[]*BillableMetricPropertyFiltersModel `tfsdk:"property_filters" json:"property_filters"`
}

type BillableMetricEventTypeFilterModel struct {
	InValues    *[]types.String `tfsdk:"in_values" json:"in_values"`
	NotInValues *[]types.String `tfsdk:"not_in_values" json:"not_in_values"`
}

type BillableMetricPropertyFiltersModel struct {
	Name        types.String    `tfsdk:"name" json:"name"`
	Exists      types.Bool      `tfsdk:"exists" json:"exists"`
	InValues    *[]types.String `tfsdk:"in_values" json:"in_values"`
	NotInValues *[]types.String `tfsdk:"not_in_values" json:"not_in_values"`
}
