# CodeSpawner
An Online Judge Platform API for Coding Contests Using Golang and Rust
Work in Progress, rewriting the whole API again

### For Development and Contribution

  1. Install Golang(>=1.13) and Rust (>= 1.38) in your system.

  2. Install PostgreSQL database. Create a user named **codespawner** in postgres db. Copy the content of *sql/1-create-database.sql* and paste it inside psql of codespawner after connecting to it.

 3. Install Air go-package for live reload of services

    ```shell
    go get -u github.com/cosmtrek/air
    ```

4. Run *setup-conf.sh* to copy sample files.

5. Finally Run **run-all.py**.

**P.S- Readlly interested join us over at** [Discord TheCodeSpawners](https://discord.gg/NySVa7A) **to walk the talk!**
