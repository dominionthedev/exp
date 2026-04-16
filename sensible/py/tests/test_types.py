import pytest
from sensible.types import Result, Option

def test_result_ok():
    r = Result.ok(10)
    assert r.is_ok
    assert not r.is_error
    assert r.unwrap() == 10
    assert r.unwrap_or(0) == 10

def test_result_error():
    r = Result.error("something went wrong")
    assert r.is_error
    assert not r.is_ok
    assert r.error_value() == "something went wrong"
    assert r.unwrap_or(0) == 0
    with pytest.raises(ValueError, match="called unwrap on an error result"):
        r.unwrap()

def test_option_some():
    o = Option.some("hello")
    assert o.is_some
    assert not o.is_none
    assert o.unwrap() == "hello"
    assert o.unwrap_or("default") == "hello"

def test_option_none():
    o = Option.none()
    assert o.is_none
    assert not o.is_some
    assert o.unwrap_or("default") == "default"
    with pytest.raises(ValueError, match="called unwrap on a None option"):
        o.unwrap()
