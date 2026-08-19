class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        d = {}
        for i in range(len(nums)):
            d[nums[i]] = d.get(nums[i],0)
            d[nums[i]] += 1
        
        for key in d:
            if d[key] > 1:
                #print(d[key])
                return True
        return False