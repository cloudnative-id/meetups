#!/bin/bash

set -e
echo $GITHUB_AUTH_SECRET > ~/.git-credentials && chmod 0600 ~/.git-credentials
git config --global credential.helper store
git config --global user.email "cncfidbot@users.noreply.github.com"
git config --global user.name "CNCF ID Bot"

git add .
git commit -m "Updating meetups list on `date`, commit ${TRAVIS_COMMIT} and job ${TRAVIS_JOB_NUMBER}" || true
git push
