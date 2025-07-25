---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metronome_v1_contract_rate_card_named_schedule Data Source - metronome"
subcategory: ""
description: |-
  
---

# metronome_v1_contract_rate_card_named_schedule (Data Source)



## Example Usage

```terraform
data "metronome_v1_contract_rate_card_named_schedule" "example_v1_contract_rate_card_named_schedule" {
  contract_id = "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"
  customer_id = "9b85c1c1-5238-4f2a-a409-61412905e1e1"
  schedule_name = "my-schedule"
  covering_date = "2022-02-15T00:00:00Z"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `contract_id` (String) ID of the contract whose named schedule is to be retrieved
- `customer_id` (String) ID of the customer whose named schedule is to be retrieved
- `schedule_name` (String) The identifier for the schedule to be retrieved

### Optional

- `covering_date` (String) If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.

### Read-Only

- `data` (Attributes List) (see [below for nested schema](#nestedatt--data))

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `ending_before` (String)
- `starting_at` (String)
- `value` (String)
