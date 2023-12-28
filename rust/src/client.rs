use hello::hello_client::HelloClient;
use hello::Request;

pub mod greeter {
    tonic::include_proto!("hello");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = HelloClient::connect("http://[::1]:10001").await?;

    let request = tonic::Request::new(Request {
        id: '0',
    });

    println!("Sending request to gRPC Server...");
    let response = client.get_info(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}
