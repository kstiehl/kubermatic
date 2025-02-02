#!/usr/bin/env bash

# Copyright 2020 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

### This script is used as a postsubmit job and updates the CI cluster
### after every commit to master.

set -euo pipefail

cd $(dirname $0)/../..
source hack/lib.sh

export DEPLOY_STACK=${DEPLOY_STACK:-kubermatic}
export GIT_HEAD_HASH="$(git rev-parse HEAD | tr -d '\n')"

if [[ "${DEPLOY_STACK}" == "kubermatic" ]]; then
  ./hack/ci/push-images.sh
fi

echodate "Getting secrets from Vault"
retry 5 vault write \
  --format=json auth/approle/login \
  role_id=${VAULT_ROLE_ID} secret_id=${VAULT_SECRET_ID} > /tmp/vault-token-response.json
export VAULT_TOKEN="$(cat /tmp/vault-token-response.json | jq .auth.client_token -r)"
export KUBECONFIG=/tmp/kubeconfig
export VALUES_FILE=/tmp/values.yaml

vault kv get -field=kubeconfig dev/seed-clusters/ci.kubermatic.io > ${KUBECONFIG}
vault kv get -field=values.yaml dev/seed-clusters/ci.kubermatic.io > ${VALUES_FILE}
echodate "Successfully got secrets for dev from Vault"

echodate "Deploying ${DEPLOY_STACK} stack to ci.kubermatic.io"
TILLER_NAMESPACE=kubermatic \
  DEPLOY_NODEPORT_PROXY=false \
  DEPLOY_ALERTMANAGER=false \
  DEPLOY_MINIO=false ./hack/ci/deploy.sh master ${VALUES_FILE}
echodate "Successfully deployed ${DEPLOY_STACK} stack to ci.kubermatic.io"
