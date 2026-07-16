# Buffy Rule Language

> Buffy doesn't own your calendar.
> Buffy keeps it faithful to reality.

---

# Purpose

The Buffy Rule Language is a declarative language for describing the
constraints, preferences, obligations, and recurring structures of a
person's life.

Users describe reality.

Buffy derives availability.

The language intentionally resembles YAML because it is easy for humans
to read, easy for machines to parse, and easy to generate from AI.

Rules are independent.

No rule should require knowledge of another rule.

---

# Core Philosophy

Users should never schedule time.

They should describe:

• things that happen
• things that matter
• things that cannot overlap
• priorities
• buffers
• relationships

Buffy computes the calendar.

---

# File Structure

A Buffy file consists of independent rules.

Example:

events:
constraints:
relationships:
preferences:

The order is irrelevant.

---

# Objects

Everything is an object.

Examples

event
calendar
person
resource
place
project
task
vehicle
holiday

Everything may have properties.

Example

person:
  name: Jane

vehicle:
  type: Car

---

# Events

Events represent things that happen.

Example

event:

  name: Dentist

  calendar: Medical

  starts:
    2026-04-14 09:00

  ends:
    2026-04-14 10:00

---

Recurring example

event:

  name: BNI

  repeat:

    weekly:

      weekday: Wednesday

      time: 11:30-13:00

---

# Constraints

Constraints prevent impossible schedules.

Example

constraint:

  if:
    calendar == Work

  then:
    unavailable

---

# Buffers

Example

buffer:

  around:
    Work

  before: 30m

  after: 30m

Meaning:

Every Work event creates
30 minutes before
30 minutes after

without editing the original calendar.

---

# Relationships

Rules may reference other objects.

Example

relationship:

  source:
    Partner Calendar

  event:
    Work

  create:

    buffer:
      before: 30m
      after: 30m

This allows derived events.

---

# Conditions

Rules may contain conditions.

Example

if:

  weekday == Friday

then:

  unavailable

---

Conditions support

equals

not equals

contains

greater than

less than

between

matches

exists

---

# Priorities

Conflicts are resolved using priorities.

Example

priority:

  Medical: 100

  Family: 90

  Work: 80

  Recreation: 40

Higher wins.

---

# Preferences

Preferences are not requirements.

Example

prefer:

  mornings

Example

avoid:

  meetings:
    Friday afternoon

Preferences may be overridden.

---

# Derived Events

Rules may create events.

Example

if:

  Partner.Work

then:

  create:

    Pickup Buffer

This is how commute rules work.

---

# Imports

Large systems may be divided into files.

Example

imports:

  work.yaml

  family.yaml

  holidays.yaml

---

# Comments

Comments begin with #

Example

# Never schedule coaching here.

---

# Human-first Design

The language is intentionally verbose.

This is preferred:

buffer:

  around:

    Partner.Work

Instead of:

buf:
 a:
   PW

Humans read rules far more often than computers.

---

# AI Compatibility

The language is designed so that AI assistants can:

• generate rules

• explain rules

• validate rules

• detect contradictions

• migrate old rules

without requiring users to understand programming.

---

# Example

calendar:

  Coaching

event:

  Coaching Session

  repeat:

    weekly:

      weekday: Monday

      time: 09:00-12:00

buffer:

  around:

    Coaching Session

  before: 15m

  after: 15m

constraint:

  if:

    calendar == Coaching

  then:

    unavailable
