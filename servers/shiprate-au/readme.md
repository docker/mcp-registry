# ShipRate AU

Australian parcel shipping rates from two independent sources in one call, plus
postcode resolution and chargeable-weight arithmetic. No account or API key.

Documentation: https://tools.profitpather.com/ship

Every rate carries a `pricing_basis`. Australia Post prices are **retail**, so a
business on an eParcel or contract account pays less; the commercial panel is
pre-negotiated pricing and a different class entirely. A lower number is
therefore not automatically a better carrier, and the server says so rather than
declaring a winner.

The server will not rank or sum across currencies, because the upstream data
carries no exchange rate. When more than one currency is returned the list is
left unsorted and `sorted_by_price` is `false`. A source that quoted nothing is
reported in `source_notes` rather than hidden, so an empty list is never
presented as "no options exist".

International prices exclude destination duty and tax.
