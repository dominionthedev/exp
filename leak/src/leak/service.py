import asyncio
import signal
from typing import Optional
from .events import EventBus, Event
from .terminal import Terminal

class Service:
    """Main lifecycle and signal management for terminal services."""
    def __init__(self, event_bus: Optional[EventBus] = None):
        self.events = event_bus or EventBus()
        self._running = False
        self._loop = asyncio.get_event_loop()

    async def start(self):
        """Initialize the service, start event bus, and trap signals."""
        if self._running:
            return
        
        self._running = True
        await self.events.start()
        self._setup_signals()
        
        await self.events.emit(Event("service_start"))

    async def stop(self):
        """Gracefully stop the service and cleanup."""
        if not self._running:
            return
            
        await self.events.emit(Event("service_stop"))
        await self.events.stop()
        self._running = False

    def _setup_signals(self):
        """Trap OS signals and convert them to internal events."""
        for sig in (signal.SIGINT, signal.SIGTERM):
            try:
                self._loop.add_signal_handler(
                    sig, 
                    lambda s=sig: self.events.emit_sync(Event("signal_stop", {"signal": s}))
                )
            except NotImplementedError:
                # Fallback for platforms without signal support in asyncio loop
                pass

        # Handle terminal resize
        try:
            self._loop.add_signal_handler(
                signal.SIGWINCH,
                lambda: self.events.emit_sync(Event("terminal_resize", {"size": Terminal.get_size()}))
            )
        except (NotImplementedError, AttributeError):
            pass

    async def run_forever(self):
        """Helper to run the service until a stop event is received."""
        stop_event = asyncio.Event()
        
        async def handle_stop(e):
            stop_event.set()

        self.events.subscribe("signal_stop", handle_stop)
        
        await self.start()
        try:
            await stop_event.wait()
        finally:
            await self.stop()
