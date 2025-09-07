import { Button } from "@/components/ui/button";
import React from "react";
import Logout from "./Logout";
import { SiVault } from "react-icons/si";

const Navbar = () => {
  return (
    <nav className="flex justify-between items-center py-4 border-b">
      <h1 className="text-2xl">
        <SiVault className="inline text-4xl mr-2 text-primary" />
        <span className="font-bold text-primary">GO-Vault</span>
      </h1>
      <div>
        <Logout />
      </div>
    </nav>
  );
};

export default Navbar;
