SLAyer is an application designed to provide reporting on Service Level Agreement criteria for issues/tickets stored in systems such as Jira, Github, or Gitlab.

SLAyer aims to be a general purpose toolkit for examining sets of tasks, issues, tickets, or TODO items for their compliance with a set of rules, ranging from time to first response, time since last response, total time to completion, or possibly more complex/compound criteria.

# Installation

# Configuration

# Usage
## Library
Using SLAyer as a library will require providing three values matching the following interfaces:

* `sla.IssueService`
* `sla.Enforcer`
* `sla.Renderer`

Once these values are instantiated, you may pass them to the `sla.Checker`, after which you may use `Report` and `Render` to generate machine and human readable reports on the issues provided by your `IssueService`. SLAyer's built-in `Enforcer` types leverage the provided `Rule` types and loading mechanics, but it is the responsibility of custom `sla.Enforcer` types to acquire their own ruleset and act accordingly.

## Service
SLAyer exposes an HTTP or gRPC API for accessing and managing reports, analytics, and configurations. SLAyer As A Service must be configured to interact with a given "provider" of issues, such as Jira or a custom service.

Once SLAyer is configured and running, it will offer on-demand metrics via a web interface and endpoints offering machine readable payloads for further processing or archival.
