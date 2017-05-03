package logger

import "testing"

func TestErrorWithFalse(t *testing.T) {
	log := New()
	if log.Error("") != false {
		t.Errorf("expected to get false.")
	}
}

func TestErrorWithTrue(t *testing.T) {
	log := New()
	if log.Error("Error Message") != true {
		t.Errorf("expected to get true.")
	}
}

func TestWarningWithFalse(t *testing.T) {
	log := New()
	if log.Warning("") != false {
		t.Errorf("expected to get false.")
	}
}

func TestWarningWithTrue(t *testing.T) {
	log := New()
	if log.Warning("Error Message") != true {
		t.Errorf("expected to get true.")
	}
}

func TestNoticeWithFalse(t *testing.T) {
	log := New()
	if log.Notice("") != false {
		t.Errorf("expected to get false.")
	}
}

func TestNoticeWithTrue(t *testing.T) {
	log := New()
	if log.Notice("Error Message") != true {
		t.Errorf("expected to get true.")
	}
}

func TestSuccessWithFalse(t *testing.T) {
	log := New()
	if log.Success("") != false {
		t.Errorf("expected to get false.")
	}
}

func TestSuccessWithTrue(t *testing.T) {
	log := New()
	if log.Success("Error Message") != true {
		t.Errorf("expected to get true.")
	}
}
