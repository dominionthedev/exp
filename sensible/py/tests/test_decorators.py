import pytest
from sensible.decorators import retry, singleton, timed

def test_retry():
    attempts = 0
    @retry(retries=3, delay=0.1)
    def fail_twice():
        nonlocal attempts
        attempts += 1
        if attempts < 3:
            raise ValueError("fail")
        return "success"

    assert fail_twice() == "success"
    assert attempts == 3

def test_singleton():
    @singleton
    class MyClass:
        pass

    a = MyClass()
    b = MyClass()
    assert a is b

def test_timed(capsys):
    @timed
    def slow_func():
        return 42
    
    slow_func()
    captured = capsys.readouterr()
    assert "Function slow_func took" in captured.out
