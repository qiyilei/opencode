"""冒泡排序（Bubble Sort）实现。"""

import random


def bubble_sort(arr):
    """返回升序排序后的新列表（非原地）。"""
    a = arr[:]
    n = len(a)
    for i in range(n - 1):
        swapped = False
        # 每轮把最大值冒到末尾，已就绪的尾部无需再比较
        for j in range(n - 1 - i):
            if a[j] > a[j + 1]:
                a[j], a[j + 1] = a[j + 1], a[j]
                swapped = True
        if not swapped:
            break
    return a


if __name__ == "__main__":
    data = [random.randint(0, 100) for _ in range(20)]
    print("原始数组:", data)
    print("排序结果:", bubble_sort(data))
