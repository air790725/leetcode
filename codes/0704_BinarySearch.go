package algorithmn

func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1 
    for left <= right {
        mid := int((right - left)/2) + left
        if nums[mid] == target {
            return mid
        }
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}