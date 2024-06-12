// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package billable_metric

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stainless-sdks/metronome-terraform/internal/apijson"
	"github.com/stainless-sdks/metronome-terraform/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BillableMetricResource{}

func NewResource() resource.Resource {
	return &BillableMetricResource{}
}

// BillableMetricResource defines the resource implementation.
type BillableMetricResource struct {
	client *metronome.Client
}

func (r *BillableMetricResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billable_metric"
}

func (r *BillableMetricResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillableMetricResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *BillableMetricModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.BillableMetrics.New(
		ctx,
		metronome.BillableMetricNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
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

func (r *BillableMetricResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *BillableMetricModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.BillableMetrics.Get(
		ctx,
		data.BillableMetricID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
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

func (r *BillableMetricResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *BillableMetricModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.BillableMetrics.New(
		ctx,
		metronome.BillableMetricNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
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

func (r *BillableMetricResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// TODO: implement not supported warning
}
