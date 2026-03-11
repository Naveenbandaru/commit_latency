# commit_latency
**Transaction Batching for Low Latency Commit Processing in Distributed Systems**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal on Science and Technology (IJSAT)
- **Publication Date:** July, 2021
- **ISSN:** E-ISSN: 2229-7677
- **DOI:**
- **Impact Factor:** 9.88

### Abstract
Distributed transaction systems often experience high commit latency because each transaction is processed independently with repeated coordination and synchronization among nodes. This work examines the impact of immediate commit processing on latency as cluster size increases. A batching based commit approach is introduced where multiple transactions are processed collectively to reduce repeated coordination overhead. Experimental analysis across different cluster sizes shows that grouped commit processing significantly lowers commit latency and improves scalability in distributed environments.

### Core Technical Contributions
- **Batch Based Commit Processing Approach:**  
Introduced a commit method that groups multiple transactions into a single coordination cycle, reducing repeated commit operations and lowering synchronization overhead in distributed transaction systems.

- **Reduced Coordination Overhead:**  
Designed a commit model that minimizes repeated communication and synchronization between coordinator and participant nodes, improving efficiency during the commit phase of transactions.

- **Distributed Transaction Simulation Model:** 
Implemented a distributed transaction processing environment using Go based concurrent workers to simulate coordinator participant communication and analyze commit behavior across multiple nodes.

- **Scalability Analysis Across Cluster Sizes:**  
Evaluated commit latency across clusters with 3, 5, 7, 9, and 11 nodes to study how batching influences scalability and transaction completion performance.

### Practical Significance and Impact
- **Lower Commit Latency:**
Batch based commit processing significantly decreases the time required to finalize transactions by reducing repeated coordination cycles that occur in conventional immediate commit protocols.

- **Improved Transaction Processing Efficiency:**  
Processing multiple transactions together reduces communication rounds and synchronization delays, enabling faster completion of distributed transactions and better utilization of system resources.

- **Better Scalability for Distributed Systems:**  
Commit latency grows gradually with cluster expansion because fewer coordination rounds are required, enabling distributed systems to maintain stable transaction processing performance.

- **Applicability to Distributed Platforms:**  
The approach benefits distributed databases, cloud transaction systems, financial platforms, and microservice architectures that require efficient transaction completion and scalable commit processing.
 
### Experimental Results (Summary)

  | Nodes | Local Telemetry CPU | Telemetry corelation CPU | Improvment (%)  |
  |-------|---------------------| -------------------------| ----------------|
  | 3     |  72                 | 54                       | 25.0           |
  | 5     |  70                 | 50                       | 28.6           |
  | 7     |  68                 | 47                       | 30.9           |
  | 9     |  67                 | 45                       | 32.8           |
  | 11    |  66                 | 43                       | 34.8           |

### Citation
Transaction Batching for Low Latency Commit Processing in Distributed Systems
* Naveen Kumar Bandaru
* International Journal on Science and Technology (IJSAT) 
* ISSN E-ISSN: 2229-7677
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com






