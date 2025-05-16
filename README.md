# Metronome Terraform Provider

The [Metronome Terraform provider](https://registry.terraform.io/providers/stainless-sdks/metronome/latest/docs) provides convenient access to
the [Metronome REST API](https://docs.metronome.com) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

```hcl
# Declare the provider and version
terraform {
  required_providers {
    metronome = {
      source  = "stainless-sdks/metronome"
      version = "~> 0.0.1-alpha.0"
    }
  }
}

# Initialize the provider
provider "metronome" {
  bearer_token = "My Bearer Token" # or set METRONOME_BEARER_TOKEN env variable
  webhook_secret = "My Webhook Secret" # or set METRONOME_WEBHOOK_SECRET env variable
}

# Configure a resource
resource "metronome_v1_contract" "example_v1_contract" {
  customer_id = "13117714-3f05-48e5-a6e9-a66093f13b4d"
  starting_at = "2020-01-01T00:00:00.000Z"
  billing_provider_configuration = {
    billing_provider = "aws_marketplace"
    billing_provider_configuration_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    delivery_method = "direct_to_billing_provider"
  }
  commits = [{
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    type = "PREPAID"
    access_schedule = {
      schedule_items = [{
        amount = 0
        ending_before = "2019-12-27T18:11:19.117Z"
        starting_at = "2019-12-27T18:11:19.117Z"
      }]
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    }
    amount = 0
    applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
    applicable_product_tags = ["string"]
    custom_fields = {
      foo = "string"
    }
    description = "description"
    invoice_schedule = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      recurring_schedule = {
        amount_distribution = "DIVIDED"
        ending_before = "2019-12-27T18:11:19.117Z"
        frequency = "MONTHLY"
        starting_at = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }
      schedule_items = [{
        timestamp = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }]
    }
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
    payment_gate_config = {
      payment_gate_type = "NONE"
      stripe_config = {
        payment_type = "INVOICE"
      }
      tax_type = "NONE"
    }
    priority = 0
    rate_type = "COMMIT_RATE"
    rollover_fraction = 0
    temporary_id = "temporary_id"
  }]
  credits = [{
    access_schedule = {
      schedule_items = [{
        amount = 0
        ending_before = "2019-12-27T18:11:19.117Z"
        starting_at = "2019-12-27T18:11:19.117Z"
      }]
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    }
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
    applicable_product_tags = ["string"]
    custom_fields = {
      foo = "string"
    }
    description = "description"
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
    priority = 0
    rate_type = "COMMIT_RATE"
  }]
  custom_fields = {
    foo = "string"
  }
  discounts = [{
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    schedule = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      recurring_schedule = {
        amount_distribution = "DIVIDED"
        ending_before = "2019-12-27T18:11:19.117Z"
        frequency = "MONTHLY"
        starting_at = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }
      schedule_items = [{
        timestamp = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }]
    }
    custom_fields = {
      foo = "string"
    }
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
  }]
  ending_before = "2019-12-27T18:11:19.117Z"
  multiplier_override_prioritization = "LOWEST_MULTIPLIER"
  name = "name"
  net_payment_terms_days = 0
  netsuite_sales_order_id = "netsuite_sales_order_id"
  overrides = [{
    starting_at = "2019-12-27T18:11:19.117Z"
    applicable_product_tags = ["string"]
    ending_before = "2019-12-27T18:11:19.117Z"
    entitled = true
    is_commit_specific = true
    multiplier = 0
    override_specifiers = [{
      commit_ids = ["string"]
      presentation_group_values = {
        foo = "string"
      }
      pricing_group_values = {
        foo = "string"
      }
      product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      product_tags = ["string"]
      recurring_commit_ids = ["string"]
      recurring_credit_ids = ["string"]
    }]
    overwrite_rate = {
      rate_type = "FLAT"
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      custom_rate = {
        foo = "bar"
      }
      is_prorated = true
      price = 0
      quantity = 0
      tiers = [{
        price = 0
        size = 0
      }]
    }
    priority = 0
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    target = "COMMIT_RATE"
    tiers = [{
      multiplier = 0
      size = 0
    }]
    type = "OVERWRITE"
  }]
  prepaid_balance_threshold_configuration = {
    commit = {
      product_id = "product_id"
      applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
      applicable_product_tags = ["string"]
      description = "description"
      name = "name"
    }
    is_enabled = true
    payment_gate_config = {
      payment_gate_type = "NONE"
      stripe_config = {
        payment_type = "INVOICE"
      }
      tax_type = "NONE"
    }
    recharge_to_amount = 0
    threshold_amount = 0
  }
  professional_services = [{
    max_amount = 0
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    quantity = 0
    unit_price = 0
    custom_fields = {
      foo = "string"
    }
    description = "description"
    netsuite_sales_order_id = "netsuite_sales_order_id"
  }]
  rate_card_alias = "rate_card_alias"
  rate_card_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  recurring_commits = [{
    access_amount = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      quantity = 0
      unit_price = 0
    }
    commit_duration = {
      unit = "PERIODS"
      value = 0
    }
    priority = 0
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    starting_at = "2019-12-27T18:11:19.117Z"
    applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
    applicable_product_tags = ["string"]
    description = "description"
    ending_before = "2019-12-27T18:11:19.117Z"
    invoice_amount = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      quantity = 0
      unit_price = 0
    }
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
    proration = "NONE"
    rate_type = "COMMIT_RATE"
    recurrence_frequency = "MONTHLY"
    rollover_fraction = 0
    temporary_id = "temporary_id"
  }]
  recurring_credits = [{
    access_amount = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      quantity = 0
      unit_price = 0
    }
    commit_duration = {
      unit = "PERIODS"
      value = 0
    }
    priority = 0
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    starting_at = "2019-12-27T18:11:19.117Z"
    applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
    applicable_product_tags = ["string"]
    description = "description"
    ending_before = "2019-12-27T18:11:19.117Z"
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
    proration = "NONE"
    rate_type = "COMMIT_RATE"
    recurrence_frequency = "MONTHLY"
    rollover_fraction = 0
    temporary_id = "temporary_id"
  }]
  reseller_royalties = [{
    fraction = 0
    netsuite_reseller_id = "netsuite_reseller_id"
    reseller_type = "AWS"
    starting_at = "2019-12-27T18:11:19.117Z"
    applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
    applicable_product_tags = ["string"]
    aws_options = {
      aws_account_number = "aws_account_number"
      aws_offer_id = "aws_offer_id"
      aws_payer_reference_id = "aws_payer_reference_id"
    }
    ending_before = "2019-12-27T18:11:19.117Z"
    gcp_options = {
      gcp_account_id = "gcp_account_id"
      gcp_offer_id = "gcp_offer_id"
    }
    reseller_contract_value = 0
  }]
  salesforce_opportunity_id = "salesforce_opportunity_id"
  scheduled_charges = [{
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    schedule = {
      credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
      recurring_schedule = {
        amount_distribution = "DIVIDED"
        ending_before = "2019-12-27T18:11:19.117Z"
        frequency = "MONTHLY"
        starting_at = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }
      schedule_items = [{
        timestamp = "2019-12-27T18:11:19.117Z"
        amount = 0
        quantity = 0
        unit_price = 0
      }]
    }
    name = "x"
    netsuite_sales_order_id = "netsuite_sales_order_id"
  }]
  scheduled_charges_on_usage_invoices = "ALL"
  spend_threshold_configuration = {
    commit = {
      product_id = "product_id"
      description = "description"
      name = "name"
    }
    is_enabled = true
    payment_gate_config = {
      payment_gate_type = "NONE"
      stripe_config = {
        payment_type = "INVOICE"
      }
      tax_type = "NONE"
    }
    threshold_amount = 0
  }
  total_contract_value = 0
  transition = {
    from_contract_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    type = "SUPERSEDE"
    future_invoice_behavior = {
      trueup = "REMOVE"
    }
  }
  uniqueness_key = "x"
  usage_filter = {
    group_key = "group_key"
    group_values = ["string"]
    starting_at = "2019-12-27T18:11:19.117Z"
  }
  usage_statement_schedule = {
    frequency = "MONTHLY"
    billing_anchor_date = "2019-12-27T18:11:19.117Z"
    day = "FIRST_OF_MONTH"
    invoice_generation_starting_at = "2019-12-27T18:11:19.117Z"
  }
}
```

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/stainless-sdks/metronome/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property       | Environment variable       | Required | Default value |
| -------------- | -------------------------- | -------- | ------------- |
| bearer_token   | `METRONOME_BEARER_TOKEN`   | true     | —             |
| webhook_secret | `METRONOME_WEBHOOK_SECRET` | false    | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/stainless-sdks/metronome-terraform/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
