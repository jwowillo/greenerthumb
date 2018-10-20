.PHONY: build

all: greenerthumb test

greenerthumb: build
	$(call subcomponent,fan)
	$(call subcomponent,bullhorn)

test: build
	$(call subcomponent_test,fan)
	$(call subcomponent_test,bullhorn)

build:
	mkdir -p build

define subcomponent
	$(MAKE) -C $(1) $(1)
	cp -rf $(1)/build/. build/$(1)
endef

define subcomponent_test
	$(MAKE) -C $(1) test
	cp -rf $(1)/build/. build/$(1)
endef
