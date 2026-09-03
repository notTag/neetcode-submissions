func groupAnagrams(strs []string) [][]string {
    var strList [][]string
    anagramMapCount := make(map[string]int)
    
    for _, curStr := range strs {
        strListRow := hasAnagramInList(curStr, strs)
        anagramMapCount[strListRow[0][0]]++
        fmt.Printf("key: %v, row: %v\n", strListRow[0][0], strListRow)
        
        if anagramMapCount[strListRow[0][0]] == 1 {
            strList = append(strList, strListRow[0])
        } 
    }
    return strList
}

func hasAnagramInList(str string, strList []string) [][]string {
    var strListRow []string
    var strListTemp [][]string

    for _, curStr := range strList {
        //compare str to each item in strList calling isAnagram in each loop to compare them. 
        if isAnagram(str, curStr) {
            //if an anagram is found
            //add it to a new list to be returned to GroupAnagarams
            strListRow = append(strListRow, curStr)
        }
    }
    
    strListTemp = append(strListTemp, strListRow)
    return strListTemp
}

func isAnagram(str1 string, str2 string) bool {
    if len(str1) != len(str2) {
        return false
    }
    //create str1 map
    var strOneMap = make(map[rune]int)
    //put characters from str1 into a map with a count of each letter
    for _, curChar := range str1 {
        strOneMap[curChar]++
    }
    //loop over characters of str2. 
    for _, curChar := range str2 {
        //if not, decrement the map count
        strOneMap[curChar]--
        //look up curChar from str2 in our strOneMap, if we are below zero then it is not an anagram
        if strOneMap[curChar] < 0 {
            return false
        }
    }
    // return true if we made it out of our loop (indicating str1 and str2 are anagrams)
    fmt.Printf("str1: %s, str2: %s\n", str1, str2)
    return true
}
