# CodeSpawner
An Online Judge Platform API for Coding Contests Using Golang and Rust
Work in Progress, rewriting the whole API again

### For Development and Contribution

  1. Install Golang(>=1.13) and Rust (>= 1.38) in your system.

  2. Install postgres(<= 11.x )

  3. Create a user named **codespawner** and database in postgres
     - sudo su - postgres 
     - CREATEUSER --interactive --pwprompt 
     - Fill the detail roleName=codespawner, password=..., superuserPermission=y
     - createdb -O codespawner databaseName(codespawner_root)
     - psql databaseName(codespawner_root)

  4. Copy the content of *sql/1-create-database.sql* and paste it inside psql of codespawner after connecting to it.

 5. Install Air go-package for live reload of services

    ```shell
    go get -u github.com/cosmtrek/air
    ```

6. Install nginx ( i.e. sudo apt-get install nginx )

7. Include nginx.conf file under following directories /etc/nginx/nginx.conf
    - root/nginx.conf
    - hurdle/nginx.conf 

8. Run *setup-conf.sh* to copy sample files.

9. Finally Run **run-all.py**.

**P.S- Readlly interested join us over at** [Discord TheCodeSpawners](https://discord.gg/NySVa7A) **to walk the talk!**
