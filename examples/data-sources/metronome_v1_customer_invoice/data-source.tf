data "metronome_v1_customer_invoice" "example_v1_customer_invoice" {
  customer_id = "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"
  invoice_id = "6a37bb88-8538-48c5-b37b-a41c836328bd"
  skip_zero_qty_line_items = true
}
