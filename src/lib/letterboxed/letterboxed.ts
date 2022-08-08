import init, {solve as lbsolve} from "$lib/letterboxed/letterboxed_wasm/pkg/letterboxed_wasm";

export async function load() {
  await init()
}

export function solve(allowed: string[], words: string[]): string[] {
  const buff = new ArrayBuffer(12);
  const arr = new Uint8Array(buff);
  
  let i = 0;
  for (let v of allowed) {
    arr[i] = v.charCodeAt(0);
    arr[i+1] = v.charCodeAt(1);
    arr[i+2] = v.charCodeAt(2);
    i += 3;
  }

  return lbsolve(arr, words);
}