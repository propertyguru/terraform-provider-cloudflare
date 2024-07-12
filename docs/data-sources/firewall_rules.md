---
page_title: "cloudflare_firewall_rules Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_firewall_rules (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_identifier` (String) Identifier

### Optional

- `action` (String) The action to search for. Must be an exact match.
- `description` (String) A case-insensitive string to find in the description.
- `id` (String) The unique identifier of the firewall rule.
- `max_items` (Number) Max items to fetch, default: 1000
- `page` (Number) Page number of paginated results.
- `paused` (Boolean) When true, indicates that the firewall rule is currently paused.
- `per_page` (Number) Number of firewall rules per page.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Optional:

- `description` (String) An informative summary of the firewall rule.
- `priority` (Number) The priority of the rule. Optional value used to define the processing order. A lower number indicates a higher priority. If not provided, rules with a defined priority will be processed before rules without a priority.
- `products` (List of String)
- `ref` (String) A short reference tag. Allows you to select related firewall rules.

Read-Only:

- `action` (String) The action to apply to a matched request. The `log` action is only available on an Enterprise plan.
- `id` (String) The unique identifier of the firewall rule.
- `paused` (Boolean) When true, indicates that the firewall rule is currently paused.

