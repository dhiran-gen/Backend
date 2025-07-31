 **Chapter 1: Fundamentals of Scalable System Design** with a pragmatic + code-driven mindset.
We’ll go **concept → real-world analogy → Golang-based exercise** to master each part.

---

# 📘 Chapter 1: Fundamentals of Scalable System Design

## 🧱 1.1 Core Characteristics of Scalable Systems

* **Availability**: Uptime, error tolerance, failover
* **Reliability**: Data correctness, retry safety
* **Scalability**: Handle more load (horizontally or vertically)
* **Performance**: Latency and throughput
* **Maintainability**: Easy to modify/debug/extend

> 🧠 *Real-World:* Think of Swiggy scaling from 100 orders/day to 1M orders/day — how would backend infra behave?

---

## 🔢 1.2 Understanding Load Metrics

* **QPS (Queries per second)**: Incoming request rate
* **Latency (P99, P95)**: Time to serve request
* **Throughput**: Data processed per second
* **SLA/SLO/SLI**: Formalized availability/performance metrics

> 🧠 *Spotify*: P95 latency for “Play” action must be < 200ms for premium users.

---

## 📶 1.3 Client-Server Protocols

* **TCP/UDP**: Transport layer, connection-oriented vs stateless
* **HTTP1.1 / HTTP2 / HTTP3 (QUIC)**
* **gRPC vs REST vs WebSocket**

> 🧠 *Example*: gRPC used internally in Google/Facebook microservices for low-latency RPC

---

## 🧰 1.4 Types of System Components

* **Frontend (Browser, Mobile)** → API Gateway → Services
* **Monolith vs Microservice** vs Serverless
* **Synchronous vs Asynchronous communication**

---

## ⚖️ 1.5 Bottlenecks and Limits

* **Vertical Scaling Limit**: CPU, Memory, Disk IOPS
* **Horizontal Scaling Challenge**: Consistency, state sharing
* **Network Latency**, **Queue Buildup**, **DB lock contention**

---

## 🔁 1.6 Lifecycle of a Request (End-to-End)

> Example: `POST /order` in Swiggy

1. Enters Load Balancer
2. Hits API Gateway
3. Auth Service checks JWT
4. Order Service validates + inserts into DB
5. Pushes event to Kafka
6. Async: Delivery Service picks it up

> Diagram will be provided during the exercise

---

## ⚙️ 1.7 Key Design Questions to Always Ask

* What’s the QPS target?
* What’s acceptable latency (P99)?
* How many users will scale up simultaneously?
* What failure scenarios must be handled?
* Is the system **eventually consistent** or **strongly consistent**?

---

# 💻 Practical Work & Golang Exercises

## ✅ Exercise 1: QPS/Latency Estimator Tool in Go

🔧 Build a CLI tool in Go that calculates:

* Expected QPS
* Latency buckets (P50/P95/P99)
* Network + DB overhead

```go
type Metrics struct {
	TotalRequests int
	AvgRespTimeMs float64
	P95RespTimeMs float64
	P99RespTimeMs float64
}

func EstimateThroughput(metrics Metrics) float64 {
	// Simple throughput estimation: requests per second
	return float64(metrics.TotalRequests) / (metrics.AvgRespTimeMs / 1000)
}
```

📌 Enhance: Add CLI support with Cobra or flag parsing

---

## ✅ Exercise 2: Simulate Real-World Latency with Go

```go
func simulateRequest(id int) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond) // simulate latency
	fmt.Printf("Request %d completed in %v\n", id, time.Since(start))
}
```

📌 Run 1000 simulated requests concurrently, collect latency histogram using a `sync.Map`.

---

## ✅ Exercise 3: Simple HTTP API with Load Metrics

🔧 Create a REST API `/ping` that returns timestamp and tracks:

* Total requests served
* Average latency (rolling window)
* Peak concurrent connections

Use:

* `sync/atomic` for counters
* `expvar` or Prometheus metrics endpoint

---

## ✅ Exercise 4: Golang Client with HTTP1.1 vs HTTP2

Create two clients to call a mock service:

* One using `net/http`
* One using gRPC (`google.golang.org/grpc`)

Measure:

* Latency over 1000 calls
* Connection reuse vs new connection overhead

---

# 📚 Your Homework (Due Before Chapter 2)

1. ✅ Build the QPS estimator (CLI) in Go.
2. ✅ Simulate 1000 concurrent requests, log latency stats.
3. ✅ Build `/ping` REST API with metrics.
4. ✅ Compare HTTP1.1 vs gRPC latency and document findings.

📝 Submit:

* `main.go` files per task
* A markdown file `report.md` with:

  * How you measured QPS
  * Latency observations
  * Trade-offs you saw between HTTP1.1 and gRPC

---
