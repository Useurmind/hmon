const merge = require('webpack-merge');
const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

const common = require('./webpack.config.js')({
    isProduction: false
});

module.exports = merge(common, {
    devtool: 'source-map',
    mode: "production",
    // throws errors with unexpected token
    plugins: [
        new UglifyJSPlugin()
    ],
});