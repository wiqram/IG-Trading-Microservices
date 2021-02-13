const locator = require('./locator');
const { isElementPresent, findElementById } = require('../../common/basePage')

class HomePage {

    async isElementPresent(id) {
        return await isElementPresent(id);
    }
    async checkMainContainer() {
        return await findElementById(locator.mainContainer);
    }
    async checkProductName_sa() {
        let result = await findElementById(locator.productName_sa);
        return await result.getText();
    }
    async checkAccountNumber_sa() {
        let result = await findElementById(locator.accountNumber_sa);
        return await result.getText();
    }
    async checkCurrentBalance_sa() {
        let result = await findElementById(locator.currentBalance_sa);
        return await result.getText();
    }
    async checkCurrentBalanceLabel_sa() {
        let result = await findElementById(locator.currentBalanceLabel_sa);
        return await result.getText();
    }
    async checkAvailableBalance_sa() {
        let result = await findElementById(locator.availableBalance_sa);
        return await result.getText();
    }
    async checkAvailableBalanceLabel_sa() {
        let result = await findElementById(locator.availableBalanceLabel_sa);
        return await result.getText();
    }

    async checkMoreInfoButton_sa() {
        let result = await findElementById(locator.moreInfoButton_sa);
        return result;
    }

}

module.exports = { HomePage };