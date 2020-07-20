import { twoOldestAges } from './solution';
import { assert } from "chai";

describe("Examples", function() {
    it("twoOldestAges([1,5,87,45,8,8]) should return [45,87]", function() {
        assert.deepEqual(twoOldestAges([1, 5, 87, 45, 8, 8]), [45, 87]);
    });
});