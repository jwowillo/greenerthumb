greenerthumb: clean build
	$(call subcomponent,fan,fan)
	$(call subcomponent,bullhorn,bullhorn)
	$(call subcomponent,plot,plot)
	$(call subcomponent,message,message)
	$(call subcomponent,log,log)
	$(call subcomponent,sense,sense)
	$(call subcomponent,process,process)
	$(call subcomponent,store,store)

air: clean build
	env GOOS=linux GOARCH=arm $(call subcomponent,sense,air)
	env GOOS=linux GOARCH=arm $(call subcomponent,message,bytes)
	env GOOS=linux GOARCH=arm $(call subcomponent,bullhorn,publish)
	env GOOS=linux GOARCH=arm $(call subcomponent,store,store)

soil: clean build
	env GOOS=linux GOARCH=arm $(call subcomponent,sense,soil)
	env GOOS=linux GOARCH=arm $(call subcomponent,message,bytes)
	env GOOS=linux GOARCH=arm $(call subcomponent,bullhorn,publish)
	env GOOS=linux GOARCH=arm $(call subcomponent,store,store)

test: clean build
	$(call subcomponent_test,fan)
	$(call subcomponent_test,bullhorn)
	$(call subcomponent_test,plot)
	$(call subcomponent_test,message)
	$(call subcomponent_test,process)
	$(call subcomponent_test,store)

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
