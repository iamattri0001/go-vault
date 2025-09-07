import { DeriveKey } from "@/utils/encryption";
import { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true); // <-- NEW
  const navigate = useNavigate();

  const redirectToAuthPage = () => {
    navigate("/auth");
  };

  // Load user from localStorage once
  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      try {
        const userObj = JSON.parse(storedUser);
        if (userObj) {
          setUser(userObj);
        } else {
          redirectToAuthPage();
        }
      } catch {
        redirectToAuthPage();
      }
    } else {
      redirectToAuthPage();
    }
    setLoading(false); // <-- mark as done
  }, []);

  // Only redirect if on /auth and logged in
  useEffect(() => {
    if (!loading && user && window.location.pathname === "/auth") {
      navigate("/dashboard");
    }
  }, [user, loading, navigate]);

  const login = async (user, password) => {
    setUser(user);
    localStorage.setItem("user", JSON.stringify(user));
    const encryptionKey = await DeriveKey(
      password,
      salts.encryption_salt,
      true
    );
    localStorage.setItem(user.username + "_encryption_key", encryptionKey);
    navigate("/dashboard");
  };

  const logout = () => {
    setUser(null);
    localStorage.clear();
    redirectToAuthPage();
  };

  return (
    <UserContext.Provider value={{ user, setUser, login, logout, loading }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUserContext = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserProvider");
  }
  return context;
};
