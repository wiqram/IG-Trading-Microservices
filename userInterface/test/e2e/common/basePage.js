
const { until, By } = require('selenium-webdriver');

const TIME_OUT = 2000;
const ERROR_MSG = 'Looking for element but not found, check locator id please !!!';


// visit a webpage
const visit = async (theUrl) => {
    return await global.getDriver.get(theUrl);
};
// wait 
const wait = async (waitTime) => {
    return await global.getDriver.manage().setTimeouts({ implicit: waitTime });
};
// quit current session
const quit = async () => {
    return await global.getDriver.quit();
};
const isElementPresent = async (id) => {
    try {
        await global.getDriver.wait(until.elementLocated(By.id(id)), TIME_OUT, ERROR_MSG);
        await global.getDriver.findElement(By.id(id));
        return true;
    }
    catch (Error) {
        return false;
    }
};

// wait and find a specific element with it's id
const findElementById = async (id) => {
    await global.getDriver.wait(until.elementLocated(By.id(id)), TIME_OUT, ERROR_MSG);
    return await global.getDriver.findElement(By.id(id));
};

// wait and find a specific element with it's name
const findElementByName = async (name) => {
    await global.getDriver.wait(until.elementLocated(By.name(name)), TIME_OUT, ERROR_MSG);
    return await global.getDriver.findElement(By.name(name));
};
// wait and find a specific element with xPath
const findElementByXPath = async (xPath) => {
    await global.getDriver.wait(until.elementLocated(By.xpath(xPath)), TIME_OUT, ERROR_MSG);
    return await global.getDriver.findElement(By.xpath(xPath));
};
// wait and find a specific element with linkText
const findElementByLink = async (linkName) => {
    await global.getDriver.wait(until.elementLocated(By.linkText(linkName)), TIME_OUT, ERROR_MSG);
    return await global.getDriver.findElement(By.linkText(linkName));
};
// fill input web elements
const write = async (el, txt) => {
    return await el.sendKeys(txt);
};


module.exports = {
    visit,
    quit,
    write,
    wait,
    findElementById,
    findElementByName,
    findElementByXPath,
    findElementByLink,
    isElementPresent
};