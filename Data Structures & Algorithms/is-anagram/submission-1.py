class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_map = {}
        t_map = {}
        if len(s) != len(t):
            return False
        
        for i in s:
            s_map[i] = s_map.get(i,0)
            s_map[i] += 1
        
        for i in t:
            t_map[i] = t_map.get(i,0)
            t_map[i] += 1

        for i in s_map:
            if s_map.get(i) != t_map.get(i):
                return False
        return True