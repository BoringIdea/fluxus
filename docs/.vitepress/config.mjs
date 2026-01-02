import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Fluxus Docs',
  description: 'Fluxus NFT Documentation - An NFT liquidity solution with intelligent pricing',
  lang: 'en-US',

  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Introduction', link: '/introduce' },
      { text: 'Getting Started', link: '/getting-started/quickstart' },
      { text: 'Basics', link: '/basics/' }
    ],

    sidebar: {
      '/': [
        {
          text: 'Introduction',
          items: [
            { text: 'Fluxus Introduction', link: '/introduce' },
            { text: 'Why Choose Fluxus', link: '/why-choose-fluxus' }
          ]
        },
        {
          text: 'Getting Started',
          items: [
            { text: 'Quickstart', link: '/getting-started/quickstart' }
          ]
        },
        {
          text: 'Basics',
          items: [
            { text: 'Architecture', link: '/basics/architecture' },
            { text: 'Core', link: '/basics/core' },
            { text: 'Roadmap', link: '/basics/roadmap' }
          ]
        }
      ],
      '/basics/': [
        {
          text: 'Basics',
          items: [
            { text: 'Architecture', link: '/basics/architecture' },
            { text: 'Core', link: '/basics/core' },
            { text: 'Roadmap', link: '/basics/roadmap' }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/HashIdea/fluxus' }
    ],

    footer: {
      message: 'Released under the BLS License.',
      copyright: 'Copyright © 2026 Fluxus'
    },

    search: {
      provider: 'local'
    }
  }
})

