class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hashMap = {}

        for i in range(0,len(nums)):
            if nums[i] not in hashMap:
                hashMap[nums[i]] = hashMap.get(nums[i],0)
            hashMap[nums[i]] += 1
        
        for key,value in hashMap.items():
            if value > 1:
                return True
        return False