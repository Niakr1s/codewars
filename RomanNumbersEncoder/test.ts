import {convertDigit, solution} from './solution';
import { assert } from 'chai';

describe('solution', () => {
    it('basic', () => {
        assert.equal(solution(1000), 'M', '1000 should == "M"')
        assert.equal(solution(4), 'IV', '4 should == "IV"')
        assert.equal(solution(1), 'I', '1 should == "I"')
        assert.equal(solution(1990), 'MCMXC', '1990 should == "MCMXC"')
        assert.equal(solution(2008), 'MMVIII', '2008 should == "MMVIII"')
        assert.equal(solution(1444), 'MCDXLIV', '1444 should == "MCDXLIV"')
    });
});

describe('convertDigit', () => {
    it ('0 to 9', () => {
        const [one, five, ten] = ['I', 'V', 'X']
        assert.equal(convertDigit(0, one, five, ten), '')
        assert.equal(convertDigit(1, one, five, ten), 'I')
        assert.equal(convertDigit(2, one, five, ten), 'II')
        assert.equal(convertDigit(3, one, five, ten), 'III')
        assert.equal(convertDigit(4, one, five, ten), 'IV')
        assert.equal(convertDigit(5, one, five, ten), 'V')
        assert.equal(convertDigit(6, one, five, ten), 'VI')
        assert.equal(convertDigit(7, one, five, ten), 'VII')
        assert.equal(convertDigit(8, one, five, ten), 'VIII')
        assert.equal(convertDigit(9, one, five, ten), 'IX')
    })
})