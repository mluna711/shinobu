return {
	"olimorris/persisted.nvim",
	-- lazy = true,
	-- cmd = {
	-- 	"SessionLoad",
	-- 	"SessionSave"
	-- },
	dependencies = {
		"goolord/alpha-nvim",
	},
	init = function()
		local wk = require("which-key")
		wk.register({
			s = {
				name = "Sessions",
				s = {
					"<cmd>SessionSave<CR>",
					"Save current session"
				},
				l = {
					"<cmd>SessionLoad<CR>",
					"Load session"
				},
        d = {
          function()
            vim.cmd("SessionDelete")
            vim.cmd("silent only")
            vim.cmd("silent bufdo bwipeout")
            vim.cmd("Alpha")
						vim.cmd("normal gg")
          end,
          "Delete current session",
        }
			}
		}, {
			prefix = "<Leader>"
		})
	end,
	config = function()
		require("persisted").setup({
			save_dir = vim.fn.expand(vim.fn.stdpath("data") .. "/sessions/"), -- directory where session files are saved
			silent = true, -- silent nvim message when sourcing session file
			use_git_branch = false, -- create session files based on the branch of the git enabled repository
			autosave = true, -- automatically save session files when exiting Neovim
			should_autosave = function()
				return vim.bo.filetype ~= "alpha"
			end, -- function to determine if a session should be autosaved
			autoload = true, -- automatically load the session for the cwd on Neovim startup
			on_autoload_no_session = nil, -- function to run when `autoload = true` but there is no session to load
			follow_cwd = true, -- change session file name to match current working directory if it changes
			allowed_dirs = nil, -- table of dirs that the plugin will auto-save and auto-load from
			ignored_dirs = nil, -- table of dirs that are ignored when auto-saving and auto-loading
			telescope = { -- options for the telescope extension
				reset_prompt_after_deletion = true, -- whether to reset prompt after session deleted
			},
		})
	end,
}
