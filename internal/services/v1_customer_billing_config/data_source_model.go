// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomerBillingConfigDataSourceModel struct {
	BillingProviderType types.String                                                         `tfsdk:"billing_provider_type" path:"billing_provider_type,required"`
	CustomerID          types.String                                                         `tfsdk:"customer_id" path:"customer_id,required"`
	Data                customfield.NestedObject[V1CustomerBillingConfigDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1CustomerBillingConfigDataSourceModel) toReadParams(_ context.Context) (params metronome.V1CustomerBillingConfigGetParams, diags diag.Diagnostics) {
	params = metronome.V1CustomerBillingConfigGetParams{
		CustomerID:          metronome.F(m.CustomerID.ValueString()),
		BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigGetParamsBillingProviderType(m.BillingProviderType.ValueString())),
	}

	return
}

type V1CustomerBillingConfigDataDataSourceModel struct {
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
