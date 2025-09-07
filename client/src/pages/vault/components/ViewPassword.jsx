import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { CapitalizeFirstLetter } from "@/utils/text";
import { useState } from "react";
import { toast } from "sonner";

import { BiShow } from "react-icons/bi";
import { Toggle } from "@/components/ui/toggle";
import { CreatePassword } from "@/api/handlers";
import { FaRegEdit } from "react-icons/fa";
import { DecryptString } from "@/utils/encryption";

const ViewPassword = ({ setPasswords, vault_id, password: passwordObject }) => {
  if (!passwordObject) {
    return null;
  }
  const [title, setTitle] = useState(passwordObject.title);
  const [description, setDescription] = useState(passwordObject.description);
  const [username, setUsername] = useState(passwordObject.username);
  const [password, setPassword] = useState(passwordObject.encrypted_password);
  const [decryptedPassword, setDecryptedPassword] = useState("");
  const [website, setWebsite] = useState(passwordObject.website);
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isEditing, setIsEditing] = useState(false);

  const handleCreatePassword = async () => {
    const response = await CreatePassword({
      title,
      description,
      username,
      password,
      website,
      vault_id,
    });
    if (response.success) {
      toast.success("Password created successfully!");
      setPasswords((prev) => [...prev, response.data.password]);
    } else {
      if (response.error !== "") {
        toast.error(CapitalizeFirstLetter(response.error));
      } else {
        toast.error("Logout failed. Please try again.");
      }
    }
  };

  const decryptPassword = async () => {
    const encryptionKey = localStorage.getItem(
      JSON.parse(localStorage.getItem("user")).username + "_encryption_key"
    );
    if (!encryptionKey) {
      toast.error("Missing encryption key. Please re-login.");
      return;
    }
    const decrypted = await DecryptString(password, encryptionKey);
    setDecryptedPassword(decrypted);
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="">Create</Button>
      </AlertDialogTrigger>
      <AlertDialogContent className={isEditing && "border-destructive"}>
        <AlertDialogHeader>
          <AlertDialogTitle className={"flex items-center justify-between"}>
            View Details
            <Button
              variant={isEditing ? "destructive" : "ghost"}
              onClick={() => {
                setIsEditing((prev) => {
                  if (prev) {
                    toast.info("Editing disabled. Changes won't be saved.");
                  } else {
                    toast.info(
                      "Editing enabled. You can now modify the fields."
                    );
                  }
                  return !prev;
                });
              }}
            >
              <FaRegEdit />
            </Button>
          </AlertDialogTitle>
          <AlertDialogDescription>
            These are the details of your saved login entry. You can review them
            securely here, and choose to edit if needed.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <div className="space-y-2">
          <Input
            placeholder="Vault entry name (e.g. Work Email, Bank Account)"
            value={title}
            onChange={(e) => {
              if (!isEditing) {
                toast.error("Enable editing to modify.");

                return;
              }
              if (e.target.value.length <= 30) {
                setTitle(e.target.value);
              } else {
                toast.error("Entry name cannot exceed 30 characters.");
              }
            }}
          />

          <Textarea
            placeholder="Notes (e.g. recovery hints, security question) – optional"
            value={description}
            onChange={(e) => {
              if (!isEditing) {
                toast.error("Enable editing to modify.");
                return;
              }
              if (e.target.value.length <= 100) {
                setDescription(e.target.value);
              } else {
                toast.error("Notes cannot exceed 100 characters.");
              }
            }}
          />

          <Input
            placeholder="Login / Email / Username"
            value={username}
            onChange={(e) => {
              if (!isEditing) {
                toast.error("Enable editing to modify.");

                return;
              }
              setUsername(e.target.value);
            }}
          />

          <div className="flex gap-x-2">
            <Input
              placeholder="Password (hidden)"
              type={isPasswordVisible ? "text" : "password"}
              value={password}
              onChange={(e) => {
                if (!isEditing) {
                  toast.error("Enable editing to modify.");
                  return;
                }
                setPassword(e.target.value);
              }}
            />
            <Toggle onClick={() => setIsPasswordVisible((prev) => !prev)}>
              <BiShow />
            </Toggle>
          </div>

          <Input
            placeholder="Website or App (e.g. https://example.com)"
            value={website}
            onChange={(e) => {
              if (!isEditing) {
                toast.error("Enable editing to modify.");
                return;
              }
              setWebsite(e.target.value);
            }}
          />
        </div>

        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={handleCreatePassword}>
            Save Password
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
};

export default ViewPassword;
