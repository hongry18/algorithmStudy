# Left Rotation

문제는 배열에 n개의 숫자들이 주어지고 d만큼 왼쪽으로 이동시켜라 이건데
queue로 반복해서 하면 풀림

NodeJS에서 queu
```javascript
var queue = [1,2,3,4,5]

var shift = queue.shift()

queue.push(shift)

console.log(queue)

// output
// [2,3,4,5,1]
```
