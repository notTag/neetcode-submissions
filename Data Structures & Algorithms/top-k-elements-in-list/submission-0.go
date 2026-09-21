func topKFrequent(nums []int, k int) []int {
    // var topKElementList []int
    //loop over nums
    //insert nums into map. Num value:freq
    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    
    //create second map freq:valList
    freqValSlice := make ([][]int, len(nums)+1)
    for curKey, curFreq := range m {
        freqValSlice[curFreq] = append(freqValSlice[curFreq], curKey)
    }

    freqSlice := []int{}
    for i := len(nums); i > 0 && len(freqSlice) < k; {
        if len(freqValSlice[i]) > 0 {
            kRemaining := k - len(freqSlice)
            for j := 0; j < kRemaining && j < len(freqValSlice[i]); j++ {
                freqSlice = append(freqSlice, freqValSlice[i][j])
            }
        }
        i--
    }

    return freqSlice
}