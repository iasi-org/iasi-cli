package cli

import (
	"strings"
	"testing"
	"time"
)

func TestFormatMessage(t *testing.T) {
	now := time.Date(2026, 8, 29, 14, 37, 22, 0, time.Local)

	got := formatMessage(now, levelSuccess, "Adapter %q installed", "codex")
	want := "14:37:22 - \033[32mAdapter \"codex\" installed\033[0m\n"

	if got != want {
		t.Fatalf("formatMessage() = %q, want %q", got, want)
	}
}

func TestMessageColors(t *testing.T) {
	tests := []struct {
		name  string
		level messageLevel
		color string
	}{
		{name: "debug", level: levelDebug, color: colorGray},
		{name: "info", level: levelInfo, color: colorBlue},
		{name: "success", level: levelSuccess, color: colorGreen},
		{name: "warning", level: levelWarning, color: colorYellow},
		{name: "error", level: levelError, color: colorRed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatMessage(time.Date(2026, 8, 29, 14, 37, 22, 0, time.Local), tt.level, "message")
			if !strings.Contains(got, tt.color+"message"+colorReset) {
				t.Fatalf("formatMessage() = %q, expected color %q", got, tt.color)
			}
		})
	}
}

func TestMessageVisibility(t *testing.T) {
	tests := []struct {
		name       string
		options    OutputOptions
		visibility messageVisibility
		want       bool
	}{
		{name: "normal", options: OutputOptions{}, visibility: visibilityNormal, want: true},
		{name: "verbose hidden by default", options: OutputOptions{}, visibility: visibilityVerbose, want: false},
		{name: "debug hidden by default", options: OutputOptions{}, visibility: visibilityDebug, want: false},
		{name: "verbose adds verbose", options: OutputOptions{Verbose: true}, visibility: visibilityVerbose, want: true},
		{name: "verbose keeps normal", options: OutputOptions{Verbose: true}, visibility: visibilityNormal, want: true},
		{name: "debug adds debug", options: OutputOptions{Debug: true}, visibility: visibilityDebug, want: true},
		{name: "debug keeps normal", options: OutputOptions{Debug: true}, visibility: visibilityNormal, want: true},
		{name: "debug does not add verbose", options: OutputOptions{Debug: true}, visibility: visibilityVerbose, want: false},
		{name: "verbose and debug", options: OutputOptions{Verbose: true, Debug: true}, visibility: visibilityVerbose, want: true},
		{name: "silent suppresses normal", options: OutputOptions{Silent: true}, visibility: visibilityNormal, want: false},
		{name: "silent suppresses verbose", options: OutputOptions{Silent: true, Verbose: true}, visibility: visibilityVerbose, want: false},
		{name: "silent suppresses debug", options: OutputOptions{Silent: true, Debug: true}, visibility: visibilityDebug, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigureOutput(tt.options)
			if got := messageVisible(tt.visibility); got != tt.want {
				t.Fatalf("messageVisible() = %t, want %t", got, tt.want)
			}
		})
	}

	ConfigureOutput(OutputOptions{})
}
