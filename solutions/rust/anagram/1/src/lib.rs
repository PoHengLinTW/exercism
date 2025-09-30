use std::collections::{HashMap, HashSet};

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut res: HashSet<&'a str> = HashSet::new();
    let mut counter: HashMap<String, i32> = HashMap::new();

    for char in word.chars() {
        let c = char.to_lowercase().to_string();
        let value = *counter.get(&c).unwrap_or(&0) + 1;
        counter.insert(c, value);
    }

    for candidate in possible_anagrams {
        let mut local_counter: HashMap<String, i32> = HashMap::new();

        for char in candidate.chars() {
            let c = char.to_lowercase().to_string();
            let value = *local_counter.get(&c).unwrap_or(&0) + 1;
            local_counter.insert(c, value);
        }

        if counter == local_counter && word.to_lowercase() != *candidate.to_lowercase() {
            res.insert(&candidate);
        }
    }

    res
}
