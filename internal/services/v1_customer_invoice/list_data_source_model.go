// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_invoice

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomerInvoicesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1CustomerInvoicesItemsDataSourceModel] `json:"data,computed"`
}

type V1CustomerInvoicesDataSourceModel struct {
	CustomerID           types.String                                                         `tfsdk:"customer_id" path:"customer_id,required"`
	CreditTypeID         types.String                                                         `tfsdk:"credit_type_id" query:"credit_type_id,optional"`
	EndingBefore         timetypes.RFC3339                                                    `tfsdk:"ending_before" query:"ending_before,optional" format:"date-time"`
	SkipZeroQtyLineItems types.Bool                                                           `tfsdk:"skip_zero_qty_line_items" query:"skip_zero_qty_line_items,optional"`
	Sort                 types.String                                                         `tfsdk:"sort" query:"sort,optional"`
	StartingOn           timetypes.RFC3339                                                    `tfsdk:"starting_on" query:"starting_on,optional" format:"date-time"`
	Status               types.String                                                         `tfsdk:"status" query:"status,optional"`
	MaxItems             types.Int64                                                          `tfsdk:"max_items"`
	Items                customfield.NestedObjectList[V1CustomerInvoicesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1CustomerInvoicesDataSourceModel) toListParams(_ context.Context) (params metronome.V1CustomerInvoiceListParams, diags diag.Diagnostics) {
	mEndingBefore, errs := m.EndingBefore.ValueRFC3339Time()
	diags.Append(errs...)
	mStartingOn, errs := m.StartingOn.ValueRFC3339Time()
	diags.Append(errs...)

	params = metronome.V1CustomerInvoiceListParams{
		CustomerID: metronome.F(m.CustomerID.ValueString()),
	}

	if !m.CreditTypeID.IsNull() {
		params.CreditTypeID = metronome.F(m.CreditTypeID.ValueString())
	}
	if !m.EndingBefore.IsNull() {
		params.EndingBefore = metronome.F(mEndingBefore)
	}
	if !m.SkipZeroQtyLineItems.IsNull() {
		params.SkipZeroQtyLineItems = metronome.F(m.SkipZeroQtyLineItems.ValueBool())
	}
	if !m.Sort.IsNull() {
		params.Sort = metronome.F(metronome.V1CustomerInvoiceListParamsSort(m.Sort.ValueString()))
	}
	if !m.StartingOn.IsNull() {
		params.StartingOn = metronome.F(mStartingOn)
	}
	if !m.Status.IsNull() {
		params.Status = metronome.F(m.Status.ValueString())
	}

	return
}

type V1CustomerInvoicesItemsDataSourceModel struct {
	ID                      types.String                                                                      `tfsdk:"id" json:"id,computed"`
	CreditType              customfield.NestedObject[V1CustomerInvoicesCreditTypeDataSourceModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomerID              types.String                                                                      `tfsdk:"customer_id" json:"customer_id,computed"`
	LineItems               customfield.NestedObjectList[V1CustomerInvoicesLineItemsDataSourceModel]          `tfsdk:"line_items" json:"line_items,computed"`
	Status                  types.String                                                                      `tfsdk:"status" json:"status,computed"`
	Total                   types.Float64                                                                     `tfsdk:"total" json:"total,computed"`
	Type                    types.String                                                                      `tfsdk:"type" json:"type,computed"`
	AmendmentID             types.String                                                                      `tfsdk:"amendment_id" json:"amendment_id,computed"`
	BillableStatus          types.String                                                                      `tfsdk:"billable_status" json:"billable_status,computed"`
	ContractCustomFields    customfield.Map[types.String]                                                     `tfsdk:"contract_custom_fields" json:"contract_custom_fields,computed"`
	ContractID              types.String                                                                      `tfsdk:"contract_id" json:"contract_id,computed"`
	CorrectionRecord        customfield.NestedObject[V1CustomerInvoicesCorrectionRecordDataSourceModel]       `tfsdk:"correction_record" json:"correction_record,computed"`
	CreatedAt               timetypes.RFC3339                                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomFields            customfield.Map[jsontypes.Normalized]                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerCustomFields    customfield.Map[types.String]                                                     `tfsdk:"customer_custom_fields" json:"customer_custom_fields,computed"`
	EndTimestamp            timetypes.RFC3339                                                                 `tfsdk:"end_timestamp" json:"end_timestamp,computed" format:"date-time"`
	ExternalInvoice         customfield.NestedObject[V1CustomerInvoicesExternalInvoiceDataSourceModel]        `tfsdk:"external_invoice" json:"external_invoice,computed"`
	InvoiceAdjustments      customfield.NestedObjectList[V1CustomerInvoicesInvoiceAdjustmentsDataSourceModel] `tfsdk:"invoice_adjustments" json:"invoice_adjustments,computed"`
	IssuedAt                timetypes.RFC3339                                                                 `tfsdk:"issued_at" json:"issued_at,computed" format:"date-time"`
	NetPaymentTermsDays     types.Float64                                                                     `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID    types.String                                                                      `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PlanCustomFields        customfield.Map[types.String]                                                     `tfsdk:"plan_custom_fields" json:"plan_custom_fields,computed"`
	PlanID                  types.String                                                                      `tfsdk:"plan_id" json:"plan_id,computed"`
	PlanName                types.String                                                                      `tfsdk:"plan_name" json:"plan_name,computed"`
	ResellerRoyalty         customfield.NestedObject[V1CustomerInvoicesResellerRoyaltyDataSourceModel]        `tfsdk:"reseller_royalty" json:"reseller_royalty,computed"`
	SalesforceOpportunityID types.String                                                                      `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	StartTimestamp          timetypes.RFC3339                                                                 `tfsdk:"start_timestamp" json:"start_timestamp,computed" format:"date-time"`
	Subtotal                types.Float64                                                                     `tfsdk:"subtotal" json:"subtotal,computed"`
}

type V1CustomerInvoicesCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CustomerInvoicesLineItemsDataSourceModel struct {
	CreditType                      customfield.NestedObject[V1CustomerInvoicesLineItemsCreditTypeDataSourceModel]            `tfsdk:"credit_type" json:"credit_type,computed"`
	Name                            types.String                                                                              `tfsdk:"name" json:"name,computed"`
	Total                           types.Float64                                                                             `tfsdk:"total" json:"total,computed"`
	AppliedCommitOrCredit           customfield.NestedObject[V1CustomerInvoicesLineItemsAppliedCommitOrCreditDataSourceModel] `tfsdk:"applied_commit_or_credit" json:"applied_commit_or_credit,computed"`
	CommitCustomFields              customfield.Map[types.String]                                                             `tfsdk:"commit_custom_fields" json:"commit_custom_fields,computed"`
	CommitID                        types.String                                                                              `tfsdk:"commit_id" json:"commit_id,computed"`
	CommitNetsuiteItemID            types.String                                                                              `tfsdk:"commit_netsuite_item_id" json:"commit_netsuite_item_id,computed"`
	CommitNetsuiteSalesOrderID      types.String                                                                              `tfsdk:"commit_netsuite_sales_order_id" json:"commit_netsuite_sales_order_id,computed"`
	CommitSegmentID                 types.String                                                                              `tfsdk:"commit_segment_id" json:"commit_segment_id,computed"`
	CommitType                      types.String                                                                              `tfsdk:"commit_type" json:"commit_type,computed"`
	CustomFields                    customfield.Map[types.String]                                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	DiscountCustomFields            customfield.Map[types.String]                                                             `tfsdk:"discount_custom_fields" json:"discount_custom_fields,computed"`
	DiscountID                      types.String                                                                              `tfsdk:"discount_id" json:"discount_id,computed"`
	EndingBefore                    timetypes.RFC3339                                                                         `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	GroupKey                        types.String                                                                              `tfsdk:"group_key" json:"group_key,computed"`
	GroupValue                      types.String                                                                              `tfsdk:"group_value" json:"group_value,computed"`
	IsProrated                      types.Bool                                                                                `tfsdk:"is_prorated" json:"is_prorated,computed"`
	ListPrice                       customfield.NestedObject[V1CustomerInvoicesLineItemsListPriceDataSourceModel]             `tfsdk:"list_price" json:"list_price,computed"`
	Metadata                        types.String                                                                              `tfsdk:"metadata" json:"metadata,computed"`
	NetsuiteInvoiceBillingEnd       timetypes.RFC3339                                                                         `tfsdk:"netsuite_invoice_billing_end" json:"netsuite_invoice_billing_end,computed" format:"date-time"`
	NetsuiteInvoiceBillingStart     timetypes.RFC3339                                                                         `tfsdk:"netsuite_invoice_billing_start" json:"netsuite_invoice_billing_start,computed" format:"date-time"`
	NetsuiteItemID                  types.String                                                                              `tfsdk:"netsuite_item_id" json:"netsuite_item_id,computed"`
	PostpaidCommit                  customfield.NestedObject[V1CustomerInvoicesLineItemsPostpaidCommitDataSourceModel]        `tfsdk:"postpaid_commit" json:"postpaid_commit,computed"`
	PresentationGroupValues         customfield.Map[types.String]                                                             `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues              customfield.Map[types.String]                                                             `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductCustomFields             customfield.Map[types.String]                                                             `tfsdk:"product_custom_fields" json:"product_custom_fields,computed"`
	ProductID                       types.String                                                                              `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags                     customfield.List[types.String]                                                            `tfsdk:"product_tags" json:"product_tags,computed"`
	ProductType                     types.String                                                                              `tfsdk:"product_type" json:"product_type,computed"`
	ProfessionalServiceCustomFields customfield.Map[types.String]                                                             `tfsdk:"professional_service_custom_fields" json:"professional_service_custom_fields,computed"`
	ProfessionalServiceID           types.String                                                                              `tfsdk:"professional_service_id" json:"professional_service_id,computed"`
	Quantity                        types.Float64                                                                             `tfsdk:"quantity" json:"quantity,computed"`
	ResellerType                    types.String                                                                              `tfsdk:"reseller_type" json:"reseller_type,computed"`
	ScheduledChargeCustomFields     customfield.Map[types.String]                                                             `tfsdk:"scheduled_charge_custom_fields" json:"scheduled_charge_custom_fields,computed"`
	ScheduledChargeID               types.String                                                                              `tfsdk:"scheduled_charge_id" json:"scheduled_charge_id,computed"`
	StartingAt                      timetypes.RFC3339                                                                         `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	SubLineItems                    customfield.NestedObjectList[V1CustomerInvoicesLineItemsSubLineItemsDataSourceModel]      `tfsdk:"sub_line_items" json:"sub_line_items,computed"`
	SubscriptionCustomFields        customfield.Map[types.String]                                                             `tfsdk:"subscription_custom_fields" json:"subscription_custom_fields,computed"`
	Tier                            customfield.NestedObject[V1CustomerInvoicesLineItemsTierDataSourceModel]                  `tfsdk:"tier" json:"tier,computed"`
	UnitPrice                       types.Float64                                                                             `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1CustomerInvoicesLineItemsCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CustomerInvoicesLineItemsAppliedCommitOrCreditDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type V1CustomerInvoicesLineItemsListPriceDataSourceModel struct {
	RateType           types.String                                                                            `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType         customfield.NestedObject[V1CustomerInvoicesLineItemsListPriceCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate         customfield.Map[jsontypes.Normalized]                                                   `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated         types.Bool                                                                              `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price              types.Float64                                                                           `tfsdk:"price" json:"price,computed"`
	PricingGroupValues customfield.Map[types.String]                                                           `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	Quantity           types.Float64                                                                           `tfsdk:"quantity" json:"quantity,computed"`
	Tiers              customfield.NestedObjectList[V1CustomerInvoicesLineItemsListPriceTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
	UseListPrices      types.Bool                                                                              `tfsdk:"use_list_prices" json:"use_list_prices,computed"`
}

type V1CustomerInvoicesLineItemsListPriceCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CustomerInvoicesLineItemsListPriceTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1CustomerInvoicesLineItemsPostpaidCommitDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1CustomerInvoicesLineItemsSubLineItemsDataSourceModel struct {
	CustomFields  customfield.Map[types.String]                                                              `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name          types.String                                                                               `tfsdk:"name" json:"name,computed"`
	Quantity      types.Float64                                                                              `tfsdk:"quantity" json:"quantity,computed"`
	Subtotal      types.Float64                                                                              `tfsdk:"subtotal" json:"subtotal,computed"`
	ChargeID      types.String                                                                               `tfsdk:"charge_id" json:"charge_id,computed"`
	CreditGrantID types.String                                                                               `tfsdk:"credit_grant_id" json:"credit_grant_id,computed"`
	EndDate       timetypes.RFC3339                                                                          `tfsdk:"end_date" json:"end_date,computed" format:"date-time"`
	Price         types.Float64                                                                              `tfsdk:"price" json:"price,computed"`
	StartDate     timetypes.RFC3339                                                                          `tfsdk:"start_date" json:"start_date,computed" format:"date-time"`
	TierPeriod    customfield.NestedObject[V1CustomerInvoicesLineItemsSubLineItemsTierPeriodDataSourceModel] `tfsdk:"tier_period" json:"tier_period,computed"`
	Tiers         customfield.NestedObjectList[V1CustomerInvoicesLineItemsSubLineItemsTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1CustomerInvoicesLineItemsSubLineItemsTierPeriodDataSourceModel struct {
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}

type V1CustomerInvoicesLineItemsSubLineItemsTiersDataSourceModel struct {
	Price      types.Float64 `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	StartingAt types.Float64 `tfsdk:"starting_at" json:"starting_at,computed"`
	Subtotal   types.Float64 `tfsdk:"subtotal" json:"subtotal,computed"`
}

type V1CustomerInvoicesLineItemsTierDataSourceModel struct {
	Level      types.Float64 `tfsdk:"level" json:"level,computed"`
	StartingAt types.String  `tfsdk:"starting_at" json:"starting_at,computed"`
	Size       types.String  `tfsdk:"size" json:"size,computed"`
}

type V1CustomerInvoicesCorrectionRecordDataSourceModel struct {
	CorrectedInvoiceID       types.String                                                                                        `tfsdk:"corrected_invoice_id" json:"corrected_invoice_id,computed"`
	Memo                     types.String                                                                                        `tfsdk:"memo" json:"memo,computed"`
	Reason                   types.String                                                                                        `tfsdk:"reason" json:"reason,computed"`
	CorrectedExternalInvoice customfield.NestedObject[V1CustomerInvoicesCorrectionRecordCorrectedExternalInvoiceDataSourceModel] `tfsdk:"corrected_external_invoice" json:"corrected_external_invoice,computed"`
}

type V1CustomerInvoicesCorrectionRecordCorrectedExternalInvoiceDataSourceModel struct {
	BillingProviderType types.String      `tfsdk:"billing_provider_type" json:"billing_provider_type,computed"`
	ExternalStatus      types.String      `tfsdk:"external_status" json:"external_status,computed"`
	InvoiceID           types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	IssuedAtTimestamp   timetypes.RFC3339 `tfsdk:"issued_at_timestamp" json:"issued_at_timestamp,computed" format:"date-time"`
}

type V1CustomerInvoicesExternalInvoiceDataSourceModel struct {
	BillingProviderType types.String      `tfsdk:"billing_provider_type" json:"billing_provider_type,computed"`
	ExternalStatus      types.String      `tfsdk:"external_status" json:"external_status,computed"`
	InvoiceID           types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	IssuedAtTimestamp   timetypes.RFC3339 `tfsdk:"issued_at_timestamp" json:"issued_at_timestamp,computed" format:"date-time"`
}

type V1CustomerInvoicesInvoiceAdjustmentsDataSourceModel struct {
	CreditType              customfield.NestedObject[V1CustomerInvoicesInvoiceAdjustmentsCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	Name                    types.String                                                                            `tfsdk:"name" json:"name,computed"`
	Total                   types.Float64                                                                           `tfsdk:"total" json:"total,computed"`
	CreditGrantCustomFields customfield.Map[types.String]                                                           `tfsdk:"credit_grant_custom_fields" json:"credit_grant_custom_fields,computed"`
	CreditGrantID           types.String                                                                            `tfsdk:"credit_grant_id" json:"credit_grant_id,computed"`
}

type V1CustomerInvoicesInvoiceAdjustmentsCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1CustomerInvoicesResellerRoyaltyDataSourceModel struct {
	Fraction           types.String                                                                         `tfsdk:"fraction" json:"fraction,computed"`
	NetsuiteResellerID types.String                                                                         `tfsdk:"netsuite_reseller_id" json:"netsuite_reseller_id,computed"`
	ResellerType       types.String                                                                         `tfsdk:"reseller_type" json:"reseller_type,computed"`
	AwsOptions         customfield.NestedObject[V1CustomerInvoicesResellerRoyaltyAwsOptionsDataSourceModel] `tfsdk:"aws_options" json:"aws_options,computed"`
	GcpOptions         customfield.NestedObject[V1CustomerInvoicesResellerRoyaltyGcpOptionsDataSourceModel] `tfsdk:"gcp_options" json:"gcp_options,computed"`
}

type V1CustomerInvoicesResellerRoyaltyAwsOptionsDataSourceModel struct {
	AwsAccountNumber    types.String `tfsdk:"aws_account_number" json:"aws_account_number,computed"`
	AwsOfferID          types.String `tfsdk:"aws_offer_id" json:"aws_offer_id,computed"`
	AwsPayerReferenceID types.String `tfsdk:"aws_payer_reference_id" json:"aws_payer_reference_id,computed"`
}

type V1CustomerInvoicesResellerRoyaltyGcpOptionsDataSourceModel struct {
	GcpAccountID types.String `tfsdk:"gcp_account_id" json:"gcp_account_id,computed"`
	GcpOfferID   types.String `tfsdk:"gcp_offer_id" json:"gcp_offer_id,computed"`
}
