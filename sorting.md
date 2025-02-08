Here's a detailed table comparing different **sorting** and **searching** algorithms, including their **time complexity** and **space complexity**:  

---

### **Sorting Algorithms Comparison**
| Algorithm         | Best Case | Average Case | Worst Case | Space Complexity | Stability  | In-Place? |
|------------------|-----------|--------------|------------|------------------|------------|-----------|
| **Bubble Sort**   | 𝑂(𝑛)       | 𝑂(𝑛²)         | 𝑂(𝑛²)      | 𝑂(1)            | ✅ Stable  | ✅ Yes      |
| **Selection Sort** | 𝑂(𝑛²)    | 𝑂(𝑛²)         | 𝑂(𝑛²)      | 𝑂(1)            | ❌ Unstable  | ✅ Yes      |
| **Insertion Sort** | 𝑂(𝑛)     | 𝑂(𝑛²)         | 𝑂(𝑛²)      | 𝑂(1)            | ✅ Stable  | ✅ Yes      |
| **Merge Sort**    | 𝑂(𝑛 log 𝑛) | 𝑂(𝑛 log 𝑛)    | 𝑂(𝑛 log 𝑛) | 𝑂(𝑛)            | ✅ Stable  | ❌ No       |
| **Quick Sort**    | 𝑂(𝑛 log 𝑛) | 𝑂(𝑛 log 𝑛)    | 𝑂(𝑛²)      | 𝑂(log 𝑛) (avg) | ❌ Unstable  | ✅ Yes      |
| **Heap Sort**     | 𝑂(𝑛 log 𝑛) | 𝑂(𝑛 log 𝑛)    | 𝑂(𝑛 log 𝑛) | 𝑂(1)            | ❌ Unstable  | ✅ Yes      |
| **Counting Sort** | 𝑂(𝑛 + 𝑘) | 𝑂(𝑛 + 𝑘)     | 𝑂(𝑛 + 𝑘)   | 𝑂(𝑛 + 𝑘)        | ✅ Stable  | ❌ No       |
| **Radix Sort**    | 𝑂(𝑛𝑘)     | 𝑂(𝑛𝑘)         | 𝑂(𝑛𝑘)      | 𝑂(𝑛 + 𝑘)        | ✅ Stable  | ❌ No       |
| **Bucket Sort**   | 𝑂(𝑛 + 𝑘) | 𝑂(𝑛 + 𝑘)     | 𝑂(𝑛²)      | 𝑂(𝑛 + 𝑘)        | ✅ Stable  | ❌ No       |

🔹 **Stable Sorting**: Maintains the order of equal elements (e.g., Merge Sort, Bubble Sort, Counting Sort).  
🔹 **In-Place Sorting**: Uses constant extra space (e.g., Quick Sort, Heap Sort, Bubble Sort).  

---

### **Searching Algorithms Comparison**
| Algorithm         | Best Case | Average Case | Worst Case | Space Complexity | Search Type |
|------------------|----------|--------------|------------|------------------|-------------|
| **Linear Search** | 𝑂(1)     | 𝑂(𝑛)         | 𝑂(𝑛)      | 𝑂(1)            | ✅ Unsorted  |
| **Binary Search** | 𝑂(1)     | 𝑂(log 𝑛)    | 𝑂(log 𝑛)  | 𝑂(1)            | ✅ Sorted    |
| **Jump Search**   | 𝑂(1)     | 𝑂(√𝑛)        | 𝑂(√𝑛)     | 𝑂(1)            | ✅ Sorted    |
| **Interpolation Search** | 𝑂(1) | 𝑂(log log 𝑛) | 𝑂(𝑛)     | 𝑂(1)            | ✅ Sorted    |
| **Exponential Search** | 𝑂(1) | 𝑂(log 𝑛)    | 𝑂(log 𝑛)  | 𝑂(1)            | ✅ Sorted    |
| **Ternary Search** | 𝑂(1)    | 𝑂(log 𝑛)    | 𝑂(log 𝑛)  | 𝑂(1)            | ✅ Sorted    |
| **Fibonacci Search** | 𝑂(1)  | 𝑂(log 𝑛)    | 𝑂(log 𝑛)  | 𝑂(1)            | ✅ Sorted    |

🔹 **Linear Search** works for **unsorted** data but is slow (𝑂(𝑛)).  
🔹 **Binary Search** is the most efficient for **sorted** data (𝑂(log 𝑛)).  
🔹 **Jump Search & Interpolation Search** optimize binary search in special cases.  

---

### **Key Takeaways**
1. **For small data sets**: Use **Insertion Sort** (𝑂(𝑛²)), it's simple and efficient.
2. **For large data sets**: Use **Merge Sort** or **Quick Sort** (𝑂(𝑛 log 𝑛)).
3. **For nearly sorted arrays**: Use **Insertion Sort** or **Bubble Sort**.
4. **For searching**:  
   - **Unsorted Data** → **Linear Search**  
   - **Sorted Data** → **Binary Search** (best general-purpose search)
   - **Large Range Data** → **Jump/Interpolation Search**
