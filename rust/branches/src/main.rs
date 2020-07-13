fn main() {
    let number = 3;

    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }

    let x = if true { 5 } else { 6 };
    println!("x: {}", x);
}
