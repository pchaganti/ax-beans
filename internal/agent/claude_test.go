package agent

import (
	"strings"
	"testing"
)

// TestReadOutputMessageOrder verifies that tool messages appear between
// the assistant text that precedes and follows them, not grouped at the end.
func TestReadOutputMessageOrder(t *testing.T) {
	// Simulate Claude Code stream-json output:
	// 1. Assistant starts typing "Before tool"
	// 2. Tool "Read" is invoked
	// 3. Assistant continues with "After tool"
	// 4. Result event closes the turn
	lines := strings.Join([]string{
		// First text block starts
		`{"type":"content_block_start","content_block":{"type":"text","text":""}}`,
		`{"type":"content_block_delta","delta":{"type":"text_delta","text":"Before tool"}}`,
		// Tool use
		`{"type":"content_block_start","content_block":{"type":"tool_use","name":"Read"}}`,
		`{"type":"content_block_delta","delta":{"type":"input_json_delta","partial_json":"{\"file_path\":\"/tmp/test\"}"}}`,
		// New text block after tool
		`{"type":"content_block_start","content_block":{"type":"text","text":""}}`,
		`{"type":"content_block_delta","delta":{"type":"text_delta","text":"After tool"}}`,
		// Turn complete
		`{"type":"result","session_id":"sess-1"}`,
	}, "\n")

	m := &Manager{
		sessions:    make(map[string]*Session),
		processes:   make(map[string]*runningProcess),
		subscribers: make(map[string][]chan struct{}),
	}

	session := &Session{
		ID:           "bean-test",
		AgentType:    "claude",
		Status:       StatusRunning,
		Messages:     []Message{{Role: RoleUser, Content: "hello"}},
		streamingIdx: -1,
	}
	m.sessions["bean-test"] = session

	m.readOutput("bean-test", strings.NewReader(lines), "")

	// Expected message order:
	// [0] USER: "hello"          (pre-existing)
	// [1] ASSISTANT: "Before tool"
	// [2] TOOL: "Read: /tmp/test"
	// [3] ASSISTANT: "After tool"
	msgs := session.Messages

	if len(msgs) != 4 {
		for i, m := range msgs {
			t.Logf("  [%d] %s: %q", i, m.Role, m.Content)
		}
		t.Fatalf("expected 4 messages, got %d", len(msgs))
	}

	tests := []struct {
		idx     int
		role    MessageRole
		contain string
	}{
		{0, RoleUser, "hello"},
		{1, RoleAssistant, "Before tool"},
		{2, RoleTool, "Read"},
		{3, RoleAssistant, "After tool"},
	}

	for _, tt := range tests {
		msg := msgs[tt.idx]
		if msg.Role != tt.role {
			t.Errorf("msgs[%d].Role = %q, want %q", tt.idx, msg.Role, tt.role)
		}
		if !strings.Contains(msg.Content, tt.contain) {
			t.Errorf("msgs[%d].Content = %q, want it to contain %q", tt.idx, msg.Content, tt.contain)
		}
	}
}

// TestReadOutputMultipleTools verifies ordering with multiple tool uses in a single turn.
func TestReadOutputMultipleTools(t *testing.T) {
	lines := strings.Join([]string{
		`{"type":"content_block_start","content_block":{"type":"text","text":""}}`,
		`{"type":"content_block_delta","delta":{"type":"text_delta","text":"Step 1"}}`,
		`{"type":"content_block_start","content_block":{"type":"tool_use","name":"Bash"}}`,
		`{"type":"content_block_start","content_block":{"type":"text","text":""}}`,
		`{"type":"content_block_delta","delta":{"type":"text_delta","text":"Step 2"}}`,
		`{"type":"content_block_start","content_block":{"type":"tool_use","name":"Read"}}`,
		`{"type":"content_block_start","content_block":{"type":"text","text":""}}`,
		`{"type":"content_block_delta","delta":{"type":"text_delta","text":"Step 3"}}`,
		`{"type":"result","session_id":"sess-2"}`,
	}, "\n")

	m := &Manager{
		sessions:    make(map[string]*Session),
		processes:   make(map[string]*runningProcess),
		subscribers: make(map[string][]chan struct{}),
	}

	session := &Session{
		ID:           "bean-multi",
		AgentType:    "claude",
		Status:       StatusRunning,
		Messages:     []Message{{Role: RoleUser, Content: "do stuff"}},
		streamingIdx: -1,
	}
	m.sessions["bean-multi"] = session

	m.readOutput("bean-multi", strings.NewReader(lines), "")

	// Expected: USER, ASSISTANT(Step 1), TOOL(Bash), ASSISTANT(Step 2), TOOL(Read), ASSISTANT(Step 3)
	msgs := session.Messages
	if len(msgs) != 6 {
		t.Fatalf("expected 6 messages, got %d", len(msgs))
	}

	expected := []struct {
		role    MessageRole
		contain string
	}{
		{RoleUser, "do stuff"},
		{RoleAssistant, "Step 1"},
		{RoleTool, "Bash"},
		{RoleAssistant, "Step 2"},
		{RoleTool, "Read"},
		{RoleAssistant, "Step 3"},
	}

	for i, tt := range expected {
		if msgs[i].Role != tt.role {
			t.Errorf("msgs[%d].Role = %q, want %q", i, msgs[i].Role, tt.role)
		}
		if !strings.Contains(msgs[i].Content, tt.contain) {
			t.Errorf("msgs[%d].Content = %q, want it to contain %q", i, msgs[i].Content, tt.contain)
		}
	}
}
