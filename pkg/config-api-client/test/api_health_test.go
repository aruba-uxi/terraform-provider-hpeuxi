package openapi

import (
	"context"
	"testing"

	"github.com/h2non/gock"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_HealthAPIService(t *testing.T) {

	configuration := config_api_client.NewConfiguration()
	configuration.Host = "localhost:80"
	configuration.Scheme = "http"
	apiClient := config_api_client.NewAPIClient(configuration)

	defer gock.Off()

	t.Run("Test HealthAPIService GetLivezHealthLivezGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/livez").Reply(200).JSON(map[string]string{"status": "OK"})

		resp, httpRes, err := apiClient.HealthAPI.GetLivezHealthLivezGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, httpRes.StatusCode, 200)
		assert.Equal(t, resp, &config_api_client.LivenessResponse{Status: "OK"})
	})

	t.Run("Test HealthAPIService GetReadyzHealthReadyzGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/readyz").Reply(200).JSON(map[string]interface{}{"status": "OK", "data": map[string]string{}})

		resp, httpRes, err := apiClient.HealthAPI.GetReadyzHealthReadyzGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.ReadinessResponse{Status: "OK", Data: map[string]string{}})
	})

	t.Run("Test HealthAPIService GetStatusHealthStatusGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/status").Reply(200).JSON(map[string]string{"name": "Configuration-API", "version": "1.0.0"})

		resp, httpRes, err := apiClient.HealthAPI.GetStatusHealthStatusGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.StatusResponse{Name: "Configuration-API", Version: "1.0.0"})
	})

}
