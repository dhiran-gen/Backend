# **🛠️ Task 3: Database Sharding Simulation**  

## **⚡ Overview**  
Database **sharding** is the practice of **splitting a large database** into smaller, **faster, and more manageable** pieces (shards) to improve **performance, scalability, and availability**.  

We will simulate **database sharding** using:  
✅ **PostgreSQL** (for relational sharding) OR  
✅ **MongoDB** (for NoSQL sharding)  

---

# **🔹 Option 1: PostgreSQL Sharding (Table Partitioning)**
### **Step 1: Install PostgreSQL**
```bash
sudo apt update && sudo apt install postgresql -y
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

### **Step 2: Create a PostgreSQL Database**
Login to PostgreSQL:
```bash
sudo -u postgres psql
```
Create a **database**:
```sql
CREATE DATABASE shard_db;
\c shard_db
```

### **Step 3: Create Parent Table**
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    region TEXT NOT NULL
) PARTITION BY LIST (region);
```

### **Step 4: Create Shard Tables**
```sql
CREATE TABLE users_us PARTITION OF users FOR VALUES IN ('US');
CREATE TABLE users_eu PARTITION OF users FOR VALUES IN ('EU');
```

### **Step 5: Insert Data into Shards**
```sql
INSERT INTO users (name, email, region) VALUES ('Alice', 'alice@example.com', 'US');
INSERT INTO users (name, email, region) VALUES ('Bob', 'bob@example.com', 'EU');
```

### **Step 6: Query Data**
```sql
SELECT * FROM users WHERE region = 'US';
```

---

# **🔹 Option 2: MongoDB Sharding (Sharded Cluster)**
### **Step 1: Install MongoDB**
```bash
sudo apt update
sudo apt install -y mongodb
sudo systemctl start mongod
sudo systemctl enable mongod
```

### **Step 2: Enable Sharding in MongoDB**
Connect to MongoDB shell:
```bash
mongo
```

Enable sharding on the database:
```javascript
sh.enableSharding("shard_db")
```

### **Step 3: Create Sharded Collection**
```javascript
sh.shardCollection("shard_db.users", { region: "hashed" })
```

### **Step 4: Insert Data into Shards**
```javascript
use shard_db

db.users.insertMany([
    { name: "Alice", email: "alice@example.com", region: "US" },
    { name: "Bob", email: "bob@example.com", region: "EU" }
])
```

### **Step 5: Query Data**
```javascript
db.users.find({ region: "US" })
```

---

# **🚀 Which One to Choose?**
| **Database** | **Best for** | **Sharding Type** | **Use Case** |
|-------------|-------------|-----------------|--------------|
| **PostgreSQL** | Relational data | Horizontal partitioning | Multi-region RDBMS |
| **MongoDB** | NoSQL & high-scale apps | Range & hash-based sharding | Large-scale NoSQL data |

Would you like a deeper dive into any step? Raise an issue and we'll resolve together.