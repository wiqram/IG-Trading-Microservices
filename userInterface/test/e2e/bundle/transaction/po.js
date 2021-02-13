const locator = require('./locator');
const { findElementByXPath } = require('../../common/basePage')

class HomePage {

    async checkMoreInfoButton_sa() {
        let result = await findElementByXPath(locator.moreInfoButton_sa);
        return result;
    }

}

module.exports = { HomePage };