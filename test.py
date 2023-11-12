def restore_array(b):
    n = len(b)
    a = [0] * n

    for i in range(n - 1, -1, -1):
        count = 0
        for j in range(n):
            if a[j] == 0 and count == b[i]:
                a[j] = i + 1
                break
            if a[j] == 0:
                count += 1

    return a

# Пример использования
b = [2, 1, 0]
result = restore_array(b)
print(result)