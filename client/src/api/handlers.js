import { DeriveKey, EncryptString, GenerateSalt } from "@/utils/encryption";
import { MakeApiCall } from "./call";

export const RegisterUser = async (username, password) => {
  const authSalt = GenerateSalt();
  const encryptionSalt = GenerateSalt();

  return await MakeApiCall("/v1/auth/register", "POST", {
    username,
    password: await DeriveKey(password, authSalt),
    encryption_salt: encryptionSalt,
    auth_salt: authSalt,
  });
};

export const LoginUser = async (username, password) => {
  const salts = await GetSalts(username);
  if (salts === null) {
    return {
      success: false,
      error: "failed to get salts",
    };
  }

  return await MakeApiCall("/v1/auth/login", "POST", {
    username,
    password: await DeriveKey(password, salts.auth_salt),
  });
};

export const GetSalts = async (username) => {
  if (!username) {
    username = (await GetUsername()) || "";
  }
  if (localStorage.getItem("salts_" + username)) {
    return JSON.parse(localStorage.getItem("salts_" + username));
  }

  const response = await MakeApiCall("/v1/salts" + "/" + username, "GET", {});

  if (response.success) {
    localStorage.setItem("salts_" + username, JSON.stringify(response.data));
    return response.data;
  }

  return null;
};

export const GetUsername = async () => {
  const user = await JSON.parse(localStorage.getItem("user"));
  return user?.username;
};

export const LogoutUser = async () => {
  return await MakeApiCall("/v1/auth/logout", "GET", {});
};

export const GetVaultsList = async () => {
  return await MakeApiCall("/v1/vault/list", "GET", {});
};

export const CreateVault = async (title, description) => {
  return await MakeApiCall("/v1/vault/", "POST", {
    title,
    description,
  });
};

export const GetPasswordsList = async (vault_id) => {
  return await MakeApiCall("v1/vault/" + vault_id, "GET", {});
};

export const CreatePassword = async ({
  title,
  description,
  username,
  password,
  website,
  vault_id,
  encryptionKey,
}) => {
  const encrypted_password = await EncryptString(password, encryptionKey);

  return await MakeApiCall("/v1/password/", "POST", {
    title,
    description,
    username,
    password: encrypted_password,
    website,
    vault_id,
  });
};

export const UpdatePassword = async ({
  id,
  title,
  description,
  username,
  password,
  website,
  vault_id,
  encryptionKey,
}) => {
  const encrypted_password = await EncryptString(password, encryptionKey);
  return await MakeApiCall("/v1/password/", "PUT", {
    id,
    title,
    vault_id,
    description,
    username,
    password: encrypted_password,
    website,
  });
};
