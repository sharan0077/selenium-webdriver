
ifndef prefix
prefix=/usr/local
endif

ifndef version
version=0.0.0
endif

PROGRAM_NAME=twist2

build:
	export GOPATH=$(GOPATH):`pwd` && cd src && go build && mv src ../selenium-webdriver-java

install:	
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver-java
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver-java/$(version)
	install -m 755 selenium-webdriver-java $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver-java/$(version)
	install -m 644 plugin.json $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver-java/$(version)

	
