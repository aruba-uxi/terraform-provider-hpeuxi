/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider"
)

const (
	ProviderConfig = `provider "uxi" {
		client_id     = "client_id"
		client_secret = "client_secret"
	}`
)

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"uxi": providerserver.NewProtocol6WithError(provider.New("test")()),
}
