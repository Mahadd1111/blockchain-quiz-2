
----- Mahad Ahmed 20i-0426 README File for Git Commands Used ------

// Initializing Git repo

git init

// Adding Files to Staging

git add .

// Committing Files in git

git commit -m "commit message"

// Creating And moving to new branch

git checkout -b dev

// Pushing to Dev remote origin

git push origin dev

// Merging Dev into Main

git checkout main
git merge dev

// Pushing to main

git add .
git commit -m "Merged with Dev"
git push origin main

