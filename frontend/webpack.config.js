// webpack.config.js

const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

// const extractSass = new ExtractTextPlugin({
//     filename: "[name].[contenthash].css",
//     disable: process.env.NODE_ENV === "development"
// });

const config = {
    entry: './src/js/index.js',
    output: {
        path: path.resolve(path.dirname(__dirname), 'static','js'),
        filename: 'bundle.js'
    },
    resolve: { extensions: [ '.js', '.scss', 'css', 'jsx'] },
    plugins: [
        new MiniCssExtractPlugin() //{ filename: path.resolve('~','[name].css')}
        //if you want to pass in options, you can do so:
        //new ExtractTextPlugin({
        //  filename: 'style.css'
        //})
      ],
    module: {
        
        rules: [
            { test: /\.txt$/, use: 'raw-loader' },
            { test: /\.jsx?$/, 
                exclude: /node_modules/,
                use: 'babel-loader',
                 exclude: /node_modules/ },
            // {
            //     test: /\.s?css$/,
            //     exclude: /node_modules/,
            //     use: ["style-loader", "css-loader", "sass-loader"]
            //   }

              {
                test: /\.s?css$/,
                use: [{
                    loader:MiniCssExtractPlugin.loader,
                    options: {
                        sourceMap: true,
                        includePaths: ['./node_modules']
                    }
                },
                {
                    loader:'css-loader',
                    options: {
                        sourceMap: true,
                        includePaths: ['./node_modules']
                    }
                },
                {
                    loader:'sass-loader',
                    options: {
                        sourceMap: true,
                        includePaths: ['./node_modules']
                    }
                }
                 ],
                
              }

            // {
            //     test: /(\.css|\.scss)$/,
            //     // include: SRC,
            //         use: [
            //             {
            //                 loader: 'css-loader',
            //                 options: {
            //                     sourceMap: true,
            //                      importLoaders: 1
            //                 }
            //             },
            //             {
            //                 loader: 'postcss-loader',
            //                 options: {
            //                     sourceMap: true,
            //                      importLoaders: 1
            //                 }
            //             },
            //             {
            //                 loader: 'sass-loader',
            //                 options: {
            //                     sourceMap: true,
            //                      importLoaders: 1
            //                 }
            //             }
            //         ],
            //     }
            // ,
            // {
            //     test: /\.scss$/,
            //     use: extractSass.extract({
            //         use: [{
            //             loader: "css-loader"
            //         }, {
            //             loader: "sass-loader"
            //         }],
            //         // use style-loader in development 
            //         fallback: "style-loader"
            //     })
            // },
            // {
            //     test: /\.(css|scss)$/,
            //     exclude: /node_modules/,
            //     use: [
            //         {
            //             loader: 'style-loader',
            //             options: {
            //                 sourceMap: true,
            //                 importLoaders: 1
            //             }
            //         },
            //         {
            //             loader: 'css-loader',
            //             options: {
            //                 sourceMap: true,
            //                 importLoaders: 1
            //             }
            //         },

                    // {
                    //     loader: ExtractTextPlugin.extract({
                    //         fallback: 'style-loader',
                    //         use: 'css-loader?minimize'
                    //       })
                    // },
            //         {
            //             loader: 'postcss-loader',
            //             options: {
            //                 sourceMap: true,
            //                 importLoaders: 1
            //             }
            //         },
            //         {
            //             loader: 'sass-loader', options: {
            //                 sourceMap: true,
            //                 includePaths: ['./node_modules']
            //             }
            //         }
            //     ],
            // },
            // { test: /\.css$/, use: 'css-loader' }

        ]
    }
};

module.exports = config;
