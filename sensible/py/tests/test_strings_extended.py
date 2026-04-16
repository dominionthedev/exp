from sensible.strings import (
    capitalize, lower, upper, swap_case, reverse, repeat,
    pad_left, pad_right, pad,
    strip_prefix, strip_suffix,
    pluralize, singularize,
    starts_with, ends_with, contains, count_occurrences,
    random_string, is_palindrome,
    mask_text, word_count, sentence_count,
    first_word, last_word
)

# String Tests

def test_capitalize():
    assert capitalize("hello") == "Hello"
    assert capitalize("HELLO") == "HELLO"
    assert capitalize("") == ""
    assert capitalize("a") == "A"

def test_lower():
    assert lower("HELLO") == "hello"
    assert lower("HeLLo") == "hello"
    assert lower("hello") == "hello"

def test_upper():
    assert upper("hello") == "HELLO"
    assert upper("HeLLo") == "HELLO"
    assert upper("HELLO") == "HELLO"

def test_swap_case():
    assert swap_case("Hello") == "hELLO"
    assert swap_case("HeLLo") == "hEllO"
    assert swap_case("hello") == "HELLO"
    assert swap_case("HELLO") == "hello"

def test_reverse():
    assert reverse("hello") == "olleh"
    assert reverse("a") == "a"
    assert reverse("") == ""

def test_repeat():
    assert repeat("ha", 3) == "hahaha"
    assert repeat("ha", 0) == ""
    assert repeat("", 5) == ""

def test_pad_left():
    assert pad_left("hi", 5, "0") == "000hi"
    assert pad_left("hello", 3, "0") == "hello"
    assert pad_left("hi", 5) == "   hi"

def test_pad_right():
    assert pad_right("hi", 5, "0") == "hi000"
    assert pad_right("hello", 3, "0") == "hello"
    assert pad_right("hi", 5) == "hi   "

def test_pad():
    assert pad("hi", 6, "0") == "00hi00"
    # 5 chars total with 2 char word = 3 padding chars (1 left, 2 right)
    assert pad("hi", 5, "0") == "0hi00"
    assert pad("hello", 3, "0") == "hello"
    assert pad("hi", 6) == "  hi  "

def test_strip_prefix():
    assert strip_prefix("hello", "he") == "llo"
    assert strip_prefix("hello", "world") == "hello"
    assert strip_prefix("hello", "") == "hello"

def test_strip_suffix():
    assert strip_suffix("hello", "lo") == "hel"
    assert strip_suffix("hello", "world") == "hello"
    assert strip_suffix("hello", "") == "hello"

def test_plurilize():
    assert pluralize("cat", 2) == "cats"
    assert pluralize("cat", 1) == "cat"
    assert pluralize("bus", 3) == "buses"
    assert pluralize("city", 2) == "cities"
    assert pluralize("box", 2) == "boxes"

def test_plurilize_custom():
    assert pluralize("mouse", 2, custom="mice") == "mice"
    assert pluralize("sheep", 2) == "sheep"  # Uncountable noun

def test_singularize():
    assert singularize("cats") == "cat"
    assert singularize("cities") == "city"
    assert singularize("boxes") == "box"
    assert singularize("sheep") == "sheep"  # Uncountable noun

def test_starts_with():
    assert starts_with("hello", "he") == True
    assert starts_with("hello", "world") == False
    assert starts_with("hello", "") == True

def test_ends_with():
    assert ends_with("hello", "lo") == True
    assert ends_with("hello", "world") == False
    assert ends_with("hello", "") == True

def test_contains():
    assert contains("hello world", "world") == True
    assert contains("hello world", "WORLD") == False
    assert contains("hello", "xyz") == False

def test_count_occurrences():
    assert count_occurrences("hello hello hello", "hello") == 3
    assert count_occurrences("aaa", "a") == 3
    assert count_occurrences("abc", "xyz") == 0

def test_random_string():
    s1 = random_string(8)
    assert len(s1) == 8
    assert s1.isalnum()
    
    s2 = random_string(10, chars="abc")
    assert len(s2) == 10
    assert all(c in "abc" for c in s2)

def test_is_palindrome():
    assert is_palindrome("racecar") == True
    assert is_palindrome("RaceCar") == True
    assert is_palindrome("A man a plan a canal Panama") == True
    assert is_palindrome("hello") == False

def test_mask_text():
    assert mask_text("12345678", visible_start=2, visible_end=2) == "12****78"
    assert mask_text("12345678", visible_start=4) == "1234****"
    assert mask_text("12345678", visible_end=4) == "****5678"
    assert mask_text("hi", visible_start=1, visible_end=1) == "hi"  # Too short

def test_word_count():
    assert word_count("hello world") == 2
    assert word_count("one") == 1
    assert word_count("") == 0
    assert word_count("   ") == 0

def test_sentence_count():
    assert sentence_count("Hello. World!") == 2
    assert sentence_count("One sentence.") == 1
    assert sentence_count("No ending punctuation") == 0

def test_first_word():
    assert first_word("hello world") == "hello"
    assert first_word("  hello  ") == "hello"
    assert first_word("") is None

def test_last_word():
    assert last_word("hello world") == "world"
    assert last_word("hello  ") == "hello"
    assert last_word("") is None