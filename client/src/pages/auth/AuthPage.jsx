import React, { useState } from "react";
import { Login } from "./components/Login";
import { Register } from "./components/Register";

const AuthPage = () => {
  const [authType, setAuthType] = useState("login");
  return (
    <section className="min-h-screen flex items-center justify-center">
      {authType === "login" ? (
        <Login setAuthType={setAuthType} />
      ) : (
        <Register setAuthType={setAuthType} />
      )}
    </section>
  );
};

export default AuthPage;
