// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v2_contract

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V2ContractModel struct {
	ContractID                                 types.String                                               `tfsdk:"contract_id" json:"contract_id,required,no_refresh"`
	CustomerID                                 types.String                                               `tfsdk:"customer_id" json:"customer_id,required,no_refresh"`
	AllowContractEndingBeforeFinalizedInvoice  types.Bool                                                 `tfsdk:"allow_contract_ending_before_finalized_invoice" json:"allow_contract_ending_before_finalized_invoice,optional,no_refresh"`
	UpdateContractEndDate                      timetypes.RFC3339                                          `tfsdk:"update_contract_end_date" json:"update_contract_end_date,optional,no_refresh" format:"date-time"`
	AddCommits                                 *[]*V2ContractAddCommitsModel                              `tfsdk:"add_commits" json:"add_commits,optional,no_refresh"`
	AddCredits                                 *[]*V2ContractAddCreditsModel                              `tfsdk:"add_credits" json:"add_credits,optional,no_refresh"`
	AddDiscounts                               *[]*V2ContractAddDiscountsModel                            `tfsdk:"add_discounts" json:"add_discounts,optional,no_refresh"`
	AddOverrides                               *[]*V2ContractAddOverridesModel                            `tfsdk:"add_overrides" json:"add_overrides,optional,no_refresh"`
	AddPrepaidBalanceThresholdConfiguration    *V2ContractAddPrepaidBalanceThresholdConfigurationModel    `tfsdk:"add_prepaid_balance_threshold_configuration" json:"add_prepaid_balance_threshold_configuration,optional,no_refresh"`
	AddProfessionalServices                    *[]*V2ContractAddProfessionalServicesModel                 `tfsdk:"add_professional_services" json:"add_professional_services,optional,no_refresh"`
	AddRecurringCommits                        *[]*V2ContractAddRecurringCommitsModel                     `tfsdk:"add_recurring_commits" json:"add_recurring_commits,optional,no_refresh"`
	AddRecurringCredits                        *[]*V2ContractAddRecurringCreditsModel                     `tfsdk:"add_recurring_credits" json:"add_recurring_credits,optional,no_refresh"`
	AddResellerRoyalties                       *[]*V2ContractAddResellerRoyaltiesModel                    `tfsdk:"add_reseller_royalties" json:"add_reseller_royalties,optional,no_refresh"`
	AddScheduledCharges                        *[]*V2ContractAddScheduledChargesModel                     `tfsdk:"add_scheduled_charges" json:"add_scheduled_charges,optional,no_refresh"`
	AddSpendThresholdConfiguration             *V2ContractAddSpendThresholdConfigurationModel             `tfsdk:"add_spend_threshold_configuration" json:"add_spend_threshold_configuration,optional,no_refresh"`
	ArchiveCommits                             *[]*V2ContractArchiveCommitsModel                          `tfsdk:"archive_commits" json:"archive_commits,optional,no_refresh"`
	ArchiveCredits                             *[]*V2ContractArchiveCreditsModel                          `tfsdk:"archive_credits" json:"archive_credits,optional,no_refresh"`
	ArchiveScheduledCharges                    *[]*V2ContractArchiveScheduledChargesModel                 `tfsdk:"archive_scheduled_charges" json:"archive_scheduled_charges,optional,no_refresh"`
	RemoveOverrides                            *[]*V2ContractRemoveOverridesModel                         `tfsdk:"remove_overrides" json:"remove_overrides,optional,no_refresh"`
	UpdateCommits                              *[]*V2ContractUpdateCommitsModel                           `tfsdk:"update_commits" json:"update_commits,optional,no_refresh"`
	UpdateCredits                              *[]*V2ContractUpdateCreditsModel                           `tfsdk:"update_credits" json:"update_credits,optional,no_refresh"`
	UpdatePrepaidBalanceThresholdConfiguration *V2ContractUpdatePrepaidBalanceThresholdConfigurationModel `tfsdk:"update_prepaid_balance_threshold_configuration" json:"update_prepaid_balance_threshold_configuration,optional,no_refresh"`
	UpdateScheduledCharges                     *[]*V2ContractUpdateScheduledChargesModel                  `tfsdk:"update_scheduled_charges" json:"update_scheduled_charges,optional,no_refresh"`
	UpdateSpendThresholdConfiguration          *V2ContractUpdateSpendThresholdConfigurationModel          `tfsdk:"update_spend_threshold_configuration" json:"update_spend_threshold_configuration,optional,no_refresh"`
	Data                                       customfield.NestedObject[V2ContractDataModel]              `tfsdk:"data" json:"data,computed"`
}

