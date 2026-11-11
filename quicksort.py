def quicksort(arr):
    """快速排序（原地，升序）。平均 O(n log n)，最坏 O(n^2)。"""
    _quicksort(arr, 0, len(arr) - 1)
    return arr


def _quicksort(arr, low, high):
    if low >= high:
        return

    pivot_index = _partition(arr, low, high)
    _quicksort(arr, low, pivot_index - 1)
    _quicksort(arr, pivot_index + 1, high)


def _partition(arr, low, high):
    # 取最后一个元素作为基准
    pivot = arr[high]
    i = low - 1

    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]

    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


if __name__ == "__main__":
    import random

    data = [random.randint(0, 100) for _ in range(20)]
    print("排序前:", data)
    print("排序后:", quicksort(data[:]))
    print("校验:", quicksort(data[:]) == sorted(data))
