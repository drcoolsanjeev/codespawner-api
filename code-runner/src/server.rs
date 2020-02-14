use tonic::{transport::Server, Request, Response, Status};

use runner::code_runner_server::{CodeRunner, CodeRunnerServer};
use runner::{HelloReply, HelloRequest};

#[macro_use]
extern crate diesel;
extern crate dotenv;

use diesel::prelude::*;
use diesel::pg::PgConnection;
use dotenv::dotenv;
use std::env;

pub mod schema;
pub mod models;

use self::models::UsersCode;

pub mod runner {
    tonic::include_proto!("runner");
}

#[derive(Default)]
pub struct MyRunner {}

#[tonic::async_trait]
impl CodeRunner for MyRunner {
    async fn say_hello(
        &self,
        request: Request<HelloRequest>,
    ) -> Result<Response<HelloReply>, Status> {
        println!("Got a request from {:?}", request.remote_addr());

        let reply = runner::HelloReply {
            message: format!("Hello {}!", request.into_inner().name),
        };
        Ok(Response::new(reply))
    }
}

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .expect(&format!("Error connecting to {}", database_url))
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {

    use self::schema::users_code::dsl::*;

    let addr = "[::1]:1405".parse().unwrap();
    let code_runner = MyRunner::default();

    let connection = establish_connection();
    let results = users_code    
        .load::<UsersCode>(&connection)
        .expect("Error loading posts");

    // println!("Displaying {:?} posts", results);
    // for post in results {
    //     println!("-----------\n");
    //     println!("{}", post.code_buffer);
    // }
    println!("Runner Server listening on {}", addr);

    Server::builder()
        .add_service(CodeRunnerServer::new(code_runner))
        .serve(addr)
        .await?;

    Ok(())
}
