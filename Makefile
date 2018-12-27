.PHONY: test

greenerthumb: | build
	$(call subcomponent,fan,fan)
	$(call subcomponent,bullhorn,bullhorn)
	$(call subcomponent,plot,plot)
	$(call subcomponent,message,message)
	$(call subcomponent,log,log)
	$(call subcomponent,sense,sense)
	$(call subcomponent,process,process)
	$(call subcomponent,store,store)
	$(call subcomponent,disclosure,disclosure)

device: | build
	env GOOS=linux GOARCH=arm $(call subcomponent,message,bytes)
	env GOOS=linux GOARCH=arm $(call subcomponent,bullhorn,publish)
	env GOOS=linux GOARCH=arm $(call subcomponent,store,store)
	env GOOS=linux GOARCH=arm $(call subcomponent,disclosure,disclosure)

disclosure: device

air: device | build
	env GOOS=linux GOARCH=arm $(call subcomponent,sense,air)

soil: device | build
	env GOOS=linux GOARCH=arm $(call subcomponent,sense,soil)

test: | build
	$(call subcomponent_test,fan)
	$(call subcomponent_test,bullhorn)
	$(call subcomponent_test,plot)
	$(call subcomponent_test,message)
	$(call subcomponent_test,process)
	$(call subcomponent_test,store)
	$(call subcomponent_test,disclosure)

build:
	mkdir -p build

clean:
	rm -rf build

define subcomponent
	$(MAKE) -C $(1) $(2)
	cp -rf $(1)/build/. build/$(1)
endef

define subcomponent_test
	$(MAKE) -C $(1) test
	cp -rf $(1)/build/. build/$(1)
endef
