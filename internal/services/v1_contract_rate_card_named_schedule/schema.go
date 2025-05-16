// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_named_schedule

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.ResourceWithConfigValidators = (*V1ContractRateCardNamedScheduleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"contract_id": schema.StringAttribute{
				Description: "ID of the contract whose named schedule is to be updated",
				Required:    true,
			},
			"customer_id": schema.StringAttribute{
				Description: "ID of the customer whose named schedule is to be updated",
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
				CustomType: customfield.NewNestedObjectListType[V1ContractRateCardNamedScheduleDataModel](ctx),
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

func (r *V1ContractRateCardNamedScheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1ContractRateCardNamedScheduleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
