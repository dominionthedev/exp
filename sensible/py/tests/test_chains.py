from sensible.chains import chain

def test_chain_then():
    res = chain(10).then(lambda x: x + 5).then(lambda x: x * 2).value()
    assert res == 30

def test_chain_pipe():
    def add_5(x): return x + 5
    def mult_2(x): return x * 2
    res = chain(10).pipe(add_5, mult_2).value()
    assert res == 30

def test_chain_complex():
    res = (
        chain("  hello  ")
        .then(str.strip)
        .then(str.upper)
        .then(lambda s: f"[{s}]")
        .value()
    )
    assert res == "[HELLO]"
