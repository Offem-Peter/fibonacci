import matplotlib.pyplot as plt
import numpy as np

tests = np.arange(1, 11) 

throughput_httprouter = [84737.77, 91066.57, 34145.17, 47860.18, 18132.93, 11897.48,6563.11, 1044.54, 994.03, 932.71]
throughput_fiber = [102817.63, 68276.34, 58027.32, 46779.77, 38616.98, 26388.19, 7309.40, 1075.71, 991.74, 898.61]
throughput_fasthttp = [105890.32, 100748.38,99810.23, 97929.78, 51813.62, 7893.20, 5186.55, 951.95, 926.03, 962.45]

plt.figure(figsize=(12, 6))

plt.plot(tests, throughput_httprouter, marker='o', label='httprouter')
plt.plot(tests, throughput_fiber, marker='x', label='fiber')
plt.plot(tests, throughput_fasthttp, marker='s', label='fasthttp')

plt.title('Throughput Comparison of frameworks without redis')
plt.xlabel('Test Number')
plt.ylabel('Throughput (requests/second)')

plt.yticks(np.arange(0, max(throughput_httprouter + throughput_fiber + throughput_fasthttp) + 1, 20000))


plt.legend()

plt.grid(True)

plt.show()
