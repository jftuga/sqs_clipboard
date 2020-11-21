
TOPTARGETS := all clean

SUBDIRS := sqscopy sqspaste sqspurge sqscopysmallfile sqspastesmallfile

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)
