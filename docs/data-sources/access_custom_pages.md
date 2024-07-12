---
page_title: "cloudflare_access_custom_pages Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_custom_pages (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Optional:

- `app_count` (Number) Number of apps the custom page is assigned to.
- `created_at` (String)
- `uid` (String) UUID
- `updated_at` (String)

Read-Only:

- `name` (String) Custom page name.
- `type` (String) Custom page type.

