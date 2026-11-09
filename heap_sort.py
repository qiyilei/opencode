"""堆排序（Heap Sort）实现。"""

import random


def heap_sort(arr):
    """返回升序排序后的新列表（基于大顶堆，非原地）。"""
    a = arr[:]
    n = len(a)

    # 建堆：从最后一个非叶子节点开始向下调整
    for i in range(n // 2 - 1, -1, -1):
        _sift_down(a, i, n)

    # 依次将堆顶最大值交换到末尾并重新调整
    for end in range(n - 1, 0, -1):
        a[0], a[end] = a[end], a[0]
        _sift_down(a, 0, end)

    return a


def _sift_down(a, root, size):
    """向下调整堆，size 为当前堆的有效长度。"""
    while True:
        largest = root
        left = 2 * root + 1
        right = 2 * root + 2

        if left < size and a[left] > a[largest]:
            largest = left
        if right < size and a[right] > a[largest]:
            largest = right

        if largest == root:
            break
        a[root], a[largest] = a[largest], a[root]
        root = largest


if __name__ == "__main__":
    data = [random.randint(0, 100) for _ in range(20)]
    print("原始数组:", data)
    print("排序结果:", heap_sort(data))
