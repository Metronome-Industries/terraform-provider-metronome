// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_rate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*V1ContractRateCardRatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
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
				CustomType:  customfield.NewNestedObjectListType[V1ContractRateCardRatesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"entitled": schema.BoolAttribute{
							Computed: true,
						},
						"product_custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"product_id": schema.StringAttribute{
							Computed: true,
						},
						"product_name": schema.StringAttribute{
							Computed: true,
						},
						"product_tags": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"rate": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1ContractRateCardRatesRateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"rate_type": schema.StringAttribute{
									Description: `Available values: "FLAT", "PERCENTAGE", "SUBSCRIPTION", "CUSTOM", "TIERED".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"FLAT",
											"PERCENTAGE",
											"SUBSCRIPTION",
											"CUSTOM",
											"TIERED",
										),
									},
								},
								"credit_type": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V1ContractRateCardRatesRateCreditTypeDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"custom_rate": schema.MapAttribute{
									Description: "Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.",
									Computed:    true,
									CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
									ElementType: jsontypes.NormalizedType{},
								},
								"is_prorated": schema.BoolAttribute{
									Description: "Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.",
									Computed:    true,
								},
								"price": schema.Float64Attribute{
									Description: "Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.",
									Computed:    true,
								},
								"pricing_group_values": schema.MapAttribute{
									Description: "if pricing groups are used, this will contain the values used to calculate the price",
									Computed:    true,
									CustomType:  customfield.NewMapType[types.String](ctx),
									ElementType: types.StringType,
								},
								"quantity": schema.Float64Attribute{
									Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
									Computed:    true,
								},
								"tiers": schema.ListNestedAttribute{
									Description: "Only set for TIERED rate_type.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V1ContractRateCardRatesRateTiersDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"price": schema.Float64Attribute{
												Computed: true,
											},
											"size": schema.Float64Attribute{
												Computed: true,
											},
										},
									},
								},
								"use_list_prices": schema.BoolAttribute{
									Description: "Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed using list prices rather than the standard rates for this product on the contract.",
									Computed:    true,
								},
							},
						},
						"starting_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"commit_rate": schema.SingleNestedAttribute{
							Description: "A distinct rate on the rate card. You can choose to use this rate rather than list rate when consuming a credit or commit.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[V1ContractRateCardRatesCommitRateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"rate_type": schema.StringAttribute{
									Description: `Available values: "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"FLAT",
											"PERCENTAGE",
											"SUBSCRIPTION",
											"TIERED",
											"CUSTOM",
										),
									},
								},
								"price": schema.Float64Attribute{
									Description: "Commit rate price. For FLAT rate_type, this must be >=0.",
									Computed:    true,
								},
								"tiers": schema.ListNestedAttribute{
									Description: "Only set for TIERED rate_type.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V1ContractRateCardRatesCommitRateTiersDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"price": schema.Float64Attribute{
												Computed: true,
											},
											"size": schema.Float64Attribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
						"ending_before": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"pricing_group_values": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *V1ContractRateCardRatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1ContractRateCardRatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
