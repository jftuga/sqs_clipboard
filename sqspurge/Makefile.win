
PROG = sqspurge

# to make a non-colsole binary, use: -ldflags "-s -w -H=windowsgui"

$(PROG) : cmd.go $(PROG).syso
	go build -ldflags "-s -w"

$(PROG).syso: $(PROG).ico
	rsrc -ico $(PROG).ico -o $(PROG).syso

clean:
	rm -f $(PROG) $(PROG).exe $(PROG).syso
