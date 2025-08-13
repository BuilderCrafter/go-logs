// Package helpers provides small CLI utilities for terminal-based apps:
// clearing the screen, printing a welcome banner, reading a line of input,
// pausing for Enter, wrapping text to a fixed width, and parsing simple
// command/menu inputs.
package helpers

/*
# Provided utilities

  • ClearScreen()
    Clears the terminal using ANSI escape sequences and moves the cursor to the
    home position. If the terminal does not support ANSI or the write fails,
    the function returns silently without error.

  • WelcomeMessage()
    Prints a short description of the application and basic usage hints.

  • WrapText(s string, limit int) string
    Wraps s on word boundaries so that lines do not exceed limit characters,
    except when a single word is longer than limit (no hyphenation). All
    whitespace in s is normalized (multiple spaces/tabs/newlines collapse into
    single spaces). If limit <= 0, s is returned unchanged.

  • ReadLine() string
    Reads a single line from standard input and trims the trailing "\r\n" or
    "\n". I/O errors are ignored; the function returns whatever was read up to
    the newline (if anything), with line endings removed.

  • WaitForEnter(prompt string)
    Prints prompt (or a default "-Press Enter to continue-" when prompt is
    empty) and then blocks until the user presses Enter. Any characters typed on
    that line are discarded.

# Parsing helpers

  • LogParseInput(s string) (string, bool)
    Validates actions in the "open log" view. Accepts "x", "e", and "r"
    case-insensitively ("x"=delete, "e"=edit, "r"=return). Returns the original
    input string and true on success, or "" and false on failure. Callers that
    need a normalized value should lowercase the returned string.

  • MenuParseInput(s string) (num int, cmd string)
    Distinguishes between a positive menu number and a non-numeric command.
    On trimmed, strictly positive integers (e.g., "1", "42"), returns (n, "").
    Otherwise returns (0, trimmedInput). The command string is not lowercased;
    callers may normalize it if needed.

# I/O and errors

All output is written to os.Stdout. Input is read from a buffered reader over
os.Stdin. The helpers ignore I/O errors by design (suitable for simple CLIs);
callers that need stricter error handling should wrap or replace these functions.

# Concurrency

These helpers are intended for single-goroutine, interactive use. The package
does not synchronize access to its internal input reader. If used concurrently,
callers must provide their own synchronization.

# Portability

ClearScreen relies on ANSI escape sequences, which are widely supported on
Unix-like terminals and modern Windows terminals/WSL. If broader compatibility
is required, consider adding platform-specific fallbacks.
*/
