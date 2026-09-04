class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hMap = {}
        resultArr = []

        for i in range(0,len(nums)):
            res = target - nums[i]
            hMap[res] = i
        
        print(hMap)
        for j in range(0, len(nums)):
            if nums[j] in hMap and j != hMap[nums[j]]:
                resultArr.append(j)
                resultArr.append(hMap[nums[j]])
                break
        return resultArr
