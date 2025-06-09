// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_commit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*V1CustomerCommitResource)(nil)

func (r *V1CustomerCommitResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
