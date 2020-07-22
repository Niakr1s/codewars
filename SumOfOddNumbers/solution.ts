export function rowSumOddNumbers(n: number): number {
    return sum(genRow(n))
}

export const sum = (gen: Generator<number, void, unknown>): number => {
    let res = 0
    for (let n of gen) res += n
    return res
}

export const getFirstNumber = (row: number): number => {
    if (row === 1) return 1
    row--
    return row * 2 + getFirstNumber(row)
}

export function* genRow(row: number) {
    let iter = getFirstNumber(row)
    for (let i = 0; i < row; i++) {
        yield iter
        iter += 2
    }
}