func (m V2ContractModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V2ContractModel) MarshalJSONForUpdate(state V2ContractModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V2ContractAddCommitsModel struct {
	ProductID             types.String                                `tfsdk:"product_id" json:"product_id,required"`
	Type                  types.String                                `tfsdk:"type" json:"type,required"`
	AccessSchedule        *V2ContractAddCommitsAccessScheduleModel    `tfsdk:"access_schedule" json:"access_schedule,optional"`
	Amount                types.Float64                               `tfsdk:"amount" json:"amount,optional"`
	ApplicableProductIDs  *[]types.String                             `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                             `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields          *map[string]types.String                    `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description           types.String                                `tfsdk:"description" json:"description,optional"`
	InvoiceSchedule       *V2ContractAddCommitsInvoiceScheduleModel   `tfsdk:"invoice_schedule" json:"invoice_schedule,optional"`
	Name                  types.String                                `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	PaymentGateConfig     *V2ContractAddCommitsPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,optional"`
	Priority              types.Float64                               `tfsdk:"priority" json:"priority,optional"`
	RateType              types.String                                `tfsdk:"rate_type" json:"rate_type,optional"`
	RolloverFraction      types.Float64                               `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                                `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V2ContractAddCommitsAccessScheduleModel struct {
	ScheduleItems *[]*V2ContractAddCommitsAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                             `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V2ContractAddCommitsAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V2ContractAddCommitsInvoiceScheduleModel struct {
	CreditTypeID      types.String                                               `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V2ContractAddCommitsInvoiceScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V2ContractAddCommitsInvoiceScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V2ContractAddCommitsInvoiceScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddCommitsInvoiceScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddCommitsPaymentGateConfigModel struct {
	PaymentGateType types.String                                            `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V2ContractAddCommitsPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                            `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V2ContractAddCommitsPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V2ContractAddCreditsModel struct {
	AccessSchedule        *V2ContractAddCreditsAccessScheduleModel `tfsdk:"access_schedule" json:"access_schedule,required"`
	ProductID             types.String                             `tfsdk:"product_id" json:"product_id,required"`
	ApplicableProductIDs  *[]types.String                          `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                          `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	CustomFields          *map[string]types.String                 `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description           types.String                             `tfsdk:"description" json:"description,optional"`
	Name                  types.String                             `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Priority              types.Float64                            `tfsdk:"priority" json:"priority,optional"`
	RateType              types.String                             `tfsdk:"rate_type" json:"rate_type,optional"`
}

type V2ContractAddCreditsAccessScheduleModel struct {
	ScheduleItems *[]*V2ContractAddCreditsAccessScheduleScheduleItemsModel `tfsdk:"schedule_items" json:"schedule_items,required"`
	CreditTypeID  types.String                                             `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
}

type V2ContractAddCreditsAccessScheduleScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V2ContractAddDiscountsModel struct {
	ProductID            types.String                         `tfsdk:"product_id" json:"product_id,required"`
	Schedule             *V2ContractAddDiscountsScheduleModel `tfsdk:"schedule" json:"schedule,required"`
	CustomFields         *map[string]types.String             `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Name                 types.String                         `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID types.String                         `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V2ContractAddDiscountsScheduleModel struct {
	CreditTypeID      types.String                                          `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V2ContractAddDiscountsScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V2ContractAddDiscountsScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V2ContractAddDiscountsScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddDiscountsScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddOverridesModel struct {
	StartingAt            timetypes.RFC3339                                 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductTags *[]types.String                                   `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	EndingBefore          timetypes.RFC3339                                 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	Entitled              types.Bool                                        `tfsdk:"entitled" json:"entitled,optional"`
	IsCommitSpecific      types.Bool                                        `tfsdk:"is_commit_specific" json:"is_commit_specific,optional"`
	Multiplier            types.Float64                                     `tfsdk:"multiplier" json:"multiplier,optional"`
	OverrideSpecifiers    *[]*V2ContractAddOverridesOverrideSpecifiersModel `tfsdk:"override_specifiers" json:"override_specifiers,optional"`
	OverwriteRate         *V2ContractAddOverridesOverwriteRateModel         `tfsdk:"overwrite_rate" json:"overwrite_rate,optional"`
	Priority              types.Float64                                     `tfsdk:"priority" json:"priority,optional"`
	ProductID             types.String                                      `tfsdk:"product_id" json:"product_id,optional"`
	Target                types.String                                      `tfsdk:"target" json:"target,optional"`
	Tiers                 *[]*V2ContractAddOverridesTiersModel              `tfsdk:"tiers" json:"tiers,optional"`
	Type                  types.String                                      `tfsdk:"type" json:"type,optional"`
}

type V2ContractAddOverridesOverrideSpecifiersModel struct {
	CommitIDs               *[]types.String          `tfsdk:"commit_ids" json:"commit_ids,optional"`
	PresentationGroupValues *map[string]types.String `tfsdk:"presentation_group_values" json:"presentation_group_values,optional"`
	PricingGroupValues      *map[string]types.String `tfsdk:"pricing_group_values" json:"pricing_group_values,optional"`
	ProductID               types.String             `tfsdk:"product_id" json:"product_id,optional"`
	ProductTags             *[]types.String          `tfsdk:"product_tags" json:"product_tags,optional"`
	RecurringCommitIDs      *[]types.String          `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,optional"`
	RecurringCreditIDs      *[]types.String          `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,optional"`
}

type V2ContractAddOverridesOverwriteRateModel struct {
	RateType     types.String                                      `tfsdk:"rate_type" json:"rate_type,required"`
	CreditTypeID types.String                                      `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	CustomRate   *map[string]jsontypes.Normalized                  `tfsdk:"custom_rate" json:"custom_rate,optional"`
	IsProrated   types.Bool                                        `tfsdk:"is_prorated" json:"is_prorated,optional"`
	Price        types.Float64                                     `tfsdk:"price" json:"price,optional"`
	Quantity     types.Float64                                     `tfsdk:"quantity" json:"quantity,optional"`
	Tiers        *[]*V2ContractAddOverridesOverwriteRateTiersModel `tfsdk:"tiers" json:"tiers,optional"`
}

type V2ContractAddOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,required"`
	Size  types.Float64 `tfsdk:"size" json:"size,optional"`
}

type V2ContractAddOverridesTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,required"`
	Size       types.Float64 `tfsdk:"size" json:"size,optional"`
}

type V2ContractAddPrepaidBalanceThresholdConfigurationModel struct {
	Commit            *V2ContractAddPrepaidBalanceThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,required"`
	IsEnabled         types.Bool                                                               `tfsdk:"is_enabled" json:"is_enabled,required"`
	PaymentGateConfig *V2ContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,required"`
	RechargeToAmount  types.Float64                                                            `tfsdk:"recharge_to_amount" json:"recharge_to_amount,required"`
	ThresholdAmount   types.Float64                                                            `tfsdk:"threshold_amount" json:"threshold_amount,required"`
}

type V2ContractAddPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String    `tfsdk:"product_id" json:"product_id,required"`
	ApplicableProductIDs  *[]types.String `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String    `tfsdk:"description" json:"description,optional"`
	Name                  types.String    `tfsdk:"name" json:"name,optional"`
}

type V2ContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                         `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V2ContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                                         `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V2ContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V2ContractAddProfessionalServicesModel struct {
	MaxAmount            types.Float64            `tfsdk:"max_amount" json:"max_amount,required"`
	ProductID            types.String             `tfsdk:"product_id" json:"product_id,required"`
	Quantity             types.Float64            `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice            types.Float64            `tfsdk:"unit_price" json:"unit_price,required"`
	CustomFields         *map[string]types.String `tfsdk:"custom_fields" json:"custom_fields,optional"`
	Description          types.String             `tfsdk:"description" json:"description,optional"`
	NetsuiteSalesOrderID types.String             `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V2ContractAddRecurringCommitsModel struct {
	AccessAmount          *V2ContractAddRecurringCommitsAccessAmountModel   `tfsdk:"access_amount" json:"access_amount,required"`
	CommitDuration        *V2ContractAddRecurringCommitsCommitDurationModel `tfsdk:"commit_duration" json:"commit_duration,required"`
	Priority              types.Float64                                     `tfsdk:"priority" json:"priority,required"`
	ProductID             types.String                                      `tfsdk:"product_id" json:"product_id,required"`
	StartingAt            timetypes.RFC3339                                 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  *[]types.String                                   `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                                   `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String                                      `tfsdk:"description" json:"description,optional"`
	EndingBefore          timetypes.RFC3339                                 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	InvoiceAmount         *V2ContractAddRecurringCommitsInvoiceAmountModel  `tfsdk:"invoice_amount" json:"invoice_amount,optional"`
	Name                  types.String                                      `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                                      `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Proration             types.String                                      `tfsdk:"proration" json:"proration,optional"`
	RateType              types.String                                      `tfsdk:"rate_type" json:"rate_type,optional"`
	RecurrenceFrequency   types.String                                      `tfsdk:"recurrence_frequency" json:"recurrence_frequency,optional"`
	RolloverFraction      types.Float64                                     `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                                      `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V2ContractAddRecurringCommitsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V2ContractAddRecurringCommitsCommitDurationModel struct {
	Unit  types.String  `tfsdk:"unit" json:"unit,required"`
	Value types.Float64 `tfsdk:"value" json:"value,required"`
}

type V2ContractAddRecurringCommitsInvoiceAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V2ContractAddRecurringCreditsModel struct {
	AccessAmount          *V2ContractAddRecurringCreditsAccessAmountModel   `tfsdk:"access_amount" json:"access_amount,required"`
	CommitDuration        *V2ContractAddRecurringCreditsCommitDurationModel `tfsdk:"commit_duration" json:"commit_duration,required"`
	Priority              types.Float64                                     `tfsdk:"priority" json:"priority,required"`
	ProductID             types.String                                      `tfsdk:"product_id" json:"product_id,required"`
	StartingAt            timetypes.RFC3339                                 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  *[]types.String                                   `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                                   `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String                                      `tfsdk:"description" json:"description,optional"`
	EndingBefore          timetypes.RFC3339                                 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	Name                  types.String                                      `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID  types.String                                      `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	Proration             types.String                                      `tfsdk:"proration" json:"proration,optional"`
	RateType              types.String                                      `tfsdk:"rate_type" json:"rate_type,optional"`
	RecurrenceFrequency   types.String                                      `tfsdk:"recurrence_frequency" json:"recurrence_frequency,optional"`
	RolloverFraction      types.Float64                                     `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
	TemporaryID           types.String                                      `tfsdk:"temporary_id" json:"temporary_id,optional"`
}

type V2ContractAddRecurringCreditsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,required"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,required"`
}

type V2ContractAddRecurringCreditsCommitDurationModel struct {
	Unit  types.String  `tfsdk:"unit" json:"unit,required"`
	Value types.Float64 `tfsdk:"value" json:"value,required"`
}

type V2ContractAddResellerRoyaltiesModel struct {
	ResellerType          types.String                                   `tfsdk:"reseller_type" json:"reseller_type,required"`
	ApplicableProductIDs  *[]types.String                                `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                                `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	AwsOptions            *V2ContractAddResellerRoyaltiesAwsOptionsModel `tfsdk:"aws_options" json:"aws_options,optional"`
	EndingBefore          timetypes.RFC3339                              `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	Fraction              types.Float64                                  `tfsdk:"fraction" json:"fraction,optional"`
	GcpOptions            *V2ContractAddResellerRoyaltiesGcpOptionsModel `tfsdk:"gcp_options" json:"gcp_options,optional"`
	NetsuiteResellerID    types.String                                   `tfsdk:"netsuite_reseller_id" json:"netsuite_reseller_id,optional"`
	ResellerContractValue types.Float64                                  `tfsdk:"reseller_contract_value" json:"reseller_contract_value,optional"`
	StartingAt            timetypes.RFC3339                              `tfsdk:"starting_at" json:"starting_at,optional" format:"date-time"`
}

type V2ContractAddResellerRoyaltiesAwsOptionsModel struct {
	AwsAccountNumber    types.String `tfsdk:"aws_account_number" json:"aws_account_number,optional"`
	AwsOfferID          types.String `tfsdk:"aws_offer_id" json:"aws_offer_id,optional"`
	AwsPayerReferenceID types.String `tfsdk:"aws_payer_reference_id" json:"aws_payer_reference_id,optional"`
}

type V2ContractAddResellerRoyaltiesGcpOptionsModel struct {
	GcpAccountID types.String `tfsdk:"gcp_account_id" json:"gcp_account_id,optional"`
	GcpOfferID   types.String `tfsdk:"gcp_offer_id" json:"gcp_offer_id,optional"`
}

type V2ContractAddScheduledChargesModel struct {
	ProductID            types.String                                `tfsdk:"product_id" json:"product_id,required"`
	Schedule             *V2ContractAddScheduledChargesScheduleModel `tfsdk:"schedule" json:"schedule,required"`
	Name                 types.String                                `tfsdk:"name" json:"name,optional"`
	NetsuiteSalesOrderID types.String                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V2ContractAddScheduledChargesScheduleModel struct {
	CreditTypeID      types.String                                                 `tfsdk:"credit_type_id" json:"credit_type_id,optional"`
	RecurringSchedule *V2ContractAddScheduledChargesScheduleRecurringScheduleModel `tfsdk:"recurring_schedule" json:"recurring_schedule,optional"`
	ScheduleItems     *[]*V2ContractAddScheduledChargesScheduleScheduleItemsModel  `tfsdk:"schedule_items" json:"schedule_items,optional"`
}

type V2ContractAddScheduledChargesScheduleRecurringScheduleModel struct {
	AmountDistribution types.String      `tfsdk:"amount_distribution" json:"amount_distribution,required"`
	EndingBefore       timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	Frequency          types.String      `tfsdk:"frequency" json:"frequency,required"`
	StartingAt         timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
	Amount             types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity           types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice          types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddScheduledChargesScheduleScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractAddSpendThresholdConfigurationModel struct {
	Commit            *V2ContractAddSpendThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,required"`
	IsEnabled         types.Bool                                                      `tfsdk:"is_enabled" json:"is_enabled,required"`
	PaymentGateConfig *V2ContractAddSpendThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,required"`
	ThresholdAmount   types.Float64                                                   `tfsdk:"threshold_amount" json:"threshold_amount,required"`
}

type V2ContractAddSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
}

type V2ContractAddSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V2ContractAddSpendThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                                `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V2ContractAddSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V2ContractArchiveCommitsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractArchiveCreditsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractArchiveScheduledChargesModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractRemoveOverridesModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractUpdateCommitsModel struct {
	CommitID              types.String                                 `tfsdk:"commit_id" json:"commit_id,required"`
	AccessSchedule        *V2ContractUpdateCommitsAccessScheduleModel  `tfsdk:"access_schedule" json:"access_schedule,optional"`
	ApplicableProductIDs  *[]types.String                              `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                              `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	InvoiceSchedule       *V2ContractUpdateCommitsInvoiceScheduleModel `tfsdk:"invoice_schedule" json:"invoice_schedule,optional"`
	NetsuiteSalesOrderID  types.String                                 `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	ProductID             types.String                                 `tfsdk:"product_id" json:"product_id,optional"`
	RolloverFraction      types.Float64                                `tfsdk:"rollover_fraction" json:"rollover_fraction,optional"`
}

type V2ContractUpdateCommitsAccessScheduleModel struct {
	AddScheduleItems    *[]*V2ContractUpdateCommitsAccessScheduleAddScheduleItemsModel    `tfsdk:"add_schedule_items" json:"add_schedule_items,optional"`
	RemoveScheduleItems *[]*V2ContractUpdateCommitsAccessScheduleRemoveScheduleItemsModel `tfsdk:"remove_schedule_items" json:"remove_schedule_items,optional"`
	UpdateScheduleItems *[]*V2ContractUpdateCommitsAccessScheduleUpdateScheduleItemsModel `tfsdk:"update_schedule_items" json:"update_schedule_items,optional"`
}

type V2ContractUpdateCommitsAccessScheduleAddScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V2ContractUpdateCommitsAccessScheduleRemoveScheduleItemsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractUpdateCommitsAccessScheduleUpdateScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,required"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,optional"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,optional" format:"date-time"`
}

type V2ContractUpdateCommitsInvoiceScheduleModel struct {
	AddScheduleItems    *[]*V2ContractUpdateCommitsInvoiceScheduleAddScheduleItemsModel    `tfsdk:"add_schedule_items" json:"add_schedule_items,optional"`
	RemoveScheduleItems *[]*V2ContractUpdateCommitsInvoiceScheduleRemoveScheduleItemsModel `tfsdk:"remove_schedule_items" json:"remove_schedule_items,optional"`
	UpdateScheduleItems *[]*V2ContractUpdateCommitsInvoiceScheduleUpdateScheduleItemsModel `tfsdk:"update_schedule_items" json:"update_schedule_items,optional"`
}

type V2ContractUpdateCommitsInvoiceScheduleAddScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractUpdateCommitsInvoiceScheduleRemoveScheduleItemsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractUpdateCommitsInvoiceScheduleUpdateScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,required"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,optional" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractUpdateCreditsModel struct {
	CreditID              types.String                                `tfsdk:"credit_id" json:"credit_id,required"`
	AccessSchedule        *V2ContractUpdateCreditsAccessScheduleModel `tfsdk:"access_schedule" json:"access_schedule,optional"`
	ApplicableProductIDs  *[]types.String                             `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String                             `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	NetsuiteSalesOrderID  types.String                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
	ProductID             types.String                                `tfsdk:"product_id" json:"product_id,optional"`
}

type V2ContractUpdateCreditsAccessScheduleModel struct {
	AddScheduleItems    *[]*V2ContractUpdateCreditsAccessScheduleAddScheduleItemsModel    `tfsdk:"add_schedule_items" json:"add_schedule_items,optional"`
	RemoveScheduleItems *[]*V2ContractUpdateCreditsAccessScheduleRemoveScheduleItemsModel `tfsdk:"remove_schedule_items" json:"remove_schedule_items,optional"`
	UpdateScheduleItems *[]*V2ContractUpdateCreditsAccessScheduleUpdateScheduleItemsModel `tfsdk:"update_schedule_items" json:"update_schedule_items,optional"`
}

type V2ContractUpdateCreditsAccessScheduleAddScheduleItemsModel struct {
	Amount       types.Float64     `tfsdk:"amount" json:"amount,required"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,required" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,required" format:"date-time"`
}

type V2ContractUpdateCreditsAccessScheduleRemoveScheduleItemsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractUpdateCreditsAccessScheduleUpdateScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,required"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,optional"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,optional" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,optional" format:"date-time"`
}

type V2ContractUpdatePrepaidBalanceThresholdConfigurationModel struct {
	Commit            *V2ContractUpdatePrepaidBalanceThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,optional"`
	IsEnabled         types.Bool                                                                  `tfsdk:"is_enabled" json:"is_enabled,optional"`
	PaymentGateConfig *V2ContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,optional"`
	RechargeToAmount  types.Float64                                                               `tfsdk:"recharge_to_amount" json:"recharge_to_amount,optional"`
	ThresholdAmount   types.Float64                                                               `tfsdk:"threshold_amount" json:"threshold_amount,optional"`
}

