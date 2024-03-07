import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "K Language Document",
  description: "K Language Document",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Get Started', link: '/get-started' }
    ],

    sidebar: [
      {
        text: 'Get Started',
        items: [
          { text: 'Get Started', link: '/get-started' },
        ]
      },
      {
        text: 'Syntax', 
        items: [
          { text: 'Loop', link: '/syntax/loop' },
        ] 
      },
      {
        text: 'Editor',
        items: [
          { text: 'Jupyter', link: '/editor/jupyter' },
          { text: 'VSCode', link: '/editor/vscode' },
        ] 
      },
    ],

    socialLinks: [
      // { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
