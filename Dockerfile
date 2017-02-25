FROM python:3.5.3

ENV ANSIBLE_VERSION 2.2.0

##### ansible #####

ADD requirements.txt requirements.txt

RUN set -ex \
	&& buildDeps=' \
		gcc \
		python-dev \
		libffi-dev \
		libssl-dev \
	' \
	&& apt-get update \
	&& apt-get install -y --no-install-recommends $buildDeps \
	&& apt-get install -y openssh-server sshpass git sudo \
	&& rm -rf /var/lib/apt/lists/* \
	&& pip install --no-cache-dir -r requirements.txt \
	&& pip install ansible==$ANSIBLE_VERSION \
	&& apt-get purge -y --auto-remove $buildDeps \
	&& mkdir -p /etc/ansible

##### setup ansible-http #####

RUN set -x \
	&& mkdir -p /ansible/work /ansible/keys /ansible/ssh \
	&& groupadd -g 1001 ansible \
	&& useradd -m -c "ansible user" -g ansible -s /bin/bash -d /home/ansible -u 1001 ansible \
	&& mkdir -p /home/ansible/.ssh \
	&& chmod 700 /home/ansible/.ssh

ADD ansible/ansible-http-go /ansible/ansible-http-go
#ADD ansible/ssh_config /home/ansible/.ssh/config
ADD ansible/sudoers_ansible /etc/sudoers.d/ansible
ADD ansible/start.sh /ansible/start.sh

RUN chown -R ansible:ansible /ansible \
	&& chown -R ansible:ansible /home/ansible \
	&& chmod +x /ansible/ansible-http-go \
	&& chmod +x /ansible/start.sh

WORKDIR /ansible
VOLUME /ansible/keys
VOLUME /ansible/ssh
USER ansible

#CMD ["/ansible/ansible-http-go", "-c", "/ansible/config.yml"]
CMD ["./start.sh"]
