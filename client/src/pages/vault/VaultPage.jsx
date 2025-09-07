import React, { useEffect, useState } from "react";
import Navbar from "../common/Navbar";
import { GetPasswordsList } from "@/api/handlers";
import { toast } from "sonner";
import { CapitalizeFirstLetter } from "@/utils/text";
import { AddNewPassword } from "./components/AddNewPassword";
import PasswordsList from "./components/PasswordsList";
import { useParams } from "react-router-dom";

const VaultPage = () => {
  const { vault_id } = useParams();
  const [passwords, setPasswords] = useState([]);
  useEffect(() => {
    const fetchPasswords = async () => {
      const response = await GetPasswordsList(vault_id);
      if (response.success) {
        setPasswords(response.data.passwords);
      } else {
        if (response.error !== "") {
          toast.error(CapitalizeFirstLetter(response.error));
        } else {
          toast.error("Failed to fetch passwords.");
        }
      }
    };
    fetchPasswords();
  }, []);

  return (
    <section className="px-4 py-2">
      <Navbar />
      <div className="mt-4 min-h-[80vh] flex flex-col">
        <div className="flex items-end justify-end">
          <AddNewPassword setPasswords={setPasswords} vault_id={vault_id} />
        </div>
        <PasswordsList
          passwords={passwords}
          setPasswords={setPasswords}
          vault_id={vault_id}
        />
      </div>
    </section>
  );
};

export default VaultPage;
