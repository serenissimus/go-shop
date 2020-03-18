const path = require('path');
const webpack = require('webpack');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const getEnvironment = () => {
  if ('production' === process.env.NODE_ENV) {
    return 'production';
  }
  return 'development';
}
const environment = getEnvironment();

const API_URL = {
  production: '',
  development: 'http://localhost:8080'
};


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
    new webpack.DefinePlugin({
      'process.env.API_URL': JSON.stringify(API_URL[environment]),
    }),
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
