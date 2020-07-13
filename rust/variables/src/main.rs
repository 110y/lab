fn main() {
    let mut x = 5;
    println!("The value of x is {}", x);
    x = 6;
    println!("The value of x is {}", x);

    let shadowed = 1;
    println!("The value of `shadowed` is {}", shadowed);

    let shadowed = 2;
    println!("The value of `shadowed` is {}", shadowed);

    let tup = (500, 1.1);
    let (x, y) = tup;

    println!("The value of `x` is {}, `y` is {}", x, y);

    println!("The value of `0` is {}, `1` is {}", tup.0, tup.1);

    let a = [1, 2];

    println!("The value of `0` is {}, `1` is {}", a[0], a[1]);
}
