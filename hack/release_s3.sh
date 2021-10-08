#!/usr/bin/env bash
set -euo pipefail

upload_prefix="s3://lstoll-lds-content-public/assets/idp"

if [[ -v CI ]]; then
    ref=${GITHUB_REF##*/}
    sha=$GITHUB_SHA
else
    ref=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
    sha=$(git rev-parse --verify HEAD)
fi

ref=$(echo "$ref" | tr / -)

echo "--> Packaging release for ref: $ref sha: $sha"

workdir=$(mktemp -d)
function cleanup {
  rm -rf "$workdir"
}
trap cleanup EXIT

echo "--> Building binary"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${workdir}/idp" .
(cd "$workdir" && zip idp.zip idp)
idp_sha=$(openssl dgst -sha256 -binary "${workdir}/idp.zip" | openssl enc -base64)

echo "--> Building terraform module"
# Update the lambda filename to match the sha we're building for. We do this
# even if it's for a "ref" named terraform bundle, to ensure what we upload
# alignes with the exact binary at time of upload.
cp -r terraform "${workdir}"
sed -i.bak "s/___LAMBDA_GIT_SHA___/${sha}/" "${workdir}/terraform/_variables.tf"
sed -i.bak "s/___LAMBDA_BASE64SHA256___/${idp_sha//\//\\/}/" "${workdir}/terraform/_variables.tf"
rm "${workdir}/terraform/_variables.tf.bak" || true
rm -r "${workdir}/terraform/.terraform"* || true
(cd "$workdir/terraform" && zip -r ../terraform.zip .)

# Disable this for now, doubles the upload amount. For now, the "entrypoint" is
# the terraform module, which just uses the sha version. If someone has their
# own entrypoint, they can source bins themselves.
#
# echo "--> Uploading idp to $upload_prefix/lambda/$ref.zip"
# aws s3 cp --acl public-read "${workdir}/idp.zip" "$upload_prefix/lambda/$ref.zip"

echo "--> Uploading idp to $upload_prefix/lambda/$sha.zip"
aws s3 cp --acl public-read "${workdir}/idp.zip" "$upload_prefix/lambda/$sha.zip"

echo "--> Uploading terraform to $upload_prefix/terraform/$ref.zip"
aws s3 cp --acl public-read "${workdir}/terraform.zip" "$upload_prefix/terraform/$ref.zip"

echo "--> Uploading terraform to $upload_prefix/terraform/$sha.zip"
aws s3 cp --acl public-read "${workdir}/terraform.zip" "$upload_prefix/terraform/$sha.zip"
