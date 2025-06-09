data "metronome_v1_contract_rate_card_rates" "example_v1_contract_rate_card_rates" {
  at = "2024-01-01T00:00:00.000Z"
  rate_card_id = "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe"
  selectors = [{
    billing_frequency = "MONTHLY"
    partial_pricing_group_values = {
      region = "us-west-2"
      cloud = "aws"
    }
    pricing_group_values = {
      foo = "string"
    }
    product_id = "d6300dbb-882e-4d2d-8dec-5125d16b65d0"
    product_tags = ["string"]
  }]
}
