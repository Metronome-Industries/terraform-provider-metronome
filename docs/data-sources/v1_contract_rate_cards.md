---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metronome_v1_contract_rate_cards Data Source - metronome"
subcategory: ""
description: |-
  
---

# metronome_v1_contract_rate_cards (Data Source)



## Example Usage

```terraform
data "metronome_v1_contract_rate_cards" "example_v1_contract_rate_cards" {
  body = {

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `body` (String)
- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `aliases` (Attributes List) (see [below for nested schema](#nestedatt--items--aliases))
- `created_at` (String)
- `created_by` (String)
- `credit_type_conversions` (Attributes List) (see [below for nested schema](#nestedatt--items--credit_type_conversions))
- `custom_fields` (Map of String)
- `description` (String)
- `fiat_credit_type` (Attributes) (see [below for nested schema](#nestedatt--items--fiat_credit_type))
- `id` (String)
- `name` (String)

<a id="nestedatt--items--aliases"></a>
### Nested Schema for `items.aliases`

Read-Only:

- `ending_before` (String)
- `name` (String)
- `starting_at` (String)


<a id="nestedatt--items--credit_type_conversions"></a>
### Nested Schema for `items.credit_type_conversions`

Read-Only:

- `custom_credit_type` (Attributes) (see [below for nested schema](#nestedatt--items--credit_type_conversions--custom_credit_type))
- `fiat_per_custom_credit` (String)

<a id="nestedatt--items--credit_type_conversions--custom_credit_type"></a>
### Nested Schema for `items.credit_type_conversions.custom_credit_type`

Read-Only:

- `id` (String)
- `name` (String)



<a id="nestedatt--items--fiat_credit_type"></a>
### Nested Schema for `items.fiat_credit_type`

Read-Only:

- `id` (String)
- `name` (String)
