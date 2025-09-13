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
import { useUserContext } from "@/contexts/UserContext";

export function AddNewPassword({ setPasswords, vault_id }) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [website, setWebsite] = useState("");
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);

  const { encKey } = useUserContext();

  const resetFields = () => {
    setTitle("");
    setDescription("");
    setUsername("");
    setPassword("");
    setWebsite("");
    setIsPasswordVisible(false);
  };

  const handleCreatePassword = async () => {
    const response = await CreatePassword({
      title,
      description,
      username,
      password,
      website,
      vault_id,
      encryptionKey: encKey,
    });
    if (response.success) {
      toast.success("Password created successfully!");
      setPasswords((prev) => [...prev, response.data.password]);
      resetFields();
    } else {
      if (response.error !== "") {
        toast.error(CapitalizeFirstLetter(response.error));
      } else {
        toast.error("Logout failed. Please try again.");
      }
    }
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="">Create</Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Add New Entry</AlertDialogTitle>
          <AlertDialogDescription>
            Give this entry a name and (optionally) a short note. Your login
            details will be encrypted and stored securely inside this vault.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <div className="space-y-2">
          <Input
            placeholder="Vault entry name (e.g. Work Email, Bank Account)"
            value={title}
            onChange={(e) => {
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
            onChange={(e) => setUsername(e.target.value)}
          />

          <div className="flex gap-x-2">
            <Input
              placeholder="Password (hidden)"
              type={isPasswordVisible ? "text" : "password"}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Toggle onClick={() => setIsPasswordVisible((prev) => !prev)}>
              <BiShow />
            </Toggle>
          </div>

          <Input
            placeholder="Website or App (e.g. https://example.com)"
            value={website}
            onChange={(e) => setWebsite(e.target.value)}
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
}
