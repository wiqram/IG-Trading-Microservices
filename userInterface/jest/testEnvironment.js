const NodeEnvironment = require('jest-environment-node');

class CustomEnvironment extends NodeEnvironment {
    constructor(config) {
        super(config);
    }

    async setup() {
        console.log('******* Setting Up Driver ******* ');

        await super.setup();
        const driver = null;
        this.global.driver = driver;
    }

    async teardown() {
        console.log('******* Shutting Down Driver ******* ');
        await super.teardown();

    }

    runScript(script) {
        return super.runScript(script);
    }
}
module.exports = CustomEnvironment