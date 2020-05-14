function smallest(n) {
    let arr = numToArray(n);

    let res = Number.MAX_VALUE;
    let from = Number.MAX_VALUE;
    let to = Number.MAX_VALUE;

    arr.forEach((it1, i) => {
        arr.forEach((it2, j) => {
            let copiedArr = [...arr];
            move(copiedArr, i, j);
            let got = arrayToNum(copiedArr);

            if (got < res) { res = got; from = i; to = j; }
        })
    })


    return { value: res, from: from, to: to }
}

function numToArray(n) {
    let res = []

    if (n <= 0) {
        return res;
    }

    do {
        let div = n % 10;
        res.push(div);
        n = Math.floor(n / 10);
    } while (n !== 0);

    res = res.reverse();

    return res;
}

function arrayToNum(arr) {
    let res = parseInt(arr.join(""), 10);
    return Number.isNaN(res) ? 0 : res;
}

function move(arr, from, to) {
    let fromValue = arr[from];

    if (from > to) {
        for (let i = from; i != to; i--) {
            arr[i] = arr[i - 1];
        }
    } else if (from < to) {
        for (let i = from; i != to; i++) {
            arr[i] = arr[i + 1];
        }
    }

    arr[to] = fromValue;

    return arr
}

exports.smallest = smallest;
exports.numToArray = numToArray;
exports.arrayToNum = arrayToNum;
exports.move = move;