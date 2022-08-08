use wasm_bindgen::prelude::*;
use letterboxed::solve as lbsolve;
use js_sys::{Array, JsString, Uint8Array};

#[wasm_bindgen]
pub fn solve(allowed: Uint8Array, words: Array) -> Array {
  let mut wordlist = Vec::new();
  for i in 0..words.length() {
    wordlist.push(words.get(i).as_string().unwrap());
  }

  let v = allowed.to_vec();
  let all = [[v[0], v[1], v[2]], [v[3], v[4], v[5]], [v[6], v[7], v[8]], [v[9], v[10], v[11]]];
  let res = lbsolve(all, &wordlist);

  let out = Array::new();
  for v in res {
    out.push(&JsString::from(v));
  }
  out
}