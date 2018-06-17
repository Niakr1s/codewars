import re
from collections import Counter

"""
parser I overlooked at solutions page, but alghoritm I wrote by myself
"""

parser = re.compile(r"(\[[^\[]+\]|\([^\(]+\)|\{[^\{]+\}|[A-Z][a-z]*)(\d*)")

def parse_molecule(formula, counter=1):
    print("took formula {} with counter {}".format(formula, counter))
    result = Counter()
    for item, count in parser.findall(formula):
        if not count:
            count = 1
        else:
            count = int(count)
        if item[0] in ("(", "[", "{"):
            print("\tgoing deeper, ", item, "=>", item[1:-1], count*counter)
            result += parse_molecule(item[1:-1], count*counter)
        else:
            print("\tsumming, ", item, count*counter)
            result += Counter([item]*count*counter)
    print(result)
    return result

