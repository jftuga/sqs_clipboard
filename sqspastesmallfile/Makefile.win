
PROG = sqspastesmallfile

$(PROG) : cmd.go
	go build -ldflags "-s -w"

clean:
	rm -f $(PROG) $(PROG).exe $(PROG).syso
