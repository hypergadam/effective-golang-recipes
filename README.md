# Effective Go Recipes

My practice solutions and notes based on the book *Effective Go Recipes: Fast Solutions*.

The exercises and recipe topics come from the book. This repository contains my
implementations of those recipes.

Each recipe has its own executable in `chapter-N/recipeN`.
With Go and Task installed, run these commands from the repository root:

```sh
task                              # List available tasks
task run                          # Run chapter 1, recipe 1
task run CHAPTER=1 RECIPE=1        # Choose a chapter and recipe
task test                         # Test all recipes
task fmt                          # Format Go code
task vet                          # Check Go code
task check                        # Run tests and vet
```
