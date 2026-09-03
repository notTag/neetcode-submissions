func isAnagram(s string, t string) bool {
    var sCharMap = make(map[byte]int)
    if len(s) != len(t){
        return false
    }

    for i:=0; i<len(s); i++ {
        // fmt.Println(string(s[i]) + " " + string(sCharMap[s[i]]))
        sCharMap[s[i]]++
        // fmt.Println(string(s[i]) + " " + string(sCharMap[s[i]]))
    }

    for i:=0; i<len(t); i++ {
        // fmt.Println(string(t[i]) + " " + string(sCharMap[t[i]]))
        sCharMap[t[i]]--
        if sCharMap[t[i]] < 0 {
            return false
        }
    }

    return true
}
