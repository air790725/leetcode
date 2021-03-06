package algorithmn

func isBadVersion(version int) bool;

func firstBadVersion(n int) int {
    left, right, last := 1, n, 0
    for left <= right {
        mid := int((right-left)/2) + left
        if (left == right) {
            if (isBadVersion(left)) {
                return left
            } else {
                return last
            }
        }
        if (isBadVersion(mid)) {
            right = mid - 1
            last = mid
        } else {
            left = mid + 1
        }
    }
    return last
}