package terminal

import (
	"os"
	"testing"
	"time"
)

func TestManagerCreateAndClose(t *testing.T) {
	mgr := NewManager()
	defer mgr.Shutdown()

	sess, err := mgr.Create("test-1", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if sess == nil {
		t.Fatal("session is nil")
	}

	// Verify we can get the session
	got := mgr.Get("test-1")
	if got != sess {
		t.Fatal("Get returned different session")
	}

	// Close should remove it
	mgr.Close("test-1")
	if mgr.Get("test-1") != nil {
		t.Fatal("session still exists after Close")
	}
}

func TestManagerCreateReplacesExisting(t *testing.T) {
	mgr := NewManager()
	defer mgr.Shutdown()

	sess1, err := mgr.Create("test-replace", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("first Create failed: %v", err)
	}

	sess2, err := mgr.Create("test-replace", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("second Create failed: %v", err)
	}

	if sess1 == sess2 {
		t.Fatal("expected different session objects")
	}

	got := mgr.Get("test-replace")
	if got != sess2 {
		t.Fatal("Get should return the new session")
	}
}

func TestSessionResize(t *testing.T) {
	mgr := NewManager()
	defer mgr.Shutdown()

	sess, err := mgr.Create("test-resize", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if err := sess.Resize(120, 40); err != nil {
		t.Fatalf("Resize failed: %v", err)
	}
}

func TestSessionWriteAndRead(t *testing.T) {
	mgr := NewManager()
	defer mgr.Shutdown()

	sess, err := mgr.Create("test-io", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Write a command to the PTY
	_, err = sess.Write([]byte("echo hello\n"))
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Read some output (should get something back within a reasonable time)
	buf := make([]byte, 4096)
	done := make(chan bool)
	go func() {
		n, err := sess.Read(buf)
		if err != nil {
			t.Errorf("Read failed: %v", err)
		}
		if n == 0 {
			t.Error("Read returned 0 bytes")
		}
		done <- true
	}()

	select {
	case <-done:
		// success
	case <-time.After(5 * time.Second):
		t.Fatal("Read timed out")
	}
}

func TestManagerShutdown(t *testing.T) {
	mgr := NewManager()

	_, err := mgr.Create("test-shutdown-1", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	_, err = mgr.Create("test-shutdown-2", os.TempDir(), 80, 24)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	mgr.Shutdown()

	if mgr.Get("test-shutdown-1") != nil {
		t.Fatal("session 1 still exists after Shutdown")
	}
	if mgr.Get("test-shutdown-2") != nil {
		t.Fatal("session 2 still exists after Shutdown")
	}
}
