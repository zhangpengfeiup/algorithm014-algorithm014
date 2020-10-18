func reverseOnlyLetters(S string) string {
    stack := make([]byte, 0)

    buf := []byte(S)

    for i := 0;i < len(buf);i++ {
        if isLetter(buf[i]) {
             stack = append(stack, buf[i])
        }
    }
    
    s := []byte{}

    for i := 0;i < len(buf);i++ {
        if isLetter(buf[i]) {
            s = append(s, stack[len(stack)-1])
            stack = stack[:len(stack) - 1]
        } else {
            s = append(s, buf[i])
        }
    }

    return string(s)
}

func isLetter(s byte) bool {
    if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z'  {
        return true
    }
    return false
}
