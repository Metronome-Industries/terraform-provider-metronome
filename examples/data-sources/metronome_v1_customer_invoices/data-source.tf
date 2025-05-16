data "metronome_v1_customer_invoices" "example_v1_customer_invoices" {
  customer_id = "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"
  credit_type_id = "credit_type_id"
  ending_before = "2019-12-27T18:11:19.117Z"
  skip_zero_qty_line_items = true
  sort = "date_asc"
  starting_on = "2019-12-27T18:11:19.117Z"
  status = "status"
}
