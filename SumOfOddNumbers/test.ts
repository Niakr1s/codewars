import { assert } from "chai";
import {getFirstNumber, rowSumOddNumbers} from "./solution";

describe("basic tests", function() {
    it("Testing for 1", () => assert.strictEqual(rowSumOddNumbers(1), 1));
    it("Testing for 2", () => assert.strictEqual(rowSumOddNumbers(2), 8));
    it("Testing for 13", () => assert.strictEqual(rowSumOddNumbers(13), 2197));
    it("Testing for 19", () => assert.strictEqual(rowSumOddNumbers(19), 6859));
    it("Testing for 41", () => assert.strictEqual(rowSumOddNumbers(41), 68921));
    it("Testing for 42", () => assert.strictEqual(rowSumOddNumbers(42), 74088));
    it("Testing for 74", () => assert.strictEqual(rowSumOddNumbers(74), 405224));
    it("Testing for 86", () => assert.strictEqual(rowSumOddNumbers(86), 636056));
    it("Testing for 93", () => assert.strictEqual(rowSumOddNumbers(93), 804357));
    it("Testing for 101", () => assert.strictEqual(rowSumOddNumbers(101), 1030301));
});

describe('getFirstNumber', () => {
    it('should return 1 for row=1', () => {assert.equal(getFirstNumber(1), 1)})
    it('should return 3 for row=2', () => {assert.equal(getFirstNumber(2), 3)})
    it('should return 7 for row=3', () => {assert.equal(getFirstNumber(3), 7)})
    it('should return 13 for row=4', () => {assert.equal(getFirstNumber(4), 13)})
    it('should return 21 for row=5', () => {assert.equal(getFirstNumber(5), 21)})
})