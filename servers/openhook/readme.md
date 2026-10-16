Docs: https://openhook.dev/docs

Source: https://github.com/openhook-dev/openhook-inbox (MIT).

Connect to the hosted endpoint to create a disposable callback inbox. No account
or global API key is needed to connect or list tools. The private token returned
by create_webhook is required as a tool argument to read or delete that inbox;
only share its public callback URL with senders. Keep tokens out of logs.

Use synthetic or authorized test data. Capturing an event confirms receipt at
Openhook, not processing by your application. Delete disposable test inboxes
when finished. Persistent hosted inboxes require a separate Openhook account key.