type V2ContractUpdatePrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String    `tfsdk:"product_id" json:"product_id,required"`
	ApplicableProductIDs  *[]types.String `tfsdk:"applicable_product_ids" json:"applicable_product_ids,optional"`
	ApplicableProductTags *[]types.String `tfsdk:"applicable_product_tags" json:"applicable_product_tags,optional"`
	Description           types.String    `tfsdk:"description" json:"description,optional"`
	Name                  types.String    `tfsdk:"name" json:"name,optional"`
}

type V2ContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                            `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V2ContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                                            `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V2ContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V2ContractUpdateScheduledChargesModel struct {
	ScheduledChargeID    types.String                                          `tfsdk:"scheduled_charge_id" json:"scheduled_charge_id,required"`
	InvoiceSchedule      *V2ContractUpdateScheduledChargesInvoiceScheduleModel `tfsdk:"invoice_schedule" json:"invoice_schedule,optional"`
	NetsuiteSalesOrderID types.String                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,optional"`
}

type V2ContractUpdateScheduledChargesInvoiceScheduleModel struct {
	AddScheduleItems    *[]*V2ContractUpdateScheduledChargesInvoiceScheduleAddScheduleItemsModel    `tfsdk:"add_schedule_items" json:"add_schedule_items,optional"`
	RemoveScheduleItems *[]*V2ContractUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemsModel `tfsdk:"remove_schedule_items" json:"remove_schedule_items,optional"`
	UpdateScheduleItems *[]*V2ContractUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemsModel `tfsdk:"update_schedule_items" json:"update_schedule_items,optional"`
}

