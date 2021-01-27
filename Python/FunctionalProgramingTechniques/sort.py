myList = [[1,2], [4,5], [1,3],[3,1]]
def test(sublist):
    return sublist[0], -sublist[1]
a = sorted(myList, key=test)
print(a)


b = sorted(dictf, key=lambda a:  for (i in a))
print(b)

def a():
        return 'I am in A'
def b():
    return 'I am in B!'
    
while True:
    choice = input('Enter your choice: ').strip()
    if not choice:
        break
    if choice in choices: 
        print(choices[choice]())
    else:
        print(f'Choice {chioce} is not valid)