function countOne(n: number): number {
    return n.toString(2).split('1').length - 1;
}

function countBits(n: number): number[] {
    let result: number[] = [];
    for (let i = 0; i <= n; i++) {
        result.push(
            i.toString(2).split('1').length - 1
        );
    }
    return result;
};

console.log(countBits(5));
