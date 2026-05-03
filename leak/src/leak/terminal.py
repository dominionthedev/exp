import sys
import tty
import termios
import contextlib
import os

class Terminal:
    """Low-level terminal utilities and state management."""
    
    @staticmethod
    @contextlib.contextmanager
    def raw_mode():
        """Context manager to put the terminal into raw mode and restore it on exit."""
        if not sys.stdin.isatty():
            yield
            return

        fd = sys.stdin.fileno()
        old_settings = termios.tcgetattr(fd)
        try:
            tty.setraw(fd)
            yield
        finally:
            termios.tcsetattr(fd, termios.TCSADRAIN, old_settings)

    @staticmethod
    def hide_cursor():
        """Hide the terminal cursor."""
        sys.stdout.write("\033[?25l")
        sys.stdout.flush()

    @staticmethod
    def show_cursor():
        """Show the terminal cursor."""
        sys.stdout.write("\033[?25h")
        sys.stdout.flush()

    @staticmethod
    def get_size():
        """Get the current terminal size (rows, cols)."""
        size = os.get_terminal_size()
        return size.lines, size.columns

    @staticmethod
    def clear():
        """Clear the terminal screen."""
        sys.stdout.write("\033[2J\033[H")
        sys.stdout.flush()

    @staticmethod
    def move_to(row: int, col: int):
        """Move the cursor to a specific position."""
        sys.stdout.write(f"\033[{row};{col}H")
        sys.stdout.flush()
