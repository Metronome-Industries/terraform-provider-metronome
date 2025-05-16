// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractModel struct {
	CustomerID                           types.String                                         `tfsdk:"customer_id" json:"customer_id,required,no_refresh"`
	StartingAt                           timetypes.RFC3339                                    `tfsdk:"starting_at" json:"starting_at,required,no_refresh" format:"date-time"`
	EndingBefore                         timetypes.RFC3339                                    `tfsdk:"ending_before" json:"ending_before,optional,no_refresh" format:"date-time"`
	MultiplierOverridePrioritization     types.String                                         `tfsdk:"multiplier_override_prioritization" json:"multiplier_override_prioritization,optional,no_refresh"`
	Name                                 types.String                                         `tfsdk:"name" json:"name,optional,no_refresh"`
	NetPaymentTermsDays                  types.Float64                                        `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,optional,no_refresh"`
	NetsuiteSalesOrderID                 types.String                                         `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional,no_refresh"`
	RateCardAlias                        types.String                                         `tfsdk:"rate_card_alias" json:"rate_card_alias,optional,no_refresh"`
	RateCardID                           types.String                                         `tfsdk:"rate_card_id" json:"rate_card_id,optional,no_refresh"`
	SalesforceOpportunityID              types.String                                         `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,optional,no_refresh"`
	ScheduledChargesOnUsageInvoices      types.String                                         `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,optional,no_refresh"`
	TotalContractValue                   types.Float64                                        `tfsdk:"total_contract_value" json:"total_contract_value,optional,no_refresh"`
	UniquenessKey                        types.String                                         `tfsdk:"uniqueness_key" json:"uniqueness_key,optional,no_refresh"`
	CustomFields                         *map[string]types.String                             `tfsdk:"custom_fields" json:"custom_fields,optional,no_refresh"`
	BillingProviderConfiguration         *V1ContractBillingProviderConfigurationModel         `tfsdk:"billing_provider_configuration" json:"billing_provider_configuration,optional,no_refresh"`
	Commits                              *[]*V1ContractCommitsModel                           `tfsdk:"commits" json:"commits,optional,no_refresh"`
	Credits                              *[]*V1ContractCreditsModel                           `tfsdk:"credits" json:"credits,optional,no_refresh"`
	Discounts                            *[]*V1ContractDiscountsModel                         `tfsdk:"discounts" json:"discounts,optional,no_refresh"`
	Overrides                            *[]*V1ContractOverridesModel                         `tfsdk:"overrides" json:"overrides,optional,no_refresh"`
	PrepaidBalanceThresholdConfiguration *V1ContractPrepaidBalanceThresholdConfigurationModel `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,optional,no_refresh"`
	ProfessionalServices                 *[]*V1ContractProfessionalServicesModel              `tfsdk:"professional_services" json:"professional_services,optional,no_refresh"`
	RecurringCommits                     *[]*V1ContractRecurringCommitsModel                  `tfsdk:"recurring_commits" json:"recurring_commits,optional,no_refresh"`
	RecurringCredits                     *[]*V1ContractRecurringCreditsModel                  `tfsdk:"recurring_credits" json:"recurring_credits,optional,no_refresh"`
	ResellerRoyalties                    *[]*V1ContractResellerRoyaltiesModel                 `tfsdk:"reseller_royalties" json:"reseller_royalties,optional,no_refresh"`
	ScheduledCharges                     *[]*V1ContractScheduledChargesModel                  `tfsdk:"scheduled_charges" json:"scheduled_charges,optional,no_refresh"`
	SpendThresholdConfiguration          *V1ContractSpendThresholdConfigurationModel          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,optional,no_refresh"`
	Transition                           *V1ContractTransitionModel                           `tfsdk:"transition" json:"transition,optional,no_refresh"`
	UsageFilter                          *V1ContractUsageFilterModel                          `tfsdk:"usage_filter" json:"usage_filter,optional,no_refresh"`
	UsageStatementSchedule               *V1ContractUsageStatementScheduleModel               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,optional,no_refresh"`
	Data                                 customfield.NestedObject[V1ContractDataModel]        `tfsdk:"data" json:"data,computed"`
}

