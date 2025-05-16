// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*V1BillableMetricDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"billable_metric_id": schema.StringAttribute{
				Required: true,
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1BillableMetricDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "ID of the billable metric",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The display name of the billable metric.",
						Computed:    true,
					},
					"aggregation_key": schema.StringAttribute{
						Description: "A key that specifies which property of the event is used to aggregate data. This key must be one of the property filter names and is not applicable when the aggregation type is 'count'.",
						Computed:    true,
					},
					"aggregation_type": schema.StringAttribute{
						Description: "Specifies the type of aggregation performed on matching events.\nAvailable values: \"COUNT\", \"LATEST\", \"MAX\", \"SUM\", \"UNIQUE\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"COUNT",
								"LATEST",
								"MAX",
								"SUM",
								"UNIQUE",
							),
						},
					},
					"archived_at": schema.StringAttribute{
						Description: "RFC 3339 timestamp indicating when the billable metric was archived. If not provided, the billable metric is not archived.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"custom_fields": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"event_type_filter": schema.SingleNestedAttribute{
						Description: "An optional filtering rule to match the 'event_type' property of an event.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[V1BillableMetricDataEventTypeFilterDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"in_values": schema.ListAttribute{
								Description: "A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"not_in_values": schema.ListAttribute{
								Description: "A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"group_keys": schema.ListAttribute{
						Description: "Property names that are used to group usage costs on an invoice. Each entry represents a set of properties used to slice events into distinct buckets.",
						Computed:    true,
						CustomType:  customfield.NewListType[customfield.List[types.String]](ctx),
						ElementType: types.ListType{
							ElemType: types.StringType,
						},
					},
					"property_filters": schema.ListNestedAttribute{
						Description: "A list of filters to match events to this billable metric. Each filter defines a rule on an event property. All rules must pass for the event to match the billable metric.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V1BillableMetricDataPropertyFiltersDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "The name of the event property.",
									Computed:    true,
								},
								"exists": schema.BoolAttribute{
									Description: "Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.",
									Computed:    true,
								},
								"in_values": schema.ListAttribute{
									Description: "Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"not_in_values": schema.ListAttribute{
									Description: "Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
					},
					"sql": schema.StringAttribute{
						Description: "The SQL query associated with the billable metric",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *V1BillableMetricDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V1BillableMetricDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
