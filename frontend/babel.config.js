module.exports = {
  plugins: ["@babel/plugin-syntax-dynamic-import"],
  presets: [["@vue/app", { useBuiltIns: "entry" }]],
};
