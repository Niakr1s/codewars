/*
Converts on digit from 0 to 9 to a romanian number
 */
export const convertDigit = (n: number, one: string, five?: string, ten?: string): string => {
    if (n <= 3) return one.repeat(n)
    if (!five || !ten) return ''
    if (n === 4) return one + five
    if (n <= 8) return five + one.repeat(n % 5)
    if (n === 9) return one + ten
    return ''
}

const romanian = {one: 'I', five: 'V', ten: 'X', fifty: 'L', hundred: 'C', fiveHundred: 'D', thousand: 'M',}

/*
returns args for convertDigit for dimensions 1000, 1000, 10
 */
const getArgs = (dimension: number): string[] => {
    switch (dimension) {
        case 1000:
            return [romanian.thousand]
        case 100:
            return [romanian.hundred, romanian.fiveHundred, romanian.thousand]
        case 10:
            return [romanian.ten, romanian.fifty, romanian.hundred]
        case 1:
            return [romanian.one, romanian.five, romanian.ten]
        default:
            return []
    }
}


export function solution(number: number): string {
    let result = ''
    for (let i = 1000; i >= 1; i /= 10) {
        const [left, right] = [Math.floor(number / i), number % i]
        const [one, five, ten] = getArgs(i)
        result += convertDigit(left, one, five, ten)
        number = right
    }
    return result
}
