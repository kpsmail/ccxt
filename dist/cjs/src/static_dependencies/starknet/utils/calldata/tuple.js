'use strict';

var cairo = require('./cairo.js');

// ----------------------------------------------------------------------------
function parseNamedTuple(namedTuple) {
    const name = namedTuple.substring(0, namedTuple.indexOf(':'));
    const type = namedTuple.substring(name.length + ':'.length);
    return { name, type };
}
function parseSubTuple(s) {
    if (!s.includes('('))
        return { subTuple: [], result: s };
    const subTuple = [];
    let result = '';
    let i = 0;
    while (i < s.length) {
        if (s[i] === '(') {
            let counter = 1;
            const lBracket = i;
            i++;
            while (counter) {
                if (s[i] === ')')
                    counter--;
                if (s[i] === '(')
                    counter++;
                i++;
            }
            subTuple.push(s.substring(lBracket, i));
            result += ' ';
            i--;
        }
        else {
            result += s[i];
        }
        i++;
    }
    return {
        subTuple,
        result,
    };
}
function extractCairo0Tuple(type) {
    const cleanType = type.replace(/\s/g, '').slice(1, -1); // remove first lvl () and spaces
    // Decompose subTuple
    const { subTuple, result } = parseSubTuple(cleanType);
    // Recompose subTuple
    let recomposed = result.split(',').map((it) => {
        return subTuple.length ? it.replace(' ', subTuple.shift()) : it;
    });
    // Parse named tuple
    if (cairo.isTypeNamedTuple(type)) {
        recomposed = recomposed.reduce((acc, it) => {
            return acc.concat(parseNamedTuple(it));
        }, []);
    }
    return recomposed;
}
function getClosureOffset(input, open, close) {
    for (let i = 0, counter = 0; i < input.length; i++) {
        if (input[i] === open) {
            counter++;
        }
        else if (input[i] === close && --counter === 0) {
            return i;
        }
    }
    return Number.POSITIVE_INFINITY;
}
function extractCairo1Tuple(type) {
    // un-named tuples support
    const input = type.slice(1, -1); // remove first lvl ()
    const result = [];
    let currentIndex = 0;
    let limitIndex;
    while (currentIndex < input.length) {
        switch (true) {
            // Tuple
            case input[currentIndex] === '(': {
                limitIndex = currentIndex + getClosureOffset(input.slice(currentIndex), '(', ')') + 1;
                break;
            }
            case input.startsWith('core::result::Result::<', currentIndex) ||
                input.startsWith('core::array::Array::<', currentIndex) ||
                input.startsWith('core::option::Option::<', currentIndex):
                {
                    limitIndex = currentIndex + getClosureOffset(input.slice(currentIndex), '<', '>') + 1;
                    break;
                }
            default: {
                const commaIndex = input.indexOf(',', currentIndex);
                limitIndex = commaIndex !== -1 ? commaIndex : Number.POSITIVE_INFINITY;
            }
        }
        result.push(input.slice(currentIndex, limitIndex));
        currentIndex = limitIndex + 2; // +2 to skip ', '
    }
    return result;
}
/**
 * Convert tuple string definition into object like definition
 * @param type tuple string definition
 * @returns object like tuple
 */
function extractTupleMemberTypes(type) {
    if (cairo.isCairo1Type(type)) {
        return extractCairo1Tuple(type);
    }
    return extractCairo0Tuple(type);
}

module.exports = extractTupleMemberTypes;
