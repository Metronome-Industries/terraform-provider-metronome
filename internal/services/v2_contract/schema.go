// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v2_contract

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*V2ContractResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"contract_id": schema.StringAttribute{
				Description: "ID of the contract being edited",
				Required:    true,
			},
			"customer_id": schema.StringAttribute{
				Description: "ID of the customer whose contract is being edited",
				Required:    true,
			},
			"allow_contract_ending_before_finalized_invoice": schema.BoolAttribute{
				Description: "If true, allows setting the contract end date earlier than the end_timestamp of existing finalized invoices. Finalized invoices will be unchanged; if you want to incorporate the new end date, you can void and regenerate finalized usage invoices. Defaults to true.",
				Optional:    true,
			},
			"update_contract_end_date": schema.StringAttribute{
				Description: "RFC 3339 timestamp indicating when the contract will end (exclusive).",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"add_commits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"type": schema.StringAttribute{
							Description: `Available values: "PREPAID", "POSTPAID".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("PREPAID", "POSTPAID"),
							},
						},
						"access_schedule": schema.SingleNestedAttribute{
							Description: `Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.`,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"schedule_items": schema.ListNestedAttribute{
									Required: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Required: true,
											},
											"ending_before": schema.StringAttribute{
												Description: "RFC 3339 timestamp (exclusive)",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"starting_at": schema.StringAttribute{
												Description: "RFC 3339 timestamp (inclusive)",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
										},
									},
								},
								"credit_type_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
						"amount": schema.Float64Attribute{
							Description: "(DEPRECATED) Use access_schedule and invoice_schedule instead.",
							Optional:    true,
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Which products the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Which tags the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"custom_fields": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "Used only in UI/API. It is not exposed to end customers.",
							Optional:    true,
						},
						"invoice_schedule": schema.SingleNestedAttribute{
							Description: `Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.`,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Description: "Defaults to USD if not passed. Only USD is supported at this time.",
									Optional:    true,
								},
								"recurring_schedule": schema.SingleNestedAttribute{
									Description: "Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"amount_distribution": schema.StringAttribute{
											Description: `Available values: "DIVIDED", "DIVIDED_ROUNDED", "EACH".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"DIVIDED",
													"DIVIDED_ROUNDED",
													"EACH",
												),
											},
										},
										"ending_before": schema.StringAttribute{
											Description: "RFC 3339 timestamp (exclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"frequency": schema.StringAttribute{
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
													"WEEKLY",
												),
											},
										},
										"starting_at": schema.StringAttribute{
											Description: "RFC 3339 timestamp (inclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"amount": schema.Float64Attribute{
											Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
											Optional:    true,
										},
										"quantity": schema.Float64Attribute{
											Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
											Optional:    true,
										},
										"unit_price": schema.Float64Attribute{
											Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
											Optional:    true,
										},
									},
								},
								"schedule_items": schema.ListNestedAttribute{
									Description: "Either provide amount or provide both unit_price and quantity.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"timestamp": schema.StringAttribute{
												Description: "timestamp of the scheduled event",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"amount": schema.Float64Attribute{
												Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
												Optional:    true,
											},
											"quantity": schema.Float64Attribute{
												Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
												Optional:    true,
											},
											"unit_price": schema.Float64Attribute{
												Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
												Optional:    true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Optional:    true,
						},
						"payment_gate_config": schema.SingleNestedAttribute{
							Description: "optionally payment gate this commit",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"payment_gate_type": schema.StringAttribute{
									Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"NONE",
											"STRIPE",
											"EXTERNAL",
										),
									},
								},
								"stripe_config": schema.SingleNestedAttribute{
									Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"payment_type": schema.StringAttribute{
											Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
											},
										},
									},
								},
								"tax_type": schema.StringAttribute{
									Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
									},
								},
							},
						},
						"priority": schema.Float64Attribute{
							Description: "If multiple commits are applicable, the one with the lower priority will apply first.",
							Optional:    true,
						},
						"rate_type": schema.StringAttribute{
							Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
							},
						},
						"rollover_fraction": schema.Float64Attribute{
							Description: "Fraction of unused segments that will be rolled over. Must be between 0 and 1.",
							Optional:    true,
						},
						"specifiers": schema.ListNestedAttribute{
							Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"presentation_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "If provided, the specifier will only apply to the product with the specified ID.",
										Optional:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "If provided, the specifier will only apply to products with all the specified tags.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"temporary_id": schema.StringAttribute{
							Description: "A temporary ID for the commit that can be used to reference the commit for commit specific overrides.",
							Optional:    true,
						},
					},
				},
			},
			"add_credits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_schedule": schema.SingleNestedAttribute{
							Description: "Schedule for distributing the credit to the customer.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"schedule_items": schema.ListNestedAttribute{
									Required: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Required: true,
											},
											"ending_before": schema.StringAttribute{
												Description: "RFC 3339 timestamp (exclusive)",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"starting_at": schema.StringAttribute{
												Description: "RFC 3339 timestamp (inclusive)",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
										},
									},
								},
								"credit_type_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"custom_fields": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "Used only in UI/API. It is not exposed to end customers.",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Optional:    true,
						},
						"priority": schema.Float64Attribute{
							Description: "If multiple credits are applicable, the one with the lower priority will apply first.",
							Optional:    true,
						},
						"rate_type": schema.StringAttribute{
							Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
							},
						},
						"specifiers": schema.ListNestedAttribute{
							Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"presentation_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "If provided, the specifier will only apply to the product with the specified ID.",
										Optional:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "If provided, the specifier will only apply to products with all the specified tags.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
				},
			},
			"add_discounts": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"schedule": schema.SingleNestedAttribute{
							Description: "Must provide either schedule_items or recurring_schedule.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Description: "Defaults to USD if not passed. Only USD is supported at this time.",
									Optional:    true,
								},
								"recurring_schedule": schema.SingleNestedAttribute{
									Description: "Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"amount_distribution": schema.StringAttribute{
											Description: `Available values: "DIVIDED", "DIVIDED_ROUNDED", "EACH".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"DIVIDED",
													"DIVIDED_ROUNDED",
													"EACH",
												),
											},
										},
										"ending_before": schema.StringAttribute{
											Description: "RFC 3339 timestamp (exclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"frequency": schema.StringAttribute{
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
													"WEEKLY",
												),
											},
										},
										"starting_at": schema.StringAttribute{
											Description: "RFC 3339 timestamp (inclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"amount": schema.Float64Attribute{
											Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
											Optional:    true,
										},
										"quantity": schema.Float64Attribute{
											Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
											Optional:    true,
										},
										"unit_price": schema.Float64Attribute{
											Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
											Optional:    true,
										},
									},
								},
								"schedule_items": schema.ListNestedAttribute{
									Description: "Either provide amount or provide both unit_price and quantity.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"timestamp": schema.StringAttribute{
												Description: "timestamp of the scheduled event",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"amount": schema.Float64Attribute{
												Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
												Optional:    true,
											},
											"quantity": schema.Float64Attribute{
												Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
												Optional:    true,
											},
											"unit_price": schema.Float64Attribute{
												Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
												Optional:    true,
											},
										},
									},
								},
							},
						},
						"custom_fields": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Optional:    true,
						},
					},
				},
			},
			"add_overrides": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"starting_at": schema.StringAttribute{
							Description: "RFC 3339 timestamp indicating when the override will start applying (inclusive)",
							Required:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "tags identifying products whose rates are being overridden",
							Optional:    true,
							ElementType: types.StringType,
						},
						"ending_before": schema.StringAttribute{
							Description: "RFC 3339 timestamp indicating when the override will stop applying (exclusive)",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"entitled": schema.BoolAttribute{
							Optional: true,
						},
						"is_commit_specific": schema.BoolAttribute{
							Description: "Indicates whether the override should only apply to commits. Defaults to `false`. If `true`, you can specify relevant commits in `override_specifiers` by passing `commit_ids`.",
							Optional:    true,
						},
						"multiplier": schema.Float64Attribute{
							Description: "Required for MULTIPLIER type. Must be >=0.",
							Optional:    true,
						},
						"override_specifiers": schema.ListNestedAttribute{
							Description: "Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"billing_frequency": schema.StringAttribute{
										Description: `Available values: "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".`,
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"MONTHLY",
												"QUARTERLY",
												"ANNUAL",
												"WEEKLY",
											),
										},
									},
									"commit_ids": schema.ListAttribute{
										Description: "If provided, the override will only apply to the specified commits. Can only be used for commit specific overrides. If not provided, the override will apply to all commits.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"presentation_group_values": schema.MapAttribute{
										Description: "A map of group names to values. The override will only apply to line items with the specified presentation group values. Can only be used for multiplier overrides.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Description: "A map of pricing group names to values. The override will only apply to products with the specified pricing group values.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "If provided, the override will only apply to the product with the specified ID.",
										Optional:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "If provided, the override will only apply to products with all the specified tags.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"recurring_commit_ids": schema.ListAttribute{
										Description: "Can only be used for commit specific overrides. Must be used in conjunction with one of product_id, product_tags, pricing_group_values, or presentation_group_values. If provided, the override will only apply to commits created by the specified recurring commit ids.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"recurring_credit_ids": schema.ListAttribute{
										Description: "Can only be used for commit specific overrides. Must be used in conjunction with one of product_id, product_tags, pricing_group_values, or presentation_group_values. If provided, the override will only apply to commits created by the specified recurring credit ids.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"overwrite_rate": schema.SingleNestedAttribute{
							Description: "Required for OVERWRITE type.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"rate_type": schema.StringAttribute{
									Description: `Available values: "FLAT", "PERCENTAGE", "SUBSCRIPTION", "TIERED", "CUSTOM".`,
									Required:    true,
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
								"credit_type_id": schema.StringAttribute{
									Optional: true,
								},
								"custom_rate": schema.MapAttribute{
									Description: "Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.",
									Optional:    true,
									ElementType: jsontypes.NormalizedType{},
								},
								"is_prorated": schema.BoolAttribute{
									Description: "Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.",
									Optional:    true,
								},
								"price": schema.Float64Attribute{
									Description: "Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.",
									Optional:    true,
								},
								"quantity": schema.Float64Attribute{
									Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
									Optional:    true,
								},
								"tiers": schema.ListNestedAttribute{
									Description: "Only set for TIERED rate_type.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"price": schema.Float64Attribute{
												Required: true,
											},
											"size": schema.Float64Attribute{
												Optional: true,
											},
										},
									},
								},
							},
						},
						"priority": schema.Float64Attribute{
							Description: "Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.",
							Optional:    true,
						},
						"product_id": schema.StringAttribute{
							Description: "ID of the product whose rate is being overridden",
							Optional:    true,
						},
						"target": schema.StringAttribute{
							Description: "Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `\"LIST_RATE\"`.\nAvailable values: \"COMMIT_RATE\", \"LIST_RATE\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
							},
						},
						"tiers": schema.ListNestedAttribute{
							Description: "Required for TIERED type. Must have at least one tier.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"multiplier": schema.Float64Attribute{
										Required: true,
									},
									"size": schema.Float64Attribute{
										Optional: true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							Description: "Overwrites are prioritized over multipliers and tiered overrides.\nAvailable values: \"OVERWRITE\", \"MULTIPLIER\", \"TIERED\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"OVERWRITE",
									"MULTIPLIER",
									"TIERED",
								),
							},
						},
					},
				},
			},
			"add_prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"commit": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"product_id": schema.StringAttribute{
								Description: "The commit product that will be used to generate the line item for commit payment.",
								Required:    true,
							},
							"applicable_product_ids": schema.ListAttribute{
								Description: "Which products the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"applicable_product_tags": schema.ListAttribute{
								Description: "Which tags the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"description": schema.StringAttribute{
								Optional: true,
							},
							"name": schema.StringAttribute{
								Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
								Optional:    true,
							},
							"specifiers": schema.ListNestedAttribute{
								Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"presentation_group_values": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
										"pricing_group_values": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
										"product_id": schema.StringAttribute{
											Description: "If provided, the specifier will only apply to the product with the specified ID.",
											Optional:    true,
										},
										"product_tags": schema.ListAttribute{
											Description: "If provided, the specifier will only apply to products with all the specified tags.",
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"is_enabled": schema.BoolAttribute{
						Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
						Required:    true,
					},
					"payment_gate_config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"payment_gate_type": schema.StringAttribute{
								Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"NONE",
										"STRIPE",
										"EXTERNAL",
									),
								},
							},
							"stripe_config": schema.SingleNestedAttribute{
								Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"payment_type": schema.StringAttribute{
										Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
										Required:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
										},
									},
								},
							},
							"tax_type": schema.StringAttribute{
								Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
								},
							},
						},
					},
					"recharge_to_amount": schema.Float64Attribute{
						Description: "Specify the amount the balance should be recharged to.",
						Required:    true,
					},
					"threshold_amount": schema.Float64Attribute{
						Description: "Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.",
						Required:    true,
					},
				},
			},
			"add_professional_services": schema.ListNestedAttribute{
				Description: "This field's availability is dependent on your client's configuration.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"max_amount": schema.Float64Attribute{
							Description: "Maximum amount for the term.",
							Required:    true,
						},
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"quantity": schema.Float64Attribute{
							Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount.",
							Required:    true,
						},
						"unit_price": schema.Float64Attribute{
							Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.",
							Required:    true,
						},
						"custom_fields": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Optional:    true,
						},
					},
				},
			},
			"add_recurring_commits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_amount": schema.SingleNestedAttribute{
							Description: "The amount of commit to grant.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Required: true,
								},
								"quantity": schema.Float64Attribute{
									Required: true,
								},
								"unit_price": schema.Float64Attribute{
									Required: true,
								},
							},
						},
						"commit_duration": schema.SingleNestedAttribute{
							Description: `Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.`,
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"value": schema.Float64Attribute{
									Required: true,
								},
								"unit": schema.StringAttribute{
									Description: `Available values: "PERIODS".`,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("PERIODS"),
									},
								},
							},
						},
						"priority": schema.Float64Attribute{
							Description: "Will be passed down to the individual commits",
							Required:    true,
						},
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"starting_at": schema.StringAttribute{
							Description: "determines the start time for the first commit",
							Required:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
						},
						"ending_before": schema.StringAttribute{
							Description: "Determines when the contract will stop creating recurring commits. optional",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"invoice_amount": schema.SingleNestedAttribute{
							Description: "The amount the customer should be billed for the commit. Not required.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Required: true,
								},
								"quantity": schema.Float64Attribute{
									Required: true,
								},
								"unit_price": schema.Float64Attribute{
									Required: true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices. will be passed through to the individual commits",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
						},
						"proration": schema.StringAttribute{
							Description: "Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"NONE",
									"FIRST",
									"LAST",
									"FIRST_AND_LAST",
								),
							},
						},
						"rate_type": schema.StringAttribute{
							Description: "Whether the created commits will use the commit rate or list rate\nAvailable values: \"COMMIT_RATE\", \"LIST_RATE\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
							},
						},
						"recurrence_frequency": schema.StringAttribute{
							Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"MONTHLY",
									"QUARTERLY",
									"ANNUAL",
									"WEEKLY",
								),
							},
						},
						"rollover_fraction": schema.Float64Attribute{
							Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
							Optional:    true,
						},
						"specifiers": schema.ListNestedAttribute{
							Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"presentation_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "If provided, the specifier will only apply to the product with the specified ID.",
										Optional:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "If provided, the specifier will only apply to products with all the specified tags.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"temporary_id": schema.StringAttribute{
							Description: "A temporary ID that can be used to reference the recurring commit for commit specific overrides.",
							Optional:    true,
						},
					},
				},
			},
			"add_recurring_credits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_amount": schema.SingleNestedAttribute{
							Description: "The amount of commit to grant.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Required: true,
								},
								"quantity": schema.Float64Attribute{
									Required: true,
								},
								"unit_price": schema.Float64Attribute{
									Required: true,
								},
							},
						},
						"commit_duration": schema.SingleNestedAttribute{
							Description: `Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.`,
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"value": schema.Float64Attribute{
									Required: true,
								},
								"unit": schema.StringAttribute{
									Description: `Available values: "PERIODS".`,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("PERIODS"),
									},
								},
							},
						},
						"priority": schema.Float64Attribute{
							Description: "Will be passed down to the individual commits",
							Required:    true,
						},
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"starting_at": schema.StringAttribute{
							Description: "determines the start time for the first commit",
							Required:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
						},
						"ending_before": schema.StringAttribute{
							Description: "Determines when the contract will stop creating recurring commits. optional",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices. will be passed through to the individual commits",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "Will be passed down to the individual commits",
							Optional:    true,
						},
						"proration": schema.StringAttribute{
							Description: "Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"NONE",
									"FIRST",
									"LAST",
									"FIRST_AND_LAST",
								),
							},
						},
						"rate_type": schema.StringAttribute{
							Description: "Whether the created commits will use the commit rate or list rate\nAvailable values: \"COMMIT_RATE\", \"LIST_RATE\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
							},
						},
						"recurrence_frequency": schema.StringAttribute{
							Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"MONTHLY",
									"QUARTERLY",
									"ANNUAL",
									"WEEKLY",
								),
							},
						},
						"rollover_fraction": schema.Float64Attribute{
							Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
							Optional:    true,
						},
						"specifiers": schema.ListNestedAttribute{
							Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"presentation_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"pricing_group_values": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"product_id": schema.StringAttribute{
										Description: "If provided, the specifier will only apply to the product with the specified ID.",
										Optional:    true,
									},
									"product_tags": schema.ListAttribute{
										Description: "If provided, the specifier will only apply to products with all the specified tags.",
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"temporary_id": schema.StringAttribute{
							Description: "A temporary ID that can be used to reference the recurring commit for commit specific overrides.",
							Optional:    true,
						},
					},
				},
			},
			"add_reseller_royalties": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"reseller_type": schema.StringAttribute{
							Description: `Available values: "AWS", "AWS_PRO_SERVICE", "GCP", "GCP_PRO_SERVICE".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"AWS",
									"AWS_PRO_SERVICE",
									"GCP",
									"GCP_PRO_SERVICE",
								),
							},
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Must provide at least one of applicable_product_ids or applicable_product_tags.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Must provide at least one of applicable_product_ids or applicable_product_tags.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"aws_options": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"aws_account_number": schema.StringAttribute{
									Optional: true,
								},
								"aws_offer_id": schema.StringAttribute{
									Optional: true,
								},
								"aws_payer_reference_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
						"ending_before": schema.StringAttribute{
							Description: "Use null to indicate that the existing end timestamp should be removed.",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"fraction": schema.Float64Attribute{
							Optional: true,
						},
						"gcp_options": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"gcp_account_id": schema.StringAttribute{
									Optional: true,
								},
								"gcp_offer_id": schema.StringAttribute{
									Optional: true,
								},
							},
						},
						"netsuite_reseller_id": schema.StringAttribute{
							Optional: true,
						},
						"reseller_contract_value": schema.Float64Attribute{
							Optional: true,
						},
						"starting_at": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
			"add_scheduled_charges": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"schedule": schema.SingleNestedAttribute{
							Description: "Must provide either schedule_items or recurring_schedule.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"credit_type_id": schema.StringAttribute{
									Description: "Defaults to USD if not passed. Only USD is supported at this time.",
									Optional:    true,
								},
								"recurring_schedule": schema.SingleNestedAttribute{
									Description: "Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"amount_distribution": schema.StringAttribute{
											Description: `Available values: "DIVIDED", "DIVIDED_ROUNDED", "EACH".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"DIVIDED",
													"DIVIDED_ROUNDED",
													"EACH",
												),
											},
										},
										"ending_before": schema.StringAttribute{
											Description: "RFC 3339 timestamp (exclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"frequency": schema.StringAttribute{
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL", "WEEKLY".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
													"WEEKLY",
												),
											},
										},
										"starting_at": schema.StringAttribute{
											Description: "RFC 3339 timestamp (inclusive).",
											Required:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"amount": schema.Float64Attribute{
											Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
											Optional:    true,
										},
										"quantity": schema.Float64Attribute{
											Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
											Optional:    true,
										},
										"unit_price": schema.Float64Attribute{
											Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
											Optional:    true,
										},
									},
								},
								"schedule_items": schema.ListNestedAttribute{
									Description: "Either provide amount or provide both unit_price and quantity.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"timestamp": schema.StringAttribute{
												Description: "timestamp of the scheduled event",
												Required:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"amount": schema.Float64Attribute{
												Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
												Optional:    true,
											},
											"quantity": schema.Float64Attribute{
												Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
												Optional:    true,
											},
											"unit_price": schema.Float64Attribute{
												Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
												Optional:    true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "displayed on invoices",
							Optional:    true,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Description: "This field's availability is dependent on your client's configuration.",
							Optional:    true,
						},
					},
				},
			},
			"add_spend_threshold_configuration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"commit": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"product_id": schema.StringAttribute{
								Description: "The commit product that will be used to generate the line item for commit payment.",
								Required:    true,
							},
							"description": schema.StringAttribute{
								Optional: true,
							},
							"name": schema.StringAttribute{
								Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
								Optional:    true,
							},
						},
					},
					"is_enabled": schema.BoolAttribute{
						Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
						Required:    true,
					},
					"payment_gate_config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"payment_gate_type": schema.StringAttribute{
								Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"NONE",
										"STRIPE",
										"EXTERNAL",
									),
								},
							},
							"stripe_config": schema.SingleNestedAttribute{
								Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"payment_type": schema.StringAttribute{
										Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
										Required:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
										},
									},
								},
							},
							"tax_type": schema.StringAttribute{
								Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
								},
							},
						},
					},
					"threshold_amount": schema.Float64Attribute{
						Description: "Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.",
						Required:    true,
					},
				},
			},
			"add_subscriptions": schema.ListNestedAttribute{
				Description: "(beta) Optional list of [subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/) to add to the contract.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"collection_schedule": schema.StringAttribute{
							Description: `Available values: "ADVANCE", "ARREARS".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ADVANCE", "ARREARS"),
							},
						},
						"initial_quantity": schema.Float64Attribute{
							Required: true,
						},
						"proration": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"invoice_behavior": schema.StringAttribute{
									Description: "Indicates how mid-period quantity adjustments are invoiced. If BILL_IMMEDIATELY is selected, the quantity increase will be billed on the scheduled date. If BILL_ON_NEXT_COLLECTION_DATE is selected, the quantity increase will be billed for in-arrears at the end of the period.\nAvailable values: \"BILL_IMMEDIATELY\", \"BILL_ON_NEXT_COLLECTION_DATE\".",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE"),
									},
								},
								"is_prorated": schema.BoolAttribute{
									Description: "Indicates if the partial period will be prorated or charged a full amount.",
									Optional:    true,
								},
							},
						},
						"subscription_rate": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"billing_frequency": schema.StringAttribute{
									Description: "Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"MONTHLY",
											"QUARTERLY",
											"ANNUAL",
											"WEEKLY",
										),
									},
								},
								"product_id": schema.StringAttribute{
									Description: "Must be subscription type product",
									Required:    true,
								},
							},
						},
						"custom_fields": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
						"ending_before": schema.StringAttribute{
							Description: "Exclusive end time for the subscription. If not provided, subscription inherits contract end date.",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"starting_at": schema.StringAttribute{
							Description: "Inclusive start time for the subscription. If not provided, defaults to contract start date",
							Optional:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
			"archive_commits": schema.ListNestedAttribute{
				Description: "IDs of commits to archive",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"archive_credits": schema.ListNestedAttribute{
				Description: "IDs of credits to archive",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"archive_scheduled_charges": schema.ListNestedAttribute{
				Description: "IDs of scheduled charges to archive",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"remove_overrides": schema.ListNestedAttribute{
				Description: "IDs of overrides to remove",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"update_commits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"commit_id": schema.StringAttribute{
							Required: true,
						},
						"access_schedule": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"add_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Required: true,
											},
											"ending_before": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"starting_at": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
								},
								"remove_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
								"update_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
											"amount": schema.Float64Attribute{
												Optional: true,
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
							},
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Which products the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Which tags the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"invoice_schedule": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"add_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"timestamp": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"amount": schema.Float64Attribute{
												Optional: true,
											},
											"quantity": schema.Float64Attribute{
												Optional: true,
											},
											"unit_price": schema.Float64Attribute{
												Optional: true,
											},
										},
									},
								},
								"remove_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
								"update_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
											"amount": schema.Float64Attribute{
												Optional: true,
											},
											"quantity": schema.Float64Attribute{
												Optional: true,
											},
											"timestamp": schema.StringAttribute{
												Optional:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"unit_price": schema.Float64Attribute{
												Optional: true,
											},
										},
									},
								},
							},
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Optional: true,
						},
						"product_id": schema.StringAttribute{
							Optional: true,
						},
						"rollover_fraction": schema.Float64Attribute{
							Optional: true,
						},
					},
				},
			},
			"update_credits": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"credit_id": schema.StringAttribute{
							Required: true,
						},
						"access_schedule": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"add_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Required: true,
											},
											"ending_before": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"starting_at": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
								},
								"remove_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
								"update_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
											"amount": schema.Float64Attribute{
												Optional: true,
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
							},
						},
						"applicable_product_ids": schema.ListAttribute{
							Description: "Which products the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "Which tags the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Optional: true,
						},
						"product_id": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"update_prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"commit": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"applicable_product_ids": schema.ListAttribute{
								Description: "Which products the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"applicable_product_tags": schema.ListAttribute{
								Description: "Which tags the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"description": schema.StringAttribute{
								Optional: true,
							},
							"name": schema.StringAttribute{
								Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
								Optional:    true,
							},
							"product_id": schema.StringAttribute{
								Description: "The commit product that will be used to generate the line item for commit payment.",
								Optional:    true,
							},
							"specifiers": schema.ListNestedAttribute{
								Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"presentation_group_values": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
										"pricing_group_values": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
										},
										"product_id": schema.StringAttribute{
											Description: "If provided, the specifier will only apply to the product with the specified ID.",
											Optional:    true,
										},
										"product_tags": schema.ListAttribute{
											Description: "If provided, the specifier will only apply to products with all the specified tags.",
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
							},
						},
					},
					"is_enabled": schema.BoolAttribute{
						Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
						Optional:    true,
					},
					"payment_gate_config": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"payment_gate_type": schema.StringAttribute{
								Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"NONE",
										"STRIPE",
										"EXTERNAL",
									),
								},
							},
							"stripe_config": schema.SingleNestedAttribute{
								Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"payment_type": schema.StringAttribute{
										Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
										Required:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
										},
									},
								},
							},
							"tax_type": schema.StringAttribute{
								Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
								},
							},
						},
					},
					"recharge_to_amount": schema.Float64Attribute{
						Description: "Specify the amount the balance should be recharged to.",
						Optional:    true,
					},
					"threshold_amount": schema.Float64Attribute{
						Description: "Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.",
						Optional:    true,
					},
				},
			},
			"update_recurring_commits": schema.ListNestedAttribute{
				Description: "Edits to these recurring commits will only affect commits whose access schedules has not started. Expired commits, and commits with an active access schedule will remain unchanged.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"recurring_commit_id": schema.StringAttribute{
							Required: true,
						},
						"access_amount": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"quantity": schema.Float64Attribute{
									Optional: true,
								},
								"unit_price": schema.Float64Attribute{
									Optional: true,
								},
							},
						},
						"ending_before": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"invoice_amount": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"quantity": schema.Float64Attribute{
									Optional: true,
								},
								"unit_price": schema.Float64Attribute{
									Optional: true,
								},
							},
						},
					},
				},
			},
			"update_recurring_credits": schema.ListNestedAttribute{
				Description: "Edits to these recurring credits will only affect credits whose access schedules has not started. Expired credits, and credits with an active access schedule will remain unchanged.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"recurring_credit_id": schema.StringAttribute{
							Required: true,
						},
						"access_amount": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"quantity": schema.Float64Attribute{
									Optional: true,
								},
								"unit_price": schema.Float64Attribute{
									Optional: true,
								},
							},
						},
						"ending_before": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
			"update_scheduled_charges": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"scheduled_charge_id": schema.StringAttribute{
							Required: true,
						},
						"invoice_schedule": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"add_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"timestamp": schema.StringAttribute{
												Required:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"amount": schema.Float64Attribute{
												Optional: true,
											},
											"quantity": schema.Float64Attribute{
												Optional: true,
											},
											"unit_price": schema.Float64Attribute{
												Optional: true,
											},
										},
									},
								},
								"remove_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
								"update_schedule_items": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
											"amount": schema.Float64Attribute{
												Optional: true,
											},
											"quantity": schema.Float64Attribute{
												Optional: true,
											},
											"timestamp": schema.StringAttribute{
												Optional:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"unit_price": schema.Float64Attribute{
												Optional: true,
											},
										},
									},
								},
							},
						},
						"netsuite_sales_order_id": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"update_spend_threshold_configuration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"commit": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Optional: true,
							},
							"name": schema.StringAttribute{
								Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
								Optional:    true,
							},
							"product_id": schema.StringAttribute{
								Description: "The commit product that will be used to generate the line item for commit payment.",
								Optional:    true,
							},
						},
					},
					"is_enabled": schema.BoolAttribute{
						Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
						Optional:    true,
					},
					"payment_gate_config": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"payment_gate_type": schema.StringAttribute{
								Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"NONE",
										"STRIPE",
										"EXTERNAL",
									),
								},
							},
							"stripe_config": schema.SingleNestedAttribute{
								Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"payment_type": schema.StringAttribute{
										Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
										Required:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
										},
									},
								},
							},
							"tax_type": schema.StringAttribute{
								Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
								},
							},
						},
					},
					"threshold_amount": schema.Float64Attribute{
						Description: "Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.",
						Optional:    true,
					},
				},
			},
			"update_subscriptions": schema.ListNestedAttribute{
				Description: "(beta) Optional list of subscriptions to update.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"subscription_id": schema.StringAttribute{
							Required: true,
						},
						"ending_before": schema.StringAttribute{
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"quantity_updates": schema.ListNestedAttribute{
							Description: "Quantity changes are applied on the effective date based on the order which they are sent. For example, if I scheduled the quantity to be 12 on May 21 and then scheduled a quantity delta change of -1, the result from that day would be 11.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"starting_at": schema.StringAttribute{
										Required:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"quantity": schema.Float64Attribute{
										Description: "The new quantity for the subscription. Must be provided if quantity_delta is not provided. Must be non-negative.",
										Optional:    true,
									},
									"quantity_delta": schema.Float64Attribute{
										Description: "The delta to add to the subscription's quantity. Must be provided if quantity is not provided. Can't be zero. It also can't result in a negative quantity on the subscription.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V2ContractDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"commits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"type": schema.StringAttribute{
									Description: `Available values: "PREPAID", "POSTPAID".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("PREPAID", "POSTPAID"),
									},
								},
								"access_schedule": schema.SingleNestedAttribute{
									Description: "The schedule that the customer will gain access to the credits purposed with this commit.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsAccessScheduleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsAccessScheduleScheduleItemsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"amount": schema.Float64Attribute{
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
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsAccessScheduleCreditTypeModel](ctx),
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
								"applicable_contract_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"archived_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"balance": schema.Float64Attribute{
									Description: "The current balance of the credit or commit. This balance reflects the amount of credit or commit that the customer has access to use at this moment - thus, expired and upcoming credit or commit segments contribute 0 to the balance. The balance will match the sum of all ledger entries with the exception of the case where the sum of negative manual ledger entries exceeds the positive amount remaining on the credit or commit - in that case, the balance will be 0. All manual ledger entries associated with active credit or commit segments are included in the balance, including future-dated manual ledger entries.",
									Computed:    true,
								},
								"contract": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsContractModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
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
								"invoice_contract": schema.SingleNestedAttribute{
									Description: "The contract that this commit will be billed on.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceContractModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"invoice_schedule": schema.SingleNestedAttribute{
									Description: "The schedule that the customer will be invoiced for this commit.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceScheduleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceScheduleCreditTypeModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsInvoiceScheduleScheduleItemsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"amount": schema.Float64Attribute{
														Computed: true,
													},
													"invoice_id": schema.StringAttribute{
														Computed: true,
													},
													"quantity": schema.Float64Attribute{
														Computed: true,
													},
													"timestamp": schema.StringAttribute{
														Computed:   true,
														CustomType: timetypes.RFC3339Type{},
													},
													"unit_price": schema.Float64Attribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
								"ledger": schema.ListNestedAttribute{
									Description: "A list of ordered events that impact the balance of a commit. For example, an invoice deduction or a rollover.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCommitsLedgerModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Computed: true,
											},
											"segment_id": schema.StringAttribute{
												Computed: true,
											},
											"timestamp": schema.StringAttribute{
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"type": schema.StringAttribute{
												Description: `Available values: "PREPAID_COMMIT_SEGMENT_START", "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION", "PREPAID_COMMIT_ROLLOVER", "PREPAID_COMMIT_EXPIRATION", "PREPAID_COMMIT_CANCELED", "PREPAID_COMMIT_CREDITED", "POSTPAID_COMMIT_INITIAL_BALANCE", "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION", "POSTPAID_COMMIT_ROLLOVER", "POSTPAID_COMMIT_TRUEUP", "PREPAID_COMMIT_MANUAL", "POSTPAID_COMMIT_MANUAL", "POSTPAID_COMMIT_EXPIRATION".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"PREPAID_COMMIT_SEGMENT_START",
														"PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION",
														"PREPAID_COMMIT_ROLLOVER",
														"PREPAID_COMMIT_EXPIRATION",
														"PREPAID_COMMIT_CANCELED",
														"PREPAID_COMMIT_CREDITED",
														"POSTPAID_COMMIT_INITIAL_BALANCE",
														"POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION",
														"POSTPAID_COMMIT_ROLLOVER",
														"POSTPAID_COMMIT_TRUEUP",
														"PREPAID_COMMIT_MANUAL",
														"POSTPAID_COMMIT_MANUAL",
														"POSTPAID_COMMIT_EXPIRATION",
													),
												},
											},
											"invoice_id": schema.StringAttribute{
												Computed: true,
											},
											"new_contract_id": schema.StringAttribute{
												Computed: true,
											},
											"reason": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"priority": schema.Float64Attribute{
									Description: "If multiple credits or commits are applicable, the one with the lower priority will apply first.",
									Computed:    true,
								},
								"rate_type": schema.StringAttribute{
									Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
									},
								},
								"rolled_over_from": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsRolledOverFromModel](ctx),
									Attributes: map[string]schema.Attribute{
										"commit_id": schema.StringAttribute{
											Computed: true,
										},
										"contract_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"rollover_fraction": schema.Float64Attribute{
									Computed: true,
								},
								"salesforce_opportunity_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"specifiers": schema.ListNestedAttribute{
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCommitsSpecifiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"presentation_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
												Computed:    true,
											},
											"product_tags": schema.ListAttribute{
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"created_by": schema.StringAttribute{
						Computed: true,
					},
					"customer_id": schema.StringAttribute{
						Computed: true,
					},
					"overrides": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"starting_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"applicable_product_tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"ending_before": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"entitled": schema.BoolAttribute{
									Computed: true,
								},
								"is_commit_specific": schema.BoolAttribute{
									Computed: true,
								},
								"multiplier": schema.Float64Attribute{
									Computed: true,
								},
								"override_specifiers": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesOverrideSpecifiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"billing_frequency": schema.StringAttribute{
												Description: `Available values: "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"MONTHLY",
														"QUARTERLY",
														"ANNUAL",
														"WEEKLY",
													),
												},
											},
											"commit_ids": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"presentation_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Computed: true,
											},
											"product_tags": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"recurring_commit_ids": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"recurring_credit_ids": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
								"override_tiers": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesOverrideTiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"multiplier": schema.Float64Attribute{
												Computed: true,
											},
											"size": schema.Float64Attribute{
												Computed: true,
											},
										},
									},
								},
								"overwrite_rate": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesOverwriteRateModel](ctx),
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
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesOverwriteRateCreditTypeModel](ctx),
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
										"quantity": schema.Float64Attribute{
											Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
											Computed:    true,
										},
										"tiers": schema.ListNestedAttribute{
											Description: "Only set for TIERED rate_type.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[V2ContractDataOverridesOverwriteRateTiersModel](ctx),
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
								"priority": schema.Float64Attribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"target": schema.StringAttribute{
									Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
									},
								},
								"type": schema.StringAttribute{
									Description: `Available values: "OVERWRITE", "MULTIPLIER", "TIERED".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"OVERWRITE",
											"MULTIPLIER",
											"TIERED",
										),
									},
								},
							},
						},
					},
					"scheduled_charges": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataScheduledChargesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"schedule": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesScheduleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesScheduleCreditTypeModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataScheduledChargesScheduleScheduleItemsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"amount": schema.Float64Attribute{
														Computed: true,
													},
													"invoice_id": schema.StringAttribute{
														Computed: true,
													},
													"quantity": schema.Float64Attribute{
														Computed: true,
													},
													"timestamp": schema.StringAttribute{
														Computed:   true,
														CustomType: timetypes.RFC3339Type{},
													},
													"unit_price": schema.Float64Attribute{
														Computed: true,
													},
												},
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
								"name": schema.StringAttribute{
									Description: "displayed on invoices",
									Computed:    true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
							},
						},
					},
					"starting_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"transitions": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataTransitionsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_contract_id": schema.StringAttribute{
									Computed: true,
								},
								"to_contract_id": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Description: `Available values: "SUPERSEDE", "RENEWAL".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("SUPERSEDE", "RENEWAL"),
									},
								},
							},
						},
					},
					"usage_filter": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataUsageFilterModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group_key": schema.StringAttribute{
									Computed: true,
								},
								"group_values": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"starting_at": schema.StringAttribute{
									Description: "This will match contract starting_at value if usage filter is active from the beginning of the contract.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"ending_before": schema.StringAttribute{
									Description: "This will match contract ending_before value if usage filter is active until the end of the contract. It will be undefined if the contract is open-ended.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
							},
						},
					},
					"usage_statement_schedule": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V2ContractDataUsageStatementScheduleModel](ctx),
						Attributes: map[string]schema.Attribute{
							"billing_anchor_date": schema.StringAttribute{
								Description: "Contract usage statements follow a selected cadence based on this date.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"frequency": schema.StringAttribute{
								Description: `Available values: "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"MONTHLY",
										"QUARTERLY",
										"ANNUAL",
										"WEEKLY",
									),
								},
							},
						},
					},
					"archived_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"credits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataCreditsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"type": schema.StringAttribute{
									Description: `Available values: "CREDIT".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("CREDIT"),
									},
								},
								"access_schedule": schema.SingleNestedAttribute{
									Description: "The schedule that the customer will gain access to the credits.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCreditsAccessScheduleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCreditsAccessScheduleScheduleItemsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"amount": schema.Float64Attribute{
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
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsAccessScheduleCreditTypeModel](ctx),
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
								"applicable_contract_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_ids": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"balance": schema.Float64Attribute{
									Description: "The current balance of the credit or commit. This balance reflects the amount of credit or commit that the customer has access to use at this moment - thus, expired and upcoming credit or commit segments contribute 0 to the balance. The balance will match the sum of all ledger entries with the exception of the case where the sum of negative manual ledger entries exceeds the positive amount remaining on the credit or commit - in that case, the balance will be 0. All manual ledger entries associated with active credit or commit segments are included in the balance, including future-dated manual ledger entries.",
									Computed:    true,
								},
								"contract": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsContractModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
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
								"ledger": schema.ListNestedAttribute{
									Description: "A list of ordered events that impact the balance of a credit. For example, an invoice deduction or an expiration.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCreditsLedgerModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"amount": schema.Float64Attribute{
												Computed: true,
											},
											"segment_id": schema.StringAttribute{
												Computed: true,
											},
											"timestamp": schema.StringAttribute{
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"type": schema.StringAttribute{
												Description: `Available values: "CREDIT_SEGMENT_START", "CREDIT_AUTOMATED_INVOICE_DEDUCTION", "CREDIT_EXPIRATION", "CREDIT_CANCELED", "CREDIT_CREDITED", "CREDIT_MANUAL".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"CREDIT_SEGMENT_START",
														"CREDIT_AUTOMATED_INVOICE_DEDUCTION",
														"CREDIT_EXPIRATION",
														"CREDIT_CANCELED",
														"CREDIT_CREDITED",
														"CREDIT_MANUAL",
													),
												},
											},
											"invoice_id": schema.StringAttribute{
												Computed: true,
											},
											"reason": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"priority": schema.Float64Attribute{
									Description: "If multiple credits or commits are applicable, the one with the lower priority will apply first.",
									Computed:    true,
								},
								"salesforce_opportunity_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"specifiers": schema.ListNestedAttribute{
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCreditsSpecifiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"presentation_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
												Computed:    true,
											},
											"product_tags": schema.ListAttribute{
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
					},
					"custom_fields": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"customer_billing_provider_configuration": schema.SingleNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[V2ContractDataCustomerBillingProviderConfigurationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"billing_provider": schema.StringAttribute{
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
							"delivery_method": schema.StringAttribute{
								Description: `Available values: "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"direct_to_billing_provider",
										"aws_sqs",
										"tackle",
										"aws_sns",
									),
								},
							},
						},
					},
					"discounts": schema.ListNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataDiscountsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"schedule": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsScheduleModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsScheduleCreditTypeModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataDiscountsScheduleScheduleItemsModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"amount": schema.Float64Attribute{
														Computed: true,
													},
													"invoice_id": schema.StringAttribute{
														Computed: true,
													},
													"quantity": schema.Float64Attribute{
														Computed: true,
													},
													"timestamp": schema.StringAttribute{
														Computed:   true,
														CustomType: timetypes.RFC3339Type{},
													},
													"unit_price": schema.Float64Attribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
								"custom_fields": schema.MapAttribute{
									Computed:    true,
									CustomType:  customfield.NewMapType[types.String](ctx),
									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
							},
						},
					},
					"ending_before": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"multiplier_override_prioritization": schema.StringAttribute{
						Description: "Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.\nAvailable values: \"LOWEST_MULTIPLIER\", \"EXPLICIT\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("LOWEST_MULTIPLIER", "EXPLICIT"),
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"net_payment_terms_days": schema.Float64Attribute{
						Computed: true,
					},
					"netsuite_sales_order_id": schema.StringAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
					},
					"prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationCommitModel](ctx),
								Attributes: map[string]schema.Attribute{
									"product_id": schema.StringAttribute{
										Description: "The commit product that will be used to generate the line item for commit payment.",
										Computed:    true,
									},
									"applicable_product_ids": schema.ListAttribute{
										Description: "Which products the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"applicable_product_tags": schema.ListAttribute{
										Description: "Which tags the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
										Computed:    true,
									},
									"specifiers": schema.ListNestedAttribute{
										Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[V2ContractDataPrepaidBalanceThresholdConfigurationCommitSpecifiersModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"presentation_group_values": schema.MapAttribute{
													Computed:    true,
													CustomType:  customfield.NewMapType[types.String](ctx),
													ElementType: types.StringType,
												},
												"pricing_group_values": schema.MapAttribute{
													Computed:    true,
													CustomType:  customfield.NewMapType[types.String](ctx),
													ElementType: types.StringType,
												},
												"product_id": schema.StringAttribute{
													Description: "If provided, the specifier will only apply to the product with the specified ID.",
													Computed:    true,
												},
												"product_tags": schema.ListAttribute{
													Description: "If provided, the specifier will only apply to products with all the specified tags.",
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
											},
										},
									},
								},
							},
							"is_enabled": schema.BoolAttribute{
								Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
								Computed:    true,
							},
							"payment_gate_config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel](ctx),
								Attributes: map[string]schema.Attribute{
									"payment_gate_type": schema.StringAttribute{
										Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NONE",
												"STRIPE",
												"EXTERNAL",
											),
										},
									},
									"stripe_config": schema.SingleNestedAttribute{
										Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
										Attributes: map[string]schema.Attribute{
											"payment_type": schema.StringAttribute{
												Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
												},
											},
										},
									},
									"tax_type": schema.StringAttribute{
										Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
										},
									},
								},
							},
							"recharge_to_amount": schema.Float64Attribute{
								Description: "Specify the amount the balance should be recharged to.",
								Computed:    true,
							},
							"threshold_amount": schema.Float64Attribute{
								Description: "Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.",
								Computed:    true,
							},
						},
					},
					"professional_services": schema.ListNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataProfessionalServicesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"max_amount": schema.Float64Attribute{
									Description: "Maximum amount for the term.",
									Computed:    true,
								},
								"product_id": schema.StringAttribute{
									Computed: true,
								},
								"quantity": schema.Float64Attribute{
									Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount.",
									Computed:    true,
								},
								"unit_price": schema.Float64Attribute{
									Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.",
									Computed:    true,
								},
								"custom_fields": schema.MapAttribute{
									Computed:    true,
									CustomType:  customfield.NewMapType[types.String](ctx),
									ElementType: types.StringType,
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
							},
						},
					},
					"rate_card_id": schema.StringAttribute{
						Computed: true,
					},
					"recurring_commits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataRecurringCommitsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"access_amount": schema.SingleNestedAttribute{
									Description: "The amount of commit to grant.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsAccessAmountModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Computed: true,
										},
										"quantity": schema.Float64Attribute{
											Computed: true,
										},
										"unit_price": schema.Float64Attribute{
											Computed: true,
										},
									},
								},
								"commit_duration": schema.SingleNestedAttribute{
									Description: "The amount of time the created commits will be valid for",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsCommitDurationModel](ctx),
									Attributes: map[string]schema.Attribute{
										"value": schema.Float64Attribute{
											Computed: true,
										},
										"unit": schema.StringAttribute{
											Description: `Available values: "PERIODS".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("PERIODS"),
											},
										},
									},
								},
								"priority": schema.Float64Attribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCommitsProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"rate_type": schema.StringAttribute{
									Description: "Whether the created commits will use the commit rate or list rate\nAvailable values: \"COMMIT_RATE\", \"LIST_RATE\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
									},
								},
								"starting_at": schema.StringAttribute{
									Description: "Determines the start time for the first commit",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"applicable_product_ids": schema.ListAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"contract": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCommitsContractModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"description": schema.StringAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"ending_before": schema.StringAttribute{
									Description: "Determines when the contract will stop creating recurring commits. Optional",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"invoice_amount": schema.SingleNestedAttribute{
									Description: "The amount the customer should be billed for the commit. Not required.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsInvoiceAmountModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Computed: true,
										},
										"quantity": schema.Float64Attribute{
											Computed: true,
										},
										"unit_price": schema.Float64Attribute{
											Computed: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "Displayed on invoices. Will be passed through to the individual commits",
									Computed:    true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"proration": schema.StringAttribute{
									Description: "Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"NONE",
											"FIRST",
											"LAST",
											"FIRST_AND_LAST",
										),
									},
								},
								"recurrence_frequency": schema.StringAttribute{
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"MONTHLY",
											"QUARTERLY",
											"ANNUAL",
											"WEEKLY",
										),
									},
								},
								"rollover_fraction": schema.Float64Attribute{
									Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
									Computed:    true,
								},
								"specifiers": schema.ListNestedAttribute{
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataRecurringCommitsSpecifiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"presentation_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
												Computed:    true,
											},
											"product_tags": schema.ListAttribute{
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
					},
					"recurring_credits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataRecurringCreditsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"access_amount": schema.SingleNestedAttribute{
									Description: "The amount of commit to grant.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCreditsAccessAmountModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Computed: true,
										},
										"quantity": schema.Float64Attribute{
											Computed: true,
										},
										"unit_price": schema.Float64Attribute{
											Computed: true,
										},
									},
								},
								"commit_duration": schema.SingleNestedAttribute{
									Description: "The amount of time the created commits will be valid for",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCreditsCommitDurationModel](ctx),
									Attributes: map[string]schema.Attribute{
										"value": schema.Float64Attribute{
											Computed: true,
										},
										"unit": schema.StringAttribute{
											Description: `Available values: "PERIODS".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("PERIODS"),
											},
										},
									},
								},
								"priority": schema.Float64Attribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCreditsProductModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"rate_type": schema.StringAttribute{
									Description: "Whether the created commits will use the commit rate or list rate\nAvailable values: \"COMMIT_RATE\", \"LIST_RATE\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
									},
								},
								"starting_at": schema.StringAttribute{
									Description: "Determines the start time for the first commit",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"applicable_product_ids": schema.ListAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"contract": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCreditsContractModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"description": schema.StringAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"ending_before": schema.StringAttribute{
									Description: "Determines when the contract will stop creating recurring commits. Optional",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"name": schema.StringAttribute{
									Description: "Displayed on invoices. Will be passed through to the individual commits",
									Computed:    true,
								},
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "Will be passed down to the individual commits",
									Computed:    true,
								},
								"proration": schema.StringAttribute{
									Description: "Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"NONE",
											"FIRST",
											"LAST",
											"FIRST_AND_LAST",
										),
									},
								},
								"recurrence_frequency": schema.StringAttribute{
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"MONTHLY",
											"QUARTERLY",
											"ANNUAL",
											"WEEKLY",
										),
									},
								},
								"rollover_fraction": schema.Float64Attribute{
									Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
									Computed:    true,
								},
								"specifiers": schema.ListNestedAttribute{
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataRecurringCreditsSpecifiersModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"presentation_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Computed:    true,
												CustomType:  customfield.NewMapType[types.String](ctx),
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
												Computed:    true,
											},
											"product_tags": schema.ListAttribute{
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
					},
					"reseller_royalties": schema.ListNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataResellerRoyaltiesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
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
								"segments": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V2ContractDataResellerRoyaltiesSegmentsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"fraction": schema.Float64Attribute{
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
											"starting_at": schema.StringAttribute{
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"applicable_product_ids": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"applicable_product_tags": schema.ListAttribute{
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"aws_account_number": schema.StringAttribute{
												Computed: true,
											},
											"aws_offer_id": schema.StringAttribute{
												Computed: true,
											},
											"aws_payer_reference_id": schema.StringAttribute{
												Computed: true,
											},
											"ending_before": schema.StringAttribute{
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
											"gcp_account_id": schema.StringAttribute{
												Computed: true,
											},
											"gcp_offer_id": schema.StringAttribute{
												Computed: true,
											},
											"reseller_contract_value": schema.Float64Attribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
					"salesforce_opportunity_id": schema.StringAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
					},
					"scheduled_charges_on_usage_invoices": schema.StringAttribute{
						Description: "Determines which scheduled and commit charges to consolidate onto the Contract's usage invoice. The charge's `timestamp` must match the usage invoice's `ending_before` date for consolidation to occur. This field cannot be modified after a Contract has been created. If this field is omitted, charges will appear on a separate invoice from usage charges.\nAvailable values: \"ALL\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ALL"),
						},
					},
					"spend_threshold_configuration": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationCommitModel](ctx),
								Attributes: map[string]schema.Attribute{
									"product_id": schema.StringAttribute{
										Description: "The commit product that will be used to generate the line item for commit payment.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
										Computed:    true,
									},
								},
							},
							"is_enabled": schema.BoolAttribute{
								Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
								Computed:    true,
							},
							"payment_gate_config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationPaymentGateConfigModel](ctx),
								Attributes: map[string]schema.Attribute{
									"payment_gate_type": schema.StringAttribute{
										Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.\nAvailable values: \"NONE\", \"STRIPE\", \"EXTERNAL\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NONE",
												"STRIPE",
												"EXTERNAL",
											),
										},
									},
									"stripe_config": schema.SingleNestedAttribute{
										Description: "Only applicable if using Stripe as your payment gateway through Metronome.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
										Attributes: map[string]schema.Attribute{
											"payment_type": schema.StringAttribute{
												Description: "If left blank, will default to INVOICE\nAvailable values: \"INVOICE\", \"PAYMENT_INTENT\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("INVOICE", "PAYMENT_INTENT"),
												},
											},
										},
									},
									"tax_type": schema.StringAttribute{
										Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.\nAvailable values: \"NONE\", \"STRIPE\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("NONE", "STRIPE"),
										},
									},
								},
							},
							"threshold_amount": schema.Float64Attribute{
								Description: "Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.",
								Computed:    true,
							},
						},
					},
					"subscriptions": schema.ListNestedAttribute{
						Description: "(beta) List of subscriptions on the contract.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataSubscriptionsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"collection_schedule": schema.StringAttribute{
									Description: `Available values: "ADVANCE", "ARREARS".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("ADVANCE", "ARREARS"),
									},
								},
								"proration": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataSubscriptionsProrationModel](ctx),
									Attributes: map[string]schema.Attribute{
										"invoice_behavior": schema.StringAttribute{
											Description: `Available values: "BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("BILL_IMMEDIATELY", "BILL_ON_NEXT_COLLECTION_DATE"),
											},
										},
										"is_prorated": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"quantity_schedule": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V2ContractDataSubscriptionsQuantityScheduleModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"quantity": schema.Float64Attribute{
												Computed: true,
											},
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
								},
								"starting_at": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"subscription_rate": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataSubscriptionsSubscriptionRateModel](ctx),
									Attributes: map[string]schema.Attribute{
										"billing_frequency": schema.StringAttribute{
											Description: `Available values: "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"ANNUAL",
													"WEEKLY",
												),
											},
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataSubscriptionsSubscriptionRateProductModel](ctx),
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
								"id": schema.StringAttribute{
									Computed: true,
								},
								"custom_fields": schema.MapAttribute{
									Computed:    true,
									CustomType:  customfield.NewMapType[types.String](ctx),
									ElementType: types.StringType,
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"ending_before": schema.StringAttribute{
									Computed:   true,
									CustomType: timetypes.RFC3339Type{},
								},
								"fiat_credit_type_id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"total_contract_value": schema.Float64Attribute{
						Computed: true,
					},
					"uniqueness_key": schema.StringAttribute{
						Description: "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *V2ContractResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V2ContractResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
