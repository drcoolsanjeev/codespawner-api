use tonic::{transport::Server, Request, Response, Status};

use runner::code_runner_server::{CodeRunner, CodeRunnerServer};
use runner::{HelloReply, HelloRequest};

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

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:1405".parse().unwrap();
    let code_runner = MyRunner::default();

    println!("Runner Server listening on {}", addr);

    Server::builder()
        .add_service(CodeRunnerServer::new(code_runner))
        .serve(addr)
        .await?;

    Ok(())
}
