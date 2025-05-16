// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_plan

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*V1CustomerPlansDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required: true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[V1CustomerPlansItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "the ID of the customer plan",
							Computed:    true,
						},
						"custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"plan_description": schema.StringAttribute{
							Computed: true,
						},
						"plan_id": schema.StringAttribute{
							Description: "the ID of the plan",
							Computed:    true,
						},
						"plan_name": schema.StringAttribute{
							Computed: true,
						},
						"starting_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"ending_before": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"net_payment_terms_days": schema.Float64Attribute{
							Computed: true,
						},
						"trial_info": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1CustomerPlansTrialInfoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ending_before": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"spending_caps": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1CustomerPlansTrialInfoSpendingCapsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Computed: true,
											},
											"amount_remaining": schema.Float64Attribute{
												Computed: true,
											},
											"credit_type": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1CustomerPlansTrialInfoSpendingCapsCreditTypeDataSourceModel](ctx),
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
							},
						},
					},
				},
			},
		},
	}
}

func (d *V1CustomerPlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1CustomerPlansDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
