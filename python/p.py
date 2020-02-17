strg = "abcdefg"

perm = set()
temp = set()

strl = list(strg)

for i in range(len(strg)-1):
	s = strg[i:i+2]
	perm.add(s)
	perm.add(s[1]+s[0])
perm = list(perm)
for x in perm:
	x=list(x)
	for i in range(len(x)):
		x.insert(i, 'a')
	perm = set(x).union(set(perm))

print(perm)

def permutations(string):
	char = string.pop()



# print(permutations(strl))