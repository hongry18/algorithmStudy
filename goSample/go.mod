module goSample

go 1.15

require (
    goSample.com/p1 v0.0.0
    goSample.com/p2 v0.0.0
)

replace (
    goSample.com/p1 v0.0.0 => ./p1
    goSample.com/p2 v0.0.0 => ./p2
)
