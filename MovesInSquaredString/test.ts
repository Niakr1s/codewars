import { vertMirror, horMirror, oper } from './solution';
import { assert } from "chai";

describe("Fixed Tests", function() {
    it("Basic tests vertMirror", function() {
        assert.strictEqual(oper(vertMirror, "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"), "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw");
    });

    it("Basic tests horMirror", function() {
        assert.strictEqual(oper(horMirror, "lVHt\nJVhv\nCSbg\nyeCt"), "yeCt\nCSbg\nJVhv\nlVHt");
    });
});


