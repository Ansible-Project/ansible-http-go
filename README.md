ansible-http
---
ansible docker container with http wrapper

- ansible with http
- multiple execution with various branches during same time

build
---
- edit requirements.txt you need more python library
- when you do the following, the docker container 'ansible-http' is registered.
```
$ ./make.sh
```

configuration
---
```
$ cat config.yml
---
port: 1323
repository_url: git@github.com:you/your-ansible-playbook-repository.git
default_inventory: inventory_path_your_repository
default_verbose: -vv
default_branch: develop
```

|key|value|description|
|---|---|---|
|port|(1323)|port number for http|
|repository_url||ansible playbook git repository|
|default_inventory||default inventory path|
|default_verbose|(-v,-vv,-vvv,-vvvv)|default verbose option|
|default_branch|(develop,master,etc..)|default branch|

run
---
- docker run
```
docker run -it --name ansible-http --rm -p 1323:1323 -v $PWD/config.yml:/ansible/config.yml ansible-http
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
