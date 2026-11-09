"""快速排序（Quick Sort）实现。"""

import random


def quick_sort(arr):
    """返回升序排序后的新列表（非原地）。"""
    if len(arr) < 2:
        return arr[:]
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    mid = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quick_sort(left) + mid + quick_sort(right)


def quick_sort_inplace(arr, lo=0, hi=None):
    """原地快速排序（迭代中位数分区）。"""
    if hi is None:
        hi = len(arr) - 1
    if lo >= hi:
        return arr

    pivot = arr[(lo + hi) // 2]
    i, j = lo, hi
    while i <= j:
        while arr[i] < pivot:
            i += 1
        while arr[j] > pivot:
            j -= 1
        if i <= j:
            arr[i], arr[j] = arr[j], arr[i]
            i += 1
            j -= 1

    quick_sort_inplace(arr, lo, j)
    quick_sort_inplace(arr, i, hi)
    return arr


if __name__ == "__main__":
    data = [random.randint(0, 100) for _ in range(20)]
    print("原始数组:", data)
    print("排序结果:", quick_sort(data))

    data2 = data[:]
    quick_sort_inplace(data2)
    print("原地排序:", data2)
