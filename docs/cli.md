

Sovereign Pathway — CLI Specification.txt
35 lignes

txt

translate to English: Sovereign Pathway — CLI Specification.txt

2:24 PM
I'll translate the document from French to English for you.


2 étapes

Creating the English translation of the Sovereign Pathway CLI Specification document
Sovereign_Pathway_CLI_Specification_EN.txt

Presented file
I've translated the document from French to English. The document appears to be incomplete (it cuts off in the middle of the sp ingest command section), but I've translated all the content that was provided. The key concepts translated include:

Cardinal principle: No implicit authority, no daemon required, no prior provisioning
Design axioms: CLI hosts nothing by default, all actions are voluntary/local/traceable, works across various devices
Sovereign ingestion: The command for transforming markdown files into immutable sovereign objects
Sovereign pathway cli specification en
TXT 





Claude est une IA et peut faire des erreurs. Veuillez vérifier les réponses.
Sovereign pathway cli specification en · TXT
Copier


cd ~/nostr-sovereign


# Sovereign Pathway — CLI Specification (v1)

Reference CLI: `sp`  
Status: **normative / frozen**

> Cardinal principle: no implicit authority, no daemon required, no prior provisioning.

---

## 0. Design axioms

- The CLI **hosts nothing by default**
- Every action is **voluntary, local, traceable**
- Works on:
  - laptop
  - recent phone
  - old edge device (degraded mode)
- No mandatory global state
- Everything that is signed can be **verified offline**

---

## 1. `sp ingest` — Sovereign ingestion

### Role
Transform a raw `.md` into an **immutable sovereign object**.

### Command
```bash
sp ingest draft.md
```
