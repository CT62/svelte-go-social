/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        'barlow': ['"Barlow Semi Condensed"', 'sans-serif'],
      },
    },
  },
  plugins: [],
}

