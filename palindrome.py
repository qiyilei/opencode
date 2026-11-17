import sys


def is_palindrome(text: str) -> bool:
    """忽略大小写、空格和标点，判断是否为回文。"""
    cleaned = [c.lower() for c in text if c.isalnum()]
    return cleaned == cleaned[::-1]


if __name__ == "__main__":
    args = sys.argv[1:]
    if not args:
        print("用法: python palindrome.py <文本>")
        sys.exit(1)
    text = " ".join(args)
    print("是回文" if is_palindrome(text) else "不是回文")
