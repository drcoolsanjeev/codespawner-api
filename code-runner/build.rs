
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let a = tonic_build::compile_protos("proto/runner.proto")?;
    println!("{:?}", a);
    Ok(())
}