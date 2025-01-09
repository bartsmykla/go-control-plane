#!/bin/bash

current_tag=$(git tag --list 'v[0-9].[0-9]*.[0-9]*' --merged origin/release --sort=-creatordate | head -n 1)
main_tag=$(git tag --list 'v[0-9].[0-9]*.[0-9]*' --merged origin/main --sort=-creatordate | head -n 1)

# Remove the prefix `v` and `-kong-*` suffix for comparison
released_tag="${current_tag%-kong-*}"
currentVersion="${released_tag#v}"

# Main branch won't have -kong-* suffix
newVersion="${main_tag#v}"

echo "Current release tag prefix: $released_tag"
echo "Upstream tag: $main_tag"
new_tag="${main_tag}-kong-1"
if [[ $current_tag != $new_tag ]]; then
  echo "New tag: $new_tag"
else
  echo "Tags are equal, no need to release"
  exit 0
fi

# Convert versions to arrays
IFS='.' read -r -a currentParts <<< "$currentVersion"
IFS='.' read -r -a newParts <<< "$newVersion"

echo "released_tag=$released_tag" >> $GITHUB_OUTPUT
echo "main_tag=$main_tag" >> $GITHUB_OUTPUT

echo "Current base tag: $currentVersion"
echo "Upstream tag: $newVersion"

# Compare each part
for i in 0 1 2; do
  if [[ ${newParts[i]:-0} -gt ${currentParts[i]:-0} ]]; then
    echo "The new tag is higher."
    echo "New version tag: $new_tag"
    echo "new_tag=$new_tag" >> $GITHUB_OUTPUT
    exit 0
  elif [[ ${newParts[i]:-0} -lt ${currentParts[i]:-0} ]]; then
    echo "The current tag is higher. That shouldn't be that case, please fix tagging."
    exit 1
  fi
done

echo "The tags are equal."
exit 0
