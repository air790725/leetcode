import (
    "math"
)

func searchInsert(nums []int, target int) int {
    insert, left, right := -1, 0, len(nums)-1

    for left <= right {
        mid := int(math.Floor(float64((right-left)/2))) + left
        if nums[mid] < target {
            left = mid + 1
            insert = left
        } else if nums[mid] == target {
            return mid
        } else {
            right = mid - 1
            insert = mid
        }
    }
	
    return insert
}