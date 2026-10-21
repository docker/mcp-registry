# Hail

Add phone numbers to your project without setting up separate carrier
accounts and integrations.

Connect Docker MCP Toolkit to Hail's hosted MCP service. Search numbers
in 120 countries, compare prices, and read registration requirements.
Your assistant can submit the details and documents you supply, check
approval, and order numbers after you confirm. Carriers make the approval
decision. Availability and local requirements vary.

Use one Hail account for inbound and outbound calls, texts, and email.
Read replies and delivery records through MCP. Signed API webhooks let
your app receive supported communication events.

Start with a read-only request:

```text
Find call-capable Hail numbers in the UK, France, and Canada. Compare
monthly fees, setup fees, and registration requirements. Do not buy
numbers or submit documents.
```

Refer to the [MCP setup guide](https://hail.so/mcp) for the hosted endpoint,
OAuth sign-in, and client setup. Hail Cloud charges for communication usage.
Number purchases and paid registrations require account credit.
Approve the required details, documents, and fees before submitting.
Sending requires account credit, a usable sender, and recipient permission.
US outbound SMS is supported. Hail helps prepare and submit A2P 10DLC
registration for US local-number texting. Test with an approved registration
and a ready SMS sender. US inbound and outbound calls are also supported.

The [source code](https://github.com/hail-hq/hail) is available under AGPL-3.0-or-later.
This entry connects to Hail Cloud. Docker does not build or distribute the
Hail server through this entry.

Refer to [pricing](https://hail.so/pricing) and the
[privacy policy](https://hail.so/legal/privacy). Contact `hi@hail.so` for help.
