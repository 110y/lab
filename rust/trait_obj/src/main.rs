pub trait Draw {
    fn draw(&self);
}

pub struct Button {
    // pub width: u32,
// pub height: u32,
// pub label: String,
}

impl Draw for Button {
    fn draw(&self) {
        // code to actually draw a button
    }
}

struct SelectBox {
    // pub width: u32,
// pub height: u32,
// pub options: Vec<String>,
}

impl Draw for SelectBox {
    fn draw(&self) {
        // code to actually draw a select box
    }
}

fn main() {
    println!("Hello, world!");

    let mut f: &dyn Draw = &Button {};

    f = &SelectBox {};
}
