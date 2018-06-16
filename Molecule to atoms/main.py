import re
from collections import Counter


pat_elem = re.compile(r"([A-Z][a-z]*)(\d*)")
pat_subitems = re.compile(r"[\(\{\[](.*)[\)\}\]](\d*)")
# pat_clean = re.compile(r"\(.*\)\d+")
pat_clean = re.compile(r"[\(\{\[].*?[\)\}\]]\d*")

def parse_molecule (formula):
    return dict(recurs(formula))

def recurs(formula, count=1):
    c = elems(formula)
    for subitem, subitem_count in subitems(formula).items():
        c += multiply(recurs(subitem, subitem_count*count), subitem_count)
    return c

def multiply(c, count):
    accum = Counter()
    for _ in range(count):
        accum += c
    return accum

def elems(s):  # returns Counter
    return to_ints(pat_elem.findall(pat_clean.sub("", s)))

def subitems(s):  # returns Counter
    return to_ints(pat_subitems.findall(s))

def to_ints(it):  # returns Counter
    cleared = ((i[0],int(i[1]) if i[1] else 1) for i in it)
    result = []
    for i in cleared:
        for _ in range(i[1]):
            result.append(i[0])
    return Counter(result)

if __name__ == "__main__":
    TEST_STRING = "K4[ON(SO3)2]2"

    assert dict(elems("K4K[ON(SO3)2]2")) == {"K":5}
    assert dict(elems("C6H12O6")) == {"C":6, "H":12, "O":6}
    assert dict(multiply(Counter({"a": 2, "b":3}), 3)) == {"a": 6, "b": 9}
    assert dict(elems(TEST_STRING)) == {"K":4}
    assert dict(subitems(TEST_STRING)) == {"ON(SO3)2":2}
    assert dict(to_ints((("b",""),("c","2")))) == {"b":1, "c":2}
    assert dict(elems("(C5H5)Fe(CO)2CH3")) == {"Fe":1, "C":1, "H":3}

    print(parse_molecule(TEST_STRING))
    print(parse_molecule("Pd[P(C6H5)3]4"))
    print(parse_molecule("C6H12O6"))
    print(parse_molecule("(C5H5)Fe(CO)2CH3"))
    # print(parse_molecule(TEST_STRING))



    CHUNK = "(SO3)2"
    parse_molecule(CHUNK)

    # print(elems(TEST_STRING))
    # print(subitems(TEST_STRING))
    # print(pat_elem.findall(TEST_STRING))
    #
    # print("\n", pat_clean.sub("", TEST_STRING))