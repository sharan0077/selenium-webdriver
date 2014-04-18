
ifndef prefix
prefix=/usr/local
endif

ifndef version
version=0.0.0
endif

PROGRAM_NAME=twist2

build:
	export GOPATH=$(GOPATH):`pwd` && cd src && go build && mv src ../selenium-webdriver

install:	
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver
	install -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)
	install -m 755 selenium-webdriver $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)
	install -m 644 plugin.json $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)
	install  -m 755 -d $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)/skel/java/
	install  -m 644 skel/WebdriverBrowserFactory.java $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)/skel/java/
	install  -m 644 skel/WebdriverHooks.java $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)/skel/java/
	install -m 644 skel/webdriver.json $(prefix)/share/$(PROGRAM_NAME)/plugins/selenium-webdriver/$(version)/skel/java/
	
