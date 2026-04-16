import time
import functools
import operator
from typing import Callable, Any, Type, Tuple, Union, Dict

def timed(func: Callable) -> Callable:
    """Decorator that prints the execution time of the function."""
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        result = func(*args, **kwargs)
        end = time.perf_counter()
        print(f"Function {func.__name__} took {end - start:.4f}s")
        return result
    return wrapper

def retry(
    retries: int = 3,
    delay: float = 1.0,
    exceptions: Union[Type[Exception], Tuple[Type[Exception], ...]] = Exception
) -> Callable:
    """Decorator that retries a function if it raises an exception."""
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            last_exception = None
            for i in range(retries):
                try:
                    return func(*args, **kwargs)
                except exceptions as e:
                    last_exception = e
                    if i < retries - 1:
                        time.sleep(delay)
            raise last_exception
        return wrapper
    return decorator

def singleton(cls: Type) -> Type:
    """Decorator that turns a class into a singleton."""
    instances = {}
    @functools.wraps(cls)
    def wrapper(*args, **kwargs):
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]
    return wrapper

def cached(max_size: int = 128) -> Callable:
    """Decorator that caches function results in memory."""
    cache: Dict[tuple, Any] = {}
    order: list = []
    
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            key = (args, tuple(sorted(kwargs.items())))
            if key in cache:
                return cache[key]
            result = func(*args, **kwargs)
            cache[key] = result
            order.append(key)
            if len(order) > max_size:
                del cache[order.pop(0)]
            return result
        return wrapper
    return decorator

def lru_cache(max_size: int = 128) -> Callable:
    """Decorator that caches function results using LRU (Least Recently Used) eviction."""
    cache: Dict[tuple, Any] = {}
    usage_order: list = []
    
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            key = (args, tuple(sorted(kwargs.items())))
            if key in cache:
                # Move to most recently used
                usage_order.remove(key)
                usage_order.append(key)
                return cache[key]
            result = func(*args, **kwargs)
            cache[key] = result
            usage_order.append(key)
            if len(usage_order) > max_size:
                oldest_key = usage_order.pop(0)
                del cache[oldest_key]
            return result
        return wrapper
    return decorator

def validate_args(*arg_types: Type) -> Callable:
    """Decorator that validates function arguments are of the correct type."""
    def decorator(func: Callable) -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            # Check positional args against provided types
            # Skip first arg if it's a class instance (self) - check by type name
            start_idx = 0
            if args and hasattr(args[0], '__class__') and not isinstance(args[0], type):
                # First arg might be self, skip it if no type specified for it
                # Only skip if we have arg_types specified
                if len(arg_types) < len(args):
                    # Check if first arg looks like self (has __dict__ and is not a type)
                    if hasattr(args[0], '__dict__'):
                        start_idx = 1
            
            for i, (arg, expected_type) in enumerate(zip(args[start_idx:], arg_types)):
                if not isinstance(arg, expected_type):
                    raise TypeError(
                        f"Argument {i+1} of {func.__name__} must be {expected_type.__name__}, "
                        f"got {type(arg).__name__}"
                    )
            return func(*args, **kwargs)
        return wrapper
    return decorator

def count_calls(func: Callable) -> Callable:
    """Decorator that counts the number of times a function is called."""
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        wrapper.calls += 1
        return func(*args, **kwargs)
    wrapper.calls = 0
    return wrapper

def memoize(func: Callable) -> Callable:
    """Decorator that memoizes function results (infinite cache)."""
    cache: Dict[tuple, Any] = {}
    
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        key = (args, tuple(sorted(kwargs.items())))
        if key in cache:
            return cache[key]
        result = func(*args, **kwargs)
        cache[key] = result
        return result
    return wrapper
