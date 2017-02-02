#Fake Aggregator
Webserver for aggragation ints from different sources.

## Install

	go get github.com/tenmozes/fake-aggregator

## Run
	fake-aggregator [params]
	Params:
 	-deadline int
 		max response time from aggregator handles in ms (default 500)
 	-delay int
    	max aggregator handlers delay in ms (default 400)
 	-factor int
    	max number limits in aggregator handlers (default 50)
 	-port int
    	server port (default 8080)

## Tests

	cd $GOPATH/src/github.com/tenmozes/fake-aggregator/
	go test ./...

##Supported Endpoints
	/fibo - returns fibonacci sequence with random length
	/odd - returns odd numbers sequence with random length
	/primes - returns sequence of prime numbers with random length
	/rand - returns rand numbers sequence
	/numbers?u=url&u=url - <u> - array of URLS
		if url contains path which supported, server does request on own enviroment (for e.g. u=http://example.com/fibo - return fibo sequence) - endpoint aggregates sequences from sources sorts them and removes duplications
