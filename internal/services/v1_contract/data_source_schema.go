// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract

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

var _ datasource.DataSourceWithConfigValidators = (*V1ContractDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1ContractDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"amendments": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"commits": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsProductDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsAccessScheduleDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"schedule_items": schema.ListNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsContractDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceContractDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"invoice_schedule": schema.SingleNestedAttribute{
												Description: "The schedule that the customer will be invoiced for this commit.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceScheduleDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCommitsLedgerDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCommitsRolledOverFromDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesCreditTypeDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverrideSpecifiersDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverrideTiersDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesOverwriteRateDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeDataSourceModel](ctx),
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
														CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesOverwriteRateTiersDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsOverridesProductDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsOverridesTiersDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsScheduledChargesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesProductDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesScheduleDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsProductDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsAccessScheduleDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"schedule_items": schema.ListNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsCreditsContractDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsCreditsLedgerDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsDiscountsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"product": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsProductDataSourceModel](ctx),
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
												CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsScheduleDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"credit_type": schema.SingleNestedAttribute{
														Computed:   true,
														CustomType: customfield.NewNestedObjectType[V1ContractDataAmendmentsDiscountsScheduleCreditTypeDataSourceModel](ctx),
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
														CustomType: customfield.NewNestedObjectListType[V1ContractDataAmendmentsDiscountsScheduleScheduleItemsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsProfessionalServicesDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[V1ContractDataAmendmentsResellerRoyaltiesDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsAccessScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsAccessScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceContractDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"invoice_schedule": schema.SingleNestedAttribute{
											Description: "The schedule that the customer will be invoiced for this commit.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCommitsLedgerDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCommitsRolledOverFromDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverrideSpecifiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverrideTiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesOverwriteRateDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesOverwriteRateCreditTypeDataSourceModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesOverwriteRateTiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentOverridesProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentOverridesTiersDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentScheduledChargesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentScheduledChargesScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentScheduledChargesScheduleScheduleItemsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentTransitionsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageStatementScheduleDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentCreditsAccessScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsAccessScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentCreditsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentCreditsLedgerDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentDiscountsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentDiscountsScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentDiscountsScheduleScheduleItemsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentProfessionalServicesDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCommitsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsAccessAmountDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsCommitDurationDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCommitsInvoiceAmountDataSourceModel](ctx),
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentRecurringCreditsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsAccessAmountDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsCommitDurationDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentRecurringCreditsContractDataSourceModel](ctx),
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataCurrentResellerRoyaltiesDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationCommitDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"current": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterCurrentDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataCurrentUsageFilterInitialDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectListType[V1ContractDataCurrentUsageFilterUpdatesDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commits": schema.ListNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsAccessScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsAccessScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceContractDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"invoice_schedule": schema.SingleNestedAttribute{
											Description: "The schedule that the customer will be invoiced for this commit.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsInvoiceScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCommitsLedgerDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCommitsRolledOverFromDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverrideSpecifiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverrideTiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesOverwriteRateDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesOverwriteRateCreditTypeDataSourceModel](ctx),
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
													CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialOverridesOverwriteRateTiersDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialOverridesProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialOverridesTiersDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialScheduledChargesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialScheduledChargesScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialScheduledChargesScheduleScheduleItemsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialTransitionsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageStatementScheduleDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCreditsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsProductDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialCreditsAccessScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"schedule_items": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialCreditsAccessScheduleScheduleItemsDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsAccessScheduleCreditTypeDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialCreditsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialCreditsLedgerDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialDiscountsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"product": schema.SingleNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsScheduleDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"credit_type": schema.SingleNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectType[V1ContractDataInitialDiscountsScheduleCreditTypeDataSourceModel](ctx),
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
													CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialDiscountsScheduleScheduleItemsDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialProfessionalServicesDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCommitsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsAccessAmountDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsCommitDurationDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsContractDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCommitsInvoiceAmountDataSourceModel](ctx),
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
								CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialRecurringCreditsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"access_amount": schema.SingleNestedAttribute{
											Description: "The amount of commit to grant.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsAccessAmountDataSourceModel](ctx),
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
											CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsCommitDurationDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsProductDataSourceModel](ctx),
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
											CustomType: customfield.NewNestedObjectType[V1ContractDataInitialRecurringCreditsContractDataSourceModel](ctx),
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
											Description: "The frequency at which the recurring commits will be created.  If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's start_date rather than the usage invoice dates.\nAvailable values: \"MONTHLY\", \"QUARTERLY\", \"ANNUAL\", \"WEEKLY\".",
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
								CustomType:  customfield.NewNestedObjectListType[V1ContractDataInitialResellerRoyaltiesDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"commit": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationCommitDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
												CustomType:  customfield.NewNestedObjectType[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"current": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterCurrentDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectType[V1ContractDataInitialUsageFilterInitialDataSourceModel](ctx),
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
										CustomType: customfield.NewNestedObjectListType[V1ContractDataInitialUsageFilterUpdatesDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[V1ContractDataCustomerBillingProviderConfigurationDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationCommitDataSourceModel](ctx),
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
								CustomType: customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel](ctx),
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
					"uniqueness_key": schema.StringAttribute{
						Description: "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *V1ContractDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V1ContractDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
