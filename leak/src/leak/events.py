import asyncio
from typing import Callable, Any, Dict, Set, Optional
import logging

logger = logging.getLogger(__name__)


class Event:
    """Base class for all leak events."""

    def __init__(self, name: str, data: Any = None):
        self.name = name
        self.data = data
        self.timestamp = asyncio.get_event_loop().time()


class EventBus:
    """An asyncio-based event dispatcher."""

    def __init__(self):
        self._subscribers: Dict[str, Set[Callable]] = {}
        self._queue: asyncio.Queue = asyncio.Queue()
        self._running = False
        self._task: Optional[asyncio.Task] = None

    def subscribe(self, event_name: str, callback: Callable):
        """Subscribe to a specific event topic."""
        if event_name not in self._subscribers:
            self._subscribers[event_name] = set()
        self._subscribers[event_name].add(callback)

    def unsubscribe(self, event_name: str, callback: Callable):
        """Unsubscribe from an event topic."""
        if event_name in self._subscribers:
            self._subscribers[event_name].discard(callback)

    async def emit(self, event: Event):
        """Put an event into the queue to be dispatched."""
        await self._queue.put(event)

    def emit_sync(self, event: Event):
        """Synchronously put an event into the queue (thread-safe)."""
        self._queue.put_nowait(event)

    async def start(self):
        """Start the background event loop."""
        if self._running:
            return
        self._running = True
        self._task = asyncio.create_task(self._worker())

    async def stop(self):
        """Stop the event loop and wait for remaining events."""
        self._running = False
        if self._task:
            await self._queue.put(None)  # Sentinel to stop worker
            await self._task

    async def _worker(self):
        while self._running or not self._queue.empty():
            event = await self._queue.get()
            if event is None:
                break

            # Dispatch to subscribers
            handlers = self._subscribers.get(event.name, set())
            for handler in handlers:
                try:
                    if asyncio.iscoroutinefunction(handler):
                        await handler(event)
                    else:
                        handler(event)
                except Exception as e:
                    logger.error(f"Error in event handler for {event.name}: {e}")

            self._queue.task_done()
