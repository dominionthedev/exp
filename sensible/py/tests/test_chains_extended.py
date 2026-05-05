from sensible.chains import chain, chain_from_iterable

# Chain Tests

def test_chain_map():
    result = chain(5).map(lambda x: x * 3).value()
    assert result == 15

def test_chain_filter():
    # Test filter passes when predicate is True
    result = chain(5).filter(lambda x: x > 3).value()
    assert result == 5
    
    # Test filter returns None when predicate is False
    result = chain(5).filter(lambda x: x > 10).value()
    assert result is None

def test_chain_tap():
    tap_result = None
    def capture(x):
        nonlocal tap_result
        tap_result = x * 2
    
    result = chain(5).tap(capture).value()
    assert result == 5  # Tap returns original value
    assert tap_result == 10  # But side effect was executed

def test_chain_if_else_true():
    result = chain(5).if_else(True, lambda x: x * 2).value()
    assert result == 10

def test_chain_if_else_false():
    result = chain(5).if_else(False, else_func=lambda x: x * 3).value()
    assert result == 15

def test_chain_if_else_no_functions():
    # Should return self when no functions provided
    result = chain(5).if_else(True).value()
    assert result == 5

def test_chain_if_else_both_functions():
    result = chain(5).if_else(True, lambda x: x + 1, lambda x: x - 1).value()
    assert result == 6  # then_func applied

def test_chain_collect():
    result = chain(42).collect()
    assert result == 42

def test_chain_if_none():
    result = chain(None).if_none(10).value()
    assert result == 10

def test_chain_or_else():
    result = chain(None).or_else(20).value()
    assert result == 20

def test_chain_value():
    result = chain(100).value()
    assert result == 100

def test_chain_chain_from_iterable():
    result = chain_from_iterable([1, 2, 3]).value()
    assert result == [1, 2, 3]

def test_complex_chain_pipeline():
    result = (
        chain([1, 2, 3, 4, 5])
        .then(sum)
        .map(lambda x: x * 2)
        .filter(lambda x: x > 10)
        .value()
    )
    assert result == 30

def test_chain_with_optional_values():
    def get_optional() -> int | None:
        return None
    
    result = chain(get_optional()).if_none(0).value()
    assert result == 0

def test_chain_conditional_processing():
    data = None
    result = (
        chain(data)
        .if_none("default")
        .then(str.upper)
        .value()
    )
    assert result == "DEFAULT"

def test_chain_fold():
    result = chain([1, 2, 3, 4, 5]).fold(lambda x, y: x + y).value()
    assert result == 15

    result = chain([1, 2, 3]).fold(lambda x, y: x * y, initial=10).value()
    assert result == 60

def test_chain_fold_error():
    import pytest
    with pytest.raises(TypeError, match=r"fold\(\) requires an iterable value"):
        chain(5).fold(lambda x, y: x + y)

def test_chain_is_none():
    assert chain(None).is_none() is True
    assert chain(10).is_none() is False

def test_chain_is_not_none():
    assert chain(None).is_not_none() is False
    assert chain(10).is_not_none() is True