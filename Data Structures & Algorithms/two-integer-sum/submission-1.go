func twoSum(nums []int, target int) []int {
    var numMap = make(map[int]int)
    //loop over nums
    for i:=0; i<len(nums); i++ {
        //put nums into map. Map (nums[i] to i)
        numMap[nums[i]] = i
    }

    for i:=0; i<len(nums); i++ {
        //get the number we are looking for to get to our target. 
        targetCompliment := target - nums[i]
        //make sure the indices are not the same. 
        if val, exists := numMap[targetCompliment]; exists && val != 0 && val != i {
            //return indices, sort by asc index
            if i > numMap[targetCompliment] {
                return []int{numMap[targetCompliment], i}
            } else {
                return []int{i, numMap[targetCompliment]}
            }
        }
    }
    return make([]int, 0)
}
