use tonic::{transport::Server, Request, Response, Status};

use hello::hello_server::{Hello, HelloServer};
use hello::{Response, Request};

// Import the generated proto-rust file into a module
pub mod hello {
    tonic::include_proto!("hello");
}

// Implement the service skeleton for the "Hello" service
// defined in the proto
#[derive(Debug, Default)]
pub struct MyHello {}

// Implement the service function(s) defined in the proto
// for the Hello service (SayHello...)
#[tonic::async_trait]
impl Hello for MyHello {
    async fn get_info(
        &self,
        request: Request<Request>,
    ) -> Result<Response<Response>, Status> {
        println!("Received request from: {:?}", request);

        let response = hello::Response {
            message: format!("Hello {}!", request.into_inner().name).into(),
        };

        Ok(Response::new(response))
    }
}

// Runtime to run our server
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:10001".parse()?;
    let hello = MyHello::default();

    println!("Starting gRPC Server...");
    Server::builder()
        .add_service(HelloServer::new(hello))
        .serve(addr)
        .await?;

    Ok(())
}
