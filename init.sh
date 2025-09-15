#!/bin/bash
echo "Enter your module path (e.g., github.com/username/project):"
read module

go mod edit -module "$module"
grep -rl 'github.com/yourusername/templaterepo' . \
  | xargs sed -i "s|github.com/yourusername/templaterepo|$module|g"

go mod tidy
echo "Module path updated to $module"
