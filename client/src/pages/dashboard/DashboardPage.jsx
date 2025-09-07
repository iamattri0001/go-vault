import { Section } from "lucide-react";
import React, { useEffect, useState } from "react";
import Navbar from "./components/Navbar";
import VaultsList from "./components/VaultsList";
import { AddNewVault } from "./components/AddNewVault";
import { GetVaultsList } from "@/api/handlers";
import { toast } from "sonner";
import { CapitalizeFirstLetter } from "@/utils/text";

const DashboardPage = () => {
  const [vaults, setVaults] = useState([]);
  console.log(vaults);
  useEffect(() => {
    const fetchVaults = async () => {
      const response = await GetVaultsList();
      if (response.success) {
        setVaults(response.data.vaults);
      } else {
        if (response.error !== "") {
          toast.error(CapitalizeFirstLetter(response.error));
        } else {
          toast.error("Failed to fetch vaults.");
        }
      }
    };
    fetchVaults();
  }, []);
  return (
    <section className="px-4 py-2">
      <Navbar />
      <div className="mt-4 min-h-[80vh] flex flex-col">
        <div className="flex items-end justify-end">
          <AddNewVault setVaults={setVaults} />
        </div>
        <VaultsList vaults={vaults} setVaults={setVaults} />
      </div>
    </section>
  );
};

export default DashboardPage;
