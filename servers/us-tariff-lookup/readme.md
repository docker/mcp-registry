# US Tariff Lookup

Find the right US tariff code for a product and the duty rate that actually
governs it, read live from the official USITC Harmonized Tariff Schedule. No
account or API key.

Documentation: https://tools.profitpather.com/duty

- `find_hts_code` - candidate HTS lines from a plain-language product description
- `lookup_us_duty_rate` - the rate that governs a code, with free-trade eligibility by origin
- `estimate_us_duty` - indicative duty on a customs value

## The two traps it exists to handle

**Rates sit on the 8-digit line.** The 10-digit statistical lines underneath
carry a blank rate and inherit from their parent, so reading the rate off the
code you were handed returns nothing and looks duty-free. Every response says
which line actually sets the rate.

**Duty follows the country of origin of the goods, not the seller's address.**
`AU` appears in the special column for most apparel, which makes "Australian
business, therefore free" a tempting and expensive mistake. For apparel the
rules of origin are yarn-forward, so Australian wool spun and sewn in China is
Chinese origin and pays the general rate.

## What is not in these numbers

Trade remedies. Section 301 on China, Section 232, antidumping and
countervailing duties, and anything imposed by executive action all live in
Chapter 99, stack on top, and routinely exceed the base rate. Chapter 98 and 99
lines are returned in a separate list so a conditional `Free` rate there can
never be read as the duty answer.

Indicative only, not customs advice. The importer of record owns the
classification and the duty payable.
