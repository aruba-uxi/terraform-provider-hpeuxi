/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &wiredNetworkResource{}
	_ resource.ResourceWithConfigure = &wiredNetworkResource{}
)

type wiredNetworkResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewWiredNetworkResource() resource.Resource {
	return &wiredNetworkResource{}
}

type wiredNetworkResource struct {
	client *config_api_client.APIClient
}

func (r *wiredNetworkResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_wired_network"
}

func (r *wiredNetworkResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *wiredNetworkResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Resource type: Wired Network. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *wiredNetworkResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan wiredNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating a wired_network is not supported; wired_networks can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state wiredNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		WiredNetworksGet(ctx).
		Id(state.ID.ValueString())
	networkResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_wired_network")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(networkResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}

	network := networkResponse.Items[0]

	state.ID = types.StringValue(network.Id)
	state.Name = types.StringValue(network.Name)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wiredNetworkResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan wiredNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"updating a wired_network is not supported; wired_networks can only be updated through the dashboard",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state wiredNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError(
		"operation not supported",
		"deleting a wired_network is not supported; wired_networks can only removed from state",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}