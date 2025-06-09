// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_named_schedule

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.ResourceWithConfigValidators = (*V1ContractNamedScheduleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"rate_card_id": schema.StringAttribute{
				Description: "ID of the rate card whose named schedule is to be updated",
				Required:    true,
			},
			"schedule_name": schema.StringAttribute{
				Description: "The identifier for the schedule to be updated",
				Required:    true,
			},
			"starting_at": schema.StringAttribute{
				Required:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"value": schema.StringAttribute{
				Description: "The value to set for the named schedule. The structure of this object is specific to the named schedule.",
				Required:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"ending_before": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"data": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[V1ContractNamedScheduleDataModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"starting_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"value": schema.StringAttribute{
							Computed:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"ending_before": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (r *V1ContractNamedScheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1ContractNamedScheduleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
