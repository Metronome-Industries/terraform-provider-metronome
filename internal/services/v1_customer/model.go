// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomerModel struct {
	CustomerID                            types.String                                             `tfsdk:"customer_id" path:"customer_id,optional"`
	Name                                  types.String                                             `tfsdk:"name" json:"name,required,no_refresh"`
	ExternalID                            types.String                                             `tfsdk:"external_id" json:"external_id,optional,no_refresh"`
	CustomFields                          *map[string]types.String                                 `tfsdk:"custom_fields" json:"custom_fields,optional,no_refresh"`
	IngestAliases                         *[]types.String                                          `tfsdk:"ingest_aliases" json:"ingest_aliases,optional,no_refresh"`
	BillingConfig                         *V1CustomerBillingConfigModel                            `tfsdk:"billing_config" json:"billing_config,optional,no_refresh"`
	CustomerBillingProviderConfigurations *[]*V1CustomerCustomerBillingProviderConfigurationsModel `tfsdk:"customer_billing_provider_configurations" json:"customer_billing_provider_configurations,optional,no_refresh"`
	Data                                  customfield.NestedObject[V1CustomerDataModel]            `tfsdk:"data" json:"data,computed"`
}

func (m V1CustomerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CustomerModel) MarshalJSONForUpdate(state V1CustomerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CustomerBillingConfigModel struct {
	BillingProviderCustomerID types.String `tfsdk:"billing_provider_customer_id" json:"billing_provider_customer_id,required"`
	BillingProviderType       types.String `tfsdk:"billing_provider_type" json:"billing_provider_type,required"`
	AwsIsSubscriptionProduct  types.Bool   `tfsdk:"aws_is_subscription_product" json:"aws_is_subscription_product,optional"`
	AwsProductCode            types.String `tfsdk:"aws_product_code" json:"aws_product_code,optional"`
	AwsRegion                 types.String `tfsdk:"aws_region" json:"aws_region,optional"`
	StripeCollectionMethod    types.String `tfsdk:"stripe_collection_method" json:"stripe_collection_method,optional"`
}

type V1CustomerCustomerBillingProviderConfigurationsModel struct {
	BillingProvider  types.String                     `tfsdk:"billing_provider" json:"billing_provider,required"`
	Configuration    *map[string]jsontypes.Normalized `tfsdk:"configuration" json:"configuration,optional"`
	DeliveryMethod   types.String                     `tfsdk:"delivery_method" json:"delivery_method,optional"`
	DeliveryMethodID types.String                     `tfsdk:"delivery_method_id" json:"delivery_method_id,optional"`
}

type V1CustomerDataModel struct {
	ID                    types.String                                                       `tfsdk:"id" json:"id,computed"`
	CreatedAt             timetypes.RFC3339                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomFields          customfield.Map[types.String]                                      `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerConfig        customfield.NestedObject[V1CustomerDataCustomerConfigModel]        `tfsdk:"customer_config" json:"customer_config,computed"`
	ExternalID            types.String                                                       `tfsdk:"external_id" json:"external_id,computed"`
	IngestAliases         customfield.List[types.String]                                     `tfsdk:"ingest_aliases" json:"ingest_aliases,computed"`
	Name                  types.String                                                       `tfsdk:"name" json:"name,computed"`
	ArchivedAt            timetypes.RFC3339                                                  `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CurrentBillableStatus customfield.NestedObject[V1CustomerDataCurrentBillableStatusModel] `tfsdk:"current_billable_status" json:"current_billable_status,computed"`
}

type V1CustomerDataCustomerConfigModel struct {
	SalesforceAccountID types.String `tfsdk:"salesforce_account_id" json:"salesforce_account_id,computed"`
}

type V1CustomerDataCurrentBillableStatusModel struct {
	Value       types.String      `tfsdk:"value" json:"value,computed"`
	EffectiveAt timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
}
