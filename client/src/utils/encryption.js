export async function DeriveKey(masterPassword, salt, shouldReturnCrypto) {
  const encoder = new TextEncoder();

  // Import raw password
  const passwordKey = await crypto.subtle.importKey(
    "raw",
    encoder.encode(masterPassword),
    { name: "PBKDF2" },
    false,
    ["deriveKey"]
  );

  // Derive a key for AES-GCM
  const derivedKey = await crypto.subtle.deriveKey(
    {
      name: "PBKDF2",
      salt: encoder.encode(salt),
      iterations: 100_000, // strong enough, but can increase further
      hash: "SHA-256",
    },
    passwordKey,
    { name: "AES-GCM", length: 256 }, // usable key
    true,
    ["encrypt", "decrypt"]
  );

  if (shouldReturnCrypto) {
    return derivedKey; // directly usable for AES-GCM
  }

  return KeyToHex(derivedKey); // hex representation for server storage
}

export async function KeyToHex(key) {
  const raw = await crypto.subtle.exportKey("raw", key);
  return Array.from(new Uint8Array(raw))
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}

export function GenerateSalt(length = 16) {
  const array = new Uint8Array(length);
  crypto.getRandomValues(array);
  return Array.from(array, (b) => b.toString(16).padStart(2, "0")).join("");
}

// Encrypts a string using AES-GCM
export async function EncryptString(plainText, key) {
  const encoder = new TextEncoder();
  const data = encoder.encode(plainText);

  // Random IV (12 bytes recommended for AES-GCM)
  const iv = crypto.getRandomValues(new Uint8Array(12));

  // Encrypt
  const encrypted = await crypto.subtle.encrypt(
    { name: "AES-GCM", iv },
    key,
    data
  );

  // Combine IV + ciphertext, then base64 encode
  const buffer = new Uint8Array(iv.byteLength + encrypted.byteLength);
  buffer.set(iv, 0);
  buffer.set(new Uint8Array(encrypted), iv.byteLength);

  return btoa(String.fromCharCode(...buffer)); // base64 string
}

// Decrypts a string using AES-GCM
export async function DecryptString(encryptedBase64, key) {
  const data = Uint8Array.from(atob(encryptedBase64), (c) => c.charCodeAt(0));

  // Extract IV (first 12 bytes) and ciphertext (rest)
  const iv = data.slice(0, 12);
  const ciphertext = data.slice(12);

  // Decrypt
  const decrypted = await crypto.subtle.decrypt(
    { name: "AES-GCM", iv },
    key,
    ciphertext
  );

  const decoder = new TextDecoder();
  return decoder.decode(decrypted);
}