type V2ContractUpdateScheduledChargesInvoiceScheduleAddScheduleItemsModel struct {
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,required" format:"date-time"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractUpdateScheduledChargesInvoiceScheduleRemoveScheduleItemsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type V2ContractUpdateScheduledChargesInvoiceScheduleUpdateScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,required"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,optional"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,optional"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,optional" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,optional"`
}

type V2ContractUpdateSpendThresholdConfigurationModel struct {
	Commit            *V2ContractUpdateSpendThresholdConfigurationCommitModel            `tfsdk:"commit" json:"commit,optional"`
	IsEnabled         types.Bool                                                         `tfsdk:"is_enabled" json:"is_enabled,optional"`
	PaymentGateConfig *V2ContractUpdateSpendThresholdConfigurationPaymentGateConfigModel `tfsdk:"payment_gate_config" json:"payment_gate_config,optional"`
	ThresholdAmount   types.Float64                                                      `tfsdk:"threshold_amount" json:"threshold_amount,optional"`
}

type V2ContractUpdateSpendThresholdConfigurationCommitModel struct {
	Description types.String `tfsdk:"description" json:"description,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
	ProductID   types.String `tfsdk:"product_id" json:"product_id,optional"`
}

type V2ContractUpdateSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                   `tfsdk:"payment_gate_type" json:"payment_gate_type,required"`
	StripeConfig    *V2ContractUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigModel `tfsdk:"stripe_config" json:"stripe_config,optional"`
	TaxType         types.String                                                                   `tfsdk:"tax_type" json:"tax_type,optional"`
}

