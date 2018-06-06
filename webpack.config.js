var Webpack = require("webpack");
var Glob = require("glob");
var CopyWebpackPlugin = require("copy-webpack-plugin");
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var ManifestPlugin = require("webpack-manifest-plugin");
var PROD = process.env.NODE_ENV || "development";
var CleanWebpackPlugin = require("clean-webpack-plugin");
var MinifyPlugin = require("babel-minify-webpack-plugin");

var entries = {
    application: [
        "babel-polyfill"
  ],
}

Glob.sync("./assets/**/*.*").reduce((_, entry) => {
  let key = entry.replace(/(\.\/assets\/(js|css|go)\/)|\.(jsx?|s[ac]ss|go)/g, '')
  if(key.startsWith("_") || (/(jsx?|s[ac]ss|go)$/i).test(entry) == false) {
    return
  }

  if( entries[key] == null) {
    entries[key] = [entry]
    return
  }

  entries[key].push(entry)
})

var mode = 'development';
if (PROD != "development") {
    mode = 'production';
}

module.exports = {
  mode: mode,
  entry: entries,
  output: {
    filename: "[name].[hash].js",
    path: `${__dirname}/public/assets`
  },

  plugins: [
    new CleanWebpackPlugin(
      [
        "public/assets"
      ], {
          verbose: false,
      }
    ),
    new ExtractTextPlugin("[name].[hash].css"),
    new CopyWebpackPlugin(
      [{
        from: "./assets",
        to: ""
      }], {
        copyUnmodified: true,
        ignore: ["css/**", "js/**"]
      }
    ),
    new Webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false
    }),
    new ManifestPlugin({
      fileName: "manifest.json"
    })
  ],

  module: {
    rules: [
      {
        test: /\.jsx?$/,
        loader: "babel-loader",
        exclude: /node_modules/
      },
      {
        test:   /\.post\.css/,
        use: ExtractTextPlugin.extract({
          use: [
            {
              loader: 'css-loader',
              options: {
                importLoaders: 1,
                // minimize: true,
                sourceMap: true,
              }
            },
            'postcss-loader',
          ]
        })
      },
      {
        test: /\.(woff|woff2|ttf|svg)(\?v=\d+\.\d+\.\d+)?$/,
        use: "url-loader"
      },
      {
        test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
        use: "file-loader"
      },
      {
        test: /\.go$/,
        use: "gopherjs-loader"
      }
    ]
  }
};

if (PROD != "development") {
  module.exports.plugins.push(
      new MinifyPlugin({
    })
  );
}
