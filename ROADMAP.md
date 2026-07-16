# ROADMAP.md

# Buffy Roadmap

> *"Build the smallest engine that faithfully represents reality... then teach it more reality."*

---

# Philosophy

Buffy will be developed in layers.

Each phase should leave the system useful, even if development stopped there.

No phase exists solely as preparation for another.

Every milestone should produce a system that someone could actually use.

---

# Phase 0 — Foundations

## Goal

Create the architectural skeleton.

### Deliverables

* Repository
* Documentation
* Core terminology
* Rule language draft
* Architecture diagrams
* Test framework
* Parser experiments

At the end of Phase 0...

Buffy understands nothing.

But everyone understands how Buffy will eventually understand.

---

# Phase 1 — Static Rules

## Goal

Produce availability from deterministic rules.

Example:

```yaml
every weekday
    unavailable
    from 09:00
    until 17:00
```

Capabilities

* recurring schedules
* weekdays
* weekends
* dates
* holidays
* fixed exclusions
* merging intervals
* conflict resolution
* calendar visualization

No external integrations.

No APIs.

Just reality generated from rules.

---

# Phase 2 — Calendar Import

## Goal

Reality can include existing calendars.

Sources might include:

* Google
* Proton
* Outlook
* Apple
* ICS

Buffy still owns the truth.

Imported calendars become observations.

Rules may override imported events.

Imported events may trigger rules.

Everything remains explainable.

---

# Phase 3 — Rule Language Expansion

Teach Buffy richer concepts.

Examples:

Family members

```yaml
when spouse is working
    unavailable
    for 30 minutes before
    and 30 minutes after
```

Weather

```yaml
when snow forecast exceeds 10 cm
    add 45 minute travel buffer
```

Travel

Driving time

Locations

School schedules

Medical appointments

Seasonal behaviour

This phase transforms Buffy from a scheduler into a reality model.

---

# Phase 4 — Explanations

Every decision becomes inspectable.

Example

```
Unavailable

because

Rule #18

AND

Google Calendar Event

AND

Driving Buffer Rule

AND

School Pickup Rule
```

Nothing is magic.

Every interval has provenance.

---

# Phase 5 — API

Expose Buffy as an engine.

Possible endpoints

```
GET /availability

GET /why

GET /events

POST /rules

POST /simulate
```

Applications become clients.

Buffy remains the authority.

---

# Phase 6 — Simulation

Ask hypothetical questions.

Examples

"What happens if I move this meeting?"

"What if school ends early?"

"What if I remove commute time?"

"What if my partner changes jobs?"

The engine predicts consequences before calendars change.

---

# Phase 7 — Learning Assistance

Buffy begins making suggestions.

Examples

"You always ignore this buffer."

"This appointment is usually moved."

"Travel time has increased."

"This rule never affects anything."

Suggestions never silently modify reality.

Humans remain responsible.

---

# Phase 8 — Ecosystem

Plugins.

Custom rule packs.

Community rules.

Specialized integrations.

Example packs

* Teachers
* Parents
* Healthcare
* Consultants
* Field technicians
* Students
* Shift workers

---

# Phase 9 — Multi-Agent Coordination

Multiple Buffys cooperate.

Examples

Families.

Businesses.

Volunteer organizations.

Sports teams.

Each participant owns their own rules.

Shared availability emerges from negotiation rather than central control.

---

# Long-Term Vision

Eventually Buffy should reason about availability as a living system.

Not merely:

> "When are you free?"

But:

> "What is actually true?"

Availability becomes a consequence of reality rather than something manually maintained.

---

# Success Criteria

Buffy succeeds when people stop maintaining calendars.

Instead...

They describe reality.

Reality produces availability.

Availability produces trustworthy schedules.

The calendar becomes an output, not the source of truth.

---

# Guiding Principle

Each phase should increase Buffy’s fidelity to reality.

Not its complexity.

When forced to choose...

Always choose the design that better represents the world as it actually is.
