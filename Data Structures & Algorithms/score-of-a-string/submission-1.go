func scoreOfString(s string) int {
    r := 0
    for i := 0; i < len(s)-1; i++ {
        d := int(s[i]) - int(s[i+1])
        if d < 0 {
            d = -d
        }
        r += d
    }
    return r
}