const path = require('path');
const CopyPlugin = require('copy-webpack-plugin');

module.exports = {
    mode: 'production',
    entry: './js/app.js',
    output: {
        path: path.resolve(__dirname, '../../public/js'),
        filename: 'app.bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.s[ac]ss$/i,
                use: [
                    // Creates `style` nodes from JS strings
                    'style-loader',
                    // Translates CSS into CommonJS
                    'css-loader',
                    // Compiles Sass to CSS
                    'sass-loader',
                ],
            },
        ],
    },
    plugins: [
        new CopyPlugin(
            [
                { from: path.resolve(__dirname, './templates'), to: path.resolve(__dirname, '../../public/templates') },
            ]
        ),
    ],
};