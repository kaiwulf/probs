"""
	Problem: Find all the permuations of str1 contained in str2
"""

import itertools

def perm(string):
	return [''.join(x) for x in itertools.permutations(string)]

str1 = 'abbc'
str2 = 'cbabadcbbabbcbabaabccbabc'

str_perm = perm(str1)

str_range = range(len(str2))

for x in str_range:
	print(str2[x:x+4])
	if str2[x:x+4] in str_perm:
		print("permutation found at location {0} - {1}, {2}".format(x, x+4, str2[x:x+4]))
