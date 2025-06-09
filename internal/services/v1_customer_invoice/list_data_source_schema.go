// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_invoice

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*V1CustomerInvoicesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required: true,
			},
			"credit_type_id": schema.StringAttribute{
				Description: "Only return invoices for the specified credit type",
				Optional:    true,
			},
			"ending_before": schema.StringAttribute{
				Description: "RFC 3339 timestamp (exclusive). Invoices will only be returned for billing periods that end before this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"skip_zero_qty_line_items": schema.BoolAttribute{
				Description: "If set, all zero quantity line items will be filtered out of the response",
				Optional:    true,
			},
			"sort": schema.StringAttribute{
				Description: "Invoice sort order by issued_at, e.g. date_asc or date_desc.  Defaults to date_asc.\nAvailable values: \"date_asc\", \"date_desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("date_asc", "date_desc"),
				},
			},
			"starting_on": schema.StringAttribute{
				Description: "RFC 3339 timestamp (inclusive). Invoices will only be returned for billing periods that start at or after this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Invoice status, e.g. DRAFT, FINALIZED, or VOID",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[V1CustomerInvoicesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"credit_type": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesCreditTypeDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"customer_id": schema.StringAttribute{
							Computed: true,
						},
						"line_items": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[V1CustomerInvoicesLineItemsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"credit_type": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsCreditTypeDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"total": schema.Float64Attribute{
										Computed: true,
									},
									"applied_commit_or_credit": schema.SingleNestedAttribute{
										Description: "Details about the credit or commit that was applied to this line item. Only present on line items with product of `USAGE`, `SUBSCRIPTION` or `COMPOSITE` types.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsAppliedCommitOrCreditDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"type": schema.StringAttribute{
												Description: `Available values: "PREPAID", "POSTPAID", "CREDIT".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"PREPAID",
														"POSTPAID",
														"CREDIT",
													),
												},
											},
										},
									},
									"commit_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"commit_id": schema.StringAttribute{
										Description: "For line items with product of `USAGE`, `SUBSCRIPTION`, or `COMPOSITE` types, the ID of the credit or commit that was applied to this line item. For line items with product type of `FIXED`, the ID of the prepaid or postpaid commit that is being paid for.",
										Computed:    true,
									},
									"commit_netsuite_item_id": schema.StringAttribute{
										Computed: true,
									},
									"commit_netsuite_sales_order_id": schema.StringAttribute{
										Computed: true,
									},
									"commit_segment_id": schema.StringAttribute{
										Computed: true,
									},
									"commit_type": schema.StringAttribute{
										Description: "`PrepaidCommit` (for commit types `PREPAID` and `CREDIT`) or `PostpaidCommit` (for commit type `POSTPAID`).",
										Computed:    true,
									},
									"custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"discount_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"discount_id": schema.StringAttribute{
										Description: "ID of the discount applied to this line item.",
										Computed:    true,
									},
									"ending_before": schema.StringAttribute{
										Description: "The line item's end date (exclusive).",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"group_key": schema.StringAttribute{
										Computed: true,
									},
									"group_value": schema.StringAttribute{
										Computed: true,
									},
									"is_prorated": schema.BoolAttribute{
										Description: "Indicates whether the line item is prorated for `SUBSCRIPTION` type product.",
										Computed:    true,
									},
									"list_price": schema.SingleNestedAttribute{
										Description: "Only present for contract invoices and when the `include_list_prices` query parameter is set to true. This will include the list rate for the charge if applicable.  Only present for usage and subscription line items.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsListPriceDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsListPriceCreditTypeDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1CustomerInvoicesLineItemsListPriceTiersDataSourceModel](ctx),
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
									"metadata": schema.StringAttribute{
										Computed: true,
									},
									"netsuite_invoice_billing_end": schema.StringAttribute{
										Description: "The end date for the billing period on the invoice.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"netsuite_invoice_billing_start": schema.StringAttribute{
										Description: "The start date for the billing period on the invoice.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"netsuite_item_id": schema.StringAttribute{
										Computed: true,
									},
									"postpaid_commit": schema.SingleNestedAttribute{
										Description: "Only present for line items paying for a postpaid commit true-up.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsPostpaidCommitDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"presentation_group_values": schema.MapAttribute{
										Description: "Includes the presentation group values associated with this line item if presentation group keys are used.",
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Description: "Includes the pricing group values associated with this line item if dimensional pricing is used.",
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"product_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "ID of the product associated with the line item.",
										Computed:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "The current product tags associated with the line item's `product_id`.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"product_type": schema.StringAttribute{
										Description: "The type of the line item's product. Possible values are `FixedProductListItem` (for `FIXED` type products), `UsageProductListItem` (for `USAGE` type products), `SubscriptionProductListItem` (for `SUBSCRIPTION` type products) or `CompositeProductListItem` (for `COMPOSITE` type products).  For scheduled charges, commit and credit payments, the value is `FixedProductListItem`.",
										Computed:    true,
									},
									"professional_service_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"professional_service_id": schema.StringAttribute{
										Computed: true,
									},
									"quantity": schema.Float64Attribute{
										Description: "The quantity associated with the line item.",
										Computed:    true,
									},
									"reseller_type": schema.StringAttribute{
										Description: `Available values: "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".`,
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"AWS",
												"AWS_PRO_SERVICE",
												"GCP",
												"GCP_PRO_SERVICE",
											),
										},
									},
									"scheduled_charge_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"scheduled_charge_id": schema.StringAttribute{
										Description: "ID of scheduled charge.",
										Computed:    true,
									},
									"starting_at": schema.StringAttribute{
										Description: "The line item's start date (inclusive).",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"sub_line_items": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[V1CustomerInvoicesLineItemsSubLineItemsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"custom_fields": schema.MapAttribute{
													Computed:    true,
													CustomType:  customfield.NewMapType[types.String](ctx),
													ElementType: types.StringType,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
												"quantity": schema.Float64Attribute{
													Computed: true,
												},
												"subtotal": schema.Float64Attribute{
													Computed: true,
												},
												"charge_id": schema.StringAttribute{
													Computed: true,
												},
												"credit_grant_id": schema.StringAttribute{
													Computed: true,
												},
												"end_date": schema.StringAttribute{
													Description: "The end date for the charge (for seats charges only).",
													Computed:    true,
													CustomType:  timetypes.RFC3339Type{},
												},
												"price": schema.Float64Attribute{
													Description: "the unit price for this charge, present only if the charge is not tiered and the quantity is nonzero",
													Computed:    true,
												},
												"start_date": schema.StringAttribute{
													Description: "The start date for the charge (for seats charges only).",
													Computed:    true,
													CustomType:  timetypes.RFC3339Type{},
												},
												"tier_period": schema.SingleNestedAttribute{
													Description: "when the current tier started and ends (for tiered charges only)",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsSubLineItemsTierPeriodDataSourceModel](ctx),
													Attributes: map[string]schema.Attribute{
														"starting_at": schema.StringAttribute{
															Computed:   true,
															CustomType: timetypes.RFC3339Type{},
														},
														"ending_before": schema.StringAttribute{
															Computed:   true,
															CustomType: timetypes.RFC3339Type{},
														},
													},
												},
												"tiers": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1CustomerInvoicesLineItemsSubLineItemsTiersDataSourceModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"price": schema.Float64Attribute{
																Computed: true,
															},
															"quantity": schema.Float64Attribute{
																Computed: true,
															},
															"starting_at": schema.Float64Attribute{
																Description: "at what metric amount this tier begins",
																Computed:    true,
															},
															"subtotal": schema.Float64Attribute{
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"subscription_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"tier": schema.SingleNestedAttribute{
										Description: "Populated if the line item has a tiered price.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesLineItemsTierDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"level": schema.Float64Attribute{
												Computed: true,
											},
											"starting_at": schema.StringAttribute{
												Computed: true,
											},
											"size": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"unit_price": schema.Float64Attribute{
										Description: "The unit price associated with the line item.",
										Computed:    true,
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"total": schema.Float64Attribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
						"amendment_id": schema.StringAttribute{
							Computed: true,
						},
						"billable_status": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.\nAvailable values: \"billable\", \"unbillable\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("billable", "unbillable"),
							},
						},
						"contract_custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"contract_id": schema.StringAttribute{
							Computed: true,
						},
						"correction_record": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesCorrectionRecordDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"corrected_invoice_id": schema.StringAttribute{
									Computed: true,
								},
								"memo": schema.StringAttribute{
									Computed: true,
								},
								"reason": schema.StringAttribute{
									Computed: true,
								},
								"corrected_external_invoice": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesCorrectionRecordCorrectedExternalInvoiceDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"billing_provider_type": schema.StringAttribute{
											Description: `Available values: "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"aws_marketplace",
													"stripe",
													"netsuite",
													"custom",
													"azure_marketplace",
													"quickbooks_online",
													"workday",
													"gcp_marketplace",
												),
											},
										},
										"external_status": schema.StringAttribute{
											Description: `Available values: "DRAFT", "FINALIZED", "PAID", "UNCOLLECTIBLE", "VOID", "DELETED", "PAYMENT_FAILED", "INVALID_REQUEST_ERROR", "SKIPPED", "SENT", "QUEUED".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"DRAFT",
													"FINALIZED",
													"PAID",
													"UNCOLLECTIBLE",
													"VOID",
													"DELETED",
													"PAYMENT_FAILED",
													"INVALID_REQUEST_ERROR",
													"SKIPPED",
													"SENT",
													"QUEUED",
												),
											},
										},
										"invoice_id": schema.StringAttribute{
											Computed: true,
										},
										"issued_at_timestamp": schema.StringAttribute{
											Computed:   true,
											CustomType: timetypes.RFC3339Type{},
										},
									},
								},
							},
						},
						"created_at": schema.StringAttribute{
							Description: "When the invoice was created (UTC). This field is present for correction invoices only.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
							ElementType: jsontypes.NormalizedType{},
						},
						"customer_custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"end_timestamp": schema.StringAttribute{
							Description: "End of the usage period this invoice covers (UTC)",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"external_invoice": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesExternalInvoiceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"billing_provider_type": schema.StringAttribute{
									Description: `Available values: "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"aws_marketplace",
											"stripe",
											"netsuite",
											"custom",
											"azure_marketplace",
											"quickbooks_online",
											"workday",
											"gcp_marketplace",
										),
									},
								},
								"external_status": schema.StringAttribute{
									Description: `Available values: "DRAFT", "FINALIZED", "PAID", "UNCOLLECTIBLE", "VOID", "DELETED", "PAYMENT_FAILED", "INVALID_REQUEST_ERROR", "SKIPPED", "SENT", "QUEUED".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"DRAFT",
											"FINALIZED",
											"PAID",
											"UNCOLLECTIBLE",
											"VOID",
											"DELETED",
											"PAYMENT_FAILED",
											"INVALID_REQUEST_ERROR",
											"SKIPPED",
											"SENT",
											"QUEUED",
										),
									},
								},
								"invoice_id": schema.StringAttribute{
									Computed: true,
								},
								"issued_at_timestamp": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
							},
						},
						"invoice_adjustments": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[V1CustomerInvoicesInvoiceAdjustmentsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"credit_type": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesInvoiceAdjustmentsCreditTypeDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"total": schema.Float64Attribute{
										Computed: true,
									},
									"credit_grant_custom_fields": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"credit_grant_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"issued_at": schema.StringAttribute{
							Description: "When the invoice was issued (UTC)",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"net_payment_terms_days": schema.Float64Attribute{
							Computed: true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Computed:    true,
						},
						"plan_custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"plan_id": schema.StringAttribute{
							Computed: true,
						},
						"plan_name": schema.StringAttribute{
							Computed: true,
						},
						"reseller_royalty": schema.SingleNestedAttribute{
							Description: "Only present for contract invoices with reseller royalties.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[V1CustomerInvoicesResellerRoyaltyDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"fraction": schema.StringAttribute{
									Computed: true,
								},
								"netsuite_reseller_id": schema.StringAttribute{
									Computed: true,
								},
								"reseller_type": schema.StringAttribute{
									Description: `Available values: "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"AWS",
											"AWS_PRO_SERVICE",
											"GCP",
											"GCP_PRO_SERVICE",
										),
									},
								},
								"aws_options": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesResellerRoyaltyAwsOptionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"aws_account_number": schema.StringAttribute{
											Computed: true,
										},
										"aws_offer_id": schema.StringAttribute{
											Computed: true,
										},
										"aws_payer_reference_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"gcp_options": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V1CustomerInvoicesResellerRoyaltyGcpOptionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"gcp_account_id": schema.StringAttribute{
											Computed: true,
										},
										"gcp_offer_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"salesforce_opportunity_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Computed:    true,
						},
						"start_timestamp": schema.StringAttribute{
							Description: "Beginning of the usage period this invoice covers (UTC)",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"subtotal": schema.Float64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *V1CustomerInvoicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1CustomerInvoicesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
