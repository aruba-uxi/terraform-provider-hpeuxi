package resources

import (
	"context"

	// "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &wirelessNetworkResource{}
	_ resource.ResourceWithConfigure = &wirelessNetworkResource{}
)

type wirelessNetworkResourceModel struct {
	ID    types.String `tfsdk:"id"`
	Alias types.String `tfsdk:"alias"`
}

// TODO: Switch this to use the Client Model when that becomes available
type WirelessNetworkResponseModel struct {
	Uid                  string // <network_uid:str>,
	Ssid                 string // <ssid:str>,
	DatetimeCreated      string // <created_at:str(isoformat(with timezone?))>,
	DatetimeUpdated      string // <updated_at:str(isoformat(with timezone?))>,
	Alias                string // <network_alias>,
	IpVersion            string // <ip_version:str>,
	Security             string // <security:str>,
	Hidden               bool   // <hidden:bool>,
	BandLocking          string // <band_locking:str>,
	DnsLookupDomain      string // <dns_lookup_domain:str> | Nullable,
	DisableEdns          bool   // <disable_edns:bool>,
	UseDns64             bool   // <use_dns64:bool>,
	ExternalConnectivity bool   // <external_connectivity:bool>
}

func NewWirelessNetworkResource() resource.Resource {
	return &wirelessNetworkResource{}
}

type wirelessNetworkResource struct{}

func (r *wirelessNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network"
}

func (r *wirelessNetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"alias": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *wirelessNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *wirelessNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a wireless_network is not supported; wireless_networks can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response := GetWirelessNetwork()

	// Update state from client response
	state.Alias = types.StringValue(response.Alias)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wirelessNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a wireless_network is not supported; wireless_networks can only be updated through the dashboard")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting a wireless_network is not supported; wireless_networks can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Get the wirelessNetwork using the configuration-api client
var GetWirelessNetwork = func() WirelessNetworkResponseModel {
	// TODO: Query the wirelessNetwork using the client

	return WirelessNetworkResponseModel{
		Uid:                  "mock_uid",
		Ssid:                 "mock_ssid",
		DatetimeCreated:      "mock_datetime_created",
		DatetimeUpdated:      "mock_datetime_updated",
		Alias:                "mock_alias",
		IpVersion:            "mock_ip_version",
		Security:             "mock_security",
		Hidden:               false,
		BandLocking:          "mock_band_locking",
		DnsLookupDomain:      "mock_dns_lookup_domain",
		DisableEdns:          false,
		UseDns64:             false,
		ExternalConnectivity: false,
	}
}