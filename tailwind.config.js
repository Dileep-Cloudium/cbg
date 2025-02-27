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
      colors: {
        'primary': '#FF6900',
        'gray': '#4E4B48'
      },
      fontFamily: {
        'Brother1816-Bold': 'Brother1816-Bold',
        'Brother1816-Book': 'Brother1816-Book'
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

