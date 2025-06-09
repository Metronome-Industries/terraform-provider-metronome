resource "metronome_v1_contract_product" "example_v1_contract_product" {
  name = "name"
  type = "FIXED"
  billable_metric_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  composite_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  composite_tags = ["string"]
  exclude_free_usage = true
  is_refundable = true
  netsuite_internal_item_id = "netsuite_internal_item_id"
  netsuite_overage_item_id = "netsuite_overage_item_id"
  presentation_group_key = ["string"]
  pricing_group_key = ["string"]
  quantity_conversion = {
    conversion_factor = 0
    operation = "MULTIPLY"
    name = "name"
  }
  quantity_rounding = {
    decimal_places = 0
    rounding_method = "ROUND_UP"
  }
  tags = ["string"]
}
