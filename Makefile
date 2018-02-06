.PHONY: spellcheck
spellcheck:
	npm test

.PHONY: fix-spelling
fix-spelling:
	npm run fix
