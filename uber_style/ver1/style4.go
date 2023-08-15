package main

import (
	"fmt"
)

type Trip struct {
	name string
}

type Driver struct {
	level string
	trips []Trip
}

func (d *Driver) SetTripsWrong(trips []Trip) {
	d.trips = trips
}

// recommend
func (d *Driver) SetTrips(trips []Trip) {
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

type Stats struct {
	mu sync.Mutex

	counters map[string]int
}

func (s *Stats) SnapshotBad() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}

func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	tmp := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		tmp[k] = v
	}

	return tmp
}

func main() {
	//slice
	var d1, d2 Driver
	fmt.Println(d1) //d1 isn't nil

	trips := []Trip{
		{name: "John"},
		{name: "Mary"},
		{name: "Jane"},
	}

	d1.SetTripsWrong(trips)
	d2.SetTrips(trips)
	fmt.Println(d1)
	fmt.Println(d2)

	// 你是要修改 d1.trips 吗？

	trips[0].name = "Jackson"
	fmt.Println(d1) //d1's trips is changed
	fmt.Println(d2) // d2's trips isn't changed

	println("====map=====")

	stats := Stats{
		counters: map[string]int{
			"John": 0,
			"Mary": 1,
		},
	}

	snapshot := stats.SnapshotBad()
	// snapshot 不再受互斥锁保护
	// 因此对 snapshot 的任何访问都将受到数据竞争的影响
	// 影响 stats.counters
}
