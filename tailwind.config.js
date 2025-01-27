/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/**/*.{html,ts,css}",
  ],
  theme: {
    extend: {
      zIndex: {
        '1000': '1000'
      },
      height: {
        'calc': 'calc(100vh - 80px)'
      },
      borderWidth: {
        '1': '1px'
      },
      borderColor: {
        'light': '#E8E8E8'
      },
      colors: {
        'primary': '#FF6900',
        'gray': '#4E4B48'
      },
      fontFamily: {
        'AvantGardEF-Bold': 'AvantGardEF-Bold',
        'AvantGardEF-Book': 'AvantGardEF-Book',
        'ITCAvantGardeStd-Demi': 'ITCAvantGardeStd-Demi'
      },
      textColor: {
        "fontBlack": '#202020'
      },
      backgroundColor: {
        'customBlue': '#262262'
      },
      visibility: ["group-hover"]
    },
  },
  plugins: [],
  important: true,
}

