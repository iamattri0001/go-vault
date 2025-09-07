import { DeriveKey, GenerateSalt } from "@/utils/encryption";
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

export const LogoutUser = async () => {
  return await MakeApiCall("/v1/auth/logout", "GET", {});
};

export const GetVaultsList = async () => {
  return await MakeApiCall("/v1/vault/list", "GET", {});
};

export const CreateVault = async (title, description) => {
  return await MakeApiCall("/v1/vault/create", "POST", {
    title,
    description,
  });
};
