var assert = require('assert')
var findTheSmallest = require('./FindTheSmallest')

describe('sallest', () => {
    let testCases = [
        { input: 261235, expected: { value: 126235, from: 2, to: 0 } },
        { input: 209917, expected: { value: 29917, from: 0, to: 1 } },
        { input: 199819884756, expected: { value: 119989884756, from: 4, to: 0 } },
        { input: 5190930055291, expected: { value: 519093055291, from: 6, to: 0 } },
        { input: 60496893958362, expected: { value: 4696893958362, from: 0, to: 2 } },
        { input: 111411111, expected: { value: 111111114, from: 3, to: 8 } },
    ]

    testCases.forEach((testCase) => {
        let got = findTheSmallest.smallest(testCase.input);
        it(`Input: ${testCase.input}`, () => assert.deepEqual(got, testCase.expected))
    })
})

describe('numToArray', () => {
    let testCases = [
        { input: 123456, expected: [1, 2, 3, 4, 5, 6] },
        { input: 0, expected: [] },
        { input: -1, expected: [] },
    ]

    testCases.forEach((testCase) => {
        let got = findTheSmallest.numToArray(testCase.input);

        it(`Input: ${testCase.input}`, () => {
            assert.deepEqual(got, testCase.expected)
        })
    })
})

describe('arrayToNum', () => {
    let testCases = [
        { expected: 123456, input: [1, 2, 3, 4, 5, 6] },
        { expected: 0, input: [] },
    ]

    testCases.forEach((testCase) => {
        let got = findTheSmallest.arrayToNum(testCase.input);

        it(`Input: ${testCase.input}`, () => {
            assert.equal(got, testCase.expected)
        })
    })
})


describe('move', () => {
    let testCases = [
        { input: [1, 2, 3, 4, 5, 6], from: 2, to: 0, expected: [3, 1, 2, 4, 5, 6], },
        { input: [1, 2, 3, 4, 5, 6], from: 2, to: 4, expected: [1, 2, 4, 5, 3, 6], },
    ]

    testCases.forEach((testCase) => {
        let got = findTheSmallest.move(testCase.input, testCase.from, testCase.to);

        it(`Input: ${testCase.input}`, () => {
            assert.deepEqual(got, testCase.expected)
        })
    })
})