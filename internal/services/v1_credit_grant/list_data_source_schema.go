// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant

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

var _ datasource.DataSourceWithConfigValidators = (*V1CreditGrantsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"effective_before": schema.StringAttribute{
				Description: "Only return credit grants that are effective before this timestamp (exclusive).",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"not_expiring_before": schema.StringAttribute{
				Description: "Only return credit grants that expire at or after this timestamp.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"credit_grant_ids": schema.ListAttribute{
				Description: "An array of credit grant IDs. If this is specified, neither credit_type_ids nor customer_ids may be specified.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"credit_type_ids": schema.ListAttribute{
				Description: "An array of credit type IDs. This must not be specified if credit_grant_ids is specified.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"customer_ids": schema.ListAttribute{
				Description: "An array of Metronome customer IDs. This must not be specified if credit_grant_ids is specified.",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[V1CreditGrantsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "the Metronome ID of the credit grant",
							Computed:    true,
						},
						"balance": schema.SingleNestedAttribute{
							Description: "The effective balance of the grant as of the end of the customer's current billing period. Expiration deductions will be included only if the grant expires before the end of the current billing period.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[V1CreditGrantsBalanceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"effective_at": schema.StringAttribute{
									Description: "The end_date of the customer's current billing period.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"excluding_pending": schema.Float64Attribute{
									Description: "The grant's current balance including all posted deductions. If the grant has expired, this amount will be 0.",
									Computed:    true,
								},
								"including_pending": schema.Float64Attribute{
									Description: "The grant's current balance including all posted and pending deductions. If the grant expires before the end of the customer's current billing period, this amount will be 0.",
									Computed:    true,
								},
							},
						},
						"custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"customer_id": schema.StringAttribute{
							Description: "the Metronome ID of the customer",
							Computed:    true,
						},
						"deductions": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[V1CreditGrantsDeductionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Float64Attribute{
										Description: "an amount representing the change to the customer's credit balance",
										Computed:    true,
									},
									"created_by": schema.StringAttribute{
										Computed: true,
									},
									"credit_grant_id": schema.StringAttribute{
										Description: "the credit grant this entry is related to",
										Computed:    true,
									},
									"effective_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"reason": schema.StringAttribute{
										Computed: true,
									},
									"running_balance": schema.Float64Attribute{
										Description: "the running balance for this credit type at the time of the ledger entry, including all preceding charges",
										Computed:    true,
									},
									"invoice_id": schema.StringAttribute{
										Description: "if this entry is a deduction, the Metronome ID of the invoice where the credit deduction was consumed; if this entry is a grant, the Metronome ID of the invoice where the grant's paid_amount was charged",
										Computed:    true,
									},
								},
							},
						},
						"effective_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"expires_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"grant_amount": schema.SingleNestedAttribute{
							Description: "the amount of credits initially granted",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[V1CreditGrantsGrantAmountDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"amount": schema.Float64Attribute{
									Computed: true,
								},
								"credit_type": schema.SingleNestedAttribute{
									Description: "the credit type for the amount granted",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1CreditGrantsGrantAmountCreditTypeDataSourceModel](ctx),
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
						"name": schema.StringAttribute{
							Computed: true,
						},
						"paid_amount": schema.SingleNestedAttribute{
							Description: "the amount paid for this credit grant",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[V1CreditGrantsPaidAmountDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"amount": schema.Float64Attribute{
									Computed: true,
								},
								"credit_type": schema.SingleNestedAttribute{
									Description: "the credit type for the amount paid",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[V1CreditGrantsPaidAmountCreditTypeDataSourceModel](ctx),
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
						"pending_deductions": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[V1CreditGrantsPendingDeductionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Float64Attribute{
										Description: "an amount representing the change to the customer's credit balance",
										Computed:    true,
									},
									"created_by": schema.StringAttribute{
										Computed: true,
									},
									"credit_grant_id": schema.StringAttribute{
										Description: "the credit grant this entry is related to",
										Computed:    true,
									},
									"effective_at": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"reason": schema.StringAttribute{
										Computed: true,
									},
									"running_balance": schema.Float64Attribute{
										Description: "the running balance for this credit type at the time of the ledger entry, including all preceding charges",
										Computed:    true,
									},
									"invoice_id": schema.StringAttribute{
										Description: "if this entry is a deduction, the Metronome ID of the invoice where the credit deduction was consumed; if this entry is a grant, the Metronome ID of the invoice where the grant's paid_amount was charged",
										Computed:    true,
									},
								},
							},
						},
						"priority": schema.Float64Attribute{
							Computed: true,
						},
						"credit_grant_type": schema.StringAttribute{
							Computed: true,
						},
						"invoice_id": schema.StringAttribute{
							Description: "the Metronome ID of the invoice with the purchase charge for this credit grant, if applicable",
							Computed:    true,
						},
						"products": schema.ListNestedAttribute{
							Description: "The products which these credits will be applied to. (If unspecified, the credits will be applied to charges for all products.)",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[V1CreditGrantsProductsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
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
						"reason": schema.StringAttribute{
							Computed: true,
						},
						"uniqueness_key": schema.StringAttribute{
							Description: "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *V1CreditGrantsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1CreditGrantsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
