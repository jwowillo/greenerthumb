.PHONY: test

all: | build
	$(call subcomponent,fan,fan)
	$(call subcomponent,bullhorn,bullhorn)
	$(call subcomponent,plot,plot)
	$(call subcomponent,message,message)
	$(call subcomponent,log,log)
	$(call subcomponent,device,device)
	$(call subcomponent,process,process)
	$(call subcomponent,store,store)
	$(call subcomponent,disclosure,disclosure)

test: | build
	$(call subcomponent_test,fan)
	$(call subcomponent_test,bullhorn)
	$(call subcomponent_test,plot)
	$(call subcomponent_test,message)
	$(call subcomponent_test,process)
	$(call subcomponent_test,store)
	$(call subcomponent_test,disclosure)

host: clean | build
	$(call subcomponent,bullhorn/pubsub,client)
	$(call subcomponent,message,json)
	$(call subcomponent,log,log)

logger: host | build
	mkdir -p build/run
	cp run/logger/logger build/run/.
	cp run/logger/README.md build/run/.

plotter: host | build
	$(call subcomponent,plot,plot)
	mkdir -p build/run
	cp run/plotter/plotter build/run/.
	cp run/plotter/README.md build/run/.

device: clean | build
	env GOOS=linux GOARCH=arm $(call subcomponent,message,bytes)
	env GOOS=linux GOARCH=arm $(call subcomponent,message,header)
	env GOOS=linux GOARCH=arm $(call subcomponent,bullhorn/pubsub,server)
	env GOOS=linux GOARCH=arm $(call subcomponent,bullhorn/broadcast,server)
	env GOOS=linux GOARCH=arm $(call subcomponent,disclosure,disclosure)
	env GOOS=linux GOARCH=arm $(call subcomponent,store,store)

air-sensor: device | build
	env GOOS=linux GOARCH=arm $(call subcomponent,device,air-sensor)
	mkdir -p build/run
	cp run/air-sensor/air-sensor build/run/.
	cp run/air-sensor/README.md build/run/.

soil-sensor: device | build
	env GOOS=linux GOARCH=arm $(call subcomponent,device,soil-sensor)
	mkdir -p build/run
	cp run/soil-sensor/soil-sensor build/run/.
	cp run/soil-sensor/README.md build/run/.

clean:
	rm -rf build
	$(MAKE) -C bullhorn clean
	$(MAKE) -C disclosure clean
	$(MAKE) -C fan clean
	$(MAKE) -C log clean
	$(MAKE) -C message clean
	$(MAKE) -C plot clean
	$(MAKE) -C process clean
	$(MAKE) -C device clean
	$(MAKE) -C store clean

build:
	mkdir -p build

define subcomponent
	$(MAKE) -C $(1) $(2)
	mkdir -p build/$(1)
	cp -rf $(1)/build/. build/$(1)
endef

define subcomponent_test
	$(MAKE) -C $(1) test
	mkdir -p build/$(1)
	cp -rf $(1)/build/. build/$(1)
endef
