// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v2_contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*V2ContractDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V2ContractDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"commits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsProductDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsAccessScheduleDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsAccessScheduleCreditTypeDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsContractDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceContractDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"invoice_schedule": schema.SingleNestedAttribute{
									Description: "The schedule that the customer will be invoiced for this commit.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceScheduleDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsInvoiceScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCommitsInvoiceScheduleScheduleItemsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCommitsLedgerDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataCommitsRolledOverFromDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesOverrideSpecifiersDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
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
									CustomType: customfield.NewNestedObjectListType[V2ContractDataOverridesOverrideTiersDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesOverwriteRateDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesOverwriteRateCreditTypeDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V2ContractDataOverridesOverwriteRateTiersDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataOverridesProductDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataScheduledChargesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesProductDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesScheduleDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataScheduledChargesScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V2ContractDataScheduledChargesScheduleScheduleItemsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataTransitionsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataUsageFilterDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V2ContractDataUsageStatementScheduleDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataCreditsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsProductDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataCreditsAccessScheduleDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"schedule_items": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[V2ContractDataCreditsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsAccessScheduleCreditTypeDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataCreditsContractDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[V2ContractDataCreditsLedgerDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[V2ContractDataCustomerBillingProviderConfigurationDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataDiscountsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"product": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsProductDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsScheduleDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"credit_type": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V2ContractDataDiscountsScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V2ContractDataDiscountsScheduleScheduleItemsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel](ctx),
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
								},
							},
							"is_enabled": schema.BoolAttribute{
								Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
								Computed:    true,
							},
							"payment_gate_config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataProfessionalServicesDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectListType[V2ContractDataRecurringCommitsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"access_amount": schema.SingleNestedAttribute{
									Description: "The amount of commit to grant.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsAccessAmountDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsCommitDurationDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCommitsProductDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCommitsContractDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCommitsInvoiceAmountDataSourceModel](ctx),
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
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
							},
						},
					},
					"recurring_credits": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V2ContractDataRecurringCreditsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"access_amount": schema.SingleNestedAttribute{
									Description: "The amount of commit to grant.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCreditsAccessAmountDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectType[V2ContractDataRecurringCreditsCommitDurationDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCreditsProductDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[V2ContractDataRecurringCreditsContractDataSourceModel](ctx),
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
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
							},
						},
					},
					"reseller_royalties": schema.ListNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[V2ContractDataResellerRoyaltiesDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V2ContractDataResellerRoyaltiesSegmentsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationCommitDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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

func (d *V2ContractDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V2ContractDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
