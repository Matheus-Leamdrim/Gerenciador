import '@testing-library/jest-dom/extend-expect';

module.exports = {
    setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
    testEnvironment: 'jsdom',
};