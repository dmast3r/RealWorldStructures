from functools import partial
from math import ceil, log
from mmh3 import hash
from random import randint

class BloomFilter:
  """
  Creates a Bloom Filter with methods to add string and check if a string exists in the data structure.
  It works by taking in input the size of the bloom filter, which represents the number of bits that we want to be
  present and the expected number of items tbat we expect to insert in the bloom filter.
  Using these two parameters, the implementation auto calculates the optimnal number of hash functions that would be needed.
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
      Adds a key to the Bloom Filter
    """
    for hash_func in self.__hash_functions:
      byte_array_index, bit_position = self.__get_byte_array_index_and_pos(hash_func, key)
      self.__byte_array[byte_array_index] |= 1 << bit_position


  def exists(self, key: str) -> bool:
    """
      Check if a key doesn't exist in the Bloom Filter, may give false positive but will never give false negative
    """
    for hash_func in self.__hash_functions:
      byte_array_index, bit_position = self.__get_byte_array_index_and_pos(hash_func, key)
      if not (self.__byte_array[byte_array_index] & 1 << bit_position):
        return False

    return True