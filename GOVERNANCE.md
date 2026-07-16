# 8.0 GOVERNANCE.md

# Buffy Governance

## Philosophy

Buffy exists to solve a specific problem:

> Reality changes.
> Calendars should reflect reality automatically.

Every decision should strengthen that purpose.

The project values simplicity, transparency, predictability, and user ownership over feature count.

---

# Project Leadership

Buffy currently follows a **Founder-Led Open Source** model.

The project maintainer is responsible for:

* overall architecture
* roadmap decisions
* release approval
* conflict resolution
* maintaining project direction

Community discussion is encouraged, but design decisions ultimately prioritize long-term coherence over popularity.

As the project grows, governance may evolve toward a broader maintainer model.

---

# Decision Making

When evaluating new ideas, Buffy asks the following questions:

1. Does this solve a real scheduling problem?
2. Can it be expressed cleanly using the rule language?
3. Does it preserve determinism?
4. Does it increase complexity?
5. Can it be implemented without surprising users?

If the answer to the last two questions is "yes" and "yes," the proposal should probably be reconsidered.

---

# Core Principles

## Reality First

Reality is authoritative.

The scheduler should adapt to reality...

...never require reality to adapt to the scheduler.

---

## User Ownership

Users own:

* their calendars
* their rules
* their data

Buffy exists to serve the user's scheduling model—not the other way around.

---

## Deterministic Behavior

Given the same:

* rules
* calendars
* inputs

Buffy should always produce the same schedule.

No hidden heuristics.

No random behavior.

No opaque AI decisions.

---

## Explicit Over Implicit

Hidden behavior creates mistrust.

Every scheduling decision should be explainable.

If Buffy blocks time, moves an event, or rejects a request, users should be able to understand exactly why.

---

## Composable Rules

Small rules should combine to create powerful scheduling behavior.

Avoid creating special cases whenever a general rule can express the same idea.

---

## Backward Compatibility

Breaking existing rule files should be rare.

When breaking changes are necessary:

* document them clearly
* provide migration guidance
* version the language appropriately

---

# Community Contributions

Everyone is welcome to:

* report bugs
* improve documentation
* propose features
* submit pull requests
* discuss architecture

Constructive disagreement is healthy.

Personal attacks are not.

---

# Code of Conduct

Buffy follows a simple principle:

> Be respectful.
> Be curious.
> Assume good intent.
> Debate ideas vigorously.
> Treat people generously.

Harassment, discrimination, or abusive behavior are not acceptable.

---

# RFC Process

Significant architectural changes should begin with an RFC (Request for Comments).

An RFC should describe:

* the problem
* proposed solution
* alternatives considered
* tradeoffs
* compatibility concerns

Discussion should focus on improving the proposal rather than defending positions.

---

# Releases

Stable releases should prioritize reliability over feature velocity.

A smaller feature set that users trust is preferable to a larger feature set with unpredictable behavior.

---

# Long-Term Vision

Buffy is intended to become dependable infrastructure.

Success is not measured by the number of features.

Success is measured by how accurately Buffy reflects reality while remaining understandable, predictable, and completely under the user's control.
