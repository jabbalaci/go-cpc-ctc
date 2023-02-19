cat:
	cat Makefile

cpc:
	go run ./cmd/cpc

ctc:
	go run ./cmd/ctc

install-cpc:
	go install ./cmd/cpc

install-ctc:
	go install ./cmd/ctc

install-both:
	go install ./cmd/cpc
	go install ./cmd/ctc
