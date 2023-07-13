return {
	"nvim-tree/nvim-web-devicons",
	config = function()
		require("nvim-web-devicons").setup({
			override = {
				rs = {
					icon = "",
					name = "rust",
				},
				zig = {
					icon = "",
					name = "zig",
				},
				hurl = {
					icon = "󱘖",
					color = "#514e60",
					name = "hurl",
				},
				lisp = {
					icon = "󰚀",
					color = "#bef5b2",
					name = "lisp",
				},
				cshtml = {
					icon = "",
					color = "#953dab",
					name = "razor",
				},
				rb = {
					icon = "",
					color = "#a91401",
					name = "ruby",
				},
				go = {
					icon = "",
					color = "#86D4DE",
					name = "go",
				},
			},
		})
	end,
}
