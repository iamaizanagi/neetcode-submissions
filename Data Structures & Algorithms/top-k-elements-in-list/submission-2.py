class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        countDict = {}

        klist = [[] for i in range(len(nums)+1)]
        final_arr = []

        for i in nums:
            countDict[i] = countDict.get(i,0)
            countDict[i] += 1
        
        for key, value in countDict.items():
            klist[value].append(key)
        
        for j in range(len(klist)-1, 0,-1):
            for item in klist[j]:
                final_arr.append(item)
                if len(final_arr) == k:
                    return final_arr

                
        return final_arr