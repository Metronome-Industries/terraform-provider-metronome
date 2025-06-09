// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CustomerBillingConfigModel struct {
	BillingProviderType       types.String                                               `tfsdk:"billing_provider_type" path:"billing_provider_type,required"`
	CustomerID                types.String                                               `tfsdk:"customer_id" path:"customer_id,required"`
	BillingProviderCustomerID types.String                                               `tfsdk:"billing_provider_customer_id" json:"billing_provider_customer_id,required,no_refresh"`
	AwsProductCode            types.String                                               `tfsdk:"aws_product_code" json:"aws_product_code,optional,no_refresh"`
	AwsRegion                 types.String                                               `tfsdk:"aws_region" json:"aws_region,optional,no_refresh"`
	StripeCollectionMethod    types.String                                               `tfsdk:"stripe_collection_method" json:"stripe_collection_method,optional,no_refresh"`
	Data                      customfield.NestedObject[V1CustomerBillingConfigDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1CustomerBillingConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1CustomerBillingConfigModel) MarshalJSONForUpdate(state V1CustomerBillingConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1CustomerBillingConfigDataModel struct {
	AwsExpirationDate         timetypes.RFC3339 `tfsdk:"aws_expiration_date" json:"aws_expiration_date,computed" format:"date-time"`
	AwsProductCode            types.String      `tfsdk:"aws_product_code" json:"aws_product_code,computed"`
	AwsRegion                 types.String      `tfsdk:"aws_region" json:"aws_region,computed"`
	AzureExpirationDate       timetypes.RFC3339 `tfsdk:"azure_expiration_date" json:"azure_expiration_date,computed" format:"date-time"`
	AzurePlanID               types.String      `tfsdk:"azure_plan_id" json:"azure_plan_id,computed"`
	AzureStartDate            timetypes.RFC3339 `tfsdk:"azure_start_date" json:"azure_start_date,computed" format:"date-time"`
	AzureSubscriptionStatus   types.String      `tfsdk:"azure_subscription_status" json:"azure_subscription_status,computed"`
	BillingProviderCustomerID types.String      `tfsdk:"billing_provider_customer_id" json:"billing_provider_customer_id,computed"`
	StripeCollectionMethod    types.String      `tfsdk:"stripe_collection_method" json:"stripe_collection_method,computed"`
}
