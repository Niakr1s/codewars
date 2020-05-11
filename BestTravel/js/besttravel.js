function ChooseBestSum(max, iterationsRemained, ls) {
    if (iterationsRemained <= 0) return -1
    if (ls.length < iterationsRemained) return -1;

    return ls.reduce((accum, value, i) => {
        if (iterationsRemained > 1) {
            sliceBest = ChooseBestSum(max - value, iterationsRemained - 1, ls.slice(i + 1));

            if (sliceBest < 0) return accum;

            value += sliceBest;
        }
        return (value > accum && value <= max) ? value : accum;
    }, -1);
}

module.exports.ChooseBestSum = ChooseBestSum