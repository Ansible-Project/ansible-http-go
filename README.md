ansible-http
---
ansible docker container with http wrapper

- ansible with http
- multiple execution with various branches during same time

usage
---
- registered docker container 'ansible-http'
```
./make.sh
```
- prepare config.yml in ansible/ dir (ref. sample_config.yml)
- edit requirements.txt you need more python library
- docker run
```
docker run -it --name ansible-http --rm -p 1323:1323 ansible-http
```
- access test
```
curl localhost:1323/version
```

(under construction...)

---

command
---
- GET /version
  - Return ansible-http version, status: 200
- GET /ansible/version
  - Return ansible version, status: 200
- POST /ansible/playbook/run
  - Execute ansible-playbook
  - Return ansible execution result
  - parameters:
    - playbook (required) - playbook file name
    - inventory (required) - inventory file path
    - limit - ansible limit option
    - tags - ansible tags option
    - skiptags - ansible skip-tags option
    - extravars - ansible extra-vars option
    - verbose - ansible verbose option
    - dir - change directory option
    - branch - git clone branch specification option

Feature
---
- test
- GET /ansible/playbook/list
- POST /ansible/command/run
- no block mode (not wait for execution result)
- notification task
- update echo version

LICENSE
---
MIT
