const expect = require('expect');
const { HomePage } = require('./po');
const { wait, quit } = require('../../common/basePage')


describe('Check Products Home Page ', async function () {
    const homePage = new HomePage();
    const mockSa = {
        id: 1,
        productId: "SA001",
        productName: "SAVING A/C",
        sortCode: "11-22-11",
        accountNumber: "012345671",
        currentBalance: 10000.59,
        availableBalance: 10000
    }
    before(() => {
        wait(20000)
    });

    it('Check the [More Info Button] for Saving Accounts', async () => {
        const buttonReturn = await homePage.checkMoreInfoButton_sa();
        expect(await buttonReturn.getText()).toEqual('More Details');
        buttonReturn.click();
        wait(20000)
    });

    after(() => {
        quit();
    });
});
