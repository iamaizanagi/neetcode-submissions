class Solution:

    def encode(self, strs: List[str]) -> str:
        enStr = ""
        for i in range(0,len(strs)):
            enStr = enStr+str(len(strs[i]))+"#"+strs[i]
        print(enStr)
        return enStr

    def decode(self, s: str) -> List[str]:
        apList = []
        i = 0
        while i < len(s):
            j = i 
            while s[j] != "#":
                j += 1
            
            length = int(s[i:j])
            apList.append(s[(j+1):(j+length+1)])
            i = j + length + 1
        return apList

