import asyncio
import pytest
from leak.service import Service


@pytest.mark.asyncio
async def test_service_lifecycle():
    service = Service()
    start_received = False
    stop_received = False

    async def on_start(e):
        nonlocal start_received
        start_received = True

    async def on_stop(e):
        nonlocal stop_received
        stop_received = True

    service.events.subscribe("service_start", on_start)
    service.events.subscribe("service_stop", on_stop)

    await service.start()
    await asyncio.sleep(0.1)  # Give event loop time to dispatch
    assert start_received is True

    await service.stop()
    await asyncio.sleep(0.1)
    assert stop_received is True
