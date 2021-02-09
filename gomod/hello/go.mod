module hello

go 1.15

require (
    myLib/sample v0.0.0
)

replace (
    myLib/sample v0.0.0 => ./sample
)
