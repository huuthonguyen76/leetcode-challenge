function longestCommonPrefix(strs: string[]): string {
    if (strs.length == 0) {
        return "";
    }

    let commonPrefix = '';

    let cursor = 0;

    while (true) {
        let curChar = strs[0]![cursor];

        if (curChar == undefined) {
            return commonPrefix;
        }

        for (let i = 0; i < strs.length; i++) {
            if ((strs[i]?.length ?? 0) <= cursor || strs[i]![cursor] != curChar) {
                return commonPrefix;                                                                            
            }
        }
        commonPrefix += curChar;

        cursor++;
    }
};

let strs = ["flower", "flow", "flight"];
console.log(longestCommonPrefix(strs));
