#!/bin/bash

if [ -f /ansible/ssh/config ]; then
  cp /ansible/ssh/config /home/ansible/.ssh/config
  chown ansible:ansible /home/ansible/.ssh/config
  chmod 600 /home/ansible/.ssh/config
fi

/ansible/ansible-http-go -c /ansible/config.yml