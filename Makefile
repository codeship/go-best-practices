.PHONY: spellcheck
spellcheck:
	yarn test

.PHONY: fix-spelling
fix-spelling:
	yarn run fix
