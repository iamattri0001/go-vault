import { GetSalts } from "@/api/handlers";
import { DeriveKey, hashKey } from "@/utils/encryption";
import { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [encKey, setEncKey] = useState(null);
  const [showUnlockModal, setShowUnlockModal] = useState(false);
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const redirectToAuthPage = () => navigate("/auth");

  // Load user from localStorage
  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      try {
        const userObj = JSON.parse(storedUser);
        if (userObj) setUser(userObj);
        else redirectToAuthPage();
      } catch {
        redirectToAuthPage();
      }
    } else {
      redirectToAuthPage();
    }
    setLoading(false);
  }, []);

  // Redirect if logged in but on /auth
  useEffect(() => {
    if (!loading && user && window.location.pathname === "/auth") {
      navigate("/dashboard");
    }
  }, [user, loading, navigate]);

  // Show unlock modal if no encKey but logged in
  useEffect(() => {
    if (user && !encKey) {
      setShowUnlockModal(true);
    }
  }, [user, encKey]);

  const login = async (loggedInUser, password) => {
    setUser(loggedInUser);
    localStorage.setItem("user", JSON.stringify(loggedInUser));

    const salts = await GetSalts(loggedInUser.username);
    const derivedKey = await DeriveKey(password, salts.encryption_salt, true);

    const keyHash = await hashKey(derivedKey);
    localStorage.setItem("encKeyHash", keyHash); // only a hash, not the key

    setEncKey(derivedKey);
    setShowUnlockModal(false);
    navigate("/dashboard");
  };

  const logout = () => {
    setUser(null);
    setEncKey(null);
    setShowUnlockModal(false);
    localStorage.clear();
    redirectToAuthPage();
  };

  const unlockVault = async () => {
    if (!user || !password) return;
    // debugger;

    const salts = await GetSalts(user.username);
    const candidateKey = await DeriveKey(password, salts.encryption_salt, true);
    const candidateHash = await hashKey(candidateKey);
    const storedHash = localStorage.getItem("encKeyHash");

    if (candidateHash === storedHash) {
      setEncKey(candidateKey);
      setShowUnlockModal(false);
      setPassword("");
      setError("");
    } else {
      setError("Invalid master password, try again");
    }
  };

  return (
    <UserContext.Provider
      value={{ user, setUser, login, logout, loading, encKey }}
    >
      {/* Blur background when locked */}
      <div className={showUnlockModal ? "blur-sm pointer-events-none" : ""}>
        {children}
      </div>

      {/* Unlock Modal */}
      <Dialog open={showUnlockModal}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>Unlock Vault</DialogTitle>
          </DialogHeader>
          <div className="space-y-3">
            <p className="text-sm text-muted-foreground">
              Enter your master password to unlock your vault.
            </p>
            <Input
              type="password"
              placeholder="Master password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && unlockVault()}
            />
            {error && <p className="text-sm text-red-500">{error}</p>}
            <Button onClick={unlockVault} className="w-full">
              Unlock
            </Button>

            {/* Logout button */}
            <Button variant="outline" className="w-full" onClick={logout}>
              Logout
            </Button>
          </div>
        </DialogContent>
      </Dialog>
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
