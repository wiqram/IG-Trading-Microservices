import webpack from 'webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';

import paths from './paths';
import rules from './rules';

module.exports = {
    entry: [
           paths.entryPath,
           'webpack-dev-server/client?http://0.0.0.0:8092/'
        ],
    module: {
        rules
    },
    stats: {
      modules: false,
    },
    resolve: {
        modules: ['src', 'node_modules'],
        extensions: ['*', '.js', '.jsx']
    },
    plugins: [
        //new webpack.ProgressPlugin(),
      new webpack.ProgressPlugin((percentage, message) => {
        // An idea to show the line only if percentage is divisible by 25
        if (percentage * 100 % 25 === 0) {
          console.log(`${(percentage * 100).toFixed()}% ${message}`);
        }
      }),
        new HtmlWebpackPlugin({
            template: paths.templatePath,
            minify: {
                collapseInlineTagWhitespace: true,
                collapseWhitespace: true,
                preserveLineBreaks: true,
                minifyURLs: true,
                removeComments: true,
                removeAttributeQuotes: true
            }
        })
    ]
};
