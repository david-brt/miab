// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface Platform {}
  }
  namespace svelteHTML {
    interface HTMLAttributes<T> {
      'on:clickoutside'?: (event: CustomEvent<T>) => void;
      'on:feltesuccess'?: (event: CustomEvent) => void;
      'on:felteerror'?: (event: CustomEvent) => void;
    }
  }
}

export {};
