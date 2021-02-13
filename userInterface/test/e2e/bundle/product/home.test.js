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

    it('Check the [Root] for the page', async () => {
        const check = await homePage.isElementPresent('root');
        expect(check).toEqual(true);
    });

    it('Check the [Product Name] for Saving Accounts', async () => {
        const result = await homePage.checkProductName_sa();
        expect(result).toEqual(mockSa.productName);
    });

    it('Check the [Account Number] for Saving Accounts', async () => {
        const result = await homePage.checkAccountNumber_sa();
        expect(result).toEqual(mockSa.accountNumber + '/' + mockSa.sortCode);
    });

    it('Check the Current Balance [Currency Format] ,[Label] & [Value] for Saving Accounts', async () => {
        const label = await homePage.checkCurrentBalanceLabel_sa();
        const value = await homePage.checkCurrentBalance_sa();
        expect(label).toEqual('Current Balance');
        expect(value).toEqual("£10,000.59");
    });

    it('Check the Available Balance [Currency Format] ,[Label] & [Value] for Saving Accounts', async () => {
        const label = await homePage.checkAvailableBalanceLabel_sa();
        const value = await homePage.checkAvailableBalance_sa();
        expect(label).toEqual('Available Balance');
        expect(value).toEqual("£10,000.00");
    });

    it('Check the [More Info Button] for Saving Accounts', async () => {
        const buttonReturn = await homePage.checkMoreInfoButton_sa();
        expect(await buttonReturn.getText()).toEqual('More Details');
        buttonReturn.click();
    });

    after(() => {
        /** Close page if you want to */
    });
});
