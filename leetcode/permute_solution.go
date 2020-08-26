package leetcode

func Permute(nums []int) [][]int {
	var res [][]int
	var track []int
	var used = make([]bool, len(nums))
	backtrack(nums, track, &res, used)
	return res
}

func backtrack(nums, track []int, res *[][]int, used []bool) {
	length := len(nums)
	if len(track) == length {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < length; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack(nums, track, res, used)
		used[i] = false
		track = track[:len(track)-1]
	}
}
