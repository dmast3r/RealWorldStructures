from functools import partial
from math import ceil, log
from mmh3 import hash
from random import randint

class BloomFilter:
  """
  This class creates a Bloom Filter with methods to add strings and check if a string exists in the data structure.
  It requires as input the size of the bloom filter, which represents the number of bits present, and the expected number
  of items that will be inserted into the bloom filter.
  Using these two parameters, the implementation automatically calculates the optimal number of hash functions needed.
  """
  def __init__(self, size: int, expected_num_items: int):
    self.size = size
    self.expected_num_items = expected_num_items
    self.__num_hash_funcs = self.__compute_optimal_hash_funcs_count()
    self.__hash_functions = self.__create_hash_functions()

    self.__byte_array_size = int(ceil(self.size / 8))
    self.__byte_array = bytearray([0] * self.__byte_array_size)

  def __compute_optimal_hash_funcs_count(self):
    return int((self.size / self.expected_num_items) * log(2))

  def __create_hash_functions(self):
    return [partial(hash, seed=randint(1, 100000)) for _ in range(self.__num_hash_funcs)]

  def __get_byte_array_index_and_pos(self, hash_func, key):
    hash_val = hash_func(key) % self.size
    return hash_val // 8, hash_val % 8
  
  def add(self, key: str) -> None:
    """
      Adds a key to the Bloom Filter.
    """
    for hash_func in self.__hash_functions:
      byte_array_index, bit_position = self.__get_byte_array_index_and_pos(hash_func, key)
      self.__byte_array[byte_array_index] |= 1 << bit_position


  def exists(self, key: str) -> bool:
    """
      Checks if a key exists in the Bloom Filter. This may yield false positives but will never yield a false negative.
    """
    for hash_func in self.__hash_functions:
      byte_array_index, bit_position = self.__get_byte_array_index_and_pos(hash_func, key)
      if not (self.__byte_array[byte_array_index] & 1 << bit_position):
        return False

    return True
