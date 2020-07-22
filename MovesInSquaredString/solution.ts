export function oper(fct: (s: string) => string, s: string): string {
    return fct(s)
}

export const vertMirror = (str: string): string => {
    return str.split('\n').map(s => s.split('').reverse().join('')).join('\n')
};

export const horMirror = (str: string): string => {
    return str.split('\n').reverse().join('\n')
};