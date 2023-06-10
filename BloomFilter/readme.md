# Bloom Filter

## Overview
Bloom Filters are a probabilistic data structure that are used to test whether an element is a member of a set. It is optimized to be space-efficient but at the cost of allowing false positives.

## Why Bloom Filters?
In many applications, the amount of data can be so large that it doesn't fit into the memory. Storing the data in secondary memory (like disk) can be an option, but it's slower compared to the main memory (RAM). Bloom Filters offer a good compromise by allowing a small rate of false positives in exchange for a very compact representation of the set.

## Operations Supported
Bloom Filters primarily support two operations:

1. `add(key)`: Adds an element to the set.
2. `exists(key)`: Checks whether an element is in the set. This operation may return false positives but it will never return a false negative. 

## Hash Algorithm
This implementation uses the MurmurHash (specifically `mmh3`) hashing algorithm. MurmurHash is a non-cryptographic hash function which is primarily used for hash-based lookups. It is renowned for its simplicity and speed and it also exhibits a good distribution of hash values which is ideal for the purpose of a Bloom Filter.

## Real-world Applications
1. **Web browsers**: Google Chrome and Mozilla Firefox use Bloom Filters for their "Safe Browsing" feature. It's used to check whether a URL is in a list of known malicious URLs.
2. **Databases**: Bloom Filters are used in databases like Cassandra and Redis for efficient membership queries and reducing disk lookups.
3. **Network Routers**: Network routers use Bloom Filters for packet routing, allowing for efficient data lookup and transfer.

## Usage
This repository contains an implementation of the Bloom Filter in Python. You can use this as a base and modify or extend it to fit your own needs.

To run the sample implementation, navigate to the `python` directory and run:

```shell
pip install -r requirements.txt
python main.py