func (m V1ContractModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1ContractModel) MarshalJSONForUpdate(state V1ContractModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1ContractBillingProviderConfigurationModel struct {
	BillingProvider                types.String `tfsdk:"billing_provider" json:"billing_provider,optional"`
	BillingProviderConfigurationID types.String `tfsdk:"billing_provider_configuration_id" json:"billing_provider_configuration_id,optional"`
	DeliveryMethod                 types.String `tfsdk:"delivery_method" json:"delivery_method,optional"`
}

type V1ContractCommitsModel struct {
	ProductID             types.String                             `tfsdk:"product_id" json:"product_id,required"`
	Type                  types.String                             `tfsdk:"type" json:"type,required"`
	AccessSchedule        *V1ContractCommitsAccessScheduleModel    `tfsdk:"access_schedule" json:"access_schedule,optional"`
	Amount                types.Float64                            `tfsdk:"amount" json:"amount,optional"`
	ApplicableProductIDs  *[]types.String                          `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                          `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields          *map[string]types.String                 `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description           types.String                             `tfsdk:"description" json:"description,optional"`
	InvoiceSchedule       *V1ContractCommitsInvoiceScheduleModel   `tfsdk:"invoice_schedule" json:"invoice_schedule,optional"`
	Name                  types.String                             `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	PaymentGateConfig     *V1ContractCommitsPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,optional"`
	Priority              types.Float64                            `tfsdk:"priority" json:"priority,optional"`
	RateType              types.String                             `tfsdk:"rate_type" json:"rate_type,optional"`
	RolloverFraction      types.Float64                            `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                             `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V1ContractCommitsAccessScheduleModel struct {
	ScheduleItems *[]*V1ContractCommitsAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                          `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V1ContractCommitsAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V1ContractCommitsInvoiceScheduleModel struct {
	CreditTypeID      types.String                                            `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V1ContractCommitsInvoiceScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V1ContractCommitsInvoiceScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V1ContractCommitsInvoiceScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractCommitsInvoiceScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractCommitsPaymentGateConfigModel struct {
	PaymentGateType types.String                                         `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V1ContractCommitsPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                         `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V1ContractCommitsPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V1ContractCreditsModel struct {
	AccessSchedule        *V1ContractCreditsAccessScheduleModel `tfsdk:"access_schedule" json:"access_schedule,required"`
	ProductID             types.String                          `tfsdk:"product_id" json:"product_id,required"`
	ApplicableProductIDs  *[]types.String                       `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                       `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields          *map[string]types.String              `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description           types.String                          `tfsdk:"description" json:"description,optional"`
	Name                  types.String                          `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Priority              types.Float64                         `tfsdk:"priority" json:"priority,optional"`
	RateType              types.String                          `tfsdk:"rate_type" json:"rate_type,optional"`
}

type V1ContractCreditsAccessScheduleModel struct {
	ScheduleItems *[]*V1ContractCreditsAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                          `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V1ContractCreditsAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V1ContractDiscountsModel struct {
	ProductID            types.String                      `tfsdk:"product_id" json:"product_id,required"`
	Schedule             *V1ContractDiscountsScheduleModel `tfsdk:"schedule" json:"schedule,required"`
	CustomFields         *map[string]types.String          `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Name                 types.String                      `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID types.String                      `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V1ContractDiscountsScheduleModel struct {
	CreditTypeID      types.String                                       `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V1ContractDiscountsScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V1ContractDiscountsScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V1ContractDiscountsScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractDiscountsScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractOverridesModel struct {
	StartingAt            timetypes.RFC3339                              `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductTags *[]types.String                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	EndingBefore          timetypes.RFC3339                              `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	Entitled              types.Bool                                     `tfsdk:"entitled" json:"entitled,optional"`
	IsCommitSpecific      types.Bool                                     `tfsdk:"is_commit_specific" json:"is_commit_specific,optional"`
	Multiplier            types.Float64                                  `tfsdk:"multiplier" json:"multiplier,optional"`
	OverrideSpecifiers    *[]*V1ContractOverridesOverrideSpecifiersModel `tfsdk:"override_specifiers" json:"override_specifiers,optional"`
	OverwriteRate         *V1ContractOverridesOverwriteRateModel         `tfsdk:"overwrite_rate" json:"overwrite_rate,optional"`
	Priority              types.Float64                                  `tfsdk:"priority" json:"priority,optional"`
	ProductID             types.String                                   `tfsdk:"product_id" json:"product_id,optional"`
	Target                types.String                                   `tfsdk:"target" json:"target,optional"`
	Tiers                 *[]*V1ContractOverridesTiersModel              `tfsdk:"tiers" json:"tiers,optional"`
	Type                  types.String                                   `tfsdk:"type" json:"type,optional"`
}

type V1ContractOverridesOverrideSpecifiersModel struct {
	CommitIDs               *[]types.String          `tfsdk:"commit_ids" json:"commit_ids,optional"`
	PresentationGroupValues *map[string]types.String `tfsdk:"presentation_group_values" json:"presentation_group_values,optional"`
	PricingGroupValues      *map[string]types.String `tfsdk:"pricing_group_values" json:"pricing_group_values,optional"`
	ProductID               types.String             `tfsdk:"product_id" json:"product_id,optional"`
	ProductTags             *[]types.String          `tfsdk:"product_tags" json:"product_tags,optional"`
	RecurringCommitIDs      *[]types.String          `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,optional"`
	RecurringCreditIDs      *[]types.String          `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,optional"`
}

type V1ContractOverridesOverwriteRateModel struct {
	RateType     types.String                                   `tfsdk:"rate_type" json:"rate_type,required"`
	CreditTypeID types.String                                   `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	CustomRate   *map[string]jsontypes.Normalized               `tfsdk:"custom_rate" json:"custom_rate,optional"`
	IsProrated   types.Bool                                     `tfsdk:"is_prorated" json:"is_prorated,optional"`
	Price        types.Float64                                  `tfsdk:"price" json:"price,optional"`
	Quantity     types.Float64                                  `tfsdk:"quantity" json:"quantity,optional"`
	Tiers        *[]*V1ContractOverridesOverwriteRateTiersModel `tfsdk:"tiers" json:"tiers,optional"`
}

type V1ContractOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,required"`
	Size  types.Float64 `tfsdk:"size" json:"size,optional"`
}

type V1ContractOverridesTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,required"`
	Size       types.Float64 `tfsdk:"size" json:"size,optional"`
}

type V1ContractPrepaidBalanceThresholdConfigurationModel struct {
	Commit            *V1ContractPrepaidBalanceThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,required"`
	IsEnabled         types.Bool                                                            `tfsdk:"is_enabled" json:"is_enabled,required"`
	PaymentGateConfig *V1ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,required"`
	RechargeToAmount  types.Float64                                                         `tfsdk:"recharge_to_amount" json:"recharge_to_amount,required"`
	ThresholdAmount   types.Float64                                                         `tfsdk:"threshold_amount" json:"threshold_amount,required"`
}

type V1ContractPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String    `tfsdk:"product_id" json:"product_id,required"`
	ApplicableProductIDs  *[]types.String `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String    `tfsdk:"description" json:"description,optional"`
	Name                  types.String    `tfsdk:"name" json:"name,optional"`
}

type V1ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                      `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V1ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                                      `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V1ContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V1ContractProfessionalServicesModel struct {
	MaxAmount            types.Float64            `tfsdk:"max_amount" json:"max_amount,required"`
	ProductID            types.String             `tfsdk:"product_id" json:"product_id,required"`
	Quantity             types.Float64            `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice            types.Float64            `tfsdk:"unit_price" json:"unit_price,required"`
	CustomFields         *map[string]types.String `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description          types.String             `tfsdk:"description" json:"description,optional"`
	NetsuiteSalesOrderID types.String             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V1ContractRecurringCommitsModel struct {
	AccessAmount          *V1ContractRecurringCommitsAccessAmountModel   `tfsdk:"access_amount" json:"access_amount,required"`
	CommitDuration        *V1ContractRecurringCommitsCommitDurationModel `tfsdk:"commit_duration" json:"commit_duration,required"`
	Priority              types.Float64                                  `tfsdk:"priority" json:"priority,required"`
	ProductID             types.String                                   `tfsdk:"product_id" json:"product_id,required"`
	StartingAt            timetypes.RFC3339                              `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  *[]types.String                                `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String                                   `tfsdk:"description" json:"description,optional"`
	EndingBefore          timetypes.RFC3339                              `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	InvoiceAmount         *V1ContractRecurringCommitsInvoiceAmountModel  `tfsdk:"invoice_amount" json:"invoice_amount,optional"`
	Name                  types.String                                   `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                                   `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Proration             types.String                                   `tfsdk:"proration" json:"proration,optional"`
	RateType              types.String                                   `tfsdk:"rate_type" json:"rate_type,optional"`
	RecurrenceFrequency   types.String                                   `tfsdk:"recurrence_frequency" json:"recurrence_frequency,optional"`
	RolloverFraction      types.Float64                                  `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                                   `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V1ContractRecurringCommitsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V1ContractRecurringCommitsCommitDurationModel struct {
	Unit  types.String  `tfsdk:"unit" json:"unit,required"`
	Value types.Float64 `tfsdk:"value" json:"value,required"`
}

type V1ContractRecurringCommitsInvoiceAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V1ContractRecurringCreditsModel struct {
	AccessAmount          *V1ContractRecurringCreditsAccessAmountModel   `tfsdk:"access_amount" json:"access_amount,required"`
	CommitDuration        *V1ContractRecurringCreditsCommitDurationModel `tfsdk:"commit_duration" json:"commit_duration,required"`
	Priority              types.Float64                                  `tfsdk:"priority" json:"priority,required"`
	ProductID             types.String                                   `tfsdk:"product_id" json:"product_id,required"`
	StartingAt            timetypes.RFC3339                              `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  *[]types.String                                `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String                                   `tfsdk:"description" json:"description,optional"`
	EndingBefore          timetypes.RFC3339                              `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	Name                  types.String                                   `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                                   `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Proration             types.String                                   `tfsdk:"proration" json:"proration,optional"`
	RateType              types.String                                   `tfsdk:"rate_type" json:"rate_type,optional"`
	RecurrenceFrequency   types.String                                   `tfsdk:"recurrence_frequency" json:"recurrence_frequency,optional"`
	RolloverFraction      types.Float64                                  `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                                   `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V1ContractRecurringCreditsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V1ContractRecurringCreditsCommitDurationModel struct {
	Unit  types.String  `tfsdk:"unit" json:"unit,required"`
	Value types.Float64 `tfsdk:"value" json:"value,required"`
}

type V1ContractResellerRoyaltiesModel struct {
	Fraction              types.Float64                               `tfsdk:"fraction" json:"fraction,required"`
	NetsuiteResellerID    types.String                                `tfsdk:"netsuite_reseller_id" json:"netsuite_reseller_id,required"`
	ResellerType          types.String                                `tfsdk:"reseller_type" json:"reseller_type,required"`
	StartingAt            timetypes.RFC3339                           `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  *[]types.String                             `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                             `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	AwsOptions            *V1ContractResellerRoyaltiesAwsOptionsModel `tfsdk:"aws_options" json:"aws_options,optional"`
	EndingBefore          timetypes.RFC3339                           `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	GcpOptions            *V1ContractResellerRoyaltiesGcpOptionsModel `tfsdk:"gcp_options" json:"gcp_options,optional"`
	ResellerContractValue types.Float64                               `tfsdk:"reseller_contract_value" json:"reseller_contract_value,optional"`
}

type V1ContractResellerRoyaltiesAwsOptionsModel struct {
	AwsAccountNumber    types.String `tfsdk:"aws_account_number" json:"aws_account_number,optional"`
	AwsOfferID          types.String `tfsdk:"aws_offer_id" json:"aws_offer_id,optional"`
	AwsPayerReferenceID types.String `tfsdk:"aws_payer_reference_id" json:"aws_payer_reference_id,optional"`
}

type V1ContractResellerRoyaltiesGcpOptionsModel struct {
	GcpAccountID types.String `tfsdk:"gcp_account_id" json:"gcp_account_id,optional"`
	GcpOfferID   types.String `tfsdk:"gcp_offer_id" json:"gcp_offer_id,optional"`
}

type V1ContractScheduledChargesModel struct {
	ProductID            types.String                             `tfsdk:"product_id" json:"product_id,required"`
	Schedule             *V1ContractScheduledChargesScheduleModel `tfsdk:"schedule" json:"schedule,required"`
	Name                 types.String                             `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID types.String                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V1ContractScheduledChargesScheduleModel struct {
	CreditTypeID      types.String                                              `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V1ContractScheduledChargesScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V1ContractScheduledChargesScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V1ContractScheduledChargesScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractScheduledChargesScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V1ContractSpendThresholdConfigurationModel struct {
	Commit            *V1ContractSpendThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,required"`
	IsEnabled         types.Bool                                                   `tfsdk:"is_enabled" json:"is_enabled,required"`
	PaymentGateConfig *V1ContractSpendThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,required"`
	ThresholdAmount   types.Float64                                                `tfsdk:"threshold_amount" json:"threshold_amount,required"`
}

type V1ContractSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
}

type V1ContractSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                             `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V1ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                             `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V1ContractSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V1ContractTransitionModel struct {
	FromContractID        types.String                                    `tfsdk:"from_contract_id" json:"from_contract_id,required"`
	Type                  types.String                                    `tfsdk:"type" json:"type,required"`
	FutureInvoiceBehavior *V1ContractTransitionFutureInvoiceBehaviorModel `tfsdk:"future_invoice_behavior" json:"future_invoice_behavior,optional"`
}

type V1ContractTransitionFutureInvoiceBehaviorModel struct {
	Trueup types.String `tfsdk:"trueup" json:"trueup,optional"`
}

type V1ContractUsageFilterModel struct {
	GroupKey    types.String      `tfsdk:"group_key" json:"group_key,required"`
	GroupValues *[]types.String   `tfsdk:"group_values" json:"group_values,required"`
	StartingAt  timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,optional" format:"date-time"`
}

type V1ContractUsageStatementScheduleModel struct {
	Frequency                   types.String      `tfsdk:"frequency" json:"frequency,required"`
	BillingAnchorDate           timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,optional" format:"date-time"`
	Day                         types.String      `tfsdk:"day" json:"day,optional"`
	InvoiceGenerationStartingAt timetypes.RFC3339 `tfsdk:"invoice_generation_starting_at" json:"invoice_generation_starting_at,optional" format:"date-time"`
}

type V1ContractDataModel struct {
	ID                                   types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Amendments                           customfield.NestedObjectList[V1ContractDataAmendmentsModel]                       `tfsdk:"amendments" json:"amendments,computed"`
	Current                              customfield.NestedObject[V1ContractDataCurrentModel]                              `tfsdk:"current" json:"current,computed"`
	CustomerID                           types.String                                                                      `tfsdk:"customer_id" json:"customer_id,computed"`
	Initial                              customfield.NestedObject[V1ContractDataInitialModel]                              `tfsdk:"initial" json:"initial,computed"`
	ArchivedAt                           timetypes.RFC3339                                                                 `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields                         customfield.Map[types.String]                                                     `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerBillingProviderConfiguration customfield.NestedObject[V1ContractDataCustomerBillingProviderConfigurationModel] `tfsdk:"customer_billing_provider_configuration" json:"customer_billing_provider_configuration,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                      `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataSpendThresholdConfigurationModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	UniquenessKey                        types.String                                                                      `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsModel struct {
	ID                      types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Commits                 customfield.NestedObjectList[V1ContractDataAmendmentsCommitsModel]              `tfsdk:"commits" json:"commits,computed"`
	CreatedAt               timetypes.RFC3339                                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy               types.String                                                                    `tfsdk:"created_by" json:"created_by,computed"`
	Overrides               customfield.NestedObjectList[V1ContractDataAmendmentsOverridesModel]            `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges        customfield.NestedObjectList[V1ContractDataAmendmentsScheduledChargesModel]     `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt              timetypes.RFC3339                                                               `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Credits                 customfield.NestedObjectList[V1ContractDataAmendmentsCreditsModel]              `tfsdk:"credits" json:"credits,computed"`
	Discounts               customfield.NestedObjectList[V1ContractDataAmendmentsDiscountsModel]            `tfsdk:"discounts" json:"discounts,computed"`
	NetsuiteSalesOrderID    types.String                                                                    `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	ProfessionalServices    customfield.NestedObjectList[V1ContractDataAmendmentsProfessionalServicesModel] `tfsdk:"professional_services" json:"professional_services,computed"`
	ResellerRoyalties       customfield.NestedObjectList[V1ContractDataAmendmentsResellerRoyaltiesModel]    `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID types.String                                                                    `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V1ContractDataAmendmentsCommitsModel struct {
	ID                      types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataAmendmentsCommitsProductModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                  `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataAmendmentsCommitsAccessScheduleModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                                 `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                                `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                                `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                             `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                                 `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataAmendmentsCommitsContractModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                 `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                  `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceContractModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceScheduleModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataAmendmentsCommitsLedgerModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                  `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                 `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                  `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataAmendmentsCommitsRolledOverFromModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                                 `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                                  `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                                  `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataAmendmentsCommitsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCommitsInvoiceScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsCommitsLedgerModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataAmendmentsCommitsRolledOverFromModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataAmendmentsOverridesModel struct {
	ID                    types.String                                                                           `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                      `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                         `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataAmendmentsOverridesCreditTypeModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                      `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                             `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                             `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                             `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                          `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverrideSpecifiersModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverrideTiersModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataAmendmentsOverridesOverwriteRateModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                          `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                          `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataAmendmentsOverridesProductModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                          `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                           `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                           `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataAmendmentsOverridesTiersModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                           `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                                  `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataAmendmentsOverridesCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesOverrideSpecifiersModel struct {
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataAmendmentsOverridesOverrideTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateModel struct {
	RateType   types.String                                                                            `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                   `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                              `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                           `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                           `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataAmendmentsOverridesOverwriteRateTiersModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsOverridesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsOverridesTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataAmendmentsScheduledChargesModel struct {
	ID                   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                               `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                   `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                    `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                    `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsScheduledChargesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsScheduledChargesScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsCreditsModel struct {
	ID                      types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataAmendmentsCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                                 `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataAmendmentsCreditsAccessScheduleModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                               `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                               `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                                `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataAmendmentsCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataAmendmentsCreditsLedgerModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                                `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                                 `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                                 `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                                 `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataAmendmentsCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataAmendmentsCreditsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataAmendmentsCreditsLedgerModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataAmendmentsDiscountsModel struct {
	ID                   types.String                                                             `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataAmendmentsDiscountsProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataAmendmentsDiscountsScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                            `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                             `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsDiscountsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataAmendmentsDiscountsScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataAmendmentsDiscountsScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataAmendmentsDiscountsScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataAmendmentsProfessionalServicesModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataAmendmentsResellerRoyaltiesModel struct {
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

type V1ContractDataCurrentModel struct {
	Commits                              customfield.NestedObjectList[V1ContractDataCurrentCommitsModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                             `tfsdk:"created_by" json:"created_by,computed"`
	Overrides                            customfield.NestedObjectList[V1ContractDataCurrentOverridesModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V1ContractDataCurrentScheduledChargesModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                        `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V1ContractDataCurrentTransitionsModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageStatementSchedule               customfield.NestedObject[V1ContractDataCurrentUsageStatementScheduleModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	Credits                              customfield.NestedObjectList[V1ContractDataCurrentCreditsModel]                          `tfsdk:"credits" json:"credits,computed"`
	Discounts                            customfield.NestedObjectList[V1ContractDataCurrentDiscountsModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                        `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                                 types.String                                                                             `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                            `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V1ContractDataCurrentProfessionalServicesModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                             `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V1ContractDataCurrentRecurringCommitsModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V1ContractDataCurrentRecurringCreditsModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V1ContractDataCurrentResellerRoyaltiesModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                             `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                             `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                            `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UsageFilter                          customfield.NestedObject[V1ContractDataCurrentUsageFilterModel]                          `tfsdk:"usage_filter" json:"usage_filter,computed"`
}

type V1ContractDataCurrentCommitsModel struct {
	ID                      types.String                                                               `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataCurrentCommitsProductModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                               `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataCurrentCommitsAccessScheduleModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                              `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                             `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                             `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                             `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                          `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                              `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataCurrentCommitsContractModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                              `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                               `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceContractModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceScheduleModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataCurrentCommitsLedgerModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                               `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                               `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                              `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                               `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataCurrentCommitsRolledOverFromModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                              `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                               `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                               `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataCurrentCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCommitsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataCurrentCommitsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataCurrentCommitsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentCommitsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCommitsInvoiceContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCommitsInvoiceScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentCommitsLedgerModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataCurrentCommitsRolledOverFromModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataCurrentOverridesModel struct {
	ID                    types.String                                                                        `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                   `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                      `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataCurrentOverridesCreditTypeModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                   `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                          `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                          `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                          `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                       `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataCurrentOverridesOverrideSpecifiersModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataCurrentOverridesOverrideTiersModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataCurrentOverridesOverwriteRateModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                       `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                       `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentOverridesProductModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                       `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                        `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataCurrentOverridesTiersModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                               `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataCurrentOverridesCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesOverrideSpecifiersModel struct {
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataCurrentOverridesOverrideTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateModel struct {
	RateType   types.String                                                                         `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataCurrentOverridesOverwriteRateCreditTypeModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                           `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                        `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                        `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataCurrentOverridesOverwriteRateTiersModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentOverridesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentOverridesTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataCurrentScheduledChargesModel struct {
	ID                   types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataCurrentScheduledChargesProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataCurrentScheduledChargesScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                            `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentScheduledChargesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentScheduledChargesScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentScheduledChargesScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentScheduledChargesScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentTransitionsModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V1ContractDataCurrentUsageStatementScheduleModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V1ContractDataCurrentCreditsModel struct {
	ID                      types.String                                                              `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataCurrentCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                              `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataCurrentCreditsAccessScheduleModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                            `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                            `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                            `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                             `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataCurrentCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                              `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataCurrentCreditsLedgerModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                              `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                              `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                             `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                              `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                              `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                              `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataCurrentCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCreditsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentCreditsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataCurrentCreditsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataCurrentCreditsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentCreditsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentCreditsLedgerModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataCurrentDiscountsModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataCurrentDiscountsProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataCurrentDiscountsScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                         `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentDiscountsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentDiscountsScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataCurrentDiscountsScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataCurrentDiscountsScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataCurrentDiscountsScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentDiscountsScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                             `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                             `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                   `tfsdk:"description" json:"description,computed"`
	Name                  types.String                   `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                          `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                          `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataCurrentPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataCurrentProfessionalServicesModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataCurrentRecurringCommitsModel struct {
	ID                    types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataCurrentRecurringCommitsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataCurrentRecurringCommitsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                      `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentRecurringCommitsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                       `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataCurrentRecurringCommitsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                       `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V1ContractDataCurrentRecurringCommitsInvoiceAmountModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                       `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                       `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                      `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V1ContractDataCurrentRecurringCommitsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCommitsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataCurrentRecurringCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentRecurringCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentRecurringCommitsInvoiceAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCreditsModel struct {
	ID                    types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataCurrentRecurringCreditsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataCurrentRecurringCreditsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                      `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataCurrentRecurringCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                       `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataCurrentRecurringCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                       `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                       `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                       `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                      `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V1ContractDataCurrentRecurringCreditsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataCurrentRecurringCreditsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataCurrentRecurringCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentRecurringCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataCurrentResellerRoyaltiesModel struct {
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

type V1ContractDataCurrentSpendThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                       `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                    `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                 `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                 `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataCurrentSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataCurrentUsageFilterModel struct {
	Current customfield.NestedObject[V1ContractDataCurrentUsageFilterCurrentModel]     `tfsdk:"current" json:"current,computed"`
	Initial customfield.NestedObject[V1ContractDataCurrentUsageFilterInitialModel]     `tfsdk:"initial" json:"initial,computed"`
	Updates customfield.NestedObjectList[V1ContractDataCurrentUsageFilterUpdatesModel] `tfsdk:"updates" json:"updates,computed"`
}

type V1ContractDataCurrentUsageFilterCurrentModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentUsageFilterInitialModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCurrentUsageFilterUpdatesModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialModel struct {
	Commits                              customfield.NestedObjectList[V1ContractDataInitialCommitsModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                             `tfsdk:"created_by" json:"created_by,computed"`
	Overrides                            customfield.NestedObjectList[V1ContractDataInitialOverridesModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V1ContractDataInitialScheduledChargesModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                        `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V1ContractDataInitialTransitionsModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageStatementSchedule               customfield.NestedObject[V1ContractDataInitialUsageStatementScheduleModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	Credits                              customfield.NestedObjectList[V1ContractDataInitialCreditsModel]                          `tfsdk:"credits" json:"credits,computed"`
	Discounts                            customfield.NestedObjectList[V1ContractDataInitialDiscountsModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                        `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                                 types.String                                                                             `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                            `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V1ContractDataInitialProfessionalServicesModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                             `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V1ContractDataInitialRecurringCommitsModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V1ContractDataInitialRecurringCreditsModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V1ContractDataInitialResellerRoyaltiesModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                             `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                             `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                            `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UsageFilter                          customfield.NestedObject[V1ContractDataInitialUsageFilterModel]                          `tfsdk:"usage_filter" json:"usage_filter,computed"`
}

type V1ContractDataInitialCommitsModel struct {
	ID                      types.String                                                               `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataInitialCommitsProductModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                               `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataInitialCommitsAccessScheduleModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	Amount                  types.Float64                                                              `tfsdk:"amount" json:"amount,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                             `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                             `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                             `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                          `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                              `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataInitialCommitsContractModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                              `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                               `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V1ContractDataInitialCommitsInvoiceContractModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V1ContractDataInitialCommitsInvoiceScheduleModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataInitialCommitsLedgerModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                               `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                               `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                              `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                               `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V1ContractDataInitialCommitsRolledOverFromModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                              `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                               `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                               `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataInitialCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCommitsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataInitialCommitsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataInitialCommitsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialCommitsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCommitsInvoiceContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialCommitsInvoiceScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCommitsInvoiceScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialCommitsLedgerModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataInitialCommitsRolledOverFromModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V1ContractDataInitialOverridesModel struct {
	ID                    types.String                                                                        `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                                   `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                                      `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	CreditType            customfield.NestedObject[V1ContractDataInitialOverridesCreditTypeModel]             `tfsdk:"credit_type" json:"credit_type,computed"`
	EndingBefore          timetypes.RFC3339                                                                   `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                          `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                          `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	IsProrated            types.Bool                                                                          `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Multiplier            types.Float64                                                                       `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V1ContractDataInitialOverridesOverrideSpecifiersModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V1ContractDataInitialOverridesOverrideTiersModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V1ContractDataInitialOverridesOverwriteRateModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Price                 types.Float64                                                                       `tfsdk:"price" json:"price,computed"`
	Priority              types.Float64                                                                       `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialOverridesProductModel]                `tfsdk:"product" json:"product,computed"`
	Quantity              types.Float64                                                                       `tfsdk:"quantity" json:"quantity,computed"`
	RateType              types.String                                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	Target                types.String                                                                        `tfsdk:"target" json:"target,computed"`
	Tiers                 customfield.NestedObjectList[V1ContractDataInitialOverridesTiersModel]              `tfsdk:"tiers" json:"tiers,computed"`
	Type                  types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Value                 customfield.Map[jsontypes.Normalized]                                               `tfsdk:"value" json:"value,computed"`
}

type V1ContractDataInitialOverridesCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesOverrideSpecifiersModel struct {
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V1ContractDataInitialOverridesOverrideTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateModel struct {
	RateType   types.String                                                                         `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V1ContractDataInitialOverridesOverwriteRateCreditTypeModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                                `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                           `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                        `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                        `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V1ContractDataInitialOverridesOverwriteRateTiersModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialOverridesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialOverridesTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V1ContractDataInitialScheduledChargesModel struct {
	ID                   types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataInitialScheduledChargesProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataInitialScheduledChargesScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                            `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                                 `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialScheduledChargesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialScheduledChargesScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialScheduledChargesScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialScheduledChargesScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialTransitionsModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V1ContractDataInitialUsageStatementScheduleModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V1ContractDataInitialCreditsModel struct {
	ID                      types.String                                                              `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V1ContractDataInitialCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                              `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V1ContractDataInitialCreditsAccessScheduleModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                            `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                            `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                            `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                             `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V1ContractDataInitialCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                              `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V1ContractDataInitialCreditsLedgerModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                              `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                              `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                             `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                              `tfsdk:"rate_type" json:"rate_type,computed"`
	SalesforceOpportunityID types.String                                                              `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	UniquenessKey           types.String                                                              `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V1ContractDataInitialCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCreditsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialCreditsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V1ContractDataInitialCreditsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V1ContractDataInitialCreditsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialCreditsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialCreditsLedgerModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V1ContractDataInitialDiscountsModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V1ContractDataInitialDiscountsProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V1ContractDataInitialDiscountsScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                         `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialDiscountsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialDiscountsScheduleModel struct {
	CreditType    customfield.NestedObject[V1ContractDataInitialDiscountsScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V1ContractDataInitialDiscountsScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V1ContractDataInitialDiscountsScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialDiscountsScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                                `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                             `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                             `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                   `tfsdk:"description" json:"description,computed"`
	Name                  types.String                   `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                          `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                          `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataInitialPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataInitialProfessionalServicesModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V1ContractDataInitialRecurringCommitsModel struct {
	ID                    types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataInitialRecurringCommitsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataInitialRecurringCommitsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                      `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialRecurringCommitsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                       `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataInitialRecurringCommitsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                       `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V1ContractDataInitialRecurringCommitsInvoiceAmountModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                       `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                       `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                      `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V1ContractDataInitialRecurringCommitsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCommitsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataInitialRecurringCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialRecurringCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialRecurringCommitsInvoiceAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCreditsModel struct {
	ID                    types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V1ContractDataInitialRecurringCreditsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V1ContractDataInitialRecurringCreditsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                                      `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V1ContractDataInitialRecurringCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                       `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                                  `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V1ContractDataInitialRecurringCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                       `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                                  `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                       `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                       `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                                      `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V1ContractDataInitialRecurringCreditsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V1ContractDataInitialRecurringCreditsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V1ContractDataInitialRecurringCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialRecurringCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V1ContractDataInitialResellerRoyaltiesModel struct {
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

type V1ContractDataInitialSpendThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                       `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                                    `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                 `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                 `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataInitialSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataInitialUsageFilterModel struct {
	Current customfield.NestedObject[V1ContractDataInitialUsageFilterCurrentModel]     `tfsdk:"current" json:"current,computed"`
	Initial customfield.NestedObject[V1ContractDataInitialUsageFilterInitialModel]     `tfsdk:"initial" json:"initial,computed"`
	Updates customfield.NestedObjectList[V1ContractDataInitialUsageFilterUpdatesModel] `tfsdk:"updates" json:"updates,computed"`
}

type V1ContractDataInitialUsageFilterCurrentModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialUsageFilterInitialModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataInitialUsageFilterUpdatesModel struct {
	GroupKey    types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt  timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V1ContractDataCustomerBillingProviderConfigurationModel struct {
	BillingProvider types.String                          `tfsdk:"billing_provider" json:"billing_provider,computed"`
	DeliveryMethod  types.String                          `tfsdk:"delivery_method" json:"delivery_method,computed"`
	ID              types.String                          `tfsdk:"id" json:"id,computed"`
	Configuration   customfield.Map[jsontypes.Normalized] `tfsdk:"configuration" json:"configuration,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                         `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                      `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                      `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                   `tfsdk:"description" json:"description,computed"`
	Name                  types.String                   `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                   `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                   `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V1ContractDataSpendThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V1ContractDataSpendThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V1ContractDataSpendThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                             `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V1ContractDataSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V1ContractDataSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                          `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                          `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V1ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}