type V2ContractUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,required"`
}

type V2ContractDataModel struct {
	ID                                   types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Commits                              customfield.NestedObjectList[V2ContractDataCommitsModel]                          `tfsdk:"commits" json:"commits,computed"`
	CreatedAt                            timetypes.RFC3339                                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                            types.String                                                                      `tfsdk:"created_by" json:"created_by,computed"`
	CustomerID                           types.String                                                                      `tfsdk:"customer_id" json:"customer_id,computed"`
	Overrides                            customfield.NestedObjectList[V2ContractDataOverridesModel]                        `tfsdk:"overrides" json:"overrides,computed"`
	ScheduledCharges                     customfield.NestedObjectList[V2ContractDataScheduledChargesModel]                 `tfsdk:"scheduled_charges" json:"scheduled_charges,computed"`
	StartingAt                           timetypes.RFC3339                                                                 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	Transitions                          customfield.NestedObjectList[V2ContractDataTransitionsModel]                      `tfsdk:"transitions" json:"transitions,computed"`
	UsageFilter                          customfield.NestedObjectList[V2ContractDataUsageFilterModel]                      `tfsdk:"usage_filter" json:"usage_filter,computed"`
	UsageStatementSchedule               customfield.NestedObject[V2ContractDataUsageStatementScheduleModel]               `tfsdk:"usage_statement_schedule" json:"usage_statement_schedule,computed"`
	ArchivedAt                           timetypes.RFC3339                                                                 `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Credits                              customfield.NestedObjectList[V2ContractDataCreditsModel]                          `tfsdk:"credits" json:"credits,computed"`
	CustomFields                         customfield.Map[types.String]                                                     `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerBillingProviderConfiguration customfield.NestedObject[V2ContractDataCustomerBillingProviderConfigurationModel] `tfsdk:"customer_billing_provider_configuration" json:"customer_billing_provider_configuration,computed"`
	Discounts                            customfield.NestedObjectList[V2ContractDataDiscountsModel]                        `tfsdk:"discounts" json:"discounts,computed"`
	EndingBefore                         timetypes.RFC3339                                                                 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	MultiplierOverridePrioritization     types.String                                                                      `tfsdk:"multiplier_override_prioritization" json:"multiplier_override_prioritization,computed"`
	Name                                 types.String                                                                      `tfsdk:"name" json:"name,computed"`
	NetPaymentTermsDays                  types.Float64                                                                     `tfsdk:"net_payment_terms_days" json:"net_payment_terms_days,computed"`
	NetsuiteSalesOrderID                 types.String                                                                      `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	PrepaidBalanceThresholdConfiguration customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationModel] `tfsdk:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,computed"`
	ProfessionalServices                 customfield.NestedObjectList[V2ContractDataProfessionalServicesModel]             `tfsdk:"professional_services" json:"professional_services,computed"`
	RateCardID                           types.String                                                                      `tfsdk:"rate_card_id" json:"rate_card_id,computed"`
	RecurringCommits                     customfield.NestedObjectList[V2ContractDataRecurringCommitsModel]                 `tfsdk:"recurring_commits" json:"recurring_commits,computed"`
	RecurringCredits                     customfield.NestedObjectList[V2ContractDataRecurringCreditsModel]                 `tfsdk:"recurring_credits" json:"recurring_credits,computed"`
	ResellerRoyalties                    customfield.NestedObjectList[V2ContractDataResellerRoyaltiesModel]                `tfsdk:"reseller_royalties" json:"reseller_royalties,computed"`
	SalesforceOpportunityID              types.String                                                                      `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
	ScheduledChargesOnUsageInvoices      types.String                                                                      `tfsdk:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,computed"`
	SpendThresholdConfiguration          customfield.NestedObject[V2ContractDataSpendThresholdConfigurationModel]          `tfsdk:"spend_threshold_configuration" json:"spend_threshold_configuration,computed"`
	TotalContractValue                   types.Float64                                                                     `tfsdk:"total_contract_value" json:"total_contract_value,computed"`
	UniquenessKey                        types.String                                                                      `tfsdk:"uniqueness_key" json:"uniqueness_key,computed"`
}

