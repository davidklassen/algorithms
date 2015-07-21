use std::io;
use std::io::prelude::*;

fn qsort(mut arr: Vec<i32>) -> Vec<i32> {
    if arr.len() == 0 {
        return vec![];
    }

    let pivot = arr.remove(0);
    let mut left = qsort(arr.iter().map(|n| n.clone()).filter(|n| *n <= pivot).collect::<Vec<_>>());
    let right = qsort(arr.iter().map(|n| n.clone()).filter(|n| *n > pivot).collect::<Vec<_>>());

    left.push(pivot);
    left.extend(right.iter().cloned());
    left
}

fn main() {
    let stdin = io::stdin();
    let arr = stdin
        .lock()
        .lines()
        .map(|n| n.unwrap().parse::<i32>().unwrap() )
        .collect::<Vec<_>>();

    for n in qsort(arr) {
        println!("{}", n);
    }
}
