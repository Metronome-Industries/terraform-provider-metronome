// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*V1CustomerBillingConfigResource)(nil)
var _ resource.ResourceWithModifyPlan = (*V1CustomerBillingConfigResource)(nil)

func NewResource() resource.Resource {
	return &V1CustomerBillingConfigResource{}
}

// V1CustomerBillingConfigResource defines the resource implementation.
type V1CustomerBillingConfigResource struct {
	client *metronome.Client
}

func (r *V1CustomerBillingConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_customer_billing_config"
}

func (r *V1CustomerBillingConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*metronome.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *metronome.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *V1CustomerBillingConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *V1CustomerBillingConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	err = r.client.V1.Customers.BillingConfig.New(
		ctx,
		metronome.V1CustomerBillingConfigNewParams{
			CustomerID:          metronome.F(data.CustomerID.ValueString()),
			BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigNewParamsBillingProviderType(data.BillingProviderType.ValueString())),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *V1CustomerBillingConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is not supported for this resource
}

func (r *V1CustomerBillingConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *V1CustomerBillingConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.V1.Customers.BillingConfig.Get(
		ctx,
		metronome.V1CustomerBillingConfigGetParams{
			CustomerID:          metronome.F(data.CustomerID.ValueString()),
			BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigGetParamsBillingProviderType(data.BillingProviderType.ValueString())),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *V1CustomerBillingConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *V1CustomerBillingConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.V1.Customers.BillingConfig.Delete(
		ctx,
		metronome.V1CustomerBillingConfigDeleteParams{
			CustomerID:          metronome.F(data.CustomerID.ValueString()),
			BillingProviderType: metronome.F(metronome.V1CustomerBillingConfigDeleteParamsBillingProviderType(data.BillingProviderType.ValueString())),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *V1CustomerBillingConfigResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
