import re
from collections import Counter


pat_elem = re.compile(r"([A-Z][a-z]*)(\d*)")
# pat_subitems = re.compile(r"[\(\{\[](.*)[\)\}\]](\d*)")
pat_subitems = re.compile(r"([\(].*[\)]\d*|[\[].*[\]]\d*|[\{].*[\}]\d*)")
# pat_clean = re.compile(r"\(.*\)\d+")
pat_clean = re.compile(r"[\(\{\[].*?[\)\}\]]\d*")
pat_stripe = re.compile(r"[\[\{\(](.*)[\]\}\)](\d*)")

def parse_molecule (formula):
    return dict(recurs(formula))

def recurs(formula, count=1):
    c = elems(formula)
    for subitem in subitems(formula):
        subitem_stripped, subitem_count = stripe_and_count(subitem)
        c += multiply(recurs(subitem_stripped, subitem_count*count), subitem_count)
    return c

def multiply(c, count):
    accum = Counter()
    for _ in range(count):
        accum += c
    return accum

def elems(s):  # returns Counter
    return to_ints(pat_elem.findall(pat_subitems.sub("", s)))

def subitems(s):  # returns Counter
    result = []
    for subitem in pat_subitems.findall(s):
        print("subitem", subitem)
        print("stripe_and_count", stripe_and_count(subitem))
        result.append(stripe_and_count(subitem))
    print(result)
    return to_ints(result)

def to_ints(it):  # returns Counter
    print("it: ", it)
    cleared = ((i[0],int(i[1]) if i[1] else 1) for i in it)
    result = []
    for i in cleared:
        for _ in range(i[1]):
            result.append(i[0])
    return Counter(result)

def stripe_and_count(chunk):  # e.g "[CO2]2" => ("CO2", "2")
    return pat_stripe.findall(chunk)[0]


if __name__ == "__main__":
    TEST_STRING = "K4[ON(SO3)2]2"

    # print(pat_subitems.findall("K4[ON(SO3)2]2(OH)"))

    # assert dict(elems("K4K[ON(SO3)2]2")) == {"K": 5}
    # assert dict(elems("C6H12O6")) == {"C":6, "H":12, "O":6}
    # assert dict(multiply(Counter({"a": 2, "b":3}), 3)) == {"a": 6, "b": 9}
    # assert dict(elems(TEST_STRING)) == {"K":4}

    # assert dict(subitems(TEST_STRING)) == {"ON(SO3)2":2}
    # assert dict(to_ints((("b",""),("c","2")))) == {"b":1, "c":2}

    print(pat_clean.findall("(C5H5)Fe(CO)2CH3"))
    print(pat_subitems.findall("(C5H5)Fe(CO)2CH3"))
    assert dict(elems("(C5H5)Fe(CO)2CH3")) == {"Fe":1, "C":1, "H":3}
    # assert stripe_and_count("[ON(SO3)]2") == ("ON(SO3)", "2")
    # assert stripe_and_count("[ON(SO3)]") == ("ON(SO3)", "")
    #
    # print(parse_molecule(TEST_STRING))
    # print(parse_molecule("Pd[P(C6H5)3]4"))
    # print(parse_molecule("C6H12O6"))
    # print(parse_molecule("(C5H5)Fe(CO)2CH3"))
    # print(parse_molecule(TEST_STRING))



    # CHUNK = "(SO3)2"
    # parse_molecule(CHUNK)

    # print(elems(TEST_STRING))
    # print(subitems(TEST_STRING))
    # print(pat_elem.findall(TEST_STRING))
    #
    # print("\n", pat_clean.sub("", TEST_STRING))