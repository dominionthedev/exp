import re
import unicodedata
import random
import string
from typing import Optional, List

def slugify(text: str, separator: str = "-") -> str:
    """Convert text to a URL-friendly slug."""
    text = unicodedata.normalize("NFKD", text).encode("ascii", "ignore").decode("ascii")
    text = re.sub(r"[^\w\s-]", "", text.lower())
    return re.sub(r"[-\s]+", separator, text).strip(separator)

def truncate(text: str, length: int, suffix: str = "...") -> str:
    """Truncate text to a specific length with an optional suffix."""
    if len(text) <= length:
        return text
    return text[:length].rstrip() + suffix

def to_snake(text: str) -> str:
    """Convert text to snake_case."""
    s1 = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", text)
    s2 = re.sub("([a-z0-9])([A-Z])", r"\1_\2", s1).lower()
    return re.sub(r"[-_\s]+", "_", s2).strip("_")

def to_camel(text: str) -> str:
    """Convert text to camelCase."""
    s = to_pascal(text)
    return s[0].lower() + s[1:] if s else ""

def to_pascal(text: str) -> str:
    """Convert text to PascalCase."""
    words = re.split(r"[-_\s]+", text)
    return "".join(word.capitalize() for word in words if word)

def to_kebab(text: str) -> str:
    """Convert text to kebab-case."""
    return to_snake(text).replace("_", "-")

def capitalize(text: str) -> str:
    """Capitalize the first letter of the text."""
    if not text:
        return text
    return text[0].upper() + text[1:]

def lower(text: str) -> str:
    """Convert text to lowercase."""
    return text.lower()

def upper(text: str) -> str:
    """Convert text to uppercase."""
    return text.upper()

def swap_case(text: str) -> str:
    """Swap the case of all letters in the text."""
    return text.swapcase()

def reverse(text: str) -> str:
    """Reverse the text."""
    return text[::-1]

def repeat(text: str, times: int) -> str:
    """Repeat the text a given number of times."""
    return text * times

def pad_left(text: str, length: int, char: str = " ") -> str:
    """Pad text on the left to reach a certain length."""
    if len(text) >= length:
        return text
    return char * (length - len(text)) + text

def pad_right(text: str, length: int, char: str = " ") -> str:
    """Pad text on the right to reach a certain length."""
    if len(text) >= length:
        return text
    return text + char * (length - len(text))

def pad(text: str, length: int, char: str = " ") -> str:
    """Pad text on both sides to reach a certain length (center padding)."""
    if len(text) >= length:
        return text
    padding = length - len(text)
    left = padding // 2
    right = padding - left
    return char * left + text + char * right

def strip_prefix(text: str, prefix: str) -> str:
    """Remove a prefix from the text if it exists."""
    if text.startswith(prefix):
        return text[len(prefix):]
    return text

def strip_suffix(text: str, suffix: str) -> str:
    """Remove a suffix from the text if it exists."""
    if text.endswith(suffix) and suffix:
        return text[:len(text)-len(suffix)]
    return text

def pluralize(word: str, count: int = 2, custom: Optional[str] = None) -> str:
    """Pluralize a word based on count, with optional custom plural form."""
    if count == 1:
        return word
    if custom is not None:
        return custom
    
    # Uncountable nouns (same singular and plural)
    uncountable = {"sheep", "fish", "deer", "species", "aircraft", "salmon", "tuna"}
    if word.lower() in uncountable:
        return word
    
    # Basic English pluralization rules
    if word.endswith("y") and not word.endswith("ay") and not word.endswith("ey") and not word.endswith("oy") and not word.endswith("uy"):
        return word[:-1] + "ies"
    if word.endswith(("s", "sh", "ch", "x", "z")):
        return word + "es"
    return word + "s"

def singularize(word: str) -> str:
    """Singularize a plural word."""
    # Uncountable nouns
    uncountable = {"sheep", "fish", "deer", "species", "aircraft", "salmon", "tuna"}
    if word.lower() in uncountable:
        return word
    
    if word.endswith("ies"):
        return word[:-3] + "y"
    if word.endswith("es") and word.endswith(("ses", "shes", "ches", "xes", "zes")):
        return word[:-2]
    if word.endswith("s"):
        return word[:-1]
    return word

def starts_with(text: str, prefix: str) -> bool:
    """Check if text starts with a prefix."""
    return text.startswith(prefix)

def ends_with(text: str, suffix: str) -> bool:
    """Check if text ends with a suffix."""
    return text.endswith(suffix)

def contains(text: str, substring: str) -> bool:
    """Check if text contains a substring."""
    return substring in text

def count_occurrences(text: str, substring: str) -> int:
    """Count occurrences of a substring in text."""
    return text.count(substring)

def random_string(length: int = 8, chars: str = string.ascii_letters + string.digits) -> str:
    """Generate a random string of a given length."""
    return "".join(random.choice(chars) for _ in range(length))

def is_palindrome(text: str) -> bool:
    """Check if text is a palindrome (ignoring case and spaces)."""
    cleaned = re.sub(r"\W+", "", text.lower())
    return cleaned == cleaned[::-1]

def mask_text(text: str, visible_start: int = 0, visible_end: int = 0, mask_char: str = "*") -> str:
    """Mask text, showing only the specified number of characters at start and end."""
    if len(text) <= visible_start + visible_end:
        return text
    masked_len = len(text) - visible_start - visible_end
    return text[:visible_start] + (mask_char * masked_len) + text[-visible_end:] if visible_end > 0 else text[:visible_start] + (mask_char * masked_len)

def word_count(text: str) -> int:
    """Count the number of words in text."""
    return len(text.split())

def sentence_count(text: str) -> int:
    """Count the number of sentences in text."""
    return len(re.findall(r"[^.!?]+[.!?]+", text))

def first_word(text: str) -> Optional[str]:
    """Get the first word from text."""
    words = text.split()
    return words[0] if words else None

def last_word(text: str) -> Optional[str]:
    """Get the last word from text."""
    words = text.split()
    return words[-1] if words else None
