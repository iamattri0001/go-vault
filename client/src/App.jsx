import React from "react";
import { Button } from "./components/ui/button";
import AuthPage from "./pages/auth/AuthPage";
import { Route, Routes } from "react-router-dom";
import DashboardPage from "./pages/dashboard/DashboardPage";

const App = () => {
  return (
    <Routes>
      <Route path="/auth" element={<AuthPage />} />
      <Route path="/dashboard" element={<DashboardPage />} />
    </Routes>
  );
};

export default App;
