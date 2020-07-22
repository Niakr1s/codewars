export const findOdd = (xs: number[]): number => {
    const times: { [k: string]: number } = {}
    for (let n of xs) {
        times[n] === undefined ? times[n] = 1 : times[n]++
    }
    for (let key of Object.keys(times)) {
        let value = times[key]
        if (isOdd(value)) return +key
    }
    return 0;
};

const isOdd = (n: number): boolean => {
    return n % 2 === 1
}