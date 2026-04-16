import pytest
import time
from sensible.decorators import (
    cached, lru_cache, validate_args, count_calls, memoize, retry, singleton, timed
)

# Decorator Tests

def test_cached():
    call_count = 0
    
    @cached(max_size=10)
    def expensive_func(x):
        nonlocal call_count
        call_count += 1
        return x * 2
    
    result1 = expensive_func(5)
    result2 = expensive_func(5)
    
    assert result1 == result2 == 10
    assert call_count == 1, f"Expected 1 call, got {call_count}"  # Should be cached

def test_lru_cache():
    call_count = 0
    
    @lru_cache(max_size=10)
    def expensive_func(x):
        nonlocal call_count
        call_count += 1
        return x * 3
    
    result1 = expensive_func(5)
    result2 = expensive_func(5)
    
    assert result1 == result2 == 15
    assert call_count == 1  # Should be cached

def test_validate_args():
    @validate_args(int, int)
    def add(x, y):
        return x + y
    
    result = add(5, 10)
    assert result == 15
    
    with pytest.raises(TypeError, match="Argument 1 of add must be int"):
        add("5", 10)

def test_validate_args_no_types_specified():
    @validate_args()
    def func():
        return "ok"
    
    assert func() == "ok"

def test_count_calls():
    @count_calls
    def my_func():
        pass
    
    my_func()
    my_func()
    my_func()
    
    assert my_func.calls == 3

def test_memoize():
    call_count = 0
    
    @memoize
    def fib(n):
        nonlocal call_count
        call_count += 1
        if n <= 1:
            return n
        return fib(n - 1) + fib(n - 2)
    
    result = fib(10)
    assert result == 55
    assert call_count < 100  # Should be much less than naive recursion due to memoization

def test_validate_args_with_kwargs():
    @validate_args(int)
    def func(x, y=10):
        return x + y
    
    result = func(5)
    assert result == 15
    
    result = func(5, y=20)
    assert result == 25

def test_validate_args_optional_type():
    @validate_args(int, int)
    def func(x, y):
        return x + y
    
    with pytest.raises(TypeError):
        func("not an int", 5)
    
    with pytest.raises(TypeError):
        func(5, "not an int")

def test_validate_args_skip_first_arg_for_class_method():
    class MyClass:
        @validate_args(int)
        def method(self, x):
            return x * 2
    
    obj = MyClass()
    result = obj.method(5)
    assert result == 10