package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/creack/pty"
)

// Session represents an active PTY session.
type Session struct {
	id   string
	cmd  *exec.Cmd
	ptyF *os.File // PTY master file descriptor
	mu   sync.Mutex
}

// Read reads PTY output into buf.
func (s *Session) Read(buf []byte) (int, error) {
	return s.ptyF.Read(buf)
}

// Write sends input to the PTY.
func (s *Session) Write(data []byte) (int, error) {
	return s.ptyF.Write(data)
}

// Resize changes the PTY window size.
func (s *Session) Resize(cols, rows uint16) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return pty.Setsize(s.ptyF, &pty.Winsize{Cols: cols, Rows: rows})
}

// Close kills the process and closes the PTY.
func (s *Session) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cmd.Process != nil {
		_ = s.cmd.Process.Kill()
	}
	_ = s.ptyF.Close()
	_ = s.cmd.Wait()
}

// Manager manages PTY sessions keyed by session ID.
type Manager struct {
	mu       sync.Mutex
	sessions map[string]*Session
}

// NewManager creates a new terminal session manager.
func NewManager() *Manager {
	return &Manager{sessions: make(map[string]*Session)}
}

// Create spawns a new PTY session in the given working directory.
func (m *Manager) Create(sessionID, workDir string, cols, rows uint16) (*Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Close existing session with same ID
	if existing, ok := m.sessions[sessionID]; ok {
		existing.Close()
		delete(m.sessions, sessionID)
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	cmd := exec.Command(shell, "-l")
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	ptmx, err := pty.StartWithSize(cmd, &pty.Winsize{Cols: cols, Rows: rows})
	if err != nil {
		return nil, fmt.Errorf("failed to start PTY: %w", err)
	}

	sess := &Session{
		id:   sessionID,
		cmd:  cmd,
		ptyF: ptmx,
	}

	m.sessions[sessionID] = sess
	return sess, nil
}

// Get retrieves an existing session.
func (m *Manager) Get(sessionID string) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sessions[sessionID]
}

// Close closes and removes a specific session.
func (m *Manager) Close(sessionID string) {
	m.mu.Lock()
	sess, ok := m.sessions[sessionID]
	if ok {
		delete(m.sessions, sessionID)
	}
	m.mu.Unlock()
	if ok {
		sess.Close()
	}
}

// Shutdown closes all active sessions.
func (m *Manager) Shutdown() {
	m.mu.Lock()
	sessions := make(map[string]*Session)
	for k, v := range m.sessions {
		sessions[k] = v
	}
	m.sessions = make(map[string]*Session)
	m.mu.Unlock()

	for _, s := range sessions {
		s.Close()
	}
}
