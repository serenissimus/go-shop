const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  devServer: {
    compress: true,
    port: 3000,
  },
  entry: './src/index.tsx',
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        exclude: /node_modules/,
        loader: ['awesome-typescript-loader'],
      },
    ],
  },
  output: {
    filename: 'bundle.min.js',
    path: path.join(__dirname, '/dist'),
  },
  plugins: [
    new CopyWebpackPlugin([{
      from: 'public' 
    }]),
    new HtmlWebpackPlugin({
      template: path.join(__dirname, './public/index.html'),
    }),
  ],
  resolve: {
    alias: {
      src: path.resolve(__dirname, 'src'),
    },
    extensions: ['.ts', '.tsx', '.js'],
  },
}
