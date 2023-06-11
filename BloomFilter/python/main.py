from random import randint
from utils.BloomFilter import BloomFilter
from utils.rand_strings import generate_random_string

# create a Bloom Filter of 1M bits (which is about 120 KB) and about 100k items.
batch_size = 200000
bloom_filter = BloomFilter(1000000, batch_size // 2)

# for testing the false positivity rate.
all_strings = []
bloom_filter_inserts = set()

# Randomly generate a total of 200k items and randomly insert about half of them.
for i in range(batch_size):
  random_string = generate_random_string(10)
  if randint(0, 1) & 1:
    bloom_filter.add(random_string)
    bloom_filter_inserts.add(random_string)
  all_strings.append(random_string)

false_positives = 0
for string in all_strings:
  if bloom_filter.exists(string) and string not in bloom_filter_inserts:
    false_positives += 1

print('Percentage of false positives observed:', false_positives / batch_size * 100)