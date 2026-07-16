# Buffy Architecture

> *"Reality is the source of truth. Calendars are merely projections."*

---

# Status

This document describes the intended architecture of Buffy.

Implementation details may evolve as the project matures, but the architectural principles should remain stable.

---

# Overview

Buffy is a rule engine that produces one or more availability calendars from a set of declarative rules.

Unlike traditional scheduling software, Buffy does not ask users to manually block every possible conflict.

Instead, Buffy continuously answers one question:

> **"Given everything that is true right now... when am I actually available?"**

The generated availability can then be consumed by scheduling systems such as:

- Cal.com
- Google Calendar
- Proton Calendar
- Outlook
- Apple Calendar
- booking systems
- automation platforms

Buffy itself owns no appointments.

Appointments belong to external calendars.

Buffy owns only the logic that determines availability.

---

# Design Philosophy

## Reality First

Reality changes.

Calendars lag behind.

Reality always wins.

---

## Rules Before Events

Most scheduling software stores events.

Buffy stores rules.

Examples:

Every Tuesday is Art Day.

Never schedule meetings after therapy.

Allow coaching only when my partner is home.


Rules generate availability automatically.

---

## Deterministic

Given the same inputs...

Buffy always produces the same output.

No randomness.

No hidden AI decisions.

---

## Explainable

Every decision must be explainable.

Instead of:

Not Available


Buffy should say:

Blocked because:

Tuesday Art Rule
Driving Buffer
Existing Appointment


Every result is traceable.

---

## Modular

Every component has one responsibility.

Nothing should know more than it needs to.

---

# High-Level Architecture

Calendars
Weather
Location
Partner Calendar
Holiday Sources
User Rules
Manual Overrides

    |
    v

INPUT LAYER

    |
    v
 NORMALIZATION
        |
    v

RULE ENGINE

    |
    v
    CONFLICT RESOLUTION
        |
        |
        |
    v
    AVAILABILITY GRAPH
        |
    v
    OUTPUT CONNECTORS
        |
    v
   Google
Proton
Cal.com
ICS
API 

---

# Core Components

## Input Layer

The input layer collects observations.

Examples:

- Google Calendar
- Proton Calendar
- ICS feeds
- Partner calendars
- GPS/location
- weather services
- public holidays
- user rule files

Input sources never make decisions.

They only report reality.

---

# Normalization Layer

The normalization layer converts different sources into a common internal representation.

For example:

Dentist Appointment
Work Shift
Meeting


all become:

Busy Time


with metadata attached.

Normalization removes platform differences.

---

# Rule Engine

The rule engine is the heart of Buffy.

It consumes:

- observations
- rules
- context

and produces:

- candidate availability

The rule engine contains no calendar-specific logic.

It only evaluates reality against rules.

---

# Conflict Resolver

Rules can disagree.

Example:

Rule A:
Always available Fridays.

Rule B:
Never schedule after 3PM.

The conflict resolver determines the final outcome using defined priority rules.

---

# Availability Graph

The availability graph is Buffy's true output.

It represents:

- available time
- unavailable time
- reasons
- priorities
- dependencies

Everything else is a projection.

Calendars are views of the availability graph.

---

# Output Connectors

Output connectors translate Buffy availability into external systems.

Examples:

- Google Calendar blocks
- Proton Calendar feeds
- ICS exports
- REST API
- Cal.com availability
- future integrations

Connectors should never modify the core engine.

---

# Rule Lifecycle

Every rule follows the same lifecycle.

Read Rule

↓

Validate

↓

Evaluate

↓

Generate Candidate Blocks

↓

Resolve Conflicts

↓

Merge

↓

Publish


Keeping every rule consistent makes Buffy extensible.

---

# Data Model

Buffy operates around a small number of core concepts.

---

## Observation

Something that is true.

Examples:

- calendar event
- sunrise
- weather condition
- location
- travel time

---

## Rule

A statement that transforms observations into decisions.

Example:

IF partner is working

THEN

block:
30 minutes before
and
30 minutes after


---

## Constraint

A rule that removes availability.

Examples:

- meetings
- sleep
- driving
- commitments

---

## Opportunity

A rule that creates availability.

Examples:

- focus time
- creative blocks
- coaching windows

---

## Override

A temporary manual intervention.

Overrides expire.

Rules remain.

---

# Layers of Truth

Buffy separates different kinds of reality.

## Physical Reality

Examples:

- location
- weather
- travel time
- daylight

---

## Calendar Reality

Examples:

- appointments
- meetings
- vacations

---

## Social Reality

Examples:

- partner schedules
- family commitments
- school
- volunteer roles

---

## User Intent

Examples:

- art day
- focus time
- lunch
- deep work

---

Each layer is evaluated independently before being merged.

---

# Plugin Architecture

External integrations are plugins.

Example:
plugins/

google/

proton/

ics/

calcom/

weather/

travel/


Plugins translate information.

They do not change the engine.

---

# Rule Language

Buffy's rule language is intentionally human-readable.

Example:

```yaml
rule:

  name: Art Tuesday

  when:

    weekday: Tuesday

  action:

    block:

      start: 08:00

      end: 17:00

      The goal is for people to describe their reality without needing programming skills.

Performance Goals

Buffy should support:

thousands of rules
hundreds of calendars
near-instant recalculation
incremental updates
deterministic output
Explainability

Every availability decision should answer:

"Why?"

Example:

Unavailable

Reason:

Partner Work Rule

↓

Driving Buffer

↓

Lunch Rule

↓

Existing Appointment

Users should never wonder why time disappeared.

Future AI

AI is optional.

Never required.

AI may:

help write rules
suggest improvements
explain conflicts

AI must never silently modify user rules.

AI is an assistant.

The rule engine remains deterministic.

Non-Goals

Buffy is not:

a calendar
a booking platform
a task manager
an AI scheduler
a project management system

Buffy computes availability.

Nothing more.

Architectural Principle

Reality → Observations → Rules → Availability → Calendars

Never the reverse.

Final Principle

Buffy doesn't predict your life. Buffy continuously reconciles your intentions with reality and projects the result as trustworthy availability.