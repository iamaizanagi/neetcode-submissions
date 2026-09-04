class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        strs_map = defaultdict(list)
        final_list = []

        for i in range(0, len(strs)):
            s= sorted(strs[i])
            s = "".join(s)
            appendList = []

            if s not in strs_map:
                strs_map[s].append(strs[i])
            else:
                strs_map[s].append(strs[i])
                
        for j in strs_map:
            final_list.append(strs_map[j])

        return final_list