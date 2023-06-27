use std::{collections::HashMap, io};


fn zenit_polar_substitutions() -> HashMap<char, char> {
    HashMap::from([

        ('z', 'p'),
        ('e', 'o'),
        ('n', 'l'),
        ('i', 'a'),
        ('t', 'r'),
        ('p', 'z'),
        ('o', 'e'),
        ('l', 'n'),
        ('a', 'i'),
        ('r', 't'),
    ])
}

fn cipher(text: &str, replacement_map: HashMap<char, char>) -> String {
    let mut message: String = String::new();

    for letter in text.chars() {
        let lowercase_letter: char = letter.to_ascii_lowercase();
        let modified_letter: Option<char> = replacement_map.get(&lowercase_letter).cloned();

        if let Some(replacement) = modified_letter {
            let final_letter: char = if letter.is_ascii_uppercase() {
                replacement.to_ascii_uppercase()
            } else {
                replacement
            };
            message.push(final_letter);
        } else {
            message.push(letter);
        }
    }

    message
}

fn prompt() -> String {
    let mut message: String = String::new();
    io::stdin()
        .read_line(&mut message)
        .expect("Could not parse your text.");
    message
}

fn main() {
    println!("Message to encrypt: ");
    let encrypted_message: String = cipher(&prompt(), zenit_polar_substitutions());

    println!("Encrypted message: {}", encrypted_message);

}
