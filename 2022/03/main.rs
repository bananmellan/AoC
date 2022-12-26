use std::fs;
use std::env;

fn input() -> Vec<String> {
	fs::read_to_string(&env::args().collect::<Vec<String>>()[1])
		.expect("poo").split("\n").map(|s| s.into()).collect()
}

fn main() {
	let lines = input();

	let bags = lines.iter().map(
		|line| line.chars().map(
			|item| u32::from(item).to_be_bytes().iter().sum::<u8>()
				- if item.is_lowercase() { 96 } else { 64 - 26 }
	));

	for bag in bags {
		for priority in bag {
			println!("{}", priority);
		}
	}
}
