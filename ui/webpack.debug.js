const webpack = require('webpack');
const merge = require('webpack-merge');

const common = require('./webpack.config.js')({
    isProduction: false
});

module.exports = merge(common, {
    // throws errors with unexpected token
    plugins: [
    ],
});