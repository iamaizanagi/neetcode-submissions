class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)
        final_arr = []
        for i in range(0,len(strs)):
            count = [0]*26
            for j in strs[i]:
                count[ord(j)-ord('a')] += 1
            if tuple(count) not in res:
                res[tuple(count)].append(strs[i])
            else:
                res[tuple(count)].append(strs[i])
        for key in res:
            final_arr.append(res[key])
        return final_arr