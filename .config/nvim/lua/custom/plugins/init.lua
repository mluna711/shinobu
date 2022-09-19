return {
  ["folke/which-key.nvim"] = {
    disable = false
  },
  ["goolord/alpha-nvim"] = {
    disable = false,
    override_options = require("custom.plugins.alpha")
  },
  ["numToStr/Comment.nvim"] = false,
  ["tpope/vim-commentary"] = {},
  ["neovim/nvim-lspconfig"] = {
    config = function ()
      require("plugins.configs.lspconfig")
      require("custom.plugins.lspconfig")
    end
  },
  ["williamboman/mason.nvim"] = {
    override_options = {
      ensure_installed = { "elixir-ls", "typescript-language-server", "solargraph" }
    }
  },
  ["sbdchd/neoformat"] = {},
  ["tpope/vim-obsession"] = {},
  ["hrsh7th/nvim-cmp"] = require("custom.plugins.cmp"),
  ["karb94/neoscroll.nvim"] = require("custom.plugins.neoscroll"),
  ["nvim-telescope/telescope.nvim"] = require("custom.plugins.telescope"),
  ["ggandor/lightspeed.nvim"] = {}
}
