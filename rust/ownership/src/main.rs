fn main() {
    let s1 = String::from("hello");
    takes_ownership(s1);

    // let s2 = s1.clone();

    // println!("{}, world!", s1);

    let s2 = String::from("hello");

    let s3 = takes_and_gives_back(s2);

    println!("s3: {}", s3)
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn takes_and_gives_back(s: String) -> String {
    s
}
