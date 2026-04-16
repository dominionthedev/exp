from sensible.strings import slugify, truncate, to_snake, to_camel, to_pascal, to_kebab

def test_slugify():
    assert slugify("Hello World!") == "hello-world"
    assert slugify("Café au Lait") == "cafe-au-lait"
    assert slugify("Some_Cool_Text", separator="_") == "some_cool_text"

def test_truncate():
    assert truncate("Hello World", 5) == "Hello..."
    assert truncate("Hello World", 20) == "Hello World"
    assert truncate("Hello World", 5, suffix="") == "Hello"

def test_case_conversions():
    text = "hello world_this-is PascalCase"
    assert to_snake(text) == "hello_world_this_is_pascal_case"
    assert to_kebab(text) == "hello-world-this-is-pascal-case"
    assert to_pascal("hello world") == "HelloWorld"
    assert to_camel("hello world") == "helloWorld"
