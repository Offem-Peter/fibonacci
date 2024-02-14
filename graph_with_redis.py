import matplotlib.pyplot as plt
import numpy as np

tests = np.arange(1, 11) 

throughput_httprouter = [96693.96, 109537.17, 100338.50, 99792.69, 107771.55, 89038.39, 91445.62, 93158.77, 88391.14, 80137.43]
throughput_fiber = [103613.28, 102754.27, 95873.50, 99266.36, 89973.71, 102298.83, 81751.11, 90566.22, 75181.90, 73596.74 ]
throughput_fasthttp = [228709.19, 199841.89, 207356.41, 203230.21, 219733.51, 215265.13, 195546.90, 194935.41, 203851.11, 225453.86]

plt.figure(figsize=(12, 6))

plt.plot(tests, throughput_httprouter, marker='o', label='httprouter')
plt.plot(tests, throughput_fiber, marker='x', label='fiber')
plt.plot(tests, throughput_fasthttp, marker='s', label='fasthttp')

plt.title('Throughput Comparison of frameworks using redis')
plt.xlabel('Test Number')
plt.ylabel('Throughput (requests/second)')

plt.yticks(np.arange(0, max(throughput_httprouter + throughput_fiber + throughput_fasthttp) + 1, 20000))


plt.legend()

plt.grid(True)

plt.show()
