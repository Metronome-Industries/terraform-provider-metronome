// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractDataSourceModel struct {
	ContractID     types.String                                            `tfsdk:"contract_id" json:"contract_id,required"`
	CustomerID     types.String                                            `tfsdk:"customer_id" json:"customer_id,required"`
	IncludeBalance types.Bool                                              `tfsdk:"include_balance" json:"include_balance,optional"`
	IncludeLedgers types.Bool                                              `tfsdk:"include_ledgers" json:"include_ledgers,optional"`
	Data           customfield.NestedObject[V1ContractDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1ContractDataSourceModel) toReadParams(_ context.Context) (params metronome.V1ContractGetParams, diags diag.Diagnostics) {
	params = metronome.V1ContractGetParams{}

	return
}

type V1ContractDataDataSourceModel struct {
	ID                                   types.String                                                                                `tfsdk:"id" json:"id,computed"`
	Amendments                           customfield.NestedObjectList[V1ContractDataAmendmentsDataSourceModel]                       `tfsdk:"amendments" json:"amendments,computed"`
	Current                              customfield.NestedObject[V1ContractDataCurrentDataSourceModel]                              `tfsdk:"current" json:"current,computed"`
	CustomerID                           types.String                                                                                `tfsdk:"customer_id" json:"customer_id,computed"`
	Initial                              customfield.NestedObject[V1ContractDataInitialDataSourceModel]                              `tfsdk:"initial" json:"initial,computed"`
	ArchivedAt                           timetypes.RFC3339                                                                           `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields                         customfield.Map[types.String]                                                               `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerBillingProviderConfiguration customfield.NestedObject[V1ContractDataCustomerBillingProviderConfigurationDataSourceModel] `tfsdk:"customer_billing_provider_configuration" json:"customer_billing_provider_configuration,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                                `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataSpendThresholdConfigurationDataSourceModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	Subscriptions                        customfield.NestedObjectList[V1ContractDataSubscriptionsDataSourceModel]                    `tfsdk:"subscriptions" json:"subscriptions,computed"`
	UniquenessKey                        types.String                                                                                `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsDataSourceModel struct {
	ID                      types.String                                                                              `tfsdk:"id" json:"id,computed"`
	Commits                 customfield.NestedObjectList[V1ContractDataAmendmentsCommitsDataSourceModel]              `tfsdk:"commits" json:"commits,computed"`
	CreatedAt               timetypes.RFC3339                                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy               types.String                                                                              `tfsdk:"created_by" json:"created_by,computed"`
	Overrides               customfield.NestedObjectList[V1ContractDataAmendmentsOverridesDataSourceModel]            `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges        customfield.NestedObjectList[V1ContractDataAmendmentsScheduledChargesDataSourceModel]     `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt              timetypes.RFC3339                                                                         `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Credits                 customfield.NestedObjectList[V1ContractDataAmendmentsCreditsDataSourceModel]              `tfsdk:"credits" json:"credits,computed"`
	Discounts               customfield.NestedObjectList[V1ContractDataAmendmentsDiscountsDataSourceModel]            `tfsdk:"discounts" json:"discounts,computed"`
	NetsuiteSalesOrderID    types.String                                                                              `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	ProfessionalServices    customfield.NestedObjectList[V1ContractDataAmendmentsProfessionalServicesDataSourceModel] `tfsdk:"professional_services" json:"professional_services,computed"`
	ResellerRoyalties       customfield.NestedObjectList[V1ContractDataAmendmentsResellerRoyaltiesDataSourceModel]    `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID types.String                                                                              `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V1ContractDataAmendmentsCommitsDataSourceModel struct {
	ID                      types.String                                                                            `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataAmendmentsCommitsProductDataSourceModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                            `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataAmendmentsCommitsAccessScheduleDataSourceModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                                           `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                          `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                          `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                          `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                                       `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                                           `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataAmendmentsCommitsContractDataSourceModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                           `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                            `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceContractDataSourceModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceScheduleDataSourceModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataAmendmentsCommitsLedgerDataSourceModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                            `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                            `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                           `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                            `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataAmendmentsCommitsRolledOverFromDataSourceModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                                           `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                                            `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataAmendmentsCommitsSpecifiersDataSourceModel]  `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                            `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsCommitsLedgerDataSourceModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataAmendmentsCommitsRolledOverFromDataSourceModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataAmendmentsCommitsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataAmendmentsOverridesDataSourceModel struct {
	ID                    types.String                                                                                     `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                                `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                                   `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataAmendmentsOverridesCreditTypeDataSourceModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                                `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                                       `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                                       `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                                       `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                                    `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverrideSpecifiersDataSourceModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverrideTiersDataSourceModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataAmendmentsOverridesOverwriteRateDataSourceModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                                    `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                                    `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataAmendmentsOverridesProductDataSourceModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                                    `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                                     `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                                     `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataAmendmentsOverridesTiersDataSourceModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                                     `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                                            `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataAmendmentsOverridesCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesOverrideSpecifiersDataSourceModel struct {
	BillingFrequency        types.String                   `tfsdk:"billing_frequency" json:"billing_frequency,computed"`
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataAmendmentsOverridesOverrideTiersDataSourceModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateDataSourceModel struct {
	RateType   types.String                                                                                      `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                             `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                                        `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                                     `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                                     `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverwriteRateTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsOverridesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsScheduledChargesDataSourceModel struct {
	ID                   types.String                                                                              `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                                         `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                              `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                              `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsScheduledChargesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsCreditsDataSourceModel struct {
	ID                      types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataAmendmentsCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                           `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataAmendmentsCreditsAccessScheduleDataSourceModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                         `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                         `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                         `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                                          `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataAmendmentsCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                          `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                           `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataAmendmentsCreditsLedgerDataSourceModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                           `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                           `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                          `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                           `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                                           `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataAmendmentsCreditsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                           `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCreditsLedgerDataSourceModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataAmendmentsCreditsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataAmendmentsDiscountsDataSourceModel struct {
	ID                   types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataAmendmentsDiscountsProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataAmendmentsDiscountsScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                                      `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsDiscountsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsDiscountsScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsDiscountsScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsProfessionalServicesDataSourceModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsResellerRoyaltiesDataSourceModel struct {
	ResellerType          types.String      `tfsdk:"reseller_type" json:"reseller_type,computed"`
	AwsAccountNumber      types.String      `tfsdk:"aws_account_number" json:"aws_account_number,computed"`
	AwsOfferID            types.String      `tfsdk:"aws_offer_id" json:"aws_offer_id,computed"`
	AwsPayerReferenceID   types.String      `tfsdk:"aws_payer_reference_id" json:"aws_payer_reference_id,computed"`
	EndingBefore          timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Fraction              types.Float64     `tfsdk:"fraction" json:"fraction,computed"`
	GcpAccountID          types.String      `tfsdk:"gcp_account_id" json:"gcp_account_id,computed"`
	GcpOfferID            types.String      `tfsdk:"gcp_offer_id" json:"gcp_offer_id,computed"`
	NetsuiteResellerID    types.String      `tfsdk:"netsuite_reseller_id" json:"netsuite_reseller_id,computed"`
	ResellerContractValue types.Float64     `tfsdk:"reseller_contract_value" json:"reseller_contract_value,computed"`
	StartingAt            timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentDataSourceModel struct {
	Commits                              customfield.NestedObjectList[V1ContractDataCurrentCommitsDataSourceModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                                       `tfsdk:"created_by" json:"created_by,computed"`
	Overrides                            customfield.NestedObjectList[V1ContractDataCurrentOverridesDataSourceModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V1ContractDataCurrentScheduledChargesDataSourceModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V1ContractDataCurrentTransitionsDataSourceModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageStatementSchedule               customfield.NestedObject[V1ContractDataCurrentUsageStatementScheduleDataSourceModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	Credits                              customfield.NestedObjectList[V1ContractDataCurrentCreditsDataSourceModel]                          `tfsdk:"credits" json:"credits,computed"`
	Discounts                            customfield.NestedObjectList[V1ContractDataCurrentDiscountsDataSourceModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                                 types.String                                                                                       `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                                      `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationDataSourceModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V1ContractDataCurrentProfessionalServicesDataSourceModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                                       `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V1ContractDataCurrentRecurringCommitsDataSourceModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V1ContractDataCurrentRecurringCreditsDataSourceModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V1ContractDataCurrentResellerRoyaltiesDataSourceModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                                       `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                                       `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationDataSourceModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                                      `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UsageFilter                          customfield.NestedObject[V1ContractDataCurrentUsageFilterDataSourceModel]                          `tfsdk:"usage_filter" json:"usage_filter,computed"`
}

type V1ContractDataCurrentCommitsDataSourceModel struct {
	ID                      types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataCurrentCommitsProductDataSourceModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                         `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataCurrentCommitsAccessScheduleDataSourceModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                                        `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                       `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                       `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                       `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                                    `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                                        `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataCurrentCommitsContractDataSourceModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                        `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                         `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceContractDataSourceModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceScheduleDataSourceModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataCurrentCommitsLedgerDataSourceModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                         `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                         `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                        `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                         `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataCurrentCommitsRolledOverFromDataSourceModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                                        `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                                         `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataCurrentCommitsSpecifiersDataSourceModel]  `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                         `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataCurrentCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCommitsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataCurrentCommitsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataCurrentCommitsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentCommitsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCommitsInvoiceContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentCommitsLedgerDataSourceModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataCurrentCommitsRolledOverFromDataSourceModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataCurrentCommitsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataCurrentOverridesDataSourceModel struct {
	ID                    types.String                                                                                  `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                             `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataCurrentOverridesCreditTypeDataSourceModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                             `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                                    `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                                    `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                                    `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                                 `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataCurrentOverridesOverrideSpecifiersDataSourceModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataCurrentOverridesOverrideTiersDataSourceModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataCurrentOverridesOverwriteRateDataSourceModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                                 `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                                 `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentOverridesProductDataSourceModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                                 `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                                  `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                                  `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataCurrentOverridesTiersDataSourceModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                                  `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                                         `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataCurrentOverridesCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesOverrideSpecifiersDataSourceModel struct {
	BillingFrequency        types.String                   `tfsdk:"billing_frequency" json:"billing_frequency,computed"`
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataCurrentOverridesOverrideTiersDataSourceModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateDataSourceModel struct {
	RateType   types.String                                                                                   `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataCurrentOverridesOverwriteRateCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                          `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                                     `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                                  `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                                  `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataCurrentOverridesOverwriteRateTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentOverridesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentScheduledChargesDataSourceModel struct {
	ID                   types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataCurrentScheduledChargesProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataCurrentScheduledChargesScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                                      `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                          `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                           `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                           `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentScheduledChargesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentScheduledChargesScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentScheduledChargesScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentTransitionsDataSourceModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V1ContractDataCurrentUsageStatementScheduleDataSourceModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V1ContractDataCurrentCreditsDataSourceModel struct {
	ID                      types.String                                                                        `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataCurrentCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                        `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataCurrentCreditsAccessScheduleDataSourceModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                      `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                      `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                      `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                                       `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataCurrentCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                       `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                        `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataCurrentCreditsLedgerDataSourceModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                        `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                        `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                       `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                                        `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataCurrentCreditsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                        `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataCurrentCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCreditsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCreditsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataCurrentCreditsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataCurrentCreditsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentCreditsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCreditsLedgerDataSourceModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataCurrentCreditsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataCurrentDiscountsDataSourceModel struct {
	ID                   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataCurrentDiscountsProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataCurrentDiscountsScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                                   `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                    `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                    `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentDiscountsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentDiscountsScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentDiscountsScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentDiscountsScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentDiscountsScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentDiscountsScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                          `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                                       `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                                       `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitDataSourceModel struct {
	ProductID             types.String                                                                                                           `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String]                                                                                         `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                                                         `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                                                                                                           `tfsdk:"description" json:"description,computed"`
	Name                  types.String                                                                                                           `tfsdk:"name" json:"name,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                                    `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                                    `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataCurrentProfessionalServicesDataSourceModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentRecurringCommitsDataSourceModel struct {
	ID                    types.String                                                                                 `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataCurrentRecurringCommitsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataCurrentRecurringCommitsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                                `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentRecurringCommitsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                                 `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                            `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataCurrentRecurringCommitsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                                 `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                            `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V1ContractDataCurrentRecurringCommitsInvoiceAmountDataSourceModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                                 `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                                 `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                                `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataCurrentRecurringCommitsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataCurrentRecurringCommitsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCommitsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataCurrentRecurringCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentRecurringCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentRecurringCommitsInvoiceAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCommitsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataCurrentRecurringCreditsDataSourceModel struct {
	ID                    types.String                                                                                 `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataCurrentRecurringCreditsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataCurrentRecurringCreditsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                                `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentRecurringCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                                 `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                            `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataCurrentRecurringCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                                 `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                            `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                                 `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                                 `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                                `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataCurrentRecurringCreditsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataCurrentRecurringCreditsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCreditsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataCurrentRecurringCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentRecurringCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentRecurringCreditsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataCurrentResellerRoyaltiesDataSourceModel struct {
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

type V1ContractDataCurrentSpendThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                 `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                              `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationCommitDataSourceModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                           `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                           `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataCurrentUsageFilterDataSourceModel struct {
	Current customfield.NestedObject[V1ContractDataCurrentUsageFilterCurrentDataSourceModel]     `tfsdk:"current" json:"current,computed"`
	Initial customfield.NestedObject[V1ContractDataCurrentUsageFilterInitialDataSourceModel]     `tfsdk:"initial" json:"initial,computed"`
	Updates customfield.NestedObjectList[V1ContractDataCurrentUsageFilterUpdatesDataSourceModel] `tfsdk:"updates" json:"updates,computed"`
}

type V1ContractDataCurrentUsageFilterCurrentDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentUsageFilterInitialDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentUsageFilterUpdatesDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialDataSourceModel struct {
	Commits                              customfield.NestedObjectList[V1ContractDataInitialCommitsDataSourceModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                                       `tfsdk:"created_by" json:"created_by,computed"`
	Overrides                            customfield.NestedObjectList[V1ContractDataInitialOverridesDataSourceModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V1ContractDataInitialScheduledChargesDataSourceModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V1ContractDataInitialTransitionsDataSourceModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageStatementSchedule               customfield.NestedObject[V1ContractDataInitialUsageStatementScheduleDataSourceModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	Credits                              customfield.NestedObjectList[V1ContractDataInitialCreditsDataSourceModel]                          `tfsdk:"credits" json:"credits,computed"`
	Discounts                            customfield.NestedObjectList[V1ContractDataInitialDiscountsDataSourceModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                                 types.String                                                                                       `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                                      `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationDataSourceModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V1ContractDataInitialProfessionalServicesDataSourceModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                                       `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V1ContractDataInitialRecurringCommitsDataSourceModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V1ContractDataInitialRecurringCreditsDataSourceModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V1ContractDataInitialResellerRoyaltiesDataSourceModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                                       `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                                       `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationDataSourceModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                                      `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UsageFilter                          customfield.NestedObject[V1ContractDataInitialUsageFilterDataSourceModel]                          `tfsdk:"usage_filter" json:"usage_filter,computed"`
}

type V1ContractDataInitialCommitsDataSourceModel struct {
	ID                      types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataInitialCommitsProductDataSourceModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                         `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataInitialCommitsAccessScheduleDataSourceModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                                        `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                       `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                       `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                       `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                                    `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                                        `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataInitialCommitsContractDataSourceModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                        `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                         `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataInitialCommitsInvoiceContractDataSourceModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataInitialCommitsInvoiceScheduleDataSourceModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataInitialCommitsLedgerDataSourceModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                         `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                         `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                        `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                         `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataInitialCommitsRolledOverFromDataSourceModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                                        `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                                         `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataInitialCommitsSpecifiersDataSourceModel]  `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                         `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataInitialCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCommitsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataInitialCommitsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataInitialCommitsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialCommitsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCommitsInvoiceContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialCommitsInvoiceScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialCommitsLedgerDataSourceModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataInitialCommitsRolledOverFromDataSourceModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataInitialCommitsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataInitialOverridesDataSourceModel struct {
	ID                    types.String                                                                                  `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                             `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataInitialOverridesCreditTypeDataSourceModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                             `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                                    `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                                    `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                                    `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                                 `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataInitialOverridesOverrideSpecifiersDataSourceModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataInitialOverridesOverrideTiersDataSourceModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataInitialOverridesOverwriteRateDataSourceModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                                 `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                                 `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialOverridesProductDataSourceModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                                 `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                                  `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                                  `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataInitialOverridesTiersDataSourceModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                                  `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                                         `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataInitialOverridesCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesOverrideSpecifiersDataSourceModel struct {
	BillingFrequency        types.String                   `tfsdk:"billing_frequency" json:"billing_frequency,computed"`
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataInitialOverridesOverrideTiersDataSourceModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateDataSourceModel struct {
	RateType   types.String                                                                                   `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataInitialOverridesOverwriteRateCreditTypeDataSourceModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                          `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                                     `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                                  `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                                  `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataInitialOverridesOverwriteRateTiersDataSourceModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialOverridesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesTiersDataSourceModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialScheduledChargesDataSourceModel struct {
	ID                   types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataInitialScheduledChargesProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataInitialScheduledChargesScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                                      `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                          `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                           `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                           `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialScheduledChargesProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialScheduledChargesScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialScheduledChargesScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialTransitionsDataSourceModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V1ContractDataInitialUsageStatementScheduleDataSourceModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V1ContractDataInitialCreditsDataSourceModel struct {
	ID                      types.String                                                                        `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataInitialCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                        `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataInitialCreditsAccessScheduleDataSourceModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                      `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                      `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                      `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                                       `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataInitialCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                       `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                        `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataInitialCreditsLedgerDataSourceModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                        `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                        `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                       `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                                        `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	Specifiers              customfield.NestedObjectList[V1ContractDataInitialCreditsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
	UniquenessKey           types.String                                                                        `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataInitialCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCreditsAccessScheduleDataSourceModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCreditsAccessScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataInitialCreditsAccessScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataInitialCreditsAccessScheduleScheduleItemsDataSourceModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialCreditsAccessScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCreditsLedgerDataSourceModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataInitialCreditsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataInitialDiscountsDataSourceModel struct {
	ID                   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataInitialDiscountsProductDataSourceModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataInitialDiscountsScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                                   `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                    `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                    `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialDiscountsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialDiscountsScheduleDataSourceModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialDiscountsScheduleCreditTypeDataSourceModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialDiscountsScheduleScheduleItemsDataSourceModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialDiscountsScheduleCreditTypeDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialDiscountsScheduleScheduleItemsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                          `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                                       `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                                       `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitDataSourceModel struct {
	ProductID             types.String                                                                                                           `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String]                                                                                         `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                                                         `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                                                                                                           `tfsdk:"description" json:"description,computed"`
	Name                  types.String                                                                                                           `tfsdk:"name" json:"name,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                                    `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                                    `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataInitialProfessionalServicesDataSourceModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialRecurringCommitsDataSourceModel struct {
	ID                    types.String                                                                                 `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataInitialRecurringCommitsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataInitialRecurringCommitsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                                `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialRecurringCommitsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                                 `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                            `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataInitialRecurringCommitsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                                 `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                            `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V1ContractDataInitialRecurringCommitsInvoiceAmountDataSourceModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                                 `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                                 `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                                `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataInitialRecurringCommitsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataInitialRecurringCommitsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCommitsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataInitialRecurringCommitsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialRecurringCommitsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialRecurringCommitsInvoiceAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCommitsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataInitialRecurringCreditsDataSourceModel struct {
	ID                    types.String                                                                                 `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataInitialRecurringCreditsAccessAmountDataSourceModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataInitialRecurringCreditsCommitDurationDataSourceModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                                `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialRecurringCreditsProductDataSourceModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                                 `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                            `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataInitialRecurringCreditsContractDataSourceModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                                 `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                            `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                                 `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                                 `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                                `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataInitialRecurringCreditsSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataInitialRecurringCreditsAccessAmountDataSourceModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCreditsCommitDurationDataSourceModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataInitialRecurringCreditsProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialRecurringCreditsContractDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialRecurringCreditsSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataInitialResellerRoyaltiesDataSourceModel struct {
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

type V1ContractDataInitialSpendThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                 `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                              `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationCommitDataSourceModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                           `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                           `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataInitialUsageFilterDataSourceModel struct {
	Current customfield.NestedObject[V1ContractDataInitialUsageFilterCurrentDataSourceModel]     `tfsdk:"current" json:"current,computed"`
	Initial customfield.NestedObject[V1ContractDataInitialUsageFilterInitialDataSourceModel]     `tfsdk:"initial" json:"initial,computed"`
	Updates customfield.NestedObjectList[V1ContractDataInitialUsageFilterUpdatesDataSourceModel] `tfsdk:"updates" json:"updates,computed"`
}

type V1ContractDataInitialUsageFilterCurrentDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialUsageFilterInitialDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialUsageFilterUpdatesDataSourceModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCustomerBillingProviderConfigurationDataSourceModel struct {
	BillingProvider types.String                          `tfsdk:"billing_provider" json:"billing_provider,computed"`
	DeliveryMethod  types.String                          `tfsdk:"delivery_method" json:"delivery_method,computed"`
	ID              types.String                          `tfsdk:"id" json:"id,computed"`
	Configuration   customfield.Map[jsontypes.Normalized] `tfsdk:"configuration" json:"configuration,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                   `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                                `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                                `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationCommitDataSourceModel struct {
	ProductID             types.String                                                                                                    `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String]                                                                                  `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                                                  `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                                                                                                    `tfsdk:"description" json:"description,computed"`
	Name                  types.String                                                                                                    `tfsdk:"name" json:"name,computed"`
	Specifiers            customfield.NestedObjectList[V1ContractDataPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel] `tfsdk:"specifiers" json:"specifiers,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationCommitSpecifiersDataSourceModel struct {
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                             `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                             `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataSpendThresholdConfigurationDataSourceModel struct {
	Commit            customfield.NestedObject[V1ContractDataSpendThresholdConfigurationCommitDataSourceModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                          `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                       `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataSpendThresholdConfigurationCommitDataSourceModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataSpendThresholdConfigurationPaymentGateConfigDataSourceModel struct {
	PaymentGateType types.String                                                                                                    `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                    `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigDataSourceModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataSubscriptionsDataSourceModel struct {
	CollectionSchedule types.String                                                                             `tfsdk:"collection_schedule" json:"collection_schedule,computed"`
	Proration          customfield.NestedObject[V1ContractDataSubscriptionsProrationDataSourceModel]            `tfsdk:"proration" json:"proration,computed"`
	QuantitySchedule   customfield.NestedObjectList[V1ContractDataSubscriptionsQuantityScheduleDataSourceModel] `tfsdk:"quantity_schedule" json:"quantity_schedule,computed"`
	StartingAt         timetypes.RFC3339                                                                        `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	SubscriptionRate   customfield.NestedObject[V1ContractDataSubscriptionsSubscriptionRateDataSourceModel]     `tfsdk:"subscription_rate" json:"subscription_rate,computed"`
	ID                 types.String                                                                             `tfsdk:"id" json:"id,computed"`
	CustomFields       customfield.Map[types.String]                                                            `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description        types.String                                                                             `tfsdk:"description" json:"description,computed"`
	EndingBefore       timetypes.RFC3339                                                                        `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	FiatCreditTypeID   types.String                                                                             `tfsdk:"fiat_credit_type_id" json:"fiat_credit_type_id,computed"`
	Name               types.String                                                                             `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataSubscriptionsProrationDataSourceModel struct {
	InvoiceBehavior types.String `tfsdk:"invoice_behavior" json:"invoice_behavior,computed"`
	IsProrated      types.Bool   `tfsdk:"is_prorated" json:"is_prorated,computed"`
}

type V1ContractDataSubscriptionsQuantityScheduleDataSourceModel struct {
	Quantity     types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}

type V1ContractDataSubscriptionsSubscriptionRateDataSourceModel struct {
	BillingFrequency types.String                                                                                `tfsdk:"billing_frequency" json:"billing_frequency,computed"`
	Product          customfield.NestedObject[V1ContractDataSubscriptionsSubscriptionRateProductDataSourceModel] `tfsdk:"product" json:"product,computed"`
}

type V1ContractDataSubscriptionsSubscriptionRateProductDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}