type V2ContractDataCommitsModel struct {
	ID                      types.String                                                        `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V2ContractDataCommitsProductModel]         `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                        `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V2ContractDataCommitsAccessScheduleModel]  `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                      `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                      `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                      `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	ArchivedAt              timetypes.RFC3339                                                   `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	Balance                 types.Float64                                                       `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V2ContractDataCommitsContractModel]        `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                       `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                        `tfsdk:"description" json:"description,computed"`
	InvoiceContract         customfield.NestedObject[V2ContractDataCommitsInvoiceContractModel] `tfsdk:"invoice_contract" json:"invoice_contract,computed"`
	InvoiceSchedule         customfield.NestedObject[V2ContractDataCommitsInvoiceScheduleModel] `tfsdk:"invoice_schedule" json:"invoice_schedule,computed"`
	Ledger                  customfield.NestedObjectList[V2ContractDataCommitsLedgerModel]      `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                        `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                        `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                       `tfsdk:"priority" json:"priority,computed"`
	RateType                types.String                                                        `tfsdk:"rate_type" json:"rate_type,computed"`
	RolledOverFrom          customfield.NestedObject[V2ContractDataCommitsRolledOverFromModel]  `tfsdk:"rolled_over_from" json:"rolled_over_from,computed"`
	RolloverFraction        types.Float64                                                       `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
	SalesforceOpportunityID types.String                                                        `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V2ContractDataCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V2ContractDataCommitsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V2ContractDataCommitsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V2ContractDataCommitsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V2ContractDataCommitsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCommitsInvoiceContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCommitsInvoiceScheduleModel struct {
	CreditType    customfield.NestedObject[V2ContractDataCommitsInvoiceScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataCommitsInvoiceScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataCommitsInvoiceScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCommitsInvoiceScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataCommitsLedgerModel struct {
	Amount        types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID     types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type          types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID     types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	NewContractID types.String      `tfsdk:"new_contract_id" json:"new_contract_id,computed"`
	Reason        types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V2ContractDataCommitsRolledOverFromModel struct {
	CommitID   types.String `tfsdk:"commit_id" json:"commit_id,computed"`
	ContractID types.String `tfsdk:"contract_id" json:"contract_id,computed"`
}

type V2ContractDataOverridesModel struct {
	ID                    types.String                                                                 `tfsdk:"id" json:"id,computed"`
	StartingAt            timetypes.RFC3339                                                            `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductTags customfield.List[types.String]                                               `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	EndingBefore          timetypes.RFC3339                                                            `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Entitled              types.Bool                                                                   `tfsdk:"entitled" json:"entitled,computed"`
	IsCommitSpecific      types.Bool                                                                   `tfsdk:"is_commit_specific" json:"is_commit_specific,computed"`
	Multiplier            types.Float64                                                                `tfsdk:"multiplier" json:"multiplier,computed"`
	OverrideSpecifiers    customfield.NestedObjectList[V2ContractDataOverridesOverrideSpecifiersModel] `tfsdk:"override_specifiers" json:"override_specifiers,computed"`
	OverrideTiers         customfield.NestedObjectList[V2ContractDataOverridesOverrideTiersModel]      `tfsdk:"override_tiers" json:"override_tiers,computed"`
	OverwriteRate         customfield.NestedObject[V2ContractDataOverridesOverwriteRateModel]          `tfsdk:"overwrite_rate" json:"overwrite_rate,computed"`
	Priority              types.Float64                                                                `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataOverridesProductModel]                `tfsdk:"product" json:"product,computed"`
	Target                types.String                                                                 `tfsdk:"target" json:"target,computed"`
	Type                  types.String                                                                 `tfsdk:"type" json:"type,computed"`
}

type V2ContractDataOverridesOverrideSpecifiersModel struct {
	CommitIDs               customfield.List[types.String] `tfsdk:"commit_ids" json:"commit_ids,computed"`
	PresentationGroupValues customfield.Map[types.String]  `tfsdk:"presentation_group_values" json:"presentation_group_values,computed"`
	PricingGroupValues      customfield.Map[types.String]  `tfsdk:"pricing_group_values" json:"pricing_group_values,computed"`
	ProductID               types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ProductTags             customfield.List[types.String] `tfsdk:"product_tags" json:"product_tags,computed"`
	RecurringCommitIDs      customfield.List[types.String] `tfsdk:"recurring_commit_ids" json:"recurring_commit_ids,computed"`
	RecurringCreditIDs      customfield.List[types.String] `tfsdk:"recurring_credit_ids" json:"recurring_credit_ids,computed"`
}

type V2ContractDataOverridesOverrideTiersModel struct {
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
	Size       types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V2ContractDataOverridesOverwriteRateModel struct {
	RateType   types.String                                                                  `tfsdk:"rate_type" json:"rate_type,computed"`
	CreditType customfield.NestedObject[V2ContractDataOverridesOverwriteRateCreditTypeModel] `tfsdk:"credit_type" json:"credit_type,computed"`
	CustomRate customfield.Map[jsontypes.Normalized]                                         `tfsdk:"custom_rate" json:"custom_rate,computed"`
	IsProrated types.Bool                                                                    `tfsdk:"is_prorated" json:"is_prorated,computed"`
	Price      types.Float64                                                                 `tfsdk:"price" json:"price,computed"`
	Quantity   types.Float64                                                                 `tfsdk:"quantity" json:"quantity,computed"`
	Tiers      customfield.NestedObjectList[V2ContractDataOverridesOverwriteRateTiersModel]  `tfsdk:"tiers" json:"tiers,computed"`
}

type V2ContractDataOverridesOverwriteRateCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataOverridesOverwriteRateTiersModel struct {
	Price types.Float64 `tfsdk:"price" json:"price,computed"`
	Size  types.Float64 `tfsdk:"size" json:"size,computed"`
}

type V2ContractDataOverridesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V2ContractDataScheduledChargesProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V2ContractDataScheduledChargesScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	ArchivedAt           timetypes.RFC3339                                                     `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CustomFields         customfield.Map[types.String]                                         `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                          `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataScheduledChargesProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesScheduleModel struct {
	CreditType    customfield.NestedObject[V2ContractDataScheduledChargesScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataScheduledChargesScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataScheduledChargesScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataScheduledChargesScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataTransitionsModel struct {
	FromContractID types.String `tfsdk:"from_contract_id" json:"from_contract_id,computed"`
	ToContractID   types.String `tfsdk:"to_contract_id" json:"to_contract_id,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
}

type V2ContractDataUsageFilterModel struct {
	GroupKey     types.String                   `tfsdk:"group_key" json:"group_key,computed"`
	GroupValues  customfield.List[types.String] `tfsdk:"group_values" json:"group_values,computed"`
	StartingAt   timetypes.RFC3339              `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	EndingBefore timetypes.RFC3339              `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
}

type V2ContractDataUsageStatementScheduleModel struct {
	BillingAnchorDate timetypes.RFC3339 `tfsdk:"billing_anchor_date" json:"billing_anchor_date,computed" format:"date-time"`
	Frequency         types.String      `tfsdk:"frequency" json:"frequency,computed"`
}

type V2ContractDataCreditsModel struct {
	ID                      types.String                                                       `tfsdk:"id" json:"id,computed"`
	Product                 customfield.NestedObject[V2ContractDataCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	Type                    types.String                                                       `tfsdk:"type" json:"type,computed"`
	AccessSchedule          customfield.NestedObject[V2ContractDataCreditsAccessScheduleModel] `tfsdk:"access_schedule" json:"access_schedule,computed"`
	ApplicableContractIDs   customfield.List[types.String]                                     `tfsdk:"applicable_contract_ids" json:"applicable_contract_ids,computed"`
	ApplicableProductIDs    customfield.List[types.String]                                     `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags   customfield.List[types.String]                                     `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Balance                 types.Float64                                                      `tfsdk:"balance" json:"balance,computed"`
	Contract                customfield.NestedObject[V2ContractDataCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	CustomFields            customfield.Map[types.String]                                      `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description             types.String                                                       `tfsdk:"description" json:"description,computed"`
	Ledger                  customfield.NestedObjectList[V2ContractDataCreditsLedgerModel]     `tfsdk:"ledger" json:"ledger,computed"`
	Name                    types.String                                                       `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID    types.String                                                       `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Priority                types.Float64                                                      `tfsdk:"priority" json:"priority,computed"`
	SalesforceOpportunityID types.String                                                       `tfsdk:"salesforce_opportunity_id" json:"salesforce_opportunity_id,computed"`
}

type V2ContractDataCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCreditsAccessScheduleModel struct {
	ScheduleItems customfield.NestedObjectList[V2ContractDataCreditsAccessScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
	CreditType    customfield.NestedObject[V2ContractDataCreditsAccessScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
}

type V2ContractDataCreditsAccessScheduleScheduleItemsModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	Amount       types.Float64     `tfsdk:"amount" json:"amount,computed"`
	EndingBefore timetypes.RFC3339 `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	StartingAt   timetypes.RFC3339 `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
}

type V2ContractDataCreditsAccessScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataCreditsLedgerModel struct {
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	SegmentID types.String      `tfsdk:"segment_id" json:"segment_id,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Reason    types.String      `tfsdk:"reason" json:"reason,computed"`
}

type V2ContractDataCustomerBillingProviderConfigurationModel struct {
	BillingProvider types.String `tfsdk:"billing_provider" json:"billing_provider,computed"`
	DeliveryMethod  types.String `tfsdk:"delivery_method" json:"delivery_method,computed"`
}

type V2ContractDataDiscountsModel struct {
	ID                   types.String                                                   `tfsdk:"id" json:"id,computed"`
	Product              customfield.NestedObject[V2ContractDataDiscountsProductModel]  `tfsdk:"product" json:"product,computed"`
	Schedule             customfield.NestedObject[V2ContractDataDiscountsScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	CustomFields         customfield.Map[types.String]                                  `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Name                 types.String                                                   `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID types.String                                                   `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataDiscountsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataDiscountsScheduleModel struct {
	CreditType    customfield.NestedObject[V2ContractDataDiscountsScheduleCreditTypeModel]        `tfsdk:"credit_type" json:"credit_type,computed"`
	ScheduleItems customfield.NestedObjectList[V2ContractDataDiscountsScheduleScheduleItemsModel] `tfsdk:"schedule_items" json:"schedule_items,computed"`
}

type V2ContractDataDiscountsScheduleCreditTypeModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataDiscountsScheduleScheduleItemsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Amount    types.Float64     `tfsdk:"amount" json:"amount,computed"`
	InvoiceID types.String      `tfsdk:"invoice_id" json:"invoice_id,computed"`
	Quantity  types.Float64     `tfsdk:"quantity" json:"quantity,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	UnitPrice types.Float64     `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                         `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	RechargeToAmount  types.Float64                                                                                      `tfsdk:"recharge_to_amount" json:"recharge_to_amount,computed"`
	ThresholdAmount   types.Float64                                                                                      `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationCommitModel struct {
	ProductID             types.String                   `tfsdk:"product_id" json:"product_id,computed"`
	ApplicableProductIDs  customfield.List[types.String] `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String] `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Description           types.String                   `tfsdk:"description" json:"description,computed"`
	Name                  types.String                   `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                                   `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                                   `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V2ContractDataPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}

type V2ContractDataProfessionalServicesModel struct {
	ID                   types.String                  `tfsdk:"id" json:"id,computed"`
	MaxAmount            types.Float64                 `tfsdk:"max_amount" json:"max_amount,computed"`
	ProductID            types.String                  `tfsdk:"product_id" json:"product_id,computed"`
	Quantity             types.Float64                 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice            types.Float64                 `tfsdk:"unit_price" json:"unit_price,computed"`
	CustomFields         customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
	Description          types.String                  `tfsdk:"description" json:"description,computed"`
	NetsuiteSalesOrderID types.String                  `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
}

type V2ContractDataRecurringCommitsModel struct {
	ID                    types.String                                                                `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V2ContractDataRecurringCommitsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V2ContractDataRecurringCommitsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                               `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataRecurringCommitsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                           `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                              `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                              `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V2ContractDataRecurringCommitsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                           `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	InvoiceAmount         customfield.NestedObject[V2ContractDataRecurringCommitsInvoiceAmountModel]  `tfsdk:"invoice_amount" json:"invoice_amount,computed"`
	Name                  types.String                                                                `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                               `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V2ContractDataRecurringCommitsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCommitsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V2ContractDataRecurringCommitsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataRecurringCommitsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataRecurringCommitsInvoiceAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCreditsModel struct {
	ID                    types.String                                                                `tfsdk:"id" json:"id,computed"`
	AccessAmount          customfield.NestedObject[V2ContractDataRecurringCreditsAccessAmountModel]   `tfsdk:"access_amount" json:"access_amount,computed"`
	CommitDuration        customfield.NestedObject[V2ContractDataRecurringCreditsCommitDurationModel] `tfsdk:"commit_duration" json:"commit_duration,computed"`
	Priority              types.Float64                                                               `tfsdk:"priority" json:"priority,computed"`
	Product               customfield.NestedObject[V2ContractDataRecurringCreditsProductModel]        `tfsdk:"product" json:"product,computed"`
	RateType              types.String                                                                `tfsdk:"rate_type" json:"rate_type,computed"`
	StartingAt            timetypes.RFC3339                                                           `tfsdk:"starting_at" json:"starting_at,computed" format:"date-time"`
	ApplicableProductIDs  customfield.List[types.String]                                              `tfsdk:"applicable_product_ids" json:"applicable_product_ids,computed"`
	ApplicableProductTags customfield.List[types.String]                                              `tfsdk:"applicable_product_tags" json:"applicable_product_tags,computed"`
	Contract              customfield.NestedObject[V2ContractDataRecurringCreditsContractModel]       `tfsdk:"contract" json:"contract,computed"`
	Description           types.String                                                                `tfsdk:"description" json:"description,computed"`
	EndingBefore          timetypes.RFC3339                                                           `tfsdk:"ending_before" json:"ending_before,computed" format:"date-time"`
	Name                  types.String                                                                `tfsdk:"name" json:"name,computed"`
	NetsuiteSalesOrderID  types.String                                                                `tfsdk:"netsuite_sales_order_id" json:"netsuite_sales_order_id,computed"`
	Proration             types.String                                                                `tfsdk:"proration" json:"proration,computed"`
	RecurrenceFrequency   types.String                                                                `tfsdk:"recurrence_frequency" json:"recurrence_frequency,computed"`
	RolloverFraction      types.Float64                                                               `tfsdk:"rollover_fraction" json:"rollover_fraction,computed"`
}

type V2ContractDataRecurringCreditsAccessAmountModel struct {
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"credit_type_id,computed"`
	Quantity     types.Float64 `tfsdk:"quantity" json:"quantity,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unit_price,computed"`
}

type V2ContractDataRecurringCreditsCommitDurationModel struct {
	Value types.Float64 `tfsdk:"value" json:"value,computed"`
	Unit  types.String  `tfsdk:"unit" json:"unit,computed"`
}

type V2ContractDataRecurringCreditsProductModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataRecurringCreditsContractModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type V2ContractDataResellerRoyaltiesModel struct {
	ResellerType types.String                                                               `tfsdk:"reseller_type" json:"reseller_type,computed"`
	Segments     customfield.NestedObjectList[V2ContractDataResellerRoyaltiesSegmentsModel] `tfsdk:"segments" json:"segments,computed"`
}

type V2ContractDataResellerRoyaltiesSegmentsModel struct {
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

type V2ContractDataSpendThresholdConfigurationModel struct {
	Commit            customfield.NestedObject[V2ContractDataSpendThresholdConfigurationCommitModel]            `tfsdk:"commit" json:"commit,computed"`
	IsEnabled         types.Bool                                                                                `tfsdk:"is_enabled" json:"is_enabled,computed"`
	PaymentGateConfig customfield.NestedObject[V2ContractDataSpendThresholdConfigurationPaymentGateConfigModel] `tfsdk:"payment_gate_config" json:"payment_gate_config,computed"`
	ThresholdAmount   types.Float64                                                                             `tfsdk:"threshold_amount" json:"threshold_amount,computed"`
}

type V2ContractDataSpendThresholdConfigurationCommitModel struct {
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type V2ContractDataSpendThresholdConfigurationPaymentGateConfigModel struct {
	PaymentGateType types.String                                                                                          `tfsdk:"payment_gate_type" json:"payment_gate_type,computed"`
	StripeConfig    customfield.NestedObject[V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel] `tfsdk:"stripe_config" json:"stripe_config,computed"`
	TaxType         types.String                                                                                          `tfsdk:"tax_type" json:"tax_type,computed"`
}

type V2ContractDataSpendThresholdConfigurationPaymentGateConfigStripeConfigModel struct {
	PaymentType types.String `tfsdk:"payment_type" json:"payment_type,computed"`
}
