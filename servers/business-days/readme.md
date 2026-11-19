# Business Days

Public holidays and working-day arithmetic for 200+ countries, including state
and province subdivisions. No account or API key.

Documentation: https://tools.profitpather.com/days

A language model's failure mode on calendar questions is confident wrongness,
and a tool that is also confidently wrong adds nothing. Holiday data is read
from three independent providers (Nager.Date, Google Calendar ICS, and
OpenHolidays) and **disagreements are reported rather than resolved**, so a
disputed date arrives as a caveat attached to the answer instead of a single
confident date.

Providers are compared by date and never by name, since the same day is named
differently by each. Supply a region code such as `AU-NSW` or `US-CA` whenever a
state or province matters, otherwise only nationwide holidays are returned.
