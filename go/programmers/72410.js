function toUpper(c) {
    if (c < 65 || c > 90) {
        return c;
    }

    return c + 32;
}

function removeChar(c) {
    if (c > 64 && c < 91) {
        return c;
    }

    if (c > 96 && c < 123) {
        return c;
    }

    if (c > 47 && c < 58) {
        return c;
    }

    if ( c > 44 && c < 47 ) {
        return c;
    }

    if ( c === 95 ) {
        return c;
    }

    return -1
}

function mergeDot(arr) {
    const res = []
    const len = arr.length
    let prev = 0
    for (var i=0; i<len; i++) {
        if (arr[i] !== 46) {
            res.push(arr[i])
            prev = 0
            continue
        }

        if (prev === 0) {
            res.push(arr[i])
            prev = arr[i]
        }
    }
    return res
}

function removeFDot(arr) {
    while (true) {
        if (arr[0] !== 46) {
            break
        }
        arr.shift()
    }
}

function removeLDot(arr) {
    while (true) {
        if (arr[arr.length-1] !== 46) break
        arr.pop()
    }
}

function ab2str(a) {
    return String.fromCharCode.apply(null, new Uint16Array(a))
}

function solution(n) {
    let ar = []
    const len = n.length;
    for (let i=0; i<len; i++) {
        let c = n.charCodeAt(i)
        c = toUpper(c)
        c = removeChar(c)
        if (c === -1) continue
        ar.push(c)
    }

    ar = mergeDot(ar)
    removeFDot(ar)
    removeLDot(ar)

    if (ar.length === 0) {
        for (var i=0; i<3; i++) {
            ar.push(97)
        }
    }

    if (ar.length > 15) {
        ar = ar.slice(0,15)
    }

    removeLDot(ar)
    if (ar.length === 1) {
        ar.push(ar[0])
        ar.push(ar[0])
    }

    if (ar.length === 2) {
        ar.push(ar[1])
    }

    return ab2str(ar)
}
