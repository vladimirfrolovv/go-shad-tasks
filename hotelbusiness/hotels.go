//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	start := make(map[int]int)
	for _, g := range guests {
		for j := g.CheckInDate; j <= g.CheckOutDate; j++ {
			start[j]++
		}
		start[g.CheckOutDate]--
	}

	load := make([]Load, len(start))
	i := 0
	for k, v := range start {
		load[i].StartDate = k
		load[i].GuestCount = v
		i++
	}
	sort.Slice(load[:], func(i, j int) bool {
		return load[i].StartDate < load[j].StartDate
	})

	tmp := make([]Load, len(load))
	j, k := 0, 0
	for j < len(load) {
		tmp[k] = load[j]
		for j+1 < len(load) && load[j].GuestCount == load[j+1].GuestCount {
			j++
		}
		j++
		k++
	}
	return tmp[:k]
}
