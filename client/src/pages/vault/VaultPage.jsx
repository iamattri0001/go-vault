import React, { useEffect, useState } from "react";
import Navbar from "../common/Navbar";
import { GetPasswordsList } from "@/api/handlers";
import { toast } from "sonner";
import { CapitalizeFirstLetter } from "@/utils/text";
import { AddNewPassword } from "./components/AddNewPassword";
import PasswordsList from "./components/PasswordsList";
import { useParams } from "react-router-dom";
import { IoMdArrowBack } from "react-icons/io";

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

  const goBack = () => {
    window.history.back();
  };

  return (
    <section className="px-6 py-2">
      <Navbar />
      <div className="mt-4 min-h-[80vh] flex flex-col">
        <div className="flex items-end justify-between">
          <div
            className="text-sm text-foreground/65 cursor-pointer"
            onClick={goBack}
          >
            <IoMdArrowBack className="inline mr-1 mb-1" />
            <span className="underline">Back to Dashboard</span>
          </div>
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
