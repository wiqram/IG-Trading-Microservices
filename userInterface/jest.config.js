module.exports = {
    coverageThreshold: {
        global: {
            branches: 85,
            functions: 85,
            lines: 85,
            statements: 85
        }
    },
    collectCoverageFrom: [
        'src/**/*.{js,jsx}'
    ],
    coveragePathIgnorePatterns: [
        'src/index.js',
    ],
    testPathIgnorePatterns: [
        '<rootDir>/test/e2e',
    ],
    collectCoverage: true,
    coverageReporters: [
        'json',
        'text',
        'html'
    ]
};