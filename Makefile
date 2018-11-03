.PHONY: build

all: greenerthumb test

greenerthumb: build
	$(call subcomponent,fan)
	$(call subcomponent,bullhorn)
	$(call subcomponent,plot)
	$(call subcomponent,message)
	$(call subcomponent,log)
	$(call subcomponent,sense)

test: build
	$(call subcomponent_test,fan)
	$(call subcomponent_test,bullhorn)
	$(call subcomponent_test,plot)
	$(call subcomponent_test,message)

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
