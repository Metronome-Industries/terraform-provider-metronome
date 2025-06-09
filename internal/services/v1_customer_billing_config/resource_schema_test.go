// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_billing_config"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1CustomerBillingConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_customer_billing_config.V1CustomerBillingConfigModel)(nil)
	schema := v1_customer_billing_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
