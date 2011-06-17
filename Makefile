# Explicitly make targets phony, just in case
.PHONY : all pkgs cmds install clean nuke test bench gofmt rfc

MAKE += -s

PKGS = gimmock
CMDS = mock

# By default, build everything
all : pkgs cmds
	@

define recurse
@echo "$(3) $(1) $(2)" | sed -e 's/^all/build/'
@[[ "$1" != "cmd" ]] || $(MAKE) -C $(2) $(3)
@[[ "$1" != "pkg" ]] || $(MAKE) -f Makefile.$(2) $(3)

endef

define all_packages
$(foreach pkg,$(PKGS),$(call recurse,pkg,$(pkg),$(1)))
endef

define all_commands
$(foreach cmd,$(CMDS),$(call recurse,cmd,$(cmd),$(1)))
endef

pkgs :
	$(call all_packages,install)

cmds :
	$(call all_commands,all)

install :
	$(call all_packages,$@)
	$(call all_commands,$@)

test bench : install
	$(call all_packages,$@)

clean nuke :
	$(call all_packages,$@)
	$(call all_commands,$@)

# Format source files
gofmt :
	@gofmt -w `find . -name "*.go"`
