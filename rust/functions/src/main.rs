fn another_function() {
    println!("hello functions")
}

fn double(x: i32) -> i32 {
    let y = { x };

    y * y
}

fn main() {
    another_function();
    println!("doubel five: {}", double(5))
}
