package monitor

import (
	"log"
	"main/resources"
	"sync"
)

// call somente
// uma unica vez
var once sync.Once

// Struct singleton
type HealthMonitor struct {
	Monitors          map[string]int
	CriticalResources []string
}

var (
	err     error
	Monitor HealthMonitor
)

type State struct {
	status int
	failed int
}

func (monitor *HealthMonitor) NewMonitor() *HealthMonitor {

	once.Do(func() {
		m := HealthMonitor{
			CriticalResources: make([]string, 0),
			Monitors:          make(map[string]int),
		}

		if err != nil {
			log.Println("Error PgConnect::", err)
		}
		return m
	})
}

func (r *resources.Resource) RegisterMonitor() *HealthMonitor {
	if !r.resourceType || !r.resourceName || !r.handler {
		return
	}

	// let resourceByType = this.monitors.get(type);
	// if (!resourceByType) {
	//   resourceByType = new Map();
	// }
	// resourceByType.set(name, handle);
	// this.monitors.set(type, resourceByType);
	// if (critical) {
	//   this.criticalResources.add(name);
	// }
}

func (hm *HealthMonitor) Check() (state State) {
	s := State{
		status: 200,
	}

	for key := range *hm.Monitors {
		r, _ := resources.GetResource(key)
		if r != nil {
			log.Println(r)
		}
	}
	//toDo: run all promises

	if err != nil {
		log.Println("error ao buscar email!")
		s.failed += 1
	}

	return s
}
