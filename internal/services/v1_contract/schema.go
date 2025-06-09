// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*V1ContractResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"starting_at": schema.StringAttribute{
				Description:   "inclusive contract start time",
				Required:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ending_before": schema.StringAttribute{
				Description:   "exclusive contract end time",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"multiplier_override_prioritization": schema.StringAttribute{
				Description: "Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first. If tiered overrides are used, prioritization must be explicit.\nAvailable values: \"LOWEST_MULTIPLIER\", \"EXPLICIT\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("LOWEST_MULTIPLIER", "EXPLICIT"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"net_payment_terms_days": schema.Float64Attribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"netsuite_sales_order_id": schema.StringAttribute{
				Description:   "This field's availability is dependent on your client's configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rate_card_alias": schema.StringAttribute{
				Description:   "Selects the rate card linked to the specified alias as of the contract's start date.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rate_card_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"salesforce_opportunity_id": schema.StringAttribute{
				Description:   "This field's availability is dependent on your client's configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"scheduled_charges_on_usage_invoices": schema.StringAttribute{
				Description: "Determines which scheduled and commit charges to consolidate onto the Contract's usage invoice. The charge's `timestamp` must match the usage invoice's `ending_before` date for consolidation to occur. This field cannot be modified after a Contract has been created. If this field is omitted, charges will appear on a separate invoice from usage charges.\nAvailable values: \"ALL\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ALL"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"total_contract_value": schema.Float64Attribute{
				Description:   "This field's availability is dependent on your client's configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Description:   "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"billing_provider_configuration": schema.SingleNestedAttribute{
				Description: "The billing provider configuration associated with a contract.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"billing_provider": schema.StringAttribute{
						Description: `Available values: "aws_marketplace", "azure_marketplace", "gcp_marketplace", "stripe", "netsuite".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"aws_marketplace",
								"azure_marketplace",
								"gcp_marketplace",
								"stripe",
								"netsuite",
							),
						},
					},
					"billing_provider_configuration_id": schema.StringAttribute{
						Description: "The Metronome ID of the billing provider configuration",
						Optional:    true,
					},
					"delivery_method": schema.StringAttribute{
						Description: `Available values: "direct_to_billing_provider", "aws_sqs", "tackle", "aws_sns".`,
						Optional:    true,
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
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"commits": schema.ListNestedAttribute{
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
									Description: "Defaults to USD (cents) if not passed",
									Optional:    true,
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
									Description: "Defaults to USD (cents) if not passed.",
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
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"credits": schema.ListNestedAttribute{
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
									Description: "Defaults to USD (cents) if not passed",
									Optional:    true,
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"discounts": schema.ListNestedAttribute{
				Description: "This field's availability is dependent on your client's configuration.",
				Optional:    true,
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
									Description: "Defaults to USD (cents) if not passed.",
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
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"overrides": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"starting_at": schema.StringAttribute{
							Description: "RFC 3339 timestamp indicating when the override will start applying (inclusive)",
							Required:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"applicable_product_tags": schema.ListAttribute{
							Description: "tags identifying products whose rates are being overridden. Cannot be used in conjunction with override_specifiers.",
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
							Description: "Indicates whether the override should only apply to commits. Defaults to `false`. If `true`, you can specify relevant commits in `override_specifiers` by passing `commit_ids`. if you do not specify `commit_ids`, then the override will apply when consuming any prepaid or postpaid commit.",
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
										Description: "Can only be used for commit specific overrides. Must be used in conjunction with one of product_id, product_tags, pricing_group_values, or presentation_group_values. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.",
										Optional:    true,
										ElementType: types.StringType,
									},
									"presentation_group_values": schema.MapAttribute{
										Description: "A map of group names to values. The override will only apply to line items with the specified presentation group values.",
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
										Description: "Can only be used for commit specific overrides. Must be used in conjunction with one of product_id, product_tags, pricing_group_values, or presentation_group_values. If provided, the override will only apply to credits created by the specified recurring credit ids.",
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
							Description: "ID of the product whose rate is being overridden. Cannot be used in conjunction with override_specifiers.",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
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
						Description: "Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"professional_services": schema.ListNestedAttribute{
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"recurring_commits": schema.ListNestedAttribute{
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
							Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
							Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"recurring_credits": schema.ListNestedAttribute{
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
							Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
							Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"reseller_royalties": schema.ListNestedAttribute{
				Description: "This field's availability is dependent on your client's configuration.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fraction": schema.Float64Attribute{
							Required: true,
						},
						"netsuite_reseller_id": schema.StringAttribute{
							Required: true,
						},
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
						"starting_at": schema.StringAttribute{
							Required:   true,
							CustomType: timetypes.RFC3339Type{},
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
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
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
						"reseller_contract_value": schema.Float64Attribute{
							Optional: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"scheduled_charges": schema.ListNestedAttribute{
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
									Description: "Defaults to USD (cents) if not passed.",
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
											Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MONTHLY",
													"QUARTERLY",
													"SEMI_ANNUAL",
													"ANNUAL",
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"spend_threshold_configuration": schema.SingleNestedAttribute{
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
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"subscriptions": schema.ListNestedAttribute{
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
							Description: "The initial quantity for the subscription. It must be non-negative value.",
							Required:    true,
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"transition": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"from_contract_id": schema.StringAttribute{
						Required: true,
					},
					"type": schema.StringAttribute{
						Description: "This field's available values may vary based on your client's configuration.\nAvailable values: \"SUPERSEDE\", \"RENEWAL\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("SUPERSEDE", "RENEWAL"),
						},
					},
					"future_invoice_behavior": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"trueup": schema.StringAttribute{
								Description: "Controls whether future trueup invoices are billed or removed. Default behavior is AS_IS if not specified.\nAvailable values: \"REMOVE\", \"AS_IS\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("REMOVE", "AS_IS"),
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"usage_filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"group_key": schema.StringAttribute{
						Required: true,
					},
					"group_values": schema.ListAttribute{
						Required:    true,
						ElementType: types.StringType,
					},
					"starting_at": schema.StringAttribute{
						Optional:   true,
						CustomType: timetypes.RFC3339Type{},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"usage_statement_schedule": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"frequency": schema.StringAttribute{
						Description: `Available values: "MONTHLY", "QUARTERLY", "ANNUAL", "WEEKLY".`,
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
					"billing_anchor_date": schema.StringAttribute{
						Description: "Required when using CUSTOM_DATE. This option lets you set a historical billing anchor date, aligning future billing cycles with a chosen cadence. For example, if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10 with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10. Subsequent statements will follow the 10th of each month.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"day": schema.StringAttribute{
						Description: "If not provided, defaults to the first day of the month.\nAvailable values: \"FIRST_OF_MONTH\", \"CONTRACT_START\", \"CUSTOM_DATE\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"FIRST_OF_MONTH",
								"CONTRACT_START",
								"CUSTOM_DATE",
							),
						},
					},
					"invoice_generation_starting_at": schema.StringAttribute{
						Description: "The date Metronome should start generating usage invoices. If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1ContractDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"amendments": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"commits": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsProductModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsAccessScheduleModel](ctx),
												Attributes: map[string]schema.Attribute{
													"schedule_items": schema.ListNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeModel](ctx),
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
											"amount": schema.Float64Attribute{
												Description: "(DEPRECATED) Use access_schedule + invoice_schedule instead.",
												Computed:    true,
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
												Description: "RFC 3339 timestamp indicating when the commit was archived. If not provided, the commit is not archived.",
												Computed:    true,
												CustomType:  timetypes.RFC3339Type{},
											},
											"balance": schema.Float64Attribute{
												Description: "The current balance of the credit or commit. This balance reflects the amount of credit or commit that the customer has access to use at this moment - thus, expired and upcoming credit or commit segments contribute 0 to the balance. The balance will match the sum of all ledger entries with the exception of the case where the sum of negative manual ledger entries exceeds the positive amount remaining on the credit or commit - in that case, the balance will be 0. All manual ledger entries associated with active credit or commit segments are included in the balance, including future-dated manual ledger entries.",
												Computed:    true,
											},
											"contract": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsContractModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceContractModel](ctx),
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"invoice_schedule": schema.SingleNestedAttribute{
												Description: "The schedule that the customer will be invoiced for this commit.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceScheduleModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsLedgerModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsRolledOverFromModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsSpecifiersModel](ctx),
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
											"uniqueness_key": schema.StringAttribute{
												Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
												Computed:    true,
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
								"overrides": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesModel](ctx),
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
											"credit_type": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesCreditTypeModel](ctx),
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"name": schema.StringAttribute{
														Computed: true,
													},
												},
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
											"is_prorated": schema.BoolAttribute{
												Description: "Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.",
												Computed:    true,
											},
											"multiplier": schema.Float64Attribute{
												Computed: true,
											},
											"override_specifiers": schema.ListNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverrideSpecifiersModel](ctx),
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
												CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverrideTiersModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesOverwriteRateModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeModel](ctx),
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
														CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverwriteRateTiersModel](ctx),
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
											"price": schema.Float64Attribute{
												Description: "Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.",
												Computed:    true,
											},
											"priority": schema.Float64Attribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesProductModel](ctx),
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
													"name": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"quantity": schema.Float64Attribute{
												Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
												Computed:    true,
											},
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
											"target": schema.StringAttribute{
												Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
												},
											},
											"tiers": schema.ListNestedAttribute{
												Description: "Only set for TIERED rate_type.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesTiersModel](ctx),
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
											"value": schema.MapAttribute{
												Description: "Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.",
												Computed:    true,
												CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
												ElementType: jsontypes.NormalizedType{},
											},
										},
									},
								},
								"scheduled_charges": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsScheduledChargesModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesProductModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesScheduleModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsModel](ctx),
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
								"credits": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsProductModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsAccessScheduleModel](ctx),
												Attributes: map[string]schema.Attribute{
													"schedule_items": schema.ListNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsContractModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsLedgerModel](ctx),
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
											"rate_type": schema.StringAttribute{
												Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
												},
											},
											"salesforce_opportunity_id": schema.StringAttribute{
												Description: "This field's availability is dependent on your client's configuration.",
												Computed:    true,
											},
											"specifiers": schema.ListNestedAttribute{
												Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsSpecifiersModel](ctx),
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
											"uniqueness_key": schema.StringAttribute{
												Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
												Computed:    true,
											},
										},
									},
								},
								"discounts": schema.ListNestedAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsDiscountsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsProductModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsScheduleModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsScheduleCreditTypeModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsDiscountsScheduleScheduleItemsModel](ctx),
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
								"netsuite_sales_order_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
								"professional_services": schema.ListNestedAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsProfessionalServicesModel](ctx),
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
								"reseller_royalties": schema.ListNestedAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsResellerRoyaltiesModel](ctx),
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
											"fraction": schema.Float64Attribute{
												Computed: true,
											},
											"gcp_account_id": schema.StringAttribute{
												Computed: true,
											},
											"gcp_offer_id": schema.StringAttribute{
												Computed: true,
											},
											"netsuite_reseller_id": schema.StringAttribute{
												Computed: true,
											},
											"reseller_contract_value": schema.Float64Attribute{
												Computed: true,
											},
											"starting_at": schema.StringAttribute{
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
								},
								"salesforce_opportunity_id": schema.StringAttribute{
									Description: "This field's availability is dependent on your client's configuration.",
									Computed:    true,
								},
							},
						},
					},
					"current": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsProductModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsAccessScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsAccessScheduleScheduleItemsModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsAccessScheduleCreditTypeModel](ctx),
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
										"amount": schema.Float64Attribute{
											Description: "(DEPRECATED) Use access_schedule + invoice_schedule instead.",
											Computed:    true,
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
											Description: "RFC 3339 timestamp indicating when the commit was archived. If not provided, the commit is not archived.",
											Computed:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"balance": schema.Float64Attribute{
											Description: "The current balance of the credit or commit. This balance reflects the amount of credit or commit that the customer has access to use at this moment - thus, expired and upcoming credit or commit segments contribute 0 to the balance. The balance will match the sum of all ledger entries with the exception of the case where the sum of negative manual ledger entries exceeds the positive amount remaining on the credit or commit - in that case, the balance will be 0. All manual ledger entries associated with active credit or commit segments are included in the balance, including future-dated manual ledger entries.",
											Computed:    true,
										},
										"contract": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceContractModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"invoice_schedule": schema.SingleNestedAttribute{
											Description: "The schedule that the customer will be invoiced for this commit.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsLedgerModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsRolledOverFromModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsSpecifiersModel](ctx),
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
										"uniqueness_key": schema.StringAttribute{
											Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
											Computed:    true,
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
							"overrides": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesModel](ctx),
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
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesCreditTypeModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
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
										"is_prorated": schema.BoolAttribute{
											Description: "Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.",
											Computed:    true,
										},
										"multiplier": schema.Float64Attribute{
											Computed: true,
										},
										"override_specifiers": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverrideSpecifiersModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverrideTiersModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesOverwriteRateModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesOverwriteRateCreditTypeModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverwriteRateTiersModel](ctx),
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
										"price": schema.Float64Attribute{
											Description: "Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.",
											Computed:    true,
										},
										"priority": schema.Float64Attribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesProductModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"quantity": schema.Float64Attribute{
											Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
											Computed:    true,
										},
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
										"target": schema.StringAttribute{
											Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
											},
										},
										"tiers": schema.ListNestedAttribute{
											Description: "Only set for TIERED rate_type.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesTiersModel](ctx),
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
										"value": schema.MapAttribute{
											Description: "Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.",
											Computed:    true,
											CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
											ElementType: jsontypes.NormalizedType{},
										},
									},
								},
							},
							"scheduled_charges": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentScheduledChargesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentScheduledChargesScheduleScheduleItemsModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentTransitionsModel](ctx),
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
							"usage_statement_schedule": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageStatementScheduleModel](ctx),
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
							"credits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsProductModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCreditsAccessScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsAccessScheduleScheduleItemsModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsAccessScheduleCreditTypeModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsLedgerModel](ctx),
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
										"rate_type": schema.StringAttribute{
											Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
											},
										},
										"salesforce_opportunity_id": schema.StringAttribute{
											Description: "This field's availability is dependent on your client's configuration.",
											Computed:    true,
										},
										"specifiers": schema.ListNestedAttribute{
											Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsSpecifiersModel](ctx),
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
										"uniqueness_key": schema.StringAttribute{
											Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
											Computed:    true,
										},
									},
								},
							},
							"discounts": schema.ListNestedAttribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentDiscountsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentDiscountsScheduleScheduleItemsModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitSpecifiersModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
										Description: "Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.",
										Computed:    true,
									},
								},
							},
							"professional_services": schema.ListNestedAttribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentProfessionalServicesModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCommitsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsAccessAmountModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsCommitDurationModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsInvoiceAmountModel](ctx),
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
											Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCommitsSpecifiersModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCreditsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsAccessAmountModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsCommitDurationModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsContractModel](ctx),
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
											Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCreditsSpecifiersModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentResellerRoyaltiesModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationCommitModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
							"total_contract_value": schema.Float64Attribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
							},
							"usage_filter": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterModel](ctx),
								Attributes: map[string]schema.Attribute{
									"current": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterCurrentModel](ctx),
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
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
									"initial": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterInitialModel](ctx),
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
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
									"updates": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentUsageFilterUpdatesModel](ctx),
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
													Computed:   true,
													CustomType: timetypes.RFC3339Type{},
												},
											},
										},
									},
								},
							},
						},
					},
					"customer_id": schema.StringAttribute{
						Computed: true,
					},
					"initial": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1ContractDataInitialModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsProductModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsAccessScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsAccessScheduleScheduleItemsModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsAccessScheduleCreditTypeModel](ctx),
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
										"amount": schema.Float64Attribute{
											Description: "(DEPRECATED) Use access_schedule + invoice_schedule instead.",
											Computed:    true,
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
											Description: "RFC 3339 timestamp indicating when the commit was archived. If not provided, the commit is not archived.",
											Computed:    true,
											CustomType:  timetypes.RFC3339Type{},
										},
										"balance": schema.Float64Attribute{
											Description: "The current balance of the credit or commit. This balance reflects the amount of credit or commit that the customer has access to use at this moment - thus, expired and upcoming credit or commit segments contribute 0 to the balance. The balance will match the sum of all ledger entries with the exception of the case where the sum of negative manual ledger entries exceeds the positive amount remaining on the credit or commit - in that case, the balance will be 0. All manual ledger entries associated with active credit or commit segments are included in the balance, including future-dated manual ledger entries.",
											Computed:    true,
										},
										"contract": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceContractModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"invoice_schedule": schema.SingleNestedAttribute{
											Description: "The schedule that the customer will be invoiced for this commit.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCommitsLedgerModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsRolledOverFromModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCommitsSpecifiersModel](ctx),
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
										"uniqueness_key": schema.StringAttribute{
											Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
											Computed:    true,
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
							"overrides": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesModel](ctx),
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
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesCreditTypeModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
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
										"is_prorated": schema.BoolAttribute{
											Description: "Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.",
											Computed:    true,
										},
										"multiplier": schema.Float64Attribute{
											Computed: true,
										},
										"override_specifiers": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverrideSpecifiersModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverrideTiersModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesOverwriteRateModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesOverwriteRateCreditTypeModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverwriteRateTiersModel](ctx),
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
										"price": schema.Float64Attribute{
											Description: "Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.",
											Computed:    true,
										},
										"priority": schema.Float64Attribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesProductModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"quantity": schema.Float64Attribute{
											Description: "Default quantity. For SUBSCRIPTION rate_type, this must be >=0.",
											Computed:    true,
										},
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
										"target": schema.StringAttribute{
											Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
											},
										},
										"tiers": schema.ListNestedAttribute{
											Description: "Only set for TIERED rate_type.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialOverridesTiersModel](ctx),
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
										"value": schema.MapAttribute{
											Description: "Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.",
											Computed:    true,
											CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
											ElementType: jsontypes.NormalizedType{},
										},
									},
								},
							},
							"scheduled_charges": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialScheduledChargesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialScheduledChargesScheduleScheduleItemsModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialTransitionsModel](ctx),
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
							"usage_statement_schedule": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageStatementScheduleModel](ctx),
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
							"credits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCreditsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsProductModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCreditsAccessScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCreditsAccessScheduleScheduleItemsModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsAccessScheduleCreditTypeModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCreditsLedgerModel](ctx),
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
										"rate_type": schema.StringAttribute{
											Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
											},
										},
										"salesforce_opportunity_id": schema.StringAttribute{
											Description: "This field's availability is dependent on your client's configuration.",
											Computed:    true,
										},
										"specifiers": schema.ListNestedAttribute{
											Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCreditsSpecifiersModel](ctx),
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
										"uniqueness_key": schema.StringAttribute{
											Description: "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
											Computed:    true,
										},
									},
								},
							},
							"discounts": schema.ListNestedAttribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialDiscountsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsScheduleModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsScheduleCreditTypeModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialDiscountsScheduleScheduleItemsModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitSpecifiersModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
										Description: "Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.",
										Computed:    true,
									},
								},
							},
							"professional_services": schema.ListNestedAttribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialProfessionalServicesModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCommitsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsAccessAmountModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsCommitDurationModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsContractModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsInvoiceAmountModel](ctx),
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
											Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCommitsSpecifiersModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCreditsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsAccessAmountModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsCommitDurationModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsProductModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsContractModel](ctx),
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
											Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).\nAvailable values: \"NONE\", \"FIRST\", \"LAST\", \"FIRST_AND_LAST\".",
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCreditsSpecifiersModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialResellerRoyaltiesModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationCommitModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
							"total_contract_value": schema.Float64Attribute{
								Description: "This field's availability is dependent on your client's configuration.",
								Computed:    true,
							},
							"usage_filter": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterModel](ctx),
								Attributes: map[string]schema.Attribute{
									"current": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterCurrentModel](ctx),
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
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
									"initial": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterInitialModel](ctx),
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
												Computed:   true,
												CustomType: timetypes.RFC3339Type{},
											},
										},
									},
									"updates": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialUsageFilterUpdatesModel](ctx),
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
													Computed:   true,
													CustomType: timetypes.RFC3339Type{},
												},
											},
										},
									},
								},
							},
						},
					},
					"archived_at": schema.StringAttribute{
						Description: "RFC 3339 timestamp indicating when the contract was archived. If not returned, the contract is not archived.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"custom_fields": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"customer_billing_provider_configuration": schema.SingleNestedAttribute{
						Description: "The billing provider configuration associated with a contract.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[V1ContractDataCustomerBillingProviderConfigurationModel](ctx),
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
							"id": schema.StringAttribute{
								Computed: true,
							},
							"configuration": schema.MapAttribute{
								Description: "Configuration for the billing provider. The structure of this object is specific to the billing provider.",
								Computed:    true,
								CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
								ElementType: jsontypes.NormalizedType{},
							},
						},
					},
					"prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationCommitModel](ctx),
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
										CustomType:  customfield.NewNestedObjectListType[V1ContractDataPrepaidBalanceThresholdConfigurationCommitSpecifiersModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
								Description: "Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.",
								Computed:    true,
							},
						},
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
						CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationCommitModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationPaymentGateConfigModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[V1ContractDataSubscriptionsModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V1ContractDataSubscriptionsProrationModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V1ContractDataSubscriptionsQuantityScheduleModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V1ContractDataSubscriptionsSubscriptionRateModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataSubscriptionsSubscriptionRateProductModel](ctx),
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
					"uniqueness_key": schema.StringAttribute{
						Description: "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *V1ContractResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1ContractResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
