# Sparse Arrays

queries 배열에 삽입된 문자열을
strings 배열에서 같은게 있는지 찾아서 리턴

O(N^2) 으로 품

```nodejs
let res = []
for ( var qi=0, qe=queries.length; qi<qe; qi++) {
    res[qi] = 0
    for (var si=0; se=strings.length; si<se; si++) {
        if (queries[qi] === strings[si]) {
            res[qi] += 1
        }
    }
}
```
