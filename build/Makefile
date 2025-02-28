export PACKER_CACHE_DIR := .cache
export PACKER_VERSION := 1.8.2
export CENTOS_ISO := 2004.01

## PACKER
fetch:
	mkdir -p ${PACKER_CACHE_DIR}/${CENTOS_ISO} || :
	test -f ${PACKER_CACHE_DIR}/id_rsa_vagrant \
	    || curl -L https://raw.githubusercontent.com/mitchellh/vagrant/master/keys/vagrant \
		-o ${PACKER_CACHE_DIR}/id_rsa_vagrant
	chmod 600 ${PACKER_CACHE_DIR}/id_rsa_vagrant
	test -f ${PACKER_CACHE_DIR}/${CENTOS_ISO}/CentOS7.ova \
		|| wget --progress=dot:giga https://app.vagrantup.com/centos/boxes/7/versions/${CENTOS_ISO}/providers/virtualbox.box \
		-O ${PACKER_CACHE_DIR}/${CENTOS_ISO}/CentOS7.ova
	test -f ${PACKER_CACHE_DIR}/${CENTOS_ISO}/box.ovf \
		|| tar -C ${PACKER_CACHE_DIR}/${CENTOS_ISO} -xf ${PACKER_CACHE_DIR}/${CENTOS_ISO}/CentOS7.ova

deps:
	mkdir -p ${PACKER_CACHE_DIR} ~/bin || :
	curl https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_amd64.zip -o ${PACKER_CACHE_DIR}/packer.zip
	unzip -o ${PACKER_CACHE_DIR}/packer.zip -d ~/bin

pmm2-ovf: fetch
	packer build -only virtualbox-ovf packer/pmm2.json

pmm2-digitalocean:
	packer build -only digitalocean -var 'single_disk=true' packer/pmm2.json

pmm2-azure:
	packer build -only azure-arm packer/pmm2.json

pmm2-ami:
	docker run --rm -v ${HOME}/.aws:/root/.aws -v `pwd`:/build -w /build hashicorp/packer:${PACKER_VERSION} \
			build -var 'pmm_client_repos=original experimental' \
                  -var 'pmm_client_repo_name=percona-experimental-x86_64' \
                  -var 'pmm2_server_repo=experimental' \
                  -only amazon-ebs -color=false \
				  packer/pmm2.json
pmm2-ami-rc:
	docker run --rm -v ${HOME}/.aws:/root/.aws -v `pwd`:/build -w /build hashicorp/packer:${PACKER_VERSION} \
			build -var 'pmm_client_repos=original testing' \
				  -var 'pmm_client_repo_name=percona-testing-x86_64' \
				  -var 'pmm2_server_repo=testing' \
				  -only amazon-ebs '-color=false' \
				  packer/pmm2.json


## //PACKER
