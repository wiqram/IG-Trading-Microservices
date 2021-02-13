const { Builder, Capabilities } = require('selenium-webdriver');
const chrome = require('selenium-webdriver/chrome');
const chromedriver = require('chromedriver');

chrome.setDefaultService(new chrome.ServiceBuilder(chromedriver.path).build());
const URL = process.env.URL ? process.env.URL : 'Test URL Missing';

this.getDriver =
    new Builder()
        .withCapabilities(Capabilities.chrome())
        .build();

this.getHomePage = this.getDriver.get(URL);



