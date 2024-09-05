package resources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &sensorGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &sensorGroupAssignmentResource{}
)

type sensorGroupAssignmentResourceModel struct {
	ID       types.String `tfsdk:"id"`
	SensorID types.String `tfsdk:"sensor_id"`
	GroupID  types.String `tfsdk:"group_id"`
}

type SensorGroupAssignmentResponseModel struct {
	UID       string `json:"uid"`
	GroupUID  string `json:"group_uid"`
	SensorUID string `json:"sensor_uid"`
}

type SensorGroupAssignmentRequestModel struct {
	GroupUID  string //  <group_uid:str>,
	SensorUID string //  <sensor_uid:str>
}

func NewSensorGroupAssignmentResource() resource.Resource {
	return &sensorGroupAssignmentResource{}
}

type sensorGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *sensorGroupAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor_group_assignment"
}

func (r *sensorGroupAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sensor_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *sensorGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Resource type: Sensor Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client

}

func (r *sensorGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client createSensorGroupAssignment method
	sensorGroupAssignment := CreateSensorGroupAssignment(SensorGroupAssignmentRequestModel{
		GroupUID:  plan.GroupID.ValueString(),
		SensorUID: plan.SensorID.ValueString(),
	})

	// Update the state to match the plan
	plan.ID = types.StringValue(sensorGroupAssignment.UID)
	plan.GroupID = types.StringValue(sensorGroupAssignment.GroupUID)
	plan.SensorID = types.StringValue(sensorGroupAssignment.SensorUID)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sensorGroupAssignmentResponse, _, err := r.client.ConfigurationAPI.
		GetConfigurationAppV1SensorGroupAssignmentsGet(context.Background()).
		Uid(state.ID.ValueString()).
		Execute()

	if err != nil || len(sensorGroupAssignmentResponse.SensorGroupAssignments) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Sensor Group Assignment",
			"Could not retrieve Sensor Group Assignment, unexpected error: "+err.Error(),
		)
		return
	}

	sensorGroupAssignment := sensorGroupAssignmentResponse.SensorGroupAssignments[0]

	// Update state from client response
	state.GroupID = types.StringValue(sensorGroupAssignment.GroupUid)
	state.SensorID = types.StringValue(sensorGroupAssignment.SensorUid)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a sensor group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing sensorGroupAssignment using the plan_id
}

func (r *sensorGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var CreateSensorGroupAssignment = func(request SensorGroupAssignmentRequestModel) SensorGroupAssignmentResponseModel {
	// TODO: Query the sensorGroupAssignment using the client

	return SensorGroupAssignmentResponseModel{
		UID:       "mock_uid",
		GroupUID:  "mock_group_uid",
		SensorUID: "mock_sensor_uid",
	}
}
