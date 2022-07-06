// main.go
package main

import (
	"fmt"
	"main/monitor"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// viper.SetConfigName("config") // name of config file (without extension)
	// viper.AddConfigPath(".")      // optionally look for config in the working directory
	// err := viper.ReadInConfig()   // Find and read the config file
	// if err != nil {               // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// listenAt := fmt.Sprintf(":%s", viper.Get("PORT"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, World!")
	// })
	// resources.GetResource()
	// log.Printf("Open the following URL in the browser: http://localhost:%\n", viper.Get("PORT"))
	// log.Fatal(http.ListenAndServe(listenAt, nil))

	// goroutine
	go func() {

		for i := 0; i < 100; i++ {

			hm := monitor.Monitor.NewMonitor()
			hm.RegisterMonitor()

			time.Sleep(time.Millisecond * 750)

			// varias goroutines
			go func() {
				time.Sleep(time.Millisecond * 300)
				monitor.Monitor.Check(&hm)
				fmt.Println("Status:", hm)
			}()
		}
	}()

	fmt.Scanln()
}
