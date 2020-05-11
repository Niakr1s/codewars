DblLinear = (n) => {
    let arr = [1]

    let yi = 0;
    let zi = 0;

    while (arr.length <= n) {
        y = 2 * arr[yi] + 1;
        z = 3 * arr[zi] + 1;

        if (y < z) {
            arr.push(y);
            yi++;
        } else if (z < y) {
            arr.push(z);
            zi++;
        } else {
            arr.push(y);
            yi++; zi++;
        }
    }

    return arr[n];
}

exports.DblLinear = DblLinear;