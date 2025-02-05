### **Day 2: Designing a URL Shortener Like TinyURL**
Today, we’ll focus on how to design a **URL Shortener** system, applying system design principles you learned yesterday.

---
## **📖 Concepts to Cover Today**
1. **Understanding Functional & Non-functional Requirements**
2. **High-level Architecture**
3. **Database Design**
4. **Choosing a Hashing Algorithm for Short URLs**
5. **Handling Redirection Efficiently**
6. **Scaling & Caching Strategies**
7. **Implementing Rate Limiting**
8. **Fault Tolerance & Monitoring**
9. **Practical Implementation (Golang)**

---

## **1️⃣ Defining Requirements**
### **📌 Functional Requirements**
- A user provides a long URL → System generates a unique short URL.
- The short URL redirects users to the original URL.
- Track analytics: number of clicks, creation timestamp, etc.

### **📌 Non-functional Requirements**
- **Scalability**: Handle millions of requests per second.
- **Availability**: Minimal downtime.
- **Performance**: Short URL resolution in milliseconds.
- **Security**: Prevent abuse & spam.

---

## **2️⃣ High-Level Architecture**
📌 **Components:**
1. **Load Balancer** - Distributes traffic to backend servers.
2. **API Layer** - Receives user requests.
3. **Database (SQL/NoSQL)** - Stores mappings between short & long URLs.
4. **Cache (Redis)** - Speeds up lookups.
5. **Background Workers** - For analytics & expired link cleanup.
6. **Monitoring & Logging** - Tracks usage.

---

## **3️⃣ Database Design**
### **Schema:**
```
Table: urls
----------------------------
id         | Primary Key
long_url   | VARCHAR(2083)
short_url  | VARCHAR(10) (Unique)
created_at | TIMESTAMP
hits       | INT
expiry     | TIMESTAMP (Optional)
```

📌 **Choosing DB Type:**
- **SQL (PostgreSQL)**: ACID compliance, relational data.
- **NoSQL (Cassandra/DynamoDB)**: High read/write throughput.
- **Hybrid Approach**: SQL for transactional data, Redis for fast lookups.

---

## **4️⃣ Choosing a Short URL Generation Strategy**
📌 **Approaches:**
1. **Base62 Encoding**  
   - Convert an integer ID to a short URL using characters [a-zA-Z0-9].
   - Example: `https://short.ly/abc123`

2. **Hashing (MD5/SHA256 + Collision Handling)**
   - Example: `hash("https://example.com") → 5d41402a`
   - Collision risk: We may need to check the DB for uniqueness.

3. **UUID-based Approach**
   - Use a **UUID** or **Snowflake ID** (Twitter's approach) for uniqueness.

✅ **Best Approach:** **Base62 encoding** is commonly used for scalability.

---

## **5️⃣ Handling URL Redirection**
- When a user accesses `short.ly/xyz123`, the system should:
  1. Look up `xyz123` in **Redis cache**.
  2. If **cache miss**, fetch from the database.
  3. Redirect to the long URL.
  4. Update analytics asynchronously.

---

## **6️⃣ Scaling & Caching Strategies**
### **📌 Strategies to Handle Load**
1. **Caching Layer (Redis)**
   - Store frequently accessed short URLs in Redis.
   - Use **LRU (Least Recently Used) eviction policy**.

2. **Sharding Strategy**
   - Store different short URLs across multiple databases.
   - Use **consistent hashing** to determine the DB partition.

3. **Read Replicas**
   - Set up read replicas to distribute read requests.

---

## **7️⃣ Rate Limiting & Security**
To prevent abuse (spamming URL generation), we need **rate limiting**:
- **IP-based throttling** using **Redis Leaky Bucket Algorithm**.
- **Authentication** for API access.
- **Expiration time for short URLs**.

---

## **8️⃣ Fault Tolerance & Monitoring**
- **Data Backups:** Periodic database snapshots.
- **Circuit Breaker Pattern:** Handle Redis failures.
- **Prometheus + Grafana**: Monitor requests, errors, and latency.

---

## **9️⃣ Practical Implementation (Golang)**
### **🛠️ Mini Project: Building a URL Shortener in Go**
We’ll implement a basic version using:
✅ **Gin (HTTP framework)**  
✅ **PostgreSQL (DB)**  
✅ **Redis (Cache layer)**  
✅ **Base62 encoding**  

Here’s a high-level breakdown:
1. **Create a Golang API with Gin**
2. **Implement URL shortener logic**
3. **Use Redis for caching**
4. **Store URLs in PostgreSQL**
5. **Redirect users from short to long URLs**

---

## **⏳ Schedule for Today**
| Time | Task |
|------|------|
| 5:00 AM - 6:00 AM | Read & understand URL Shortener design concepts |
| 6:00 AM - 7:00 AM | Plan DB schema & API endpoints |
| 7:00 AM - 8:00 AM | Implement Base62 encoding in Golang |
| 8:00 AM - 9:00 AM | Set up PostgreSQL & Redis, store URLs |
| 9:00 AM - 10:00 AM | Build redirection logic & test API |
| 9:00 PM - 11:00 PM | Optimize caching, rate limiting, and monitoring |

---

## **💻 Starter Code for URL Shortener**
You can find The full Golang project with API routes, database integration, and Redis caching.