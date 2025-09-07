import React from "react";
import { Route, Routes } from "react-router-dom";

import AuthPage from "./pages/auth/AuthPage";
import DashboardPage from "./pages/dashboard/DashboardPage";
import VaultPage from "./pages/vault/VaultPage";
import ProtectedRoute from "./pages/common/ProtectedRoute";

const App = () => {
  return (
    <Routes>
      <Route path="/auth" element={<AuthPage />} />

      <Route
        path="/dashboard"
        element={
          <ProtectedRoute>
            <DashboardPage />
          </ProtectedRoute>
        }
      />

      <Route
        path="/vault/:vault_id"
        element={
          <ProtectedRoute>
            <VaultPage />
          </ProtectedRoute>
        }
      />
    </Routes>
  );
};

export default App;
