에라토스테네스의 체를 미리 채워놓고 카운팅한다

```go
func eratos() {
    var prime []int = make([]int, 100000)
    prime[0]=1
    prime[1]=1

    for i:=2; i*i <= 1000; i++ {
        if prime[i] != 0 {
            continue
        }

        for j:= i*i; j<=1000; j+=i {
            print[j] = 1
        }
    }
}
```
