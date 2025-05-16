// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v2_contract

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V2ContractDataSourceModel struct {
	Data customfield.NestedObject[V2ContractDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V2ContractDataSourceModel) toReadParams(_ context.Context) (params metronome.V2ContractGetParams, diags diag.Diagnostics) {
	params = metronome.V2ContractGetParams{}

	return
}

type V2ContractDataDataSourceModel struct {
	ID                                   types.String                                                                                `tfsdk:"id" json:"id,computed"`
	Commits                              customfield.NestedObjectList[V2ContractDataCommitsDataSourceModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                                `tfsdk:"created_by" json:"created_by,computed"`
	CustomerID                           types.String                                                                                `tfsdk:"customer_id" json:"customer_id,computed"`
	Overrides                            customfield.NestedObjectList[V2ContractDataOverridesDataSourceModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V2ContractDataScheduledChargesDataSourceModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                           `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V2ContractDataTransitionsDataSourceModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageFilter                          customfield.NestedObjectList[V2ContractDataUsageFilterDataSourceModel]                      `tfsdk:"usage_filter" json:"usage_filter,computed"`
	UsageStatementSchedule               customfield.NestedObject[V2ContractDataUsageStatementScheduleDataSourceModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	ArchivedAt                           timetypes.RFC3339                                                                           `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Credits                              customfield.NestedObjectList[V2ContractDataCreditsDataSourceModel]                          `tfsdk:"credits" json:"credits,computed"`
	CustomFields                         customfield.Map[types.String]                                                               `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerBillingProviderConfiguration customfield.NestedObject[V2ContractDataCustomerBillingProviderConfigurationDataSourceModel] `tfsdk:"customer_billing_provider_configuration" json:"customer_billing_provider_configuration,computed"`
	Discounts                            customfield.NestedObjectList[V2ContractDataDiscountsDataSourceModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                           `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	MultiplierOverridePrioritization     types.String                                                                                `tfsdk:"multiplier_override_prioritization" json:"multiplier_override_prioritization,computed"`
	Name                                 types.String                                                                                `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                               `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V2ContractDataProfessionalServicesDataSourceModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                                `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V2ContractDataRecurringCommitsDataSourceModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V2ContractDataRecurringCreditsDataSourceModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V2ContractDataResellerRoyaltiesDataSourceModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                                `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                                `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V2ContractDataSpendThresholdConfigurationDataSourceModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                               `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UniquenessKey                        types.String                                                                                `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V2ContractDataCommitsDataSourceModel struct {
	ID                      types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V2ContractDataCommitsProductDataSourceModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                  `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V2ContractDataCommitsAccessScheduleDataSourceModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                             `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                                 `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V2ContractDataCommitsContractDataSourceModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                 `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                  `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V2ContractDataCommitsInvoiceContractDataSourceModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V2ContractDataCommitsInvoiceScheduleDataSourceModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V2ContractDataCommitsLedgerDataSourceModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                  `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                 `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                  `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V2ContractDataCommitsRolledOverFromDataSourceModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                                 `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                                  `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V2ContractDataCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V2ContractDataCommitsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V2ContractDataCommitsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V2ContractDataCommitsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V2ContractDataCommitsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCommitsInvoiceContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCommitsInvoiceScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V2ContractDataCommitsInvoiceScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataCommitsInvoiceScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataCommitsInvoiceScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsInvoiceScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataCommitsLedgerDataSourceModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V2ContractDataCommitsRolledOverFromDataSourceModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V2ContractDataOverridesDataSourceModel struct {
	ID                    types.String                                                                           `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                      `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                         `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	EndingBefore          timetypes.RFC3339                                                                      `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                             `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                             `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	Multiplier            types.Float64                                                                          `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V2ContractDataOverridesOverrideSpecifiersDataSourceModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V2ContractDataOverridesOverrideTiersDataSourceModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V2ContractDataOverridesOverwriteRateDataSourceModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Priority              types.Float64                                                                          `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataOverridesProductDataSourceModel]                `tfsdk:"product" json:"product,computed"`
	Target                types.String                                                                           `tfsdk:"target" json:"target,computed"`
	Type                  types.String                                                                           `tfsdk:"type" json:"type,computed"`
}

type V2ContractDataOverridesOverrideSpecifiersDataSourceModel struct {
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V2ContractDataOverridesOverrideTiersDataSourceModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V2ContractDataOverridesOverwriteRateDataSourceModel struct {
	RateType   types.String                                                                            `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V2ContractDataOverridesOverwriteRateCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                   `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                              `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                           `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                           `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V2ContractDataOverridesOverwriteRateTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V2ContractDataOverridesOverwriteRateCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataOverridesOverwriteRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V2ContractDataOverridesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesDataSourceModel struct {
	ID                   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V2ContractDataScheduledChargesProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V2ContractDataScheduledChargesScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                               `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                   `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                    `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                    `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataScheduledChargesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V2ContractDataScheduledChargesScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataScheduledChargesScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataScheduledChargesScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataTransitionsDataSourceModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V2ContractDataUsageFilterDataSourceModel struct {
	GroupKey     types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues  customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt   timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	EndingBefore timetypes.RFC3339              `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}

type V2ContractDataUsageStatementScheduleDataSourceModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V2ContractDataCreditsDataSourceModel struct {
	ID                      types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V2ContractDataCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                 `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V2ContractDataCreditsAccessScheduleDataSourceModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                               `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                                `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V2ContractDataCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V2ContractDataCreditsLedgerDataSourceModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                `tfsdk:"priority" json:"priority,computed"`
	SalesforceOpportunityID types.String                                                                 `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V2ContractDataCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCreditsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V2ContractDataCreditsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V2ContractDataCreditsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V2ContractDataCreditsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V2ContractDataCreditsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCreditsLedgerDataSourceModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V2ContractDataCustomerBillingProviderConfigurationDataSourceModel struct {
	BillingProvider types.String `tfsdk:"billing_provider" json:"billing_provider,computed"`
	DeliveryMethod  types.String `tfsdk:"delivery_method" json:"delivery_method,computed"`
}

type V2ContractDataDiscountsDataSourceModel struct {
	ID                   types.String                                                             `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V2ContractDataDiscountsProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V2ContractDataDiscountsScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                            `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                             `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataDiscountsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataDiscountsScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V2ContractDataDiscountsScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataDiscountsScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataDiscountsScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataDiscountsScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                   `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                                `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                                `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel struct {
	ProductID             types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                   `tfsdk:"description" json:"description,computed"`
	Name                  types.String                   `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                             `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                             `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V2ContractDataProfessionalServicesDataSourceModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataRecurringCommitsDataSourceModel struct {
	ID                    types.String                                                                          `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V2ContractDataRecurringCommitsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V2ContractDataRecurringCommitsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                         `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataRecurringCommitsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                          `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                     `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                        `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                        `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V2ContractDataRecurringCommitsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                          `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                     `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V2ContractDataRecurringCommitsInvoiceAmountDataSourceModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                          `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                          `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                          `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                         `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V2ContractDataRecurringCommitsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCommitsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V2ContractDataRecurringCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataRecurringCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataRecurringCommitsInvoiceAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCreditsDataSourceModel struct {
	ID                    types.String                                                                          `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V2ContractDataRecurringCreditsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V2ContractDataRecurringCreditsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                         `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataRecurringCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                          `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                     `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                        `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                        `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V2ContractDataRecurringCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                          `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                     `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                          `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                          `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                          `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                         `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V2ContractDataRecurringCreditsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCreditsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V2ContractDataRecurringCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataRecurringCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataResellerRoyaltiesDataSourceModel struct {
	ResellerType types.String                                                                         `tfsdk:"reseller_type" json:"reseller_type,computed"`
	Segments     customfield.NestedObjectList[V2ContractDataResellerRoyaltiesSegmentsDataSourceModel] `tfsdk:"segments" json:"segments,computed"`
}

type V2ContractDataResellerRoyaltiesSegmentsDataSourceModel struct {
	Fraction              types.Float64                  `tfsdk:"fraction" json:"fraction,computed"`
	NetsuiteResellerID    types.String                   `tfsdk:"netsuite_reseller_id" json:"netsuite_reseller_id,computed"`
	ResellerType          types.String                   `tfsdk:"reseller_type" json:"reseller_type,computed"`
	StartingAt            timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	AwsAccountNumber      types.String                   `tfsdk:"aws_account_number" json:"aws_account_number,computed"`
	AwsOfferID            types.String                   `tfsdk:"aws_offer_id" json:"aws_offer_id,computed"`
	AwsPayerReferenceID   types.String                   `tfsdk:"aws_payer_reference_id" json:"aws_payer_reference_id,computed"`
	EndingBefore          timetypes.RFC3339              `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	GcpAccountID          types.String                   `tfsdk:"gcp_account_id" json:"gcp_account_id,computed"`
	GcpOfferID            types.String                   `tfsdk:"gcp_offer_id" json:"gcp_offer_id,computed"`
	ResellerContractValue types.Float64                  `tfsdk:"reseller_contract_value" json:"reseller_contract_value,computed"`
}

type V2ContractDataSpendThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V2ContractDataSpendThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                          `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V2ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                       `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V2ContractDataSpendThresholdConfigurationCommitDataSourceModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                    `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                    `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}
