---
# beans-qc6g
title: Integrate Agent Communication Protocol (ACP) as agent communication layer
status: draft
type: feature
tags:
    - idea
created_at: 2026-03-10T11:22:05Z
updated_at: 2026-03-10T11:22:05Z
---

Add ACP (Agent Communication Protocol) support to beans-serve, enabling standardized REST+SSE communication with agents alongside the existing GraphQL API.

## Context

ACP (https://agentcommunicationprotocol.dev/) is an open protocol (v0.2.0, Apache 2.0, Linux Foundation) that standardizes how clients communicate with AI agents via REST. Currently, beans-serve talks to agents exclusively through a custom GraphQL API backed by Claude Code CLI processes. Adding ACP support would allow any ACP-compatible client to interact with beans-managed agents, and potentially allow beans to orchestrate agents hosted on external ACP servers.

## ACP Protocol Summary

**Endpoints:**
- `GET /agents` — List available agents (with manifests describing capabilities)
- `GET /agents/{name}` — Get agent manifest
- `POST /runs` — Create a run (sync, async, or streaming via SSE)
- `GET /runs/{run_id}` — Get run status
- `POST /runs/{run_id}` — Resume a paused/awaiting run
- `POST /runs/{run_id}/cancel` — Cancel a run
- `GET /runs/{run_id}/events` — List run events
- `GET /session/{session_id}` — Get session details

**Key Concepts:**
- **Agent Manifests**: Describe agent capabilities, supported content types, metadata
- **Runs**: A single agent execution with input messages, supporting sync/async/stream modes
- **Messages**: Multipart (text, images, files) with roles (`user`, `agent`, `agent/{name}`)
- **Sessions**: Stateful conversation context across multiple runs
- **Await**: Agent pauses execution requesting external input (maps to beans' pending interactions — permission requests, plan mode switches, ask-user)
- **Streaming**: SSE events (`message.created`, `message.part`, `message.completed`, `run.created`, `run.in-progress`, `run.awaiting`, `run.completed`, `run.failed`, etc.)
- **Trajectory Metadata**: Tool execution and reasoning step tracking

**OpenAPI Spec:** https://github.com/i-am-bee/acp/blob/main/docs/spec/openapi.yaml

## Mapping to beans' Current Architecture

| ACP Concept | beans Equivalent |
|---|---|
| Agent | Bean worktree agent session (keyed by beanID) |
| Run | A message-send-to-response cycle within a session |
| Message | AgentMessage (user/assistant/tool) |
| Await | PendingInteraction (permission_request, ask_user, plan mode) |
| Session | Session (with JSONL persistence, --resume) |
| SSE streaming | GraphQL subscriptions (pub/sub channels) |
| Agent manifest | Could describe Claude Code capabilities per-worktree |

## Potential Scope

### Phase 1: ACP Server (beans-serve exposes ACP endpoints)
- Mount ACP REST endpoints on beans-serve alongside GraphQL
- Translate between ACP runs/messages and internal agent sessions
- Map ACP await/resume to beans' permission and plan-mode interactions
- SSE streaming for real-time agent output
- Agent discovery endpoint listing available bean-scoped agents

### Phase 2: ACP Client (beans consumes external ACP agents)
- Allow beans to delegate to agents hosted on external ACP servers
- Agent type selection (not just Claude Code) per bean/worktree
- External agent discovery and registration

## Open Questions
- Should ACP endpoints replace or coexist with the GraphQL agent API?
- How to map bean IDs to ACP agent names (RFC 1123 DNS label format)?
- Should each bean worktree be a separate "agent", or should there be one agent with sessions per bean?
- How to expose tool invocations via ACP's trajectory metadata?
- Authentication/authorization for the ACP endpoints?

## References
- Protocol spec: https://agentcommunicationprotocol.dev/
- OpenAPI spec: https://github.com/i-am-bee/acp/blob/main/docs/spec/openapi.yaml
- GitHub: https://github.com/i-am-bee/acp
- Python and TypeScript SDKs available
