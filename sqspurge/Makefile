
PROG = sqspurge

# to make a non-colsole binary, use: -ldflags "-s -w -H=windowsgui"

$(PROG) : cmd.go
	go build -ldflags "-s -w"


clean:
	rm -f $(PROG) $(PROG).exe
