## Introduction

In the search for a webframework, I embarked on a journey of research to identify the most suitable web framework that meets the need for performance and high throughput. The goal was to ascertain which framework could deliver the fastest responses under heavy load.

### My System Specification for Benchmarking

The benchmarks were conducted on my personal machine, which has the following specifications:

- **CPU**: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
- **RAM**: 32GB
- **Operating System**: Linux
- **Benchmarking Tool**: [Bombardier](https://github.com/codesenberg/bombardier)

### Research and Selection Process

Before commencing the actual development phase, I devoted considerable effort to researching various web frameworks. The landscape is diverse, with many frameworks offering a range of benefits. Some shine with their rich features such as embedded routes, advanced debugging capabilities, and extensive middleware support, which are invaluable for certain project requirements.

However, my focus was on finding a framework that not only performs exceptionally in high-throughput scenarios but also promotes clean, concise, and maintainable code. I was particularly interested in a framework that would not impose extensive boilerplate code, allowing for swift and straightforward development without compromising on performance.

After reviewing extensive benchmarks done by the community and also [here](https://github.com/smallnest/go-web-framework-benchmark) and analyzing performance metrics, three frameworks stood out for use:

- `httprouter`
- `fiber`
- `fasthttp/router`

I chose to focus on these three contenders for their well-documented performance advantages.

## Benchmarking Without Caching

Initially, I implemented the same Fibonacci sequence generation logic within each framework and stored the data in application memory. During this phase, I discovered that, despite `fasthttp` showcasing superior initial throughput, the performance of `fasthttp` and the other frameworks degraded with subsequent benchmark runs. This was due to the increased memory allocation for the variables storing the Fibonacci numbers.

![Performance Graph Without Caching](/images/benchamark.png)

To view the graphs data in detail, you can run in the project directory

```
make graph-without-redis
```

---

###

So no no no, We have to look for a better way to do this, or else in a matter of time, our application will stop functioning. And we don't want that.

---

## Benchmarking With Caching

To address the performance degradation, I introduced caching using Redis. By caching the growing Fibonacci data, I aimed to achieve more consistent throughput across tests.

Here are the results after the introduction of Redis for caching:

![Performance Graph With Caching](/images/bencmark-redis.png)

To view the graphs data in detail, you can run in the project directory

```
make graph-with-redis
```

---

## Conclusion

The benchmark tests were insightful, guiding me to a data-driven decision rather than a subjective preference for a particular framework. `fasthttp` emerged as what I leader based on data driven decisions; however, the introduction of caching proved to enhance the consistency of the API's throughput.

---

### How to Run the application

To run the final application, you need to run this docker command, in the project directory to start a redis server.

```
docker compose up -d
```

Then you can start the server by running

```
make run/api
```

you can also run the other servers which I built to tests out the benchmark.
You can see a list of commands you can run by running just

```
make
```
