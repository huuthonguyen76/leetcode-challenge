def solution(n: int):
    l_result = []
    def recursive(l, r, result):
        if l == n:
            result += ')' * (n - r)
            l_result.append(result)
            return
        
        recursive(l + 1, r, result + '(')
        if r < l:
            recursive(l, r + 1, result + ')')
    
    recursive(0, 0, '')
    print(l_result)
            

n = 3

solution(n)
