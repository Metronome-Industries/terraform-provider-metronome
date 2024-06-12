// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billable_metric

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r BillableMetricResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"billable_metric_id": schema.StringAttribute{
				Optional: true,
			},
			"aggregation_type": schema.StringAttribute{
				Description: "Specifies the type of aggregation performed on matching events.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("count", "Count", "COUNT", "latest", "Latest", "LATEST", "max", "Max", "MAX", "sum", "Sum", "SUM", "unique", "Unique", "UNIQUE"),
				},
			},
			"name": schema.StringAttribute{
				Description: "The display name of the billable metric.",
				Required:    true,
			},
			"aggregation_key": schema.StringAttribute{
				Description: "A key that specifies which property of the event is used to aggregate data. This key must be one of the property filter names and is not applicable when the aggregation type is 'count'.",
				Optional:    true,
			},
			"custom_fields": schema.MapAttribute{
				Description: "Custom fields to attach to the billable metric.",
				Optional:    true,
				ElementType: types.StringType,
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
			},
			"group_keys": schema.ListAttribute{
				Description: "Property names that are used to group usage costs on an invoice. Each entry represents a set of properties used to slice events into distinct buckets.",
				Optional:    true,
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
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
			},
		},
	}
}
