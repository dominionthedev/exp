import pytest
from sensible.types import Result, Option

# Result Tests

def test_result_map():
    # Test mapping on Ok
    r = Result.ok(5)
    result = r.map(lambda x: x * 2)
    assert result.is_ok
    assert result.unwrap() == 10
    
    # Test mapping on Error (should not apply function)
    r = Result.error("error message")
    result = r.map(lambda x: x * 2)
    assert result.is_error
    assert result.error_value() == "error message"

def test_result_map_error():
    # Test mapping error on Ok (should not apply function)
    r = Result.ok(5)
    result = r.map_error(lambda e: "new error")
    assert result.is_ok
    assert result.unwrap() == 5
    
    # Test mapping on Error
    r = Result.error("old error")
    result = r.map_error(lambda e: f"new: {e}")
    assert result.is_error
    assert result.error_value() == "new: old error"

def test_result_and_then():
    # Test and_then on Ok
    r = Result.ok(5)
    result = r.and_then(lambda x: Result.ok(x * 2))
    assert result.is_ok
    assert result.unwrap() == 10
    
    # Test and_then on Error (should skip function)
    r = Result.error("error")
    result = r.and_then(lambda x: Result.ok(x * 2))
    assert result.is_error
    assert result.error_value() == "error"
    
    # Test and_then with error in chain
    r = Result.ok(5)
    result = r.and_then(lambda x: Result.error("chain error"))
    assert result.is_error
    assert result.error_value() == "chain error"

def test_result_unwrap_or_else():
    # Test on Ok (should not apply function)
    r = Result.ok(10)
    result = r.unwrap_or_else(lambda e: 0)
    assert result == 10
    
    # Test on Error (should apply function)
    r = Result.error("error occurred")
    result = r.unwrap_or_else(lambda e: len(e))
    assert result == 14

def test_result_str_repr():
    r1 = Result.ok("value")
    assert str(r1) == "Ok(value)"
    assert repr(r1) == "Ok(value)"
    
    r2 = Result.error("err")
    assert str(r2) == "Error(err)"
    assert repr(r2) == "Error(err)"

def test_option_map():
    # Test mapping on Some
    o = Option.some(5)
    result = o.map(lambda x: x * 3)
    assert result.is_some
    assert result.unwrap() == 15
    
    # Test mapping on None
    o = Option.none()
    result = o.map(lambda x: x * 3)
    assert result.is_none

def test_option_map_or():
    # Test on Some
    o = Option.some(10)
    result = o.map_or(0, lambda x: x + 5)
    assert result == 15
    
    # Test on None
    o = Option.none()
    result = o.map_or(20, lambda x: x + 5)
    assert result == 20

def test_option_filter():
    # Test filtering Some with matching predicate
    o = Option.some(5)
    result = o.filter(lambda x: x > 3)
    assert result.is_some
    assert result.unwrap() == 5
    
    # Test filtering Some with non-matching predicate
    o = Option.some(5)
    result = o.filter(lambda x: x > 10)
    assert result.is_none
    
    # Test filtering None
    o = Option.none()
    result = o.filter(lambda x: x > 3)
    assert result.is_none

def test_option_and_then():
    # Test and_then on Some
    o = Option.some(5)
    result = o.and_then(lambda x: Option.some(x * 2))
    assert result.is_some
    assert result.unwrap() == 10
    
    # Test and_then on None
    o = Option.none()
    result = o.and_then(lambda x: Option.some(x * 2))
    assert result.is_none
    
    # Test and_then with None in chain
    o = Option.some(5)
    result = o.and_then(lambda x: Option.none())
    assert result.is_none

def test_option_or_else():
    # Test or_else on Some (should not apply function)
    o = Option.some(10)
    result = o.or_else(lambda: Option.some(20))
    assert result.is_some
    assert result.unwrap() == 10
    
    # Test or_else on None (should apply function)
    o = Option.none()
    result = o.or_else(lambda: Option.some(30))
    assert result.is_some
    assert result.unwrap() == 30

def test_option_unwrap_or_else():
    # Test on Some (should not apply function)
    o = Option.some(10)
    result = o.unwrap_or_else(lambda: 0)
    assert result == 10
    
    # Test on None (should apply function)
    o = Option.none()
    result = o.unwrap_or_else(lambda: 42)
    assert result == 42

def test_option_str_repr():
    o1 = Option.some("value")
    assert str(o1) == "Some(value)"
    assert repr(o1) == "Some(value)"
    
    o2 = Option.none()
    assert str(o2) == "None"
    assert repr(o2) == "None"