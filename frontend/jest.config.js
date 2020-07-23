module.exports = {
    preset: "@vue/cli-plugin-unit-jest",
    testMatch: ["**/?(*.)+(spec|test).[jt]s?(x)"],
    collectCoverageFrom: [
        "src/**/*.{js,jsx,vue}",
        "!**/node_modules/**",
        "!**/vendor/**",
        "!**/dist/**",
    ],
};
