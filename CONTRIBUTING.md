# CONTRIBUTING.md

# Contributing to Buffy

Thank you for your interest in Buffy.

Buffy is an attempt to solve a surprisingly difficult problem:

> Maintain a faithful representation of a person's availability based on reality rather than manual calendar management.

Although the project may begin with a single contributor, it is designed from the beginning to become an open, well-engineered system. This document exists to ensure that every contribution moves the project toward that goal.

---

# Philosophy

Before writing code, understand the architecture.

Buffy is not "another scheduling application."

It is a deterministic rule engine that produces availability.

Everything else is an interface.

If a proposed feature requires special cases, exceptions, or hidden state, the design should be questioned before implementation begins.

The architecture should become simpler as the project grows.

---

# Core Principles

Every contribution should support these principles.

## Reality First

Reality is the source of truth.

Calendars are generated representations of reality—not the other way around.

---

## Deterministic

The same rules applied to the same inputs must always produce the same outputs.

No randomness.

No hidden state.

No unexplained behaviour.

---

## Explicit over Implicit

Rules should state exactly what they mean.

Avoid "magic."

Avoid hidden defaults.

Avoid surprising behaviour.

---

## Human Readable

Rules should be understandable without reading source code.

Configuration is part of the product.

If users cannot understand their own scheduling rules, Buffy has failed.

---

## Small Composable Pieces

Prefer many simple rules over one complicated rule.

Composition scales.

Complexity does not.

---

# Before Opening an Issue

Please determine whether the problem is:

* a bug
* a missing capability
* an implementation detail
* a documentation issue
* an architectural concern

Architecture discussions are encouraged before implementation.

The most valuable contribution is often discovering that something should **not** be built.

---

# Before Writing Code

Ask yourself:

* Does this make the engine simpler?
* Does this preserve determinism?
* Does this increase faithfulness to reality?
* Can this be expressed as a rule instead of code?
* Can an existing abstraction solve this?

If the answer is "no" to several of these questions, open a design discussion first.

---

# Coding Guidelines

## Keep Functions Small

Functions should do one thing.

If explaining a function requires several paragraphs, it is probably doing too much.

---

## Avoid Hidden State

Inputs should be explicit.

Outputs should be explicit.

Dependencies should be obvious.

---

## Pure Logic First

Whenever possible, the scheduling engine should be pure.

External systems (Google Calendar, CalDAV, APIs, databases, etc.) belong at the edges.

Core scheduling logic should remain independent of infrastructure.

---

## Write Tests

Every behavioural rule should have tests.

Tests are executable documentation.

A bug should become a regression test before it becomes a fix.

---

# Documentation Matters

Architecture documents are first-class artifacts.

When behaviour changes, documentation should change with it.

If code and documentation disagree, one of them is wrong.

---

# Commit Messages

Prefer clear commit messages.

Examples:

```
Add recurring exclusion rules

Fix timezone normalization

Refactor rule evaluator

Implement interval merging

Document precedence resolution
```

Avoid messages like:

```
stuff

fixes

updates

misc

working now
```

---

# Pull Requests

A good pull request explains:

* What changed
* Why it changed
* What alternatives were considered
* Any architectural implications

Small pull requests are preferred over large ones.

---

# Project Priorities

In order:

1. Correctness
2. Simplicity
3. Predictability
4. Performance
5. Convenience

Performance matters.

Correctness matters more.

---

# What Buffy Is Not

Buffy is not:

* another calendar application
* a to-do list
* a project manager
* an AI assistant
* a productivity system

It is a scheduling engine.

Everything else should remain optional.

---

# The Buffy Standard

A contribution is successful when it makes the system:

* simpler
* more predictable
* more faithful to reality
* easier to understand
* easier to extend

If a feature makes Buffy more complicated than the problem it solves, it probably doesn't belong.

---

# Mission Statement

Buffy doesn't own your calendar.

Buffy keeps it faithful to reality.
