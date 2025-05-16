// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*V1ContractProductsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[V1ContractProductsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"current": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1ContractProductsCurrentDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
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
								"billable_metric_id": schema.StringAttribute{
									Computed: true,
								},
								"composite_product_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"composite_tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"exclude_free_usage": schema.BoolAttribute{
									Computed: true,
								},
								"is_refundable": schema.BoolAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"netsuite_internal_item_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"netsuite_overage_item_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"presentation_group_key": schema.ListAttribute{
									Description: "For USAGE products only. Groups usage line items on invoices. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"pricing_group_key": schema.ListAttribute{
									Description: "For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"quantity_conversion": schema.SingleNestedAttribute{
									Description: `Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".`,
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1ContractProductsCurrentQuantityConversionDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"conversion_factor": schema.Float64Attribute{
											Description: "The factor to multiply or divide the quantity by.",
											Computed:    true,
										},
										"operation": schema.StringAttribute{
											Description: "The operation to perform on the quantity\nAvailable values: \"MULTIPLY\", \"DIVIDE\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("MULTIPLY", "DIVIDE"),
											},
										},
										"name": schema.StringAttribute{
											Description: "Optional name for this conversion.",
											Computed:    true,
										},
									},
								},
								"quantity_rounding": schema.SingleNestedAttribute{
									Description: `Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.`,
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1ContractProductsCurrentQuantityRoundingDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"decimal_places": schema.Float64Attribute{
											Computed: true,
											Validators: []validator.Float64{
												float64validator.AtLeast(0),
											},
										},
										"rounding_method": schema.StringAttribute{
											Description: `Available values: "ROUND_UP", "ROUND_DOWN", "ROUND_HALF_UP".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"ROUND_UP",
													"ROUND_DOWN",
													"ROUND_HALF_UP",
												),
											},
										},
									},
								},
								"starting_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"initial": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1ContractProductsInitialDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
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
								"billable_metric_id": schema.StringAttribute{
									Computed: true,
								},
								"composite_product_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"composite_tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"exclude_free_usage": schema.BoolAttribute{
									Computed: true,
								},
								"is_refundable": schema.BoolAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"netsuite_internal_item_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"netsuite_overage_item_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"presentation_group_key": schema.ListAttribute{
									Description: "For USAGE products only. Groups usage line items on invoices. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"pricing_group_key": schema.ListAttribute{
									Description: "For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"quantity_conversion": schema.SingleNestedAttribute{
									Description: `Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".`,
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1ContractProductsInitialQuantityConversionDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"conversion_factor": schema.Float64Attribute{
											Description: "The factor to multiply or divide the quantity by.",
											Computed:    true,
										},
										"operation": schema.StringAttribute{
											Description: "The operation to perform on the quantity\nAvailable values: \"MULTIPLY\", \"DIVIDE\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("MULTIPLY", "DIVIDE"),
											},
										},
										"name": schema.StringAttribute{
											Description: "Optional name for this conversion.",
											Computed:    true,
										},
									},
								},
								"quantity_rounding": schema.SingleNestedAttribute{
									Description: `Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.`,
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1ContractProductsInitialQuantityRoundingDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"decimal_places": schema.Float64Attribute{
											Computed: true,
											Validators: []validator.Float64{
												float64validator.AtLeast(0),
											},
										},
										"rounding_method": schema.StringAttribute{
											Description: `Available values: "ROUND_UP", "ROUND_DOWN", "ROUND_HALF_UP".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"ROUND_UP",
													"ROUND_DOWN",
													"ROUND_HALF_UP",
												),
											},
										},
									},
								},
								"starting_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"type": schema.StringAttribute{
							Description: `Available values: "USAGE", "SUBSCRIPTION", "COMPOSITE", "FIXED", "PRO_SERVICE".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"USAGE",
									"SUBSCRIPTION",
									"COMPOSITE",
									"FIXED",
									"PRO_SERVICE",
								),
							},
						},
						"updates": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[V1ContractProductsUpdatesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"created_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"created_by": schema.StringAttribute{
										Computed: true,
									},
									"billable_metric_id": schema.StringAttribute{
										Computed: true,
									},
									"composite_product_ids": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"composite_tags": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"exclude_free_usage": schema.BoolAttribute{
										Computed: true,
									},
									"is_refundable": schema.BoolAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"netsuite_internal_item_id": schema.StringAttribute{
										Description: "This field's availability is dependent on your client's configuration.",
										Computed:    true,
									},
									"netsuite_overage_item_id": schema.StringAttribute{
										Description: "This field's availability is dependent on your client's configuration.",
										Computed:    true,
									},
									"presentation_group_key": schema.ListAttribute{
										Description: "For USAGE products only. Groups usage line items on invoices. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pricing_group_key": schema.ListAttribute{
										Description: "For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole. The superset of values in the pricing group key and presentation group key must be set as one compound group key on the billable metric.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"quantity_conversion": schema.SingleNestedAttribute{
										Description: `Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".`,
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1ContractProductsUpdatesQuantityConversionDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"conversion_factor": schema.Float64Attribute{
												Description: "The factor to multiply or divide the quantity by.",
												Computed:    true,
											},
											"operation": schema.StringAttribute{
												Description: "The operation to perform on the quantity\nAvailable values: \"MULTIPLY\", \"DIVIDE\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("MULTIPLY", "DIVIDE"),
												},
											},
											"name": schema.StringAttribute{
												Description: "Optional name for this conversion.",
												Computed:    true,
											},
										},
									},
									"quantity_rounding": schema.SingleNestedAttribute{
										Description: `Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.`,
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1ContractProductsUpdatesQuantityRoundingDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"decimal_places": schema.Float64Attribute{
												Computed: true,
												Validators: []validator.Float64{
													float64validator.AtLeast(0),
												},
											},
											"rounding_method": schema.StringAttribute{
												Description: `Available values: "ROUND_UP", "ROUND_DOWN", "ROUND_HALF_UP".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"ROUND_UP",
														"ROUND_DOWN",
														"ROUND_HALF_UP",
													),
												},
											},
										},
									},
									"starting_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"tags": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
						},
						"archived_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"custom_fields": schema.MapAttribute{
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

func (d *V1ContractProductsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1ContractProductsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
