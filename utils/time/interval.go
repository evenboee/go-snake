package time

import "time"

func SetInterval(f func(chan bool), duration time.Duration) {
	quit := make(chan bool)

	go func() {
		for {
			time.Sleep(duration)
			go f(quit)
		}
	}()

	for v := false; !v; v = <-quit {

	}
}
