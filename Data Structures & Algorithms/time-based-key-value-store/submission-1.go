type entry struct{
	value string
	time int
}
type TimeMap struct {
	m map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key]=append(this.m[key], entry{value,timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.m[key]; !exists{
		return ""
	}

	pairs := this.m[key]
	l, r := 0, len(pairs) - 1

	for l <= r{
		mid := (l+r)/2
		if pairs[mid].time <= timestamp{
			if mid == len(pairs) - 1 || pairs[mid+1].time > timestamp{
				return pairs[mid].value
			}
			l = mid + 1
		}else{
			r = mid - 1
		}
	}
	return ""
}
