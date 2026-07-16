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
