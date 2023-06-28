local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
	vim.fn.system({
		"git",
		"clone",
		"--filter=blob:none",
		"https://github.com/folke/lazy.nvim.git",
		"--branch=stable", -- latest stable release
		lazypath,
	})
end
vim.opt.rtp:prepend(lazypath)

vim.g.mapleader = " "

require("lazy").setup({
	require("plugins.config.lsp"),
	require("plugins.config.mason"),
	require("plugins.config.nvim-cmp"),
	require("plugins.config.dashboard"),
	require("plugins.config.conjure"),
	require("plugins.config.telescope"),
	require("plugins.config.trouble"),
	require("plugins.config.which_key"),
	require("plugins.config.toggleterm"),
	require("plugins.config.ultisnips"),
	require("plugins.config.mini"),
	require("plugins.config.neoscroll"),
	require("plugins.config.tags"),
	require("plugins.config.emmet"),
	require("plugins.config.neotest"),
	require("plugins.config.todo"),
	require("plugins.config.lightspeed"),
	require("plugins.config.treesitter"),
	require("plugins.config.null-ls"),
	require("plugins.config.devicons"),
	require("plugins.config.bufferline"),
	require("plugins.config.gitsigns"),
	require("plugins.config.ufo"),
	require("plugins.config.autopairs"),
	require("plugins.config.wordmotion"),
	require("plugins.config.glow"),
	require("plugins.config.diff"),
	require("plugins.config.neotree"),
	require("plugins.config.comment"),
	require("plugins.config.sleuth"),
	require("plugins.config.blame"),
	require("plugins.config.colors"),
	require("plugins.config.kanagawa"),
	require("plugins.config.lualine"),
	require("plugins.config.dap"),
	require("plugins.config.neodev"),
	require("plugins.config.notify"),
	require("plugins.config.lsp-signature"),
	require("plugins.config.oil"),
	require("plugins.config.vim-test"),
	require("plugins.config.better-escape"),
	require("plugins.config.tint"),
	require("plugins.config.noice"),
	"RRethy/nvim-treesitter-endwise",
	"editorconfig/editorconfig-vim",
	"nvim-lua/popup.nvim",
	"nvim-lua/plenary.nvim",
	"junegunn/goyo.vim",
	"tpope/vim-obsession",
	"junegunn/fzf.vim",
	"junegunn/fzf",
	"alvan/vim-closetag",
	"moll/vim-bbye",
	"xolox/vim-misc",
	"tpope/vim-repeat",
	"elixir-editors/vim-elixir",
})
