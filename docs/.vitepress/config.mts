import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  locales: {
    root: {
      label: 'English',
      lang: 'en'
    },
    zh: {
      label: '中文',
      lang: 'zh',
      title: "K 语言文档",
      themeConfig: {
        nav: [
          { text: '文档主页', link: '/' },
          { text: '快速开始', link: '/get-started' },
        ],
        sidebar: [
        {
          text: '快速开始', items: [
            { text: '快速开始', link: '/zh/get-started' },
          ]
        },
        {
        text: '语法', 
        items: [
          { text: '定义变量', link: '/zh/syntax/define-var' },
          { text: '循环', link: '/zh/syntax/loop' },
          { text: '基本的内置函数', link: '/zh/syntax/ess-internal-fx'}
        ] 
      },
      {
        text: '编辑器集成',
        items: [
          { text: 'Jupyter', link: '/zh/editor/jupyter' },
          { text: 'VSCode', link: '/zh/editor/vscode' },
        ] 
      },
    ],
      }
    }
  },
  title: "K Language Document",
  description: "K Language Document",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Get Started', link: '/get-started' },
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
          { text: 'Define Variable', link: '/syntax/define-var' },
          { text: 'Loop', link: '/syntax/loop' },
          { text: 'Essential Intenal Functions', link: '/syntax/ess-internal-fx'}
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
