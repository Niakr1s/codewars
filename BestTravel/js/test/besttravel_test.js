var assert = require('assert');
var besttravel = require('../besttravel')

describe('ChooseBestSum', function () {
    testCases = [
        { arr: [0, 1, 2, 3], max: 20, towns: 0, expected: -1, },
        { arr: [0, 1, 2, 3], max: 20, towns: 1, expected: 3, },
        { arr: [0, 1, 2, 3], max: 20, towns: 2, expected: 5, },
        { arr: [0, 1, 2, 3], max: 20, towns: 3, expected: 6, },
        { arr: [50, 55, 56, 57, 58], max: 163, towns: 3, expected: 163, },
        { arr: [50], max: 163, towns: 3, expected: -1, },
        { arr: [91, 74, 73, 85, 73, 81, 87], max: 230, towns: 3, expected: 228, },
        { arr: [91, 74, 73, 85, 73, 81, 87], max: 331, towns: 2, expected: 178, },
        { arr: [91, 74, 73, 85, 73, 81, 87], max: 331, towns: 4, expected: 331, },
        { arr: [91, 74, 73, 85, 73, 81, 87], max: 331, towns: 5, expected: -1, },
        {
            arr: [100, 76, 56, 44, 89, 73, 68, 56, 64, 123, 2333, 144, 50, 132, 123, 34, 89],
            max: 880, towns: 8, expected: 876,
        },
    ]

    testCases.forEach((testCase) => {
        describe(`Arr: [${testCase.arr}],MaxDistance: ${testCase.max},Towns:${testCase.towns}`, function () {
            let got = besttravel.ChooseBestSum(testCase.max, testCase.towns, testCase.arr);
            it(`Expected: ${testCase.expected}, got: ${got}`, () => {
                assert.equal(got, testCase.expected);
            });
        });
    });
});