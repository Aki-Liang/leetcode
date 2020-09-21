package leetcode
import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	for i:=0;i<len(nums)-3;{
		for j:=i+1;j<len(nums)-2;{
			m:= j+1
			n := len(nums)-1
			for m<n {
				sum := nums[i] + nums[j]+nums[m]+nums[n]
				if sum == target {
					res = append(res, []int{nums[i],nums[j],nums[m],nums[n]})
					for m < n {
						n--
						if nums[n] != nums[n+1] {
							break
						}
					}
					for m < n {
						m++
						if nums[m] != nums[m-1] {
							break
						}
					}
				} else if sum > target {
					for m < n {
						n--
						if nums[n] != nums[n+1] {
							break
						}
					}
				} else if sum < target {
					for m < n {
						m++
						if nums[m] != nums[m-1] {
							break
						}
					}
				}
			}

			for j<len(nums)-2 {
				j++
				if nums[j] != nums[j-1] {
					break
				}
			}
		}
		
		for i<len(nums)-3 {
			i++
			if nums[i] != nums[i-1] {
				break
			}
		}
	}

	return res
}