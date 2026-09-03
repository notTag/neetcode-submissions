func hasDuplicate(nums []int) bool {
    var numMap = make(map[int]bool)
    result := false

    for i:=0; i<len(nums); i++ {
        if numMap[nums[i]] {
            result = true
            break;
        }
        numMap[nums[i]] = true
    }
    
    fmt.Println(result)
    return result
}
