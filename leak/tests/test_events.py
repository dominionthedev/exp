import asyncio
import pytest
from leak.events import EventBus, Event

@pytest.mark.asyncio
async def test_event_bus_dispatch():
    bus = EventBus()
    received_data = []

    async def callback(event):
        received_data.append(event.data)

    bus.subscribe("test_event", callback)
    await bus.start()
    
    await bus.emit(Event("test_event", "hello"))
    await asyncio.sleep(0.1)  # Give worker time to process
    
    await bus.stop()
    
    assert received_data == ["hello"]

@pytest.mark.asyncio
async def test_event_bus_unsubscribe():
    bus = EventBus()
    received_count = 0

    def callback(event):
        nonlocal received_count
        received_count += 1

    bus.subscribe("test", callback)
    bus.unsubscribe("test", callback)
    
    await bus.start()
    await bus.emit(Event("test"))
    await asyncio.sleep(0.1)
    await bus.stop()
    
    assert received_count == 0
