nostr-sovereign

A sovereign content, storage, and negotiation stack — designed to survive outages, coercion, and disappearance.

This repository defines the foundational specifications and intent for a self-sovereign publishing and relay pathway built on:

Markdown as the canonical human artifact

Deterministic objects and receipts

Voluntary hosting (not provisioning)

Nostr as a public, censorship-resistant anchoring layer

Onion + clearnet dual-homing

Edge devices as first-class citizens

This is not a product.
This is a pathway.

Core Components (Conceptual)

The stack is composed of three distinct but interoperable roles:

1. sovereign_pathway.sovereignrelay.xyz — Storage Pathway

What exists, endures.

Defines how .md becomes an immutable object

Governs object identity, lifecycle, and receipts

Makes no assumptions about hosting permanence

Can live on laptops, phones, disks, or disappear entirely

📁 Specs live in /specs
📁 Lifecycle docs live in /docs

2. sovereignrelay.xyz — Relay Layer

What moves, when possible.

Onion and clearnet reachable

Voluntary relaying only

No requirement to store long-term

Receipts prove passage, not trust

⚠️ Relay implementation is intentionally deferred
This repo defines how a relay behaves, not who runs it.

3. provisioner.quest — Negotiation Layer

What asks, before taking.

Negotiates PayJoin v2 and Silent Payments

Negotiates object fetch and hosting

Never auto-provisions

Never assumes authority

Provisioning is opt-in, contextual, and reversible.

What This Repository Contains
/specs

Hard, frozen specifications:

object.schema.json
Canonical object format (v1)

receipt.md
Receipt issuance, signing, and chaining rules

nostr-mapping.md
How objects and receipts anchor to Nostr
(NIP-01 + NIP-23)

These files are intent anchors.
They change slowly and deliberately.

/docs

Human-readable doctrine and workflow:

object-lifecycle.md
.md → object → receipt → anchoring → re-hosting

dual-home.md
Clearnet ↔ Onion coexistence model

cli.md
Frozen CLI primitives (sp ingest, sp materialize, etc.)

Design Principles

Markdown First
If it can’t live as .md, it doesn’t belong here.

Receipts Over Trust
No reputation systems. Only signed facts.

Voluntary Hosting
Nodes declare willingness — never obligation.

Edge-First
Old phones must work. Laptops are optional.

Negotiation Before Transfer
Especially for money. Especially for bandwidth.

Disappearability
The system must function if the author vanishes.

What This Is Not

Not a blockchain

Not a DAO

Not a cloud service

Not a promise of uptime

Not a custodial system

This is a sovereign grammar, not a platform.

Status

✅ Object schema frozen (v1)

✅ Receipt model frozen

✅ Nostr anchoring defined

✅ CLI semantics frozen

⏳ Reference implementation (minimal)

⏳ Edge-device field testing

⏳ Relay packaging

Nothing here is theoretical anymore —
but not everything is implemented yet.

Who Is This For?

Writers who don’t want to disappear when platforms do

Developers who want clean specs before code

Operators who believe “hosting” should be voluntary

Readers who value provenance over polish

License & Ethos

This repository favors:

Forkability

Attribution

Non-coercion

Licensing will be clarified once the minimal implementation lands.

Until then:
Read the specs. Respect the receipts. Don’t assume authority.

Next Steps (Explicitly Deferred)

Minimal sp reference implementation

Receipt verification tooling

Provisioner negotiation flows

PayJoin / Silent Payments wiring

These proceed only after this README is committed.

Pirate First.
Sovereign, without spectacle.
