module.exports = {
    client: {
        service: {
            name: 'my-app',
            url: 'http://localhost:4000/query',
        },
        includes: [
            "src/**/*.ts",
            "src/**/*.tsx",
            "src/**/*.vue",
            "tests/**/*.ts",
            "tests/**/*.tsx"
        ],
    },
}