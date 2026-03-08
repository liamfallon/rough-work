# Question 16(a)
# Examination Number:

# function definition used in part (v)
def is_anagram(w1, w2):
  if sorted(w1.lower()) == sorted(w2.lower()):
    return True
  else:
    return False

word1 = input("Enter the first word: ")
word2 = input("Enter the second word: ")

# test whether the sorted strings are the same as each other
# if the sorted strings are the same then they must be anagrams
if (sorted(word1.lower()) == sorted(word2.lower())):
  print(word1, "is an anagram of", word2)
else: 
  print(word1, "is NOT an anagram of", word2)

if (is_anagram (word1, word2)):
 print(word1, "is an anagram of", word2)
else: 
  print(word1, "is NOT an anagram of", word2)

phrase1 = input("Enter a phrase: ")
if (is_anagram (word1, phrase1.replace(" ",""))):
 print(word1, "is an anagram of", phrase1)
else: 
  print(word1, "is NOT an anagram of", phrase1)