function ChooseBestSum(max, iterationsRemained, ls) {
    if (iterationsRemained <= 0) return -1
    if (ls.length < iterationsRemained) return -1;

    let best = -1;
    for (let i = 0; i != ls.length; i++) {
        let innerbest = ls[i]
        if (iterationsRemained > 1) {
            sliceBest = ChooseBestSum(max - ls[i], iterationsRemained - 1, ls.slice(i + 1));

            if (sliceBest < 0) continue;
            innerbest += sliceBest;
        }

        if (innerbest > best && innerbest <= max) best = innerbest;
    };

    return best
}

module.exports.ChooseBestSum = ChooseBestSum