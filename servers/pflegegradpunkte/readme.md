Docs: https://pflegegradpunkte.de/ki-assistenten.html

# Pflegegradpunkte

Helps relatives estimate the German care grade (Pflegegrad 1 to 5) before the assessment by the Medizinischer Dienst. The assistant asks the questions of the official assessment (Annex 1 to Section 15 SGB XI), and the server calculates points per area, weighted points and the resulting grade. A facts tool returns benefits per grade, deadlines and appeal rules, each with its source and check date. Answers are in German. It asks for scores only, never names or diagnoses, and stores nothing.

## Tools

- `get_assessment_questions`: the assessment questions per area
- `calculate_pflegegrad`: points per area, total and resulting grade
- `get_pflegegrad_facts`: sourced facts on benefits, deadlines and appeals

All tools are read-only. The server is public: no account, no login and no API key.
