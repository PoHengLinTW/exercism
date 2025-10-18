use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let lowercase_word = word.to_lowercase();
    let sorted_word = sort(&lowercase_word);

    possible_anagrams
        .iter()
        .filter(|&&candidate| {
            let lowercase_candidate = candidate.to_lowercase();

            lowercase_word != lowercase_candidate && sorted_word == sort(&lowercase_candidate)
        })
        .cloned()
        .collect()
}

fn sort(word: &str) -> String {
    let mut v: Vec<char> = word.chars().collect();

    v.sort();

    v.into_iter().collect()
}
