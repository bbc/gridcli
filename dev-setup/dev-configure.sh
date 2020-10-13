#!/bin/bash
set -eou
YELLOW='\033[1;33m'
NC='\033[0m' # no colour - reset console colour

generateDevProperties() {
    echo -e "${YELLOW}Installing grid dependencies in grid/Brewfile...${NC}"
    brew bundle -v

    #Generate property files
    echo -e "${YELLOW}Generating properties files in /etc/gu/${NC}"
    echo -e "${YELLOW}This requires sudo access, enter password if prompted ...${NC}"
    sudo mkdir /etc/gu
    sudo chown -R "$(whoami)" /etc/gu
    cd scripts/generate-dot-properties || echo "cannot find path scripts/generate-dot-properties" && exit 1
    sed -i '' "s/domainRoot: '',/domainRoot: 'example.com',/g" config.json5
    sed -i '' '48,52d' config.json5
    sed -i '' "47i\\
    google: { tracking: { id: 'idbbc$(whoami)' } }," config.json5
    npm install
    npm run generate-properties

    # stash the changes to config.json5
    git stash
    popd || echo "cannot find path scripts/generate-dot-properties" && exit 1
}

amendProperties() {
    pushd "/etc/gu" || echo "cannot find path /etc/gu" && exit 1

    # Add properties
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a auth.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a cropper.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a kahuna.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a collections.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a image-loader.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a metadata-editor.properties > /dev/null
    echo 'permissions.bucket=media-service-dev-permissions-bucket' | sudo tee -a thrall.properties > /dev/null

    sudo echo 'transcoded.mime.types=image/tiff' | sudo tee -a image-loader.properties

    # Edit properties
    sed -i '' "s/origin.crops=media-service-dev-imageoriginbucket-pxo4mg24rknq/origin.crops=media-service-dev-imageoriginbucket-pxo4mg24rknq.s3-eu-west-1.amazonaws.com/" kahuna.properties
    sed -i '' "s/publishing.image.host=media-service-dev-imageoriginbucket-pxo4mg24rknq.s3.amazonaws.com/publishing.image.host=media-service-dev-imageoriginbucket-pxo4mg24rknq.s3-eu-west-1.amazonaws.com/" cropper.properties
    popd "/etc/gu" || echo "cannot find path /etc/gu" &&  exit 1
}

configureNginx() {
    brew tap guardian/homebrew-devtools
    brew install guardian/devtools/dev-nginx
    echo -e "${YELLOW}Setting up nginx servers and certs ...${NC}"
    dev-nginx setup-app nginx-mappings.yml
}

main() {
    generateDevProperties
    amendProperties
    configureNginx
}

main