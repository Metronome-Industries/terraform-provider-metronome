// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_alert_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_alert"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1AlertModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_alert.V1AlertModel)(nil)
	schema := v1_alert.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
