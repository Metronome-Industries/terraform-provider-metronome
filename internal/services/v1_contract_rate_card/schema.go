// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*V1ContractRateCardResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"fiat_credit_type_id": schema.StringAttribute{
				Description:   "The Metronome ID of the credit type to associate with the rate card, defaults to USD (cents) if not passed.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"credit_type_conversions": schema.ListNestedAttribute{
				Description: "Required when using custom pricing units in rates.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"custom_credit_type_id": schema.StringAttribute{
							Required: true,
						},
						"fiat_per_custom_credit": schema.Float64Attribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Used only in UI/API. It is not exposed to end customers.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"rate_card_id": schema.StringAttribute{
				Description: "ID of the rate card to update",
				Optional:    true,
			},
			"aliases": schema.ListNestedAttribute{
				Description: "Reference this alias when creating a contract. If the same alias is assigned to multiple rate cards, it will reference the rate card to which it was most recently assigned. It is not exposed to end customers.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"ending_before": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"starting_at": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1ContractRateCardDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"created_by": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"aliases": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V1ContractRateCardDataAliasesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Computed: true,
								},
								"ending_before": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"starting_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
							},
						},
					},
					"credit_type_conversions": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V1ContractRateCardDataCreditTypeConversionsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"custom_credit_type": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V1ContractRateCardDataCreditTypeConversionsCustomCreditTypeModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"fiat_per_custom_credit": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"custom_fields": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"fiat_credit_type": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1ContractRateCardDataFiatCreditTypeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *V1ContractRateCardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1ContractRateCardResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
