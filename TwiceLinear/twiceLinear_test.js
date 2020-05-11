var assert = require('assert')
var twiceLinear = require('./twiceLinear')

describe("Twice linear", () => {
    [[10, 22], [50, 175], [100, 447], [500, 3355], [1000, 8488], [100000, 2911582]].forEach((iter) => {
        let got = twiceLinear.DblLinear(iter[0]);
        it(`Expected: ${iter[1]}, got: ${got}`, () => {
            assert.equal(got, iter[1]);
        });
    })
})