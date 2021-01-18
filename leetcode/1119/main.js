/**
 * @param {string}
 * @return {string}
 */
// Runtime: 76 ms, faster than 73.24% of JavaScript online submissions for Remove Vowels from a String.
// Memory Usage: 38.6 MB, less than 47.49% of JavaScript online submissions for Remove Vowels from a String.
var removeVowels = function(s) {
    
    // let result = ""
    let result = []
    var p = 0
    for (var i = 0; i < s.length; i++) {
        switch (s[i]) {
            case 'a':
            case 'e':
            case 'i':
            case 'o':
            case 'u':
                break;
            default:
                result[p] = s[i];
                p++;
        }
    }

    result = result.join("");
    
    return result
};

/**
 * @param {string}
 * @return {string}
 */
// Runtime: 80 ms, faster than 43.26% of JavaScript online submissions for Remove Vowels from a String.
// Memory Usage: 38.8 MB, less than 19.92% of JavaScript online submissions for Remove Vowels from a String.
var removeVowels = function(s) {
    
    let result = ""
    for (var i = 0; i < s.length; i++) {
        switch (s[i]) {
            case 'a':
            case 'e':
            case 'i':
            case 'o':
            case 'u':
                break;
            default:
                result += s[i]
        }
    }
    
    return result
};