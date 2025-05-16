// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*V1BillableMetricResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"billable_metric_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The display name of the billable metric.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aggregation_key": schema.StringAttribute{
				Description:   "Specifies the type of aggregation performed on matching events. Required if `sql` is not provided.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aggregation_type": schema.StringAttribute{
				Description: "Specifies the type of aggregation performed on matching events.\nAvailable values: \"COUNT\", \"LATEST\", \"MAX\", \"SUM\", \"UNIQUE\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"COUNT",
						"LATEST",
						"MAX",
						"SUM",
						"UNIQUE",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"sql": schema.StringAttribute{
				Description:   "The SQL query associated with the billable metric. This field is mutually exclusive with aggregation_type, event_type_filter, property_filters, aggregation_key, and group_keys. If provided, these other fields must be omitted.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Description:   "Custom fields to attach to the billable metric.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"group_keys": schema.ListAttribute{
				Description: "Property names that are used to group usage costs on an invoice. Each entry represents a set of properties used to slice events into distinct buckets.",
				Optional:    true,
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"event_type_filter": schema.SingleNestedAttribute{
				Description: "An optional filtering rule to match the 'event_type' property of an event.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"in_values": schema.ListAttribute{
						Description: "A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"not_in_values": schema.ListAttribute{
						Description: "A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"property_filters": schema.ListNestedAttribute{
				Description: "A list of filters to match events to this billable metric. Each filter defines a rule on an event property. All rules must pass for the event to match the billable metric.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The name of the event property.",
							Required:    true,
						},
						"exists": schema.BoolAttribute{
							Description: "Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.",
							Optional:    true,
						},
						"in_values": schema.ListAttribute{
							Description: "Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"not_in_values": schema.ListAttribute{
							Description: "Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.",
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1BillableMetricDataModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[V1BillableMetricDataEventTypeFilterModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[V1BillableMetricDataPropertyFiltersModel](ctx),
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

func (r *V1BillableMetricResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1BillableMetricResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
