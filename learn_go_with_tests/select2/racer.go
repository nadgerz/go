package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

/*
If you recall from the concurrency chapter, you can wait for
values to be sent to a channel with myVar := <-ch.

This is a blocking call, as you're waiting for a value.

What select lets you do is wait on multiple channels.

The first one to send a value "wins" and the code underneath the case is executed.

We use ping in our select to set up two channels for each of our URLs.

Whichever one writes to its channel first will have its code executed in the select,
which results in its URL being returned (and being the winner).
*/
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
We have defined a function ping which creates a chan bool and returns it.

In our case, we don't really care what the type sent in the channel,
we just want to send a signal to say we're finished so booleans are fine.

Inside the same function, we start a goroutine which will send a signal
into that channel once we have completed http.Get(url).
*/
func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